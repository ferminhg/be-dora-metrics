package bootstrap

import (
	"github.com/ferminhg/be-dora-metrics/internal/domain"
	"github.com/ferminhg/be-dora-metrics/internal/platform/server"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	metricRepository := domain.NewInMemoryMetricRepository([]domain.Metric{})
	srv := server.New(host, port, metricRepository)
	return srv.Run()
}
