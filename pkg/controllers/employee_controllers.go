package controllers

import (
	"github.com/Muha113/task-service/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AddNewEmployee(context *gin.Context) {
	handlers.AddNewEmployeeHandler(context.Writer, context.Request)
}

func UpdateEmployee(context *gin.Context) {
	handlers.UpdateEmployeeHandler(context.Writer, context.Request)
}

func FindEmployee(context *gin.Context) {
	handlers.FindEmployeeHandler(context.Writer, context.Request)
}

func UpdateEmployeeWithFormData(context *gin.Context) {
	handlers.UpdateEmployeeWithFormDataHandler(context.Writer, context.Request)
}

func DeleteEmployee(context *gin.Context) {
	handlers.DeleteEmployeeHandler(context.Writer, context.Request)
}
