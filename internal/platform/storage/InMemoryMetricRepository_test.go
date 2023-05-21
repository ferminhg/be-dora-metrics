package storage

import (
	"context"
	"testing"
)

func TestMetricRepository(t *testing.T) {
	var emptyMetric = RandMetricsMother(0)
	repository := NewInMemoryMetricRepository(emptyMetric)
	results, error := repository.FindAll(context.Background())
	if error != nil {
		t.Errorf("error in FindAll: %s", error.Error())
	}
	if len(emptyMetric) != len(results) {
		t.Errorf("emptyMetric and results are not the same length")
	}
}
