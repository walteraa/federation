load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["foo.go"],
    importpath = "github.com/docker/docker/pkg/plugins/pluginrpc-gen/fixtures",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/plugins/pluginrpc-gen/fixtures/otherfixture:go_default_library"],
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
        "//vendor/github.com/docker/docker/pkg/plugins/pluginrpc-gen/fixtures/otherfixture:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
