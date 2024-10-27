package service

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/transport/grpc/rpc/protogen"
	"cmd/main.go/models"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func CreateOrder(cfg config.Config, doc models.Document) error {
	address := fmt.Sprintf("%s", cfg.HttpServer.CourierService)

	// Устанавливаем соединение с gRPC-сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // Используйте Insecure только для локальной разработки
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	// Создаём клиента
	client := protogen.NewOrderServiceClient(conn)

	// Создаём контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Преобразуем модель Document в нужный формат для gRPC
	request := &protogen.CreateOrderRequest{
		Cost:      doc.Cost,
		OverPrice: doc.OverPrice,
		Courier: &protogen.Courier{
			Id:   doc.Courier.ID,
			Name: doc.Courier.Name,
			Dist: doc.Courier.Dist,
			Type: doc.Courier.Type,
		},
		Couriers: make([]*protogen.Courier, len(doc.Couriers)),
	}

	// Конвертация массива курьеров
	for i, courier := range doc.Couriers {
		request.Couriers[i] = &protogen.Courier{
			Id:   courier.ID,
			Name: courier.Name,
			Dist: courier.Dist,
			Type: courier.Type,
		}
	}

	// Вызываем метод CreateOrder
	_, err = client.CreateOrder(ctx, request)
	if err != nil {
		return fmt.Errorf("could not create order: %w", err)
	}

	return nil
}
