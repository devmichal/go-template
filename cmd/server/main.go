// @title notification-service
// @version 1.0
// @description microservice of notification email, sms and other notification
// @BasePath /
package server

import (
	"fmt"
	"github.com/chemiq/notification/internal/api/healthCheck"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func Command(c *cli.Context) (err error) {
	// Gin instance
	router := gin.Default()

	//newTemplate := template.New()

	template := router.Group("/template")
	{
		template.POST("/create")
	}
	// Route => handler
	router.GET("/health", healthCheck.Handler)

	//router.POST("/notification/create/template", newTemplate.Handler)

	// Start server
	fmt.Println(router.Run())

	return nil
}
