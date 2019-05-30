package application

import (
	"github.com/sirupsen/logrus"
	"net/http"

)


func HomeHandler(logger *logrus.Logger) http.HandlerFunc {

  return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Catch a home request")
		w.WriteHeader(http.StatusOK)
	}

}

