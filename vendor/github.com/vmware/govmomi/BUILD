load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/vmware/govmomi",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vmware/govmomi/property:go_default_library",
        "//vendor/github.com/vmware/govmomi/session:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/soap:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["client_test.go"],
    importpath = "github.com/vmware/govmomi",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/vmware/govmomi/test:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/mo:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/types:go_default_library",
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
        "//vendor/github.com/vmware/govmomi/event:all-srcs",
        "//vendor/github.com/vmware/govmomi/examples/datastores:all-srcs",
        "//vendor/github.com/vmware/govmomi/find:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc:all-srcs",
        "//vendor/github.com/vmware/govmomi/guest:all-srcs",
        "//vendor/github.com/vmware/govmomi/license:all-srcs",
        "//vendor/github.com/vmware/govmomi/list:all-srcs",
        "//vendor/github.com/vmware/govmomi/object:all-srcs",
        "//vendor/github.com/vmware/govmomi/ovf:all-srcs",
        "//vendor/github.com/vmware/govmomi/pbm:all-srcs",
        "//vendor/github.com/vmware/govmomi/performance:all-srcs",
        "//vendor/github.com/vmware/govmomi/property:all-srcs",
        "//vendor/github.com/vmware/govmomi/session:all-srcs",
        "//vendor/github.com/vmware/govmomi/task:all-srcs",
        "//vendor/github.com/vmware/govmomi/test:all-srcs",
        "//vendor/github.com/vmware/govmomi/units:all-srcs",
        "//vendor/github.com/vmware/govmomi/view:all-srcs",
        "//vendor/github.com/vmware/govmomi/vim25:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
