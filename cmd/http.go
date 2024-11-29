package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	r := gin.Default()

	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	r.GET("/health", healthcheckAPI.HealthcheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}
