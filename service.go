package goco2

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// Service interface defines the service layer of the microservice
type Service interface {
	GetCO2Saving(ctx context.Context) (*CO2Saving, error)
	AddIntervention(ctx context.Context, intervention Intervention) error
}

// CO2Service struct implements Service interface
type CO2Service struct {
	savingCoefficient uint64        // multiplication factor for the CO2 saving computation
	interventions     Interventions // efficiency upgrade interventions - in a more complex scenario we would inject a repository interface
}

func NewCO2Service(savingCoefficient uint64) Service {
	return &CO2Service{savingCoefficient: savingCoefficient, interventions: make(Interventions)}
}

// GetCO2Saving computes CO2 emission reduction for the day, week, month and year
// CO2 saving equals the saving coefficient (expressed in kg) multiplied by the number of interventions.
func (c *CO2Service) GetCO2Saving(ctx context.Context) (*CO2Saving, error) {
	return &CO2Saving{
		Day:   c.interventions.Today() * c.savingCoefficient,
		Week:  c.interventions.ThisWeek() * c.savingCoefficient,
		Month: c.interventions.ThisMonth() * c.savingCoefficient,
		Year:  c.interventions.ThisYear() * c.savingCoefficient,
	}, nil
}

// AddIntervention adds a new house energy upgrade intervention
// it generates an uuid if not provided
// it sets current time if not provided
func (c *CO2Service) AddIntervention(ctx context.Context, intervention Intervention) error {
	if intervention.ID == uuid.Nil {
		intervention.ID = uuid.New()
	}
	if intervention.Date.IsZero() {
		intervention.Date = time.Now()
	}
	c.interventions[intervention.ID.String()] = intervention
	return nil
}

// GetInterventions retrieves all energy efficiency upgrade interventions
func (c *CO2Service) GetInterventions() Interventions {
	return c.interventions
}

// GetInterventionByID retrieves the intervention information given its ID
// Returns ErrInterventionNotFound error if no intervention with specified ID is found.
func (c *CO2Service) GetInterventionByID(uid uuid.UUID) (*Intervention, error) {
	intervention, found := c.interventions[uid.String()]
	if !found {
		return nil, ErrInterventionNotFound
	}
	return &intervention, nil
}
