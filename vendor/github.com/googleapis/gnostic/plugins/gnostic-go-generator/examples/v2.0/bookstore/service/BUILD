load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "init.go",
        "main.go",
        "service.go",
    ],
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore/service",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore/bookstore:go_default_library"],
)

go_binary(
    name = "service",
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore/service",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
