package v1alpha1

import (
	_ "github.com/gogo/protobuf/gogoproto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
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

// NodeResourceTopology is a specification for a Foo resource
type NodeResourceTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TopologyPolicies []string `json:"topologyPolicies"`
	Zones            ZoneList `json:"zones"`
}

// Zone is the spec for a NodeResourceTopology resource
// +protobuf=true
type Zone struct {
	Name       string           `json:"name" protobuf:"bytes,1,opt,name=name"`
	Type       string           `json:"type" protobuf:"bytes,2,opt,name=type"`
	Parent     string           `json:"parent,omitempty" protobuf:"bytes,3,opt,name=parent"`
	Costs      CostList         `json:"costs,omitempty" protobuf:"bytes,4,rep,name=costs"`
	Attributes AttributeList    `json:"attributes,omitempty" protobuf:"bytes,5,rep,name=attributes"`
	Resources  ResourceInfoList `json:"resources,omitempty" protobuf:"bytes,6,rep,name=resources"`
}

// +protobuf=true
type ZoneList []Zone

// +protobuf=true
type ResourceInfo struct {
	Name        string             `json:"name" protobuf:"bytes,1,opt,name=name"`
	Allocatable intstr.IntOrString `json:"allocatable" protobuf:"bytes,2,opt,name=allocatable"`
	Capacity    intstr.IntOrString `json:"capacity" protobuf:"bytes,3,opt,name=capacity"`
}

// +protobuf=true
type ResourceInfoList []ResourceInfo

// +protobuf=true
type CostInfo struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value int    `json:"value" protobuf:"varint,2,opt,name=value"`
}

// +protobuf=true
type CostList []CostInfo

// +protobuf=true
type AttributeInfo struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +protobuf=true
type AttributeList []AttributeInfo

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourceTopologyList is a list of NodeResourceTopology resources
type NodeResourceTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeResourceTopology `json:"items"`
}
