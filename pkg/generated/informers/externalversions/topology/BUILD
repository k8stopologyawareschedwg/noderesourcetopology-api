load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["interface.go"],
    importmap = "k8s.io/kubernetes/vendor/github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology",
    importpath = "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology",
    visibility = ["//visibility:public"],
    deps = [
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/internalinterfaces:go_default_library",
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology/v1alpha1:go_default_library",
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
        "//github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology/v1alpha1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
