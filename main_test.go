package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	router := chi.NewRouter()
	addRoutes(router)

	testCases := append([]Route{}, routes...)

	// Add a 404 test case that only exists in the test
	testCases = append(testCases, Route{
		Name:       "404 Test",
		Path:       "/this-does-not-exist",
		Type:       HandlerRoute,
		TestStatus: http.StatusNotFound,
		Handler:    nil,
	})

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tc.Path, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if rr.Code != tc.TestStatus {
				t.Errorf("Expected status %d; got %d for route %s",
					tc.TestStatus, rr.Code, tc.Path)
			}
		})
	}
}
