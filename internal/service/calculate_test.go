package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_CalculateHandler(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantStatus int
		wantResp   string
	}{
		{
			name:       "Valid expression",
			expression: "1 + 2",
			wantStatus: http.StatusOK,
			wantResp:   `{"result":3}`,
		},
		{
			name:       "Complex expression",
			expression: "(1 + 2) * (3 + 4)",
			wantStatus: http.StatusOK,
			wantResp:   `{"result":21}`,
		},
		{
			name:       "Float numbers",
			expression: "3.5 + 4.2",
			wantStatus: http.StatusOK,
			wantResp:   `{"result":7.7}`,
		},
		{
			name:       "Invalid expression",
			expression: "2 +",
			wantStatus: http.StatusUnprocessableEntity,
			wantResp:   `{"error":"Expression is not valid"}`,
		},
		{
			name:       "Division by zero",
			expression: "4 / 0",
			wantStatus: http.StatusInternalServerError,
			wantResp:   `{"error":"Internal server error"}`,
		},
		{
			name:       "Empty expression",
			expression: "",
			wantStatus: http.StatusUnprocessableEntity,
			wantResp:   `{"error":"Expression is not valid"}`,
		},
		{
			name:       "Invalid JSON body",
			expression: `{"invalid_json"}`,
			wantStatus: http.StatusInternalServerError,
			wantResp:   `{"error":"Internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.name == "Invalid JSON body" {
				body = []byte(tt.expression)
			} else {
				body, err = json.Marshal(calculateRequest{Expression: tt.expression})
				assert.NoError(t, err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			s := &Service{}
			s.CalculateHandler(rr, req)

			assert.Equal(t, tt.wantStatus, rr.Code)
			assert.JSONEq(t, tt.wantResp, strings.TrimSpace(rr.Body.String()))
		})
	}
}
