load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "defaults.go",
        "devices_unsupported.go",
        "namespaces.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "devices_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/oci",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
        "//vendor/github.com/opencontainers/runtime-spec/specs-go:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/opencontainers/runc/libcontainer/devices:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
