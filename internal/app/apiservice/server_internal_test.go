package apiservice

import (
	"bytes"
	"encoding/json"
	"github.com/Rumpelstiltski1/restapi/store/testStore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(testStore.New())
	testCases := []struct {
		name     string
		payload  interface{}
		expected int
	}{{
		name: "valid",
		payload: map[string]string{
			"email":    "user@xample.org",
			"password": "password",
		},
		expected: http.StatusCreated,
	}, {
		name:     "invalid payload",
		payload:  "invalid",
		expected: http.StatusBadRequest,
	}, {
		name: "invalid parces",
		payload: map[string]string{
			"email": "invalid",
		},
		expected: http.StatusUnprocessableEntity,
	},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expected, rec.Code)
		})
	}

}
