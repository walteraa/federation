load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "download.go",
        "transfer.go",
        "upload.go",
    ],
    importpath = "github.com/docker/docker/distribution/xfer",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/docker/image:go_default_library",
        "//vendor/github.com/docker/docker/layer:go_default_library",
        "//vendor/github.com/docker/docker/pkg/archive:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/progress:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "download_test.go",
        "transfer_test.go",
        "upload_test.go",
    ],
    importpath = "github.com/docker/docker/distribution/xfer",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/docker/image:go_default_library",
        "//vendor/github.com/docker/docker/layer:go_default_library",
        "//vendor/github.com/docker/docker/pkg/progress:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
