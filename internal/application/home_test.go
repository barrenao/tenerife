package application

import (
	"github.com/sirupsen/logrus"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	req := httptest.NewRequest("Get","http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler := HomeHandler(logger)
	handler(w,req)
}