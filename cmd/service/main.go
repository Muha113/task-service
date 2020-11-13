package main

import (
	"fmt"

	"github.com/Muha113/task-service/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	srv, err := server.NewServer("")
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("Starting server on ->", srv.ListenAddr)
	if srv.Start() != nil {
		logrus.Fatal(err)
	}
}
