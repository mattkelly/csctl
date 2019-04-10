package resource

import (
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/table"
)

// NodePools is a list of the associated cloud resource with additional functionality
type NodePools struct {
	resource
	items []types.NodePool
}

// NewNodePools constructs a new NodePools wrapping the given cloud type
func NewNodePools(items []types.NodePool) *NodePools {
	return &NodePools{
		resource: resource{
			name:    "node-pool",
			plural:  "node-pools",
			aliases: []string{"nodepool", "nodepools", "np", "nps"},
		},
		items: items,
	}
}

// NodePool constructs a new NodePools with no underlying items, useful for
// interacting with the metadata itself.
func NodePool() *NodePools {
	return NewNodePools(nil)
}

func (p *NodePools) columns() []string {
	return []string{
		"Name",
		"ID",
		"Mode",
		"Status",
		"Kubernetes Version",
		"Etcd Version",
		"Docker Version",
	}
}

// Table outputs the table representation to the given writer
func (p *NodePools) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, np := range p.items {
		var etcdVersion string
		if np.EtcdVersion == nil || *np.EtcdVersion == "" {
			etcdVersion = "N/A"
		} else {
			etcdVersion = *np.EtcdVersion
		}

		var status string
		if np.Status == nil || np.Status.Type == nil ||
			*np.Status.Type == "" {
			status = "UNKNOWN"
		} else {
			status = *np.Status.Type
		}

		table.Append([]string{
			*np.Name,
			string(np.ID),
			*np.KubernetesMode,
			status,
			*np.KubernetesVersion,
			etcdVersion,
			*np.DockerVersion,
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *NodePools) JSON(w io.Writer, listView bool) error {
	if !listView && len(p.items) == 1 {
		return displayJSON(w, p.items[0])
	}

	return displayJSON(w, p.items)
}

// YAML outputs the YAML representation to the given writer
func (p *NodePools) YAML(w io.Writer, listView bool) error {
	if !listView && len(p.items) == 1 {
		return displayYAML(w, p.items[0])
	}

	return displayYAML(w, p.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *NodePools) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
