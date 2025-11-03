package metric

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/octohelm/storage/pkg/session"
	singleflight "github.com/octohelm/x/sync/singleflight"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

var collectors = make([]prometheus.Collector, 0)

func register(c prometheus.Collector) {
	collectors = append(collectors, c)
}

const namespace = "pg"

// +gengo:injectable
type Metric struct {
	Namespace string
	SubSystem string
	Help      string
	ValueType prometheus.ValueType
	Labels    []string
	Names     []string

	Gather func(ctx context.Context, a session.Adapter, yield func(v float64, labelValues ...string)) error

	c *pgctl.Controller `inject:""`

	cache singleflight.GroupValue[string, *prometheus.Desc]
}

func (m *Metric) afterInit(ctx context.Context) error {
	if m.Names == nil {
		if m.ValueType == prometheus.GaugeValue {
			m.Names = []string{
				"count",
			}
		}
	}

	return nil
}

func (m *Metric) Describe(ch chan<- *prometheus.Desc) {
	labels, _ := m.labels()

	for _, name := range m.Names {
		ch <- m.descOrNew(name, labels)
	}
}

func (m *Metric) labels() ([]string, bool) {
	labels := make([]string, 0, len(m.Labels))
	includeName := false

	for _, label := range m.Labels {
		if label == "__name__" {
			includeName = true
			continue
		}
		labels = append(labels, label)
	}
	return labels, includeName
}

func (m *Metric) descOrNew(name string, labels []string) *prometheus.Desc {
	desc, _, _ := m.cache.Do(name, func() (*prometheus.Desc, error) {
		return prometheus.NewDesc(
			prometheus.BuildFQName(namespace, m.SubSystem, name),
			m.Help,
			labels,
			nil,
		), nil
	})

	return desc
}

func (m *Metric) Collect(ch chan<- prometheus.Metric) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	a, err := m.c.DBController(ctx).Open(ctx)
	if err != nil {
		return
	}
	defer a.Close()

	labels, includeName := m.labels()

	if err := m.Gather(ctx, a, func(v float64, labelValues ...string) {
		nameFilter := ""

		if includeName {
			nameFilter = labelValues[0]
			labelValues = labelValues[1:]
		}

		for _, name := range m.Names {
			if nameFilter != "" && name != nameFilter {
				continue
			}

			ch <- prometheus.MustNewConstMetric(m.descOrNew(name, labels), m.ValueType, v, labelValues...)
		}
	}); err != nil {
		return
	}
}
