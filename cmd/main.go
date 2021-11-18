package main

import (
	"database/sql"
	"fmt"
	"net"

	"bitbucket.org/udevs/ur_go_user_service/config"
	"bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/ur_go_user_service/pkg/logger"
	"bitbucket.org/udevs/ur_go_user_service/service"
	"bitbucket.org/udevs/ur_go_user_service/storage/sqlc"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "ur_go_user_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)
	// conStr = `host=localhost port=5432 user=postgres password=20072003 dbname=user_service sslmode=disable`
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}
	// client, err := grpc_client.NewGrpcClients(&cfg)
	// if err != nil {
	// 	log.Error("error while connecting to clients", logger.Error(err))
	// 	return
	// }
	// fmt.Println(db)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	userService := service.NewBookService(sqlc.New(db), log)

	s := grpc.NewServer()
	reflection.Register(s)

	user_service.RegisterBookServiceServer(s, userService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
