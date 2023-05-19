package domain

import (
	"context"
	"testing"
	"time"

	"github.com/ferminhg/be-dora-metrics/internal/shared"
)

func TestMetricID(t *testing.T) {
	id := NewMetricID()
	gen := &shared.GoogleUUIDGenerator{}
	_, error := gen.Parse(id.value)
	if error != nil {
		t.Errorf("metricID.value is not valid UUID: %s", id.value)
	}
}

func TestNewMetric(t *testing.T) {
	name := "test"
	value := 1.0
	timestamp := time.Now()
	metric := NewMetric(name, value, timestamp)
	if metric.name != name {
		t.Errorf("metric.name is not valid: %s", metric.name)
	}
	if metric.value != value {
		t.Errorf("metric.value is not valid: %f", metric.value)
	}
	if metric.timestamp.String() != timestamp.String() {
		t.Errorf("metric.date is not valid: %s", metric.timestamp.String())
	}
}

func TestMetricRepository(t *testing.T) {
	var emptyMetric []Metric
	repository := NewInMemoryMetricRepository(emptyMetric)
	results, error := repository.FindAll(context.Background())
	if error != nil {
		t.Errorf("error in FindAll: %s", error.Error())
	}
	if len(emptyMetric) != len(results) {
		t.Errorf("emptyMetric and results are not the same length")
	}
}
