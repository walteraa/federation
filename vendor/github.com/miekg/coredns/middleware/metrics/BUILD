load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "handler.go",
        "metrics.go",
        "setup.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/metrics",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mholt/caddy:go_default_library",
        "//vendor/github.com/miekg/coredns/core/dnsserver:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/metrics/vars:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/pkg/dnsrecorder:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/pkg/rcode:go_default_library",
        "//vendor/github.com/miekg/coredns/request:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "metrics_test.go",
        "setup_test.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/metrics",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/mholt/caddy:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/metrics/test:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/pkg/dnsrecorder:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/test:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
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
        "//vendor/github.com/miekg/coredns/middleware/metrics/test:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/metrics/vars:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
