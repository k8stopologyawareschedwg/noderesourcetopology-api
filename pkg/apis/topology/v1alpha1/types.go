package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TopologyManagerPolicy string

const (
	// Constants of type TopologyManagerPolicy represent policy of the worker
	// node's resource management component. It's TopologyManager in kubele.
	// SingleNUMANodeContainerLevel represent single-numa-node policy of
	// the TopologyManager
	SingleNUMANodeContainerLevel TopologyManagerPolicy = "SingleNUMANodeContainerLevel"
	// SingleNUMANodePodLevel enables pod level resource counting, this policy assumes
	// TopologyManager policy single-numa-node also was set on the node
	SingleNUMANodePodLevel TopologyManagerPolicy = "SingleNUMANodePodLevel"
	// Restricted TopologyManager policy was set on the node
	Restricted TopologyManagerPolicy = "Restricted"
	// BestEffort TopologyManager policy was set on the node
	BestEffort TopologyManagerPolicy = "BestEffort"
	// None policy is the default policy and does not perform any topology alignment.
	None TopologyManagerPolicy = "None"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourceTopology describes node resources and their topology.
type NodeResourceTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TopologyPolicies []string `json:"topologyPolicies"`
	Zones            ZoneList `json:"zones"`
}

// Zone represents a resource topology zone, e.g. socket, node, die or core.
type Zone struct {
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	Parent     string           `json:"parent,omitempty"`
	Costs      CostList         `json:"costs,omitempty"`
	Attributes AttributeList    `json:"attributes,omitempty"`
	Resources  ResourceInfoList `json:"resources,omitempty"`
}

// ZoneList contains an array of Zone objects.
type ZoneList []Zone

// ResourceInfo contains information about one resource type.
type ResourceInfo struct {
	Name        string            `json:"name"`
	Allocatable resource.Quantity `json:"allocatable"`
	Capacity    resource.Quantity `json:"capacity"`
}

// ResourceInfoList contains an array of ResourceInfo objects.
type ResourceInfoList []ResourceInfo

// CostInfo describes the cost (or distance) between two Zones.
type CostInfo struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

// CostList contains an array of CostInfo objects.
type CostList []CostInfo

// AttributeInfo contains one attribute of a Zone.
type AttributeInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AttributeList contains an array of AttributeInfo objects.
type AttributeList []AttributeInfo

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourceTopologyList is a list of NodeResourceTopology resources
type NodeResourceTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeResourceTopology `json:"items"`
}
