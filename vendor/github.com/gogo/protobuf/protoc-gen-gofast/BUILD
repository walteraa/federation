load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/gogo/protobuf/protoc-gen-gofast",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/gogo/protobuf/vanity:go_default_library",
        "//vendor/github.com/gogo/protobuf/vanity/command:go_default_library",
    ],
)

go_binary(
    name = "protoc-gen-gofast",
    importpath = "github.com/gogo/protobuf/protoc-gen-gofast",
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
