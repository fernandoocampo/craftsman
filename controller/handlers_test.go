package controller_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fernandoocampo/craftsman/controller"
)

func TestFindByID(t *testing.T) {
	// GIVEN this craftsman id
	// Create a Request to pass to the handler
	req := httptest.NewRequest("GET", "/1", nil)
	req.Header.Add("Content-Type", "application/json")
	expected := "{ \"username\": \"theuser\", \"firstname\": \"The\", \"lastname\": \"user\", \"email\": \"theuser@domain.com\", \"email2\": \"theuser2@domain.com\", \"state\": 1, \"kwds\": \"cook chef ramen\", \"skills\": [{\"name\": \"cook\",	 \"stats\": {\"feedback\": 200, \"credits\": 50, \"rating\": 234}}] }"

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	resp := httptest.NewRecorder()

	// WHEN some user makes a search from http graphql protocol and its handled by the Router
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	controller.NewRouter().ServeHTTP(resp, req)

	// THEN the response should be to expected
	result := strings.TrimSpace(resp.Body.String())
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
