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
        "//pkg/apis/topology:all-srcs",
        "//pkg/generated/clientset/versioned:all-srcs",
        "//pkg/generated/informers/externalversions:all-srcs",
        "//pkg/generated/listers/topology/v1alpha1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
