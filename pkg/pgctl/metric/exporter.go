package metric

import (
	"context"

	"github.com/innoai-tech/infra/pkg/configuration"
	"github.com/innoai-tech/postgres-operator/pkg/openmetrics"
	"github.com/prometheus/client_golang/prometheus"
)

// +gengo:injectable:provider
type Exporter struct {
	registry *prometheus.Registry
}

func (e *Exporter) afterInit(ctx context.Context) error {
	if e.registry != nil {
		return nil
	}

	e.registry = prometheus.NewRegistry()

	for _, c := range collectors {
		if canInit, ok := c.(configuration.CanInit); ok {
			if err := canInit.Init(ctx); err != nil {
				return err
			}
		}

		if err := e.registry.Register(c); err != nil {
			return err
		}
	}

	return nil
}

func (e *Exporter) MetricFamilySet(ctx context.Context) (openmetrics.MetricFamilySet, error) {
	metricFamilies, err := e.registry.Gather()
	if err != nil {
		return nil, err
	}

	mfs := openmetrics.MetricFamilySet{}
	for _, mf := range metricFamilies {
		mfs[*mf.Name] = mf
	}
	return mfs, nil
}
