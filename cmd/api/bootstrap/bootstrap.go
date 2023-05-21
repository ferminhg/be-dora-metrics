package bootstrap

import (
	"github.com/ferminhg/be-dora-metrics/internal/domain"
	"github.com/ferminhg/be-dora-metrics/internal/platform/server"
	"math/rand"
	"time"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	n := domain.NUM_RANDOM_METRICS // Replace with the number of times you want to repeat
	metrics := make([]domain.Metric, n)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	min := 10.0
	max := 20.0
	for i := 0; i < n; i++ {
		randomFloat := min + r.Float64()*(max-min)
		metrics[i] = domain.NewMetric("tc", randomFloat, time.Now())
	}

	metricRepository := domain.NewInMemoryMetricRepository(metrics)
	srv := server.New(host, port, metricRepository)
	return srv.Run()
}
