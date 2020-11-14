package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Muha113/task-service/pkg/apiproto"
	"github.com/Muha113/task-service/pkg/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func AddNewCompanyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connecting to task-repo...")
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := apiproto.NewTaskRepoClient(conn)

	buff := make([]byte, 10000)
	fmt.Println("Body ---> ", r.Body)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error("This shit error ->", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("dasdasdasdasd...")

	var compReq *apiproto.Company
	json.Unmarshal(buff, compReq)

	compResp, err := c.AddNewCompany(context.Background(), compReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(compResp.Status))
		return
	}

	msg, err := json.Marshal(compResp.Company)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.ApiResponse{
		Code:    200,
		Type:    "POST",
		Message: string(msg[:]),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseJSON); err != nil {
		logrus.Error(err)
		// sedn code status
	}
}

func UpdateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := apiproto.NewTaskRepoClient(conn)

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var compReq *apiproto.Company
	json.Unmarshal(buff, compReq)

	compResp, err := c.UpdateCompany(context.Background(), compReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(compResp.Status))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func FindCompanyHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := apiproto.NewTaskRepoClient(conn)

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var compReq *apiproto.Company
	json.Unmarshal(buff, compReq)

	compResp, err := c.AddNewCompany(context.Background(), compReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(compResp.Status))
		return
	}

	msg, err := json.Marshal(compResp.Company)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.ApiResponse{
		Code:    200,
		Type:    "GET",
		Message: string(msg[:]),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseJSON); err != nil {
		logrus.Error(err)
		// sedn code status
	}
}

func UpdateCompanyWithFormDataHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := apiproto.NewTaskRepoClient(conn)

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	optReq := &apiproto.OptionalFields{
		Data: string(buff[:]),
	}

	optResp, err := c.UpdateCompanyWithFormData(context.Background(), optReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(optResp.Status))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := apiproto.NewTaskRepoClient(conn)

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var compReq *apiproto.Company
	json.Unmarshal(buff, compReq)

	compResp, err := c.DeleteCompany(context.Background(), compReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(compResp.Status))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetCompanyEmployeesListHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := apiproto.NewTaskRepoClient(conn)

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var compReq *apiproto.Company
	json.Unmarshal(buff, compReq)

	compResp, err := c.GetCompanyEmployeesList(context.Background(), compReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(compResp.Status))
		return
	}

	msg, err := json.Marshal(compResp.Employee)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.ApiResponse{
		Code:    200,
		Type:    "GET",
		Message: string(msg[:]),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseJSON); err != nil {
		logrus.Error(err)
		// sedn code status
	}
}
