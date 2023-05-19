package bootstrap

import (
	"github.com/ferminhg/be-dora-metrics/internal/domain"
	"github.com/ferminhg/be-dora-metrics/internal/platform/server"
	"time"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	metrics := []domain.Metric{
		domain.NewMetric("tc", 0.5, time.Now()),
		domain.NewMetric("tc", 0.4, time.Now()),
		domain.NewMetric("tc", 0.6, time.Now()),
		domain.NewMetric("tc", 0.3, time.Now()),
		domain.NewMetric("tc", 0.7, time.Now()),
	}
	metricRepository := domain.NewInMemoryMetricRepository(metrics)
	srv := server.New(host, port, metricRepository)
	return srv.Run()
}
