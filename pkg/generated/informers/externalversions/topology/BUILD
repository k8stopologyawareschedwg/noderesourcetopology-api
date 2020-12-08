load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["interface.go"],
    importmap = "k8s.io/kubernetes/vendor/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology",
    importpath = "k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology",
    visibility = ["//visibility:public"],
    deps = [
        "//k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/internalinterfaces:go_default_library",
        "//k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology/v1alpha1:go_default_library",
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
        "//k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions/topology/v1alpha1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
