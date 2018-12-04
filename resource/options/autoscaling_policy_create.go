package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

// AutoscalingPolicyCreate is the set of options required to create an autoscaling policy
type AutoscalingPolicyCreate struct {
	// Defaultable
	Name                string
	MetricsBackend      string
	Metric              string
	MetricConfiguration map[string]interface{}
	PollInterval        int32
	SamplePeriod        int32
}

// DefaultAndValidate defaults and validates all options
func (o *AutoscalingPolicyCreate) DefaultAndValidate() error {
	if err := o.defaultAndValidateName(); err != nil {
		return errors.Wrap(err, "name")
	}

	if err := o.defaultAndValidateMetric(); err != nil {
		return errors.Wrap(err, "metric")
	}

	if err := o.defaultAndValidateMetricsBackend(); err != nil {
		return errors.Wrap(err, "metrics backend")
	}

	if err := o.defaultAndValidatePollInterval(); err != nil {
		return errors.Wrap(err, "poll interval")
	}

	if err := o.defaultAndValidateSamplePeriod(); err != nil {
		return errors.Wrap(err, "sample period")
	}

	// TODO needs to be user-settable somehow
	// Maybe --metric-configuration=aggregation=avg,range=1m,...
	o.MetricConfiguration = map[string]interface{}{
		"aggregation": "avg",
	}

	return nil
}

// CreateAutoscalingPolicyRequest constructs a create autoscaling policy request from these options
func (o *AutoscalingPolicyCreate) CreateAutoscalingPolicyRequest() types.AutoscalingPolicy {
	// TODO all of this needs to be user-settable somehow
	gt := ">"
	lt := "<"
	thresholdUp := float32(0.8)
	thresholdDown := float32(0.3)
	absolute := "absolute"
	adjustBy := int32(1)
	return types.AutoscalingPolicy{
		Name:                &o.Name,
		MetricsBackend:      o.MetricsBackend,
		Metric:              &o.Metric,
		MetricConfiguration: o.MetricConfiguration,
		Policy: &types.ScalingPolicy{
			ScaleUp: &types.ScalingPolicyConfiguration{
				ComparisonOperator: &gt,
				Threshold:          &thresholdUp,
				AdjustmentType:     &absolute,
				AdjustmentValue:    &adjustBy,
			},
			ScaleDown: &types.ScalingPolicyConfiguration{
				ComparisonOperator: &lt,
				Threshold:          &thresholdDown,
				AdjustmentType:     &absolute,
				AdjustmentValue:    &adjustBy,
			},
		},
		PollInterval: &o.PollInterval,
		SamplePeriod: &o.SamplePeriod,
	}
}

func (o *AutoscalingPolicyCreate) defaultAndValidateName() error {
	if o.Name == "" {
		o.Name = "none"
	}
	return nil
}

func (o *AutoscalingPolicyCreate) defaultAndValidateMetric() error {
	if o.Metric == "" {
		o.Metric = "cpu"
	}

	switch o.Metric {
	case "cpu", "memory":
		return nil
	default:
		return errors.Errorf("unknown metric %q", o.Metric)
	}
}

func (o *AutoscalingPolicyCreate) defaultAndValidateMetricsBackend() error {
	if o.MetricsBackend == "" {
		o.MetricsBackend = "prometheus"
	}

	// TODO support others
	if o.MetricsBackend != "prometheus" {
		return errors.Errorf("unknown metrics backend %q", o.MetricsBackend)
	}

	return nil
}

func (o *AutoscalingPolicyCreate) defaultAndValidatePollInterval() error {
	if o.PollInterval == -1 {
		o.PollInterval = 15
	}
	return nil
}

func (o *AutoscalingPolicyCreate) defaultAndValidateSamplePeriod() error {
	if o.SamplePeriod == -1 {
		o.SamplePeriod = 600
	}
	return nil
}
