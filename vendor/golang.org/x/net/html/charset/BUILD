load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["charset.go"],
    importpath = "golang.org/x/net/html/charset",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/net/html:go_default_library",
        "//vendor/golang.org/x/text/encoding:go_default_library",
        "//vendor/golang.org/x/text/encoding/charmap:go_default_library",
        "//vendor/golang.org/x/text/encoding/htmlindex:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["charset_test.go"],
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/net/html/charset",
    library = ":go_default_library",
    deps = ["//vendor/golang.org/x/text/transform:go_default_library"],
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
