load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["splunk.go"],
    importpath = "github.com/docker/docker/daemon/logger/splunk",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/daemon/logger:go_default_library",
        "//vendor/github.com/docker/docker/daemon/logger/loggerutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/urlutil:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "splunk_test.go",
        "splunkhecmock_test.go",
    ],
    importpath = "github.com/docker/docker/daemon/logger/splunk",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/docker/daemon/logger:go_default_library"],
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
