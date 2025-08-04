package apiservice

import (
	"bytes"
	"encoding/json"
	"github.com/Rumpelstiltski1/restapi/internal/app/model"
	"github.com/Rumpelstiltski1/restapi/store/testStore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(testStore.New(), sessions.NewCookieStore([]byte("secret")))
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

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := testStore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name     string
		payload  interface{}
		expected int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expected: http.StatusOK,
		}, {
			name:     "invalid payload",
			payload:  "invalid",
			expected: http.StatusBadRequest,
		}, {
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expected: http.StatusUnauthorized,
		}, {
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expected: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expected, rec.Code)
		})
	}
}
