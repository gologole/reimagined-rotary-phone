package service

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/transport/grpc/protogen"
	"cmd/main.go/models"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

// FetchAllCouriers получает всех курьеров и возвращает их в виде массива
func FetchAllCouriers(cfg config.Config) ([]models.Courier, error) {
	// Формируем адрес подключения
	address := fmt.Sprintf("%s", cfg.HttpServer.CourierService)

	// Устанавливаем соединение с gRPC-сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // Используйте Insecure только для локальной разработки
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	// Создаём клиента
	client := protogen.NewCouriersClient(conn)

	// Создаём контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Создаём запрос для получения всех курьеров
	getAllRequest := &protogen.GetAllCouriersRequest{}
	response, err := client.GetAllCouriers(ctx, getAllRequest)
	if err != nil {
		return nil, fmt.Errorf("could not get all couriers: %w", err)
	}

	// Парсим ответ в массив курьеров
	couriers := make([]models.Courier, len(response.Couriers))
	for i, c := range response.Couriers {
		couriers[i] = models.Courier{
			ID:   int64(c.Id),
			Name: c.FullName,
			Dist: int64(c.Dist),
			Type: c.Type,
		}
	}

	return couriers, nil
}
