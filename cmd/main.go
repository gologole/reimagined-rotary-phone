package main

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/service"
	"cmd/main.go/internal/transport/grpc/rpc/protogen"
	pb "cmd/main.go/internal/transport/grpc/server/rpc/protogen"
	mylogger "cmd/main.go/pkg/logger"
	"cmd/main.go/pkg/logstash"
	"cmd/main.go/server/grpcserver"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

func main() {
	cfg, err := config.InitConfig("./config/config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	mylogger.NewLogger()
	logstash.NewLogstashLogger(*cfg)

	// Создание сервиса
	svc := service.NewService(cfg)

	// Создание gRPC сервера
	grpcServer := grpc.NewServer()

	orderService := &grpcserver.OrderServiceServerImpl{
		S: svc, // Предполагаем, что у вас есть функция для создания сервиса
	}

	// Регистрируем OrderServiceServer
	protogen.RegisterOrderServiceServer(grpcServer, orderService)

	// Регистрация вашего gRPC сервера
	pb.RegisterOrderServiceServer(grpcServer, orderService)

	reflection.Register(grpcServer)
	// Прослушивание входящих соединений
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.GrpcServer.Port))
	if err != nil {
		logstash.GlobalLogstash.LogstashError("failed to listen: %v", err)
	}

	logstash.GlobalLogstash.LogstashInfo((fmt.Sprintf("Server is listening on port %s", cfg.GrpcServer.Port)))

	// Запуск gRPC сервера
	if err := grpcServer.Serve(listener); err != nil {
		logstash.GlobalLogstash.LogstashError("failed to serve: %v", err)
	}

}
