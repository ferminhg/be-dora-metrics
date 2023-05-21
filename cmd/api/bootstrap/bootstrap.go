package bootstrap

import (
	"github.com/ferminhg/be-dora-metrics/internal/platform/server"
	"github.com/ferminhg/be-dora-metrics/internal/platform/storage"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	metricRepository := storage.NewInMemoryMetricRepository(storage.RandMetricsMother(50))
	srv := server.New(host, port, metricRepository)
	return srv.Run()
}
