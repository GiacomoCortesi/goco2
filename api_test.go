package goco2

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetCO2SavingHandler(t *testing.T) {
	tests := []struct {
		name   string
		svc    Service
		status int
		want   *CO2Saving
	}{
		{
			name: "Success",
			svc: &MockService{
				Saving: &CO2Saving{
					Day:   100,
					Week:  700,
					Month: 3000,
					Year:  12000,
				},
			},
			status: http.StatusOK,
			want: &CO2Saving{
				Day:   100,
				Week:  700,
				Month: 3000,
				Year:  12000,
			},
		},
		// Add more test cases for other scenarios if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiServer := NewAPIServer(tt.svc)

			req, err := http.NewRequest("GET", "/saving", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := apiServer.GetCO2SavingHandler()

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.status, rr.Code)

			var saving CO2Saving
			err = json.Unmarshal(rr.Body.Bytes(), &saving)
			assert.NoError(t, err)

			assert.Equal(t, tt.want, &saving)
		})
	}
}

func TestAddInterventionHandler(t *testing.T) {
	tests := []struct {
		name   string
		svc    Service
		input  interface{}
		status int
		want   interface{}
	}{
		{
			name: "Success",
			svc: &MockService{
				Saving: &CO2Saving{
					Day:   100,
					Week:  700,
					Month: 3000,
					Year:  12000,
				},
			},
			input: Intervention{
				ID:   uuid.New(),
				Date: time.Now(),
			},
			status: http.StatusNoContent,
			want:   nil,
		},
		{
			name:   "InvalidInput",
			svc:    &MockService{},
			input:  "invalid json",
			status: http.StatusBadRequest,
			want:   map[string]interface{}{"error": "json: cannot unmarshal string into Go value of type goco2.Intervention"},
		},
		// Add more test cases for other scenarios if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiServer := NewAPIServer(tt.svc)

			var body []byte
			if tt.input != nil {
				var err error
				body, err = json.Marshal(tt.input)
				assert.NoError(t, err)
			}

			req, err := http.NewRequest("POST", "/intervention", bytes.NewBuffer(body))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := apiServer.AddInterventionHandler()

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.status, rr.Code)

			var response interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, tt.want, response)
		})
	}
}

type MockService struct {
	Saving *CO2Saving
}

func (m *MockService) GetCO2Saving(ctx context.Context) (*CO2Saving, error) {
	return m.Saving, nil
}

func (m *MockService) AddIntervention(ctx context.Context, intervention Intervention) error {
	return nil
}
