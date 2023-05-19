package domain

import (
	"context"
	"time"

	"github.com/ferminhg/be-dora-metrics/internal/shared"
)

type metricID struct {
	value string
}

func NewMetricID() metricID {
	gen := &shared.GoogleUUIDGenerator{}
	return metricID{value: gen.New().String()}
}

func (id metricID) String() string {
	return id.value
}

type Metric struct {
	id        metricID
	name      string
	value     float64
	timestamp time.Time
}

func NewMetric(name string, value float64, timestamp time.Time) Metric {
	return Metric{
		id:        NewMetricID(),
		name:      name,
		value:     value,
		timestamp: timestamp,
	}
}

type MetricRepository interface {
	FindAll(ctx context.Context) ([]Metric, error)
}

type InMemoryMetricRepository struct {
	metric []Metric
}

func NewInMemoryMetricRepository(metric []Metric) *InMemoryMetricRepository {
	return &InMemoryMetricRepository{
		metric: metric,
	}
}

func (InMemoryMetricRepository) FindAll(ctx context.Context) ([]Metric, error) {
	return []Metric{}, nil
}
