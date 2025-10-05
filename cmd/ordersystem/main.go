package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/marcosvlima/clean-arch-go-sample/configs"
	"github.com/marcosvlima/clean-arch-go-sample/internal/event/handler"
	"github.com/marcosvlima/clean-arch-go-sample/internal/infra/grpc/pb"
	"github.com/marcosvlima/clean-arch-go-sample/internal/infra/grpc/service"
	"github.com/marcosvlima/clean-arch-go-sample/internal/infra/web/webserver"
	"github.com/marcosvlima/clean-arch-go-sample/pkg/events"
	"github.com/streadway/amqp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitConn, rabbitMQChannel := getRabbitMQChannel()
	defer func() {
		if rabbitMQChannel != nil {
			_ = rabbitMQChannel.Close()
		}
		if rabbitConn != nil {
			_ = rabbitConn.Close()
		}
	}()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// webserver
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go func() {
		if err := webserver.Start(); err != nil {
			fmt.Println("web server stopped:", err)
		}
	}()

	// gRPC server
	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("gRPC server stopped:", err)
		}
	}()

	// wait for termination signal to gracefully shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("Shutting down...")
	grpcServer.GracefulStop()
}

func getRabbitMQChannel() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		panic(err)
	}
	return conn, ch
}
