load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "json.go",
        "range.go",
        "semver.go",
        "sort.go",
        "sql.go",
    ],
    importpath = "github.com/blang/semver",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "json_test.go",
        "range_test.go",
        "semver_test.go",
        "sort_test.go",
        "sql_test.go",
    ],
    importpath = "github.com/blang/semver",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/blang/semver/examples:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
