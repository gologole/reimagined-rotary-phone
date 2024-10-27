package service

import (
	"cmd/main.go/config"
	"cmd/main.go/models"
	"sort"
)

var queue int = 0

func selectCourier(couriers []models.Courier, thresholds config.DistanceThresholds) models.Courier {
	near, middle, far := AgragateByDist(couriers, thresholds)

	near = sortCouriersOrder1(near)
	middle = sortCouriersOrder1(middle)
	far = sortCouriersOrder1(far)

	var selectedCourier models.Courier

	switch queue {
	case 0:
		if len(near) > 0 {
			selectedCourier = near[0]
		} else if len(middle) > 0 {
			selectedCourier = middle[0]
		} else if len(far) > 0 {
			selectedCourier = far[0]
		}
		queue = 1 // Увеличиваем queue после выбора

	case 1:
		if len(middle) > 0 {
			selectedCourier = middle[0]
		} else if len(far) > 0 {
			selectedCourier = far[0]
		} else if len(near) > 0 {
			selectedCourier = near[0]
		}
		queue = 2

	case 2:
		if len(far) > 0 {
			selectedCourier = far[0]
		} else if len(near) > 0 {
			selectedCourier = near[0]
		} else if len(middle) > 0 {
			selectedCourier = middle[0]
		}
		queue = 0 // Сбрасываем queue до 0

	default:
		return models.Courier{} // Если по какой-то причине queue не в пределах 0-2
	}

	return selectedCourier
}

func AgragateByDist(couriers []models.Courier, thresholds config.DistanceThresholds) ([]models.Courier, []models.Courier, []models.Courier) {
	var near []models.Courier
	var middle []models.Courier
	var far []models.Courier

	for _, courier := range couriers {
		if int(courier.Dist) < thresholds.Near {
			near = append(near, courier)
		}
		if int(courier.Dist) >= thresholds.Near && int(courier.Dist) <= thresholds.Medium {
			middle = append(middle, courier)
		}
		if int(courier.Dist) > thresholds.Medium && int(courier.Dist) < thresholds.Far {
			far = append(far, courier)
		}
	}
	return near, middle, far
}

// sortCouriersOrder1 сортирует курьеров по порядку: авто, вело, пеший
func sortCouriersOrder1(couriers []models.Courier) []models.Courier {
	order := []string{"авто", "вело", "пеший"}
	orderMap := make(map[string]int)

	for i, transport := range order {
		orderMap[transport] = i
	}

	sort.Slice(couriers, func(i, j int) bool {
		return orderMap[couriers[i].Type] < orderMap[couriers[j].Type]
	})

	return couriers
}

// sortCouriersOrder2 сортирует курьеров по порядку: пеший, вело, авто
func sortCouriersOrder2(couriers []models.Courier) []models.Courier {
	order := []string{"пеший", "вело", "авто"}
	orderMap := make(map[string]int)

	for i, transport := range order {
		orderMap[transport] = i
	}

	sort.Slice(couriers, func(i, j int) bool {
		return orderMap[couriers[i].Type] < orderMap[couriers[j].Type]
	})

	return couriers
}
