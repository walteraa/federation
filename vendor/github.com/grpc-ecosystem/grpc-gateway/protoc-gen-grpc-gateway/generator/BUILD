load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["generator.go"],
    importpath = "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/generator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin:go_default_library",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
