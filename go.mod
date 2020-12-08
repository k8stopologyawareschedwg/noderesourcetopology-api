module github.com/k8stopologyawareschedwg/noderesourcetopology-api

go 1.15

require (
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v0.19.3
	k8s.io/code-generator v0.19.3
	k8stopologyawareschedwg/noderesourcetopology-api v0.0.0-00010101000000-000000000000
)

replace k8stopologyawareschedwg/noderesourcetopology-api => ../noderesourcetopology-api
