load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "benchmarker.go",
        "interfaces.go",
        "it_node.go",
        "measure_node.go",
        "runner.go",
        "setup_nodes.go",
        "suite_nodes.go",
        "synchronized_after_suite_node.go",
        "synchronized_before_suite_node.go",
    ],
    importpath = "github.com/onsi/ginkgo/internal/leafnodes",
    visibility = ["//vendor/github.com/onsi/ginkgo:__subpackages__"],
    deps = [
        "//vendor/github.com/onsi/ginkgo/internal/codelocation:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/failer:go_default_library",
        "//vendor/github.com/onsi/ginkgo/types:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "it_node_test.go",
        "leaf_node_suite_test.go",
        "measure_node_test.go",
        "setup_nodes_test.go",
        "shared_runner_test.go",
        "suite_nodes_test.go",
        "synchronized_after_suite_node_test.go",
        "synchronized_before_suite_node_test.go",
    ],
    importpath = "github.com/onsi/ginkgo/internal/leafnodes_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/codelocation:go_default_library",
        "//vendor/github.com/onsi/ginkgo/internal/failer:go_default_library",
        "//vendor/github.com/onsi/ginkgo/types:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
        "//vendor/github.com/onsi/gomega/ghttp:go_default_library",
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
