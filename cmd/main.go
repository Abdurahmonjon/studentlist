package main

import (
	"fmt"
	"github.com/Abdurahmonjon/studentcrud/studentservice/config"
	pb "github.com/Abdurahmonjon/studentcrud/studentservice/genproto/gitlab.com/Abdurahmonjon/studentproto"
	"github.com/Abdurahmonjon/studentcrud/studentservice/pkg/db"
	"github.com/Abdurahmonjon/studentcrud/studentservice/pkg/logger"
	"github.com/Abdurahmonjon/studentcrud/studentservice/service"
	"github.com/Abdurahmonjon/studentcrud/studentservice/storage"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "student-crud")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	pg := storage.NewStoragePg(connDB)
	studentServiceServer := service.NewService(pg, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	log.Info(fmt.Sprintf("%v", lis))

	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, studentServiceServer)

	if err = s.Serve(lis); err != nil {
		fmt.Println(errors.New("error while serving grpc"))
		return
	}
}
