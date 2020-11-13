package controllers

import (
	"github.com/Muha113/task-service/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AddNewCompany(context *gin.Context) {
	handlers.AddNewCompanyHandler(context.Writer, context.Request)
}

func UpdateCompany(context *gin.Context) {
	handlers.UpdateCompanyHandler(context.Writer, context.Request)
}

func FindCompany(context *gin.Context) {
	handlers.FindCompanyHandler(context.Writer, context.Request)
}

func UpdateCompanyWithFormData(context *gin.Context) {
	handlers.UpdateCompanyWithFormDataHandler(context.Writer, context.Request)
}

func DeleteCompany(context *gin.Context) {
	handlers.DeleteCompanyHandler(context.Writer, context.Request)
}

func GetCompanyEmployeesList(context *gin.Context) {
	handlers.GetCompanyEmployeesListHandler(context.Writer, context.Request)
}
