module github.com/k8stopologyawareschedwg/noderesourcetopology-api

go 1.15

require (
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/code-generator v0.19.0
)

replace github.com/k8stopologyawareschedwg/noderesourcetopology-api => ../noderesourcetopology-api
