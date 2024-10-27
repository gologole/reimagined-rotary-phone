package main

import (
	"cmd/main.go/config"
	"cmd/main.go/models"
	"fmt"
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

func main() {
	me := config.DistanceThresholds{
		Near:   5,
		Medium: 7,
		Far:    20,
	}

	couriers := []models.Courier{
		{ID: 1, Name: "Курьер 1", Dist: 8, Type: "авто"},  // 15 км
		{ID: 2, Name: "Курьер 2", Dist: 1, Type: "вело"},  // 10 км
		{ID: 3, Name: "Курьер 3", Dist: 5, Type: "пеший"}, // 18 км

	}

	selectedCourier := selectCourier(couriers, me)
	fmt.Printf("Выбранный курьер: %s (ID: %d)\n", selectedCourier.Name, selectedCourier.ID)
	selectedCourier = selectCourier(couriers, me)
	fmt.Printf("Выбранный курьер: %s (ID: %d)\n", selectedCourier.Name, selectedCourier.ID)
	selectedCourier = selectCourier(couriers, me)
	fmt.Printf("Выбранный курьер: %s (ID: %d)\n", selectedCourier.Name, selectedCourier.ID)

}
