package diagostics

import (
	"github.com/sirupsen/logrus"
	"net/http"
)



func RedianessHandler(logger *logrus.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Readianess Response")
		w.WriteHeader(http.StatusOK)
	}

}


func LiveanessHandler(logger *logrus.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Liveaness Response")
		w.WriteHeader(http.StatusOK)
	}

}
