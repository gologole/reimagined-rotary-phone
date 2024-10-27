package service

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/transport/grpc/rpc/protogen"
	"cmd/main.go/models"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func FetchOrder(cfg config.Config) (*protogen.NewOrder, error) {
	address := fmt.Sprintf("%s", cfg.HttpServer.CourierService)

	// Устанавливаем соединение с gRPC-сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // Используйте Insecure только для локальной разработки
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	// Создаём клиента
	client := protogen.NewOrderServiceClient(conn)

	// Создаём контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Создаём пустой запрос
	req := &emptypb.Empty{}
	// Вызываем метод GetOrder
	order, err := client.GetOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not get order: %w", err)
	}

	return order, nil
}
func UploadDocument(cfg config.Config, doc models.Document) error {
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
	response := &protogen.Response{
		Price:     doc.Cost,
		OverPrice: doc.OverPrice,
		Yourcourier: &protogen.Courier{ // Преобразование Courier
			Id:   doc.Courier.ID,
			Name: doc.Courier.Name,
			Dist: doc.Courier.Dist,
			Type: doc.Courier.Type,
		},
		Courierlist: make([]*protogen.Courier, len(doc.Couriers)), // Создание массива для курьеров
	}

	// Конвертация массива курьеров
	for i, courier := range doc.Couriers {
		response.Courierlist[i] = &protogen.Courier{
			Id:   courier.ID,
			Name: courier.Name,
			Dist: courier.Dist,
			Type: courier.Type,
		}
	}

	// Вызываем метод UploadDocument
	_, err = client.UploadDocument(ctx, response)
	if err != nil {
		return fmt.Errorf("could not upload document: %w", err)
	}

	return nil
}
