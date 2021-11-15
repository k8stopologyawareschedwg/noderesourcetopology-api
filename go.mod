module github.com/k8stopologyawareschedwg/noderesourcetopology-api

go 1.16

require (
	github.com/gogo/protobuf v1.3.2
	k8s.io/apimachinery v0.22.3
	k8s.io/client-go v0.22.3
	k8s.io/code-generator v0.22.3
)

replace github.com/k8stopologyawareschedwg/noderesourcetopology-api => ../noderesourcetopology-api
