package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pojntfx/go-support/pkg/controllers"
	"github.com/pojntfx/go-support/pkg/converters"
	"github.com/pojntfx/go-support/pkg/proto"
	"github.com/pojntfx/go-support/pkg/svc"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	dbFile := os.Getenv("DB")

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userController := controllers.UserController{
		DB: db,
	}
	userConverter := converters.UserConverter{}

	userService := svc.UserService{
		UserController: &userController,
		UserConverter:  &userConverter,
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, &userService)

	fmt.Printf("Server started on port %v", port)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
