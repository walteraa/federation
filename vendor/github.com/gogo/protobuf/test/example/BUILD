load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

filegroup(
    name = "go_default_library_protos",
    srcs = ["example.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["example.pb.go"],
    importpath = "github.com/gogo/protobuf/test/example",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/descriptor:go_default_library",
        "//vendor/github.com/gogo/protobuf/test:go_default_library",
        "//vendor/github.com/gogo/protobuf/test/custom:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "example_test.go",
        "examplepb_test.go",
    ],
    importpath = "github.com/gogo/protobuf/test/example",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/jsonpb:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
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
