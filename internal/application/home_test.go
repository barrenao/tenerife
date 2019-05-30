package application

import (
	"github.com/sirupsen/logrus/hooks/test"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	logger, _ := test.NewNullLogger()
	req := httptest.NewRequest("Get","http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler := HomeHandler(logger)
	handler(w,req)
}