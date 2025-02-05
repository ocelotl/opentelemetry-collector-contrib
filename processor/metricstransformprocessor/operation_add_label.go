// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metricstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"

import (
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func (mtp *metricsTransformProcessor) addLabelOp(metric *metricspb.Metric, op internalOperation) {
	var lb = metricspb.LabelKey{
		Key: op.configOperation.NewLabel,
	}
	metric.MetricDescriptor.LabelKeys = append(metric.MetricDescriptor.LabelKeys, &lb)
	for _, ts := range metric.Timeseries {
		lv := &metricspb.LabelValue{
			Value:    op.configOperation.NewValue,
			HasValue: true,
		}
		ts.LabelValues = append(ts.LabelValues, lv)
	}
}

// addLabelOp add a new attribute to metric data points.
func addLabelOp(metric pmetric.Metric, op internalOperation) {
	rangeDataPointAttributes(metric, func(attrs pcommon.Map) bool {
		attrs.InsertString(op.configOperation.NewLabel, op.configOperation.NewValue)
		return true
	})
}
