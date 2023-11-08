package goco2

import (
	"context"
	"log"
	"time"
)

// LoggingService implements basic middleware logging for the microservice
type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCO2Saving(ctx context.Context) (saving *CO2Saving, err error) {
	defer func(start time.Time) {
		log.Printf("CO2 Saving: %+v, error: %v, took: %v\n", saving, err, time.Since(start))
	}(time.Now())
	return s.next.GetCO2Saving(ctx)
}

func (s *LoggingService) AddIntervention(ctx context.Context, intervention Intervention) (err error) {
	defer func(start time.Time) {
		log.Printf("Intervention: %+v, error: %v, took: %v\n", intervention, err, time.Since(start))
	}(time.Now())
	return s.next.AddIntervention(ctx, intervention)
}
