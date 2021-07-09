module github.com/k8stopologyawareschedwg/noderesourcetopology-api

go 1.15

require (
	github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1 v0.0.0-00010101000000-000000000000
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/code-generator v0.19.0
)

replace github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1 => ./pkg/apis/topology/v1alpha1
