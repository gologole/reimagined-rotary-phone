package service

import (
	"cmd/main.go/config"
	"cmd/main.go/models"
)

func CalculateCost(courier models.Courier, cfg *config.Config) int {
	var pricePerKm int

	switch courier.Type {
	case "Авто":
		pricePerKm = cfg.CourierScores.Price.Auto
	case "Велосипед":
		pricePerKm = cfg.CourierScores.Price.Bike
	case "Пеший":
		pricePerKm = cfg.CourierScores.Price.Foot
	default:
		return 0 // Возвращаем 0, если тип курьера неизвестен
	}

	return int(courier.Dist) * pricePerKm
}
