load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "factory.go",
        "generic.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions",
    importpath = "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//staging/src/k8s.io/client-go/tools/cache:go_default_library",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1:go_default_library",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/clientset/versioned:go_default_library",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/internalinterfaces:go_default_library",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/internalinterfaces:all-srcs",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
