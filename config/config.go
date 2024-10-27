package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Структура для хранения коэффициентов для типов курьеров
type CourierScores struct {
	Auto  int               `yaml:"auto"`
	Bike  int               `yaml:"bike"`
	Foot  int               `yaml:"foot"`
	Price PricePerKilometer `yaml:"price_per_km"`
}

// Структура для хранения цен за километр
type PricePerKilometer struct {
	Auto int `yaml:"auto"`
	Bike int `yaml:"bike"`
	Foot int `yaml:"foot"`
}

// Структура для хранения пороговых значений расстояний
type DistanceThresholds struct {
	Near   int `yaml:"near"`   // Ближнее расстояние
	Medium int `yaml:"medium"` // Среднее расстояние
	Far    int `yaml:"far"`    // Дальнее расстояние
}

type Config struct {
	Postgres           postgres           `yaml:"postgres"`
	Redis              redis              `yaml:"redis"`
	HttpServer         httpServer         `yaml:"http_server"`
	GrpcServer         grpcServer         `yaml:"grpc_server"`
	CourierScores      CourierScores      `yaml:"courier_scores"`
	DistanceThresholds DistanceThresholds `yaml:"distance_thresholds"`
}

type httpServer struct {
	Port           int    `yaml:"port"`
	ElkDomain      string `yaml:"elk_domain"`
	CourierService string `yaml:"courier_service"`
	OrderService   string `yaml:"order_service"`
}

type grpcServer struct {
	Port int `yaml:"port"`
}

type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type postgres struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// InitConfig загружает конфигурацию из файла и обрабатывает переменные окружения
func InitConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML config: %w", err)
	}

	return &config, nil
}
