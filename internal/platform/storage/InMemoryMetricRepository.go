package storage

import (
	"context"
	"math/rand"
	"time"

	domain "github.com/ferminhg/be-dora-metrics/internal/domain"
)

func RandMetricsMother(size int) []domain.Metric {
	metrics := make([]domain.Metric, size)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	min := 10.0
	max := 20.0
	for i := 0; i < size; i++ {
		randomFloat := min + r.Float64()*(max-min)
		metrics[i] = domain.NewMetric("tc", randomFloat, time.Now())
	}
	return metrics
}

type InMemoryMetricRepository struct {
	metric []domain.Metric
}

func NewInMemoryMetricRepository(metric []domain.Metric) *InMemoryMetricRepository {
	return &InMemoryMetricRepository{
		metric: metric,
	}
}

func (r InMemoryMetricRepository) FindAll(ctx context.Context) ([]domain.Metric, error) {
	return r.metric, nil
}
