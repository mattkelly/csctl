package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

var (
	kubernetesVersion string
)

// upgradeNodePoolCmd represents the upgradeNodePool command
var upgradeNodePoolCmd = &cobra.Command{
	Use:     "node-pool",
	Short:   "Upgrade the Kubernetes version for a node pool",
	Aliases: resource.NodePool().Aliases(),
	Long: `Upgrade the Kubernetes version for a node pool.

Node pool upgrades are performed on nodes in-place. Nodes are upgraded one-by-one.

There are several Kubernetes version skew requirements that Containership enforces.

For more information, please see https://docs.containership.io/cluster-management/upgrading-kubernetes-versions.`,

	Args: cobra.NoArgs,

	PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		req := &types.NodePoolUpgradeRequest{
			KubernetesVersion: &kubernetesVersion,
		}

		_, err := clientset.Provision().NodePools(organizationID, clusterID).Upgrade(nodePoolID, req)
		if err != nil {
			return errors.Wrapf(err, "patching node pool %q", nodePoolID)
		}

		fmt.Printf("Node pool %s upgrade to Kubernetes %s successfully initiated\n", nodePoolID, kubernetesVersion)

		return nil
	},
}

func init() {
	upgradeCmd.AddCommand(upgradeNodePoolCmd)

	bindCommandToNodePoolScope(upgradeNodePoolCmd, false)

	upgradeNodePoolCmd.Flags().StringVar(&kubernetesVersion, "kubernetes-version", "", "Kubernetes version to upgrade to")
	_ = upgradeNodePoolCmd.MarkFlagRequired("kubernetes-version")
}
