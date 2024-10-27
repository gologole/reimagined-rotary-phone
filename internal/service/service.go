package service

import (
	"cmd/main.go/config"
	"cmd/main.go/models"
)

// // Структура для хранения коэффициентов для типов курьеров
//
//	type CourierScores struct {
//		Auto map[string]int
//		Bike map[string]int
//		Foot map[string]int
//	}
//
// // Структура для хранения пороговых значений расстояний
//
//	type DistanceThresholds struct {
//		Near   int // Ближнее расстояние
//		Medium int // Среднее расстояние
//		Far    int // Дальнее расстояние
//	}
type service struct {
	cfg *config.Config
}
type Service interface {
	CalculateDeafultCost(order models.Order) models.Document
	GetCourier(order models.Order) models.Document
}

func NewService(cfg *config.Config) Service {
	return service{
		cfg,
	}
}

func (s service) CalculateDeafultCost(order models.Order) models.Document {
	couriers, _ := FetchAllCouriers(*s.cfg)

	courier := selectCourier(couriers, s.cfg.DistanceThresholds)

	cost := CalculateCost(courier, s.cfg)
	document := models.Document{
		int64(cost),
		0,
		models.Courier{}, //при отправке результата расчетов дефолт цены возвращаю пустого
		couriers,
	}
	//UploadDocument(*s.cfg, document)
	return document
}

func (s service) GetCourier(order models.Order) models.Document {

	couriers, _ := FetchAllCouriers(*s.cfg)
	courier := selectCourier(couriers, s.cfg.DistanceThresholds)

	cost := CalculateCost(courier, s.cfg)

	return models.Document{
		int64(cost),
		order.OverPrice,
		courier,
		couriers,
	}
	//UploadDocument(*s.cfg, document)
}
