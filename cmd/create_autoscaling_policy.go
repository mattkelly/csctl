package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var createAutoscalingPolicyOpts options.AutoscalingPolicyCreate

// createAutoscalingPolicyCmd represents the createAutoscalingPolicy command
var createAutoscalingPolicyCmd = &cobra.Command{
	Use:     "autoscaling-policy",
	Short:   "Create an autoscaling policy",
	Aliases: resource.AutoscalingPolicy().Aliases(),

	Args: cobra.NoArgs,

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := createAutoscalingPolicyOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := createAutoscalingPolicyOpts.CreateAutoscalingPolicyRequest()

		asp, err := clientset.Provision().AutoscalingPolicies(organizationID, clusterID).Create(&req)
		if err != nil {
			return err
		}

		t := resource.NewAutoscalingPolicies([]types.AutoscalingPolicy{*asp})
		return t.Table(os.Stdout)
	},
}

func init() {
	createCmd.AddCommand(createAutoscalingPolicyCmd)

	bindCommandToClusterScope(createAutoscalingPolicyCmd, false)

	// No defaulting is performed here because the logic in many cases is nontrivial,
	// and we'd like to be consistent with where and how we default.
	createAutoscalingPolicyCmd.Flags().StringVarP(&createAutoscalingPolicyOpts.Name,
		"name", "n", "", "name of this autoscaling policy")
	createAutoscalingPolicyCmd.Flags().StringVarP(&createAutoscalingPolicyOpts.MetricsBackend,
		"metrics-backend", "b", "", "name of the metrics backend to use")
	createAutoscalingPolicyCmd.Flags().StringVarP(&createAutoscalingPolicyOpts.Metric,
		"metric", "m", "", "name of the metric")
	createAutoscalingPolicyCmd.Flags().Int32VarP(&createAutoscalingPolicyOpts.PollInterval,
		"poll-interval", "p", -1, "poll interval")
	createAutoscalingPolicyCmd.Flags().Int32VarP(&createAutoscalingPolicyOpts.SamplePeriod,
		"sample-period", "s", -1, "sample period")
}
