package domain

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ferminhg/be-dora-metrics/internal/shared"
)

type metricID struct {
	Value string
}

// Added to return custom json on marshal metricID
func (m metricID) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value)
}

func (m *metricID) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.Value)
}

func NewMetricID() metricID {
	gen := &shared.GoogleUUIDGenerator{}
	return metricID{Value: gen.New().String()}
}

func (id metricID) String() string {
	return id.Value
}

type Metric struct {
	ID        metricID  `json:"id"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func NewMetric(name string, value float64, timestamp time.Time) Metric {
	return Metric{
		ID:        NewMetricID(),
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
}

type MetricRepository interface {
	FindAll(ctx context.Context) ([]Metric, error)
}

const NUM_RANDOM_METRICS = 50

type InMemoryMetricRepository struct {
	metric []Metric
}

func NewInMemoryMetricRepository(metric []Metric) *InMemoryMetricRepository {
	return &InMemoryMetricRepository{
		metric: metric,
	}
}

func (r InMemoryMetricRepository) FindAll(ctx context.Context) ([]Metric, error) {
	return r.metric, nil
}
