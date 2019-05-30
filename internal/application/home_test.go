package application

import (
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {

	req := httptest.NewRequest("Get","http://example.com/foo", nil)
	w := httptest.NewRecorder()
	HomeHandler(w, req)
}