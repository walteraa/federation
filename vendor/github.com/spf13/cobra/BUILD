load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bash_completions.go",
        "cobra.go",
        "command.go",
        "command_notwin.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "command_win.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/spf13/cobra",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/spf13/pflag:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/inconshreveable/mousetrap:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = [
        "bash_completions_test.go",
        "cobra_test.go",
        "command_test.go",
    ],
    importpath = "github.com/spf13/cobra",
    library = ":go_default_library",
    deps = ["//vendor/github.com/spf13/pflag:go_default_library"],
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
        "//vendor/github.com/spf13/cobra/cobra:all-srcs",
        "//vendor/github.com/spf13/cobra/doc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
