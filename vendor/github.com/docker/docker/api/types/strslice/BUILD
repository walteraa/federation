load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["strslice.go"],
    importpath = "github.com/docker/docker/api/types/strslice",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["strslice_test.go"],
    importpath = "github.com/docker/docker/api/types/strslice",
    library = ":go_default_library",
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
