load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "windows.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "ansi_reader.go",
            "ansi_writer.go",
            "console.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/pkg/term/windows",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-ansiterm:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/Azure/go-ansiterm/winterm:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = ["windows_test.go"],
    importpath = "github.com/docker/docker/pkg/term/windows",
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
