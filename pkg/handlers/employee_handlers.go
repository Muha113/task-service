package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Muha113/task-service/pkg/apiproto"
	"github.com/Muha113/task-service/pkg/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func AddNewEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var emplReq *apiproto.Employee
	json.Unmarshal(buff, emplReq)

	c := apiproto.NewTaskRepoClient(conn)
	emplResp, err := c.AddNewEmployee(context.Background(), emplReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(emplResp.Status))
		return
	}

	// remake
	status := emplResp.Status
	if status == 405 {
		w.WriteHeader(int(status))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var emplReq *apiproto.Employee
	json.Unmarshal(buff, emplReq)

	c := apiproto.NewTaskRepoClient(conn)
	emplResp, err := c.UpdateEmployee(context.Background(), emplReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(emplResp.Status))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func FindEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	buff := make([]byte, 4096)
	_, err = r.Body.Read(buff)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var emplReq *apiproto.Employee
	json.Unmarshal(buff, emplReq)

	c := apiproto.NewTaskRepoClient(conn)
	emplResp, err := c.FindEmployee(context.Background(), emplReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(emplResp.Status))
		return
	}

	msg, err := json.Marshal(emplResp)
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

func UpdateEmployeeWithFormDataHandler(w http.ResponseWriter, r *http.Request) {
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

	optResp, err := c.UpdateEmployeeWithFormData(context.Background(), optReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(optResp.Status))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
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

	var emplReq *apiproto.Employee
	json.Unmarshal(buff, emplReq)

	emplResp, err := c.DeleteEmployee(context.Background(), emplReq)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(int(emplResp.Status))
		return
	}

	w.WriteHeader(http.StatusOK)
}
