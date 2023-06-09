package domain

import (
	"testing"
	"time"

	"github.com/ferminhg/be-dora-metrics/internal/shared"
)

func TestMetricID(t *testing.T) {
	id := NewMetricID()
	gen := &shared.GoogleUUIDGenerator{}
	_, error := gen.Parse(id.Value)
	if error != nil {
		t.Errorf("metricID.value is not valid UUID: %s", id.Value)
	}
}

func TestNewMetric(t *testing.T) {
	name := NewMetricName("DF")
	value := 1.0
	timestamp := time.Now()
	metric := NewMetric(name, value, timestamp)
	if metric.Name != name {
		t.Errorf("metric.name is not valid: %s", metric.Name)
	}
	if metric.Value != value {
		t.Errorf("metric.value is not valid: %f", metric.Value)
	}
	if metric.Timestamp.String() != timestamp.String() {
		t.Errorf("metric.date is not valid: %s", metric.Timestamp.String())
	}
}
