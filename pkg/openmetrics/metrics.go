package openmetrics

import (
	"context"
	"fmt"
	"io"
	"maps"
	"net/http"
	"slices"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"

	"github.com/octohelm/courier/pkg/courierhttp"
	"github.com/octohelm/courier/pkg/openapi/jsonschema"
)

type (
	LabelPair    = dto.LabelPair
	Metric       = dto.Metric
	MetricFamily = dto.MetricFamily
)

type MetricFamilySet map[string]*MetricFamily

func (s MetricFamilySet) OpenAPISchemaType() []string {
	return []string{"string"}
}

var _ jsonschema.OpenAPISchemaTypeGetter = MetricFamilySet{}

func (MetricFamilySet) ContentType() string {
	return string(expfmt.FmtOpenMetrics_1_0_0)
}

var _ courierhttp.ResponseWriter = MetricFamilySet{}

func (s MetricFamilySet) WriteResponse(ctx context.Context, rw http.ResponseWriter, req courierhttp.RequestInfo) error {
	rw.Header().Set("Content-Type", string(expfmt.FmtOpenMetrics_1_0_0))
	rw.WriteHeader(http.StatusOK)

	_, err := s.WriteTo(rw)
	return err
}

var _ io.WriterTo = MetricFamilySet{}

func (s MetricFamilySet) WriteTo(w io.Writer) (n int64, finalErr error) {
	for _, name := range slices.Sorted(maps.Keys(s)) {
		written, err := expfmt.MetricFamilyToOpenMetrics(w, s[name])
		if err != nil {
			return -1, fmt.Errorf("write metric family %s failed: %w", name, err)
		}
		n += int64(written)
	}

	written, err := expfmt.FinalizeOpenMetrics(w)
	if err != nil {
		return -1, fmt.Errorf("finalize open metrics failed: %w", err)
	}
	n += int64(written)

	return n, finalErr
}

type (
	Gauge   = dto.Gauge
	Counter = dto.Counter
)
