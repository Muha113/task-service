package routers

import (
	"github.com/Muha113/task-service/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.New()

	router.Use(gin.Recovery())

	v1 := router.Group("/employee")

	v1.POST("/", controllers.AddNewEmployee)
	v1.PUT("/", controllers.UpdateEmployee)
	v1.GET("/:employee_id", controllers.FindEmployee)
	v1.POST("/:employee_id", controllers.UpdateEmployeeWithFormData)
	v1.DELETE("/:employee_id", controllers.DeleteEmployee)

	v2 := router.Group("/company")

	v2.POST("/", controllers.AddNewCompany)
	v2.PUT("/", controllers.UpdateCompany)
	v2.GET("/:company_id", controllers.FindCompany)
	v2.POST("/:company_id", controllers.UpdateCompanyWithFormData)
	v2.DELETE("/:company_id", controllers.DeleteCompany)
	v2.GET("/:company_id/employee", controllers.GetCompanyEmployeesList)

	return router
}
