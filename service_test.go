package goco2

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCO2Service_GetCO2Saving(t *testing.T) {
	tests := []struct {
		name          string
		savingCoeff   uint64
		interventions Interventions
		want          *CO2Saving
	}{
		{
			name:        "Success - same time",
			savingCoeff: 100,
			interventions: Interventions{
				uuid.New().String(): {Date: time.Now()},
				uuid.New().String(): {Date: time.Now()},
				uuid.New().String(): {Date: time.Now()},
				uuid.New().String(): {Date: time.Now()},
			},
			want: &CO2Saving{
				Day:   400,
				Week:  400,
				Month: 400,
				Year:  400,
			},
		},
		{
			name:        "Success - different times",
			savingCoeff: 100,
			interventions: Interventions{
				uuid.New().String(): {Date: time.Now().AddDate(0, 0, -1)},
				uuid.New().String(): {Date: time.Now().AddDate(0, -1, -2)},
				uuid.New().String(): {Date: time.Now().AddDate(-1, -1, 0)},
				uuid.New().String(): {Date: time.Now()},
			},
			want: &CO2Saving{
				Day:   100,
				Week:  200,
				Month: 200,
				Year:  300,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCO2Service(tt.savingCoeff)
			service.(*CO2Service).interventions = tt.interventions

			saving, err := service.GetCO2Saving(context.TODO())
			assert.NoError(t, err)
			assert.Equal(t, tt.want, saving)
		})
	}
}

func TestCO2Service_AddIntervention(t *testing.T) {
	tests := []struct {
		name        string
		savingCoeff uint64
		input       Intervention
		errExpected bool
	}{
		{
			name:        "Success",
			savingCoeff: 100,
			input:       Intervention{ID: uuid.New(), Date: time.Now()},
			errExpected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCO2Service(tt.savingCoeff)

			err := service.AddIntervention(context.TODO(), tt.input)

			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
