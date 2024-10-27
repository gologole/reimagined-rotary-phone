package grpcserver

import (
	"cmd/main.go/internal/service"
	"cmd/main.go/internal/transport/grpc/rpc/protogen"
	pb "cmd/main.go/internal/transport/grpc/server/rpc/protogen"
	"cmd/main.go/models"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

// OrderServiceServerImpl реализует интерфейс OrderServiceServer
type OrderServiceServerImpl struct {
	protogen.UnimplementedOrderServiceServer
	s service.Service
}

// Реализуйте метод CalculateDefaultCost
func (s *OrderServiceServerImpl) CalculateDefaultCost(ctx context.Context, in *pb.NewOrder) (*emptypb.Empty, error) {
	order := models.Order{
		ID:          0,              // Обновите это значение в зависимости от вашей логики
		OverPrice:   in.Overprice,   // Предполагаем, что NewOrder имеет поле Overprice
		Description: in.Description, // Предполагаем, что NewOrder имеет поле Description
	}

	// Вызов метода для расчета стоимости
	s.s.CalculateDeafultCost(order)

	return &emptypb.Empty{}, nil
}

func (s *OrderServiceServerImpl) GetDocument(ctx context.Context, order *pb.NewOrder) (*pb.Response, error) {
	document := s.s.GetCourier(models.Order{OverPrice: order.Overprice})

	// Преобразование списка курьеров
	var courierList []*pb.Courier
	for _, c := range document.Couriers {
		courierList = append(courierList, &pb.Courier{
			Id:   c.ID,
			Name: c.Name,
			Type: c.Type,
			Dist: c.Dist,
		})
	}

	// Преобразование вашего курьера
	var yourCourier *pb.Courier
	if document.Courier != (models.Courier{}) { // Проверяем, что ваш курьер не пустой
		yourCourier = &pb.Courier{
			Id:   document.Courier.ID,
			Name: document.Courier.Name,
			Type: document.Courier.Type,
			Dist: document.Courier.Dist,
		}
	}

	// Возвращение ответа
	return &pb.Response{
		Id:          0,
		Price:       document.Cost,
		Courierlist: courierList, // Теперь это []*pb.Courier
		Yourcourier: yourCourier, // Теперь это *pb.Courier
	}, nil
}
