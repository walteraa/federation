load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "bigquery.go",
        "books.go",
        "calendar.go",
        "compute.go",
        "debug.go",
        "drive.go",
        "fitness.go",
        "gmail.go",
        "main.go",
        "mirror.go",
        "prediction.go",
        "pubsub.go",
        "storage.go",
        "tasks.go",
        "urlshortener.go",
        "youtube.go",
    ],
    importpath = "google.golang.org/api/examples",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/golang.org/x/oauth2/google:go_default_library",
        "//vendor/google.golang.org/api/bigquery/v2:go_default_library",
        "//vendor/google.golang.org/api/books/v1:go_default_library",
        "//vendor/google.golang.org/api/calendar/v3:go_default_library",
        "//vendor/google.golang.org/api/compute/v1:go_default_library",
        "//vendor/google.golang.org/api/drive/v2:go_default_library",
        "//vendor/google.golang.org/api/fitness/v1:go_default_library",
        "//vendor/google.golang.org/api/gmail/v1:go_default_library",
        "//vendor/google.golang.org/api/googleapi:go_default_library",
        "//vendor/google.golang.org/api/mirror/v1:go_default_library",
        "//vendor/google.golang.org/api/prediction/v1.6:go_default_library",
        "//vendor/google.golang.org/api/pubsub/v1beta2:go_default_library",
        "//vendor/google.golang.org/api/storage/v1:go_default_library",
        "//vendor/google.golang.org/api/tasks/v1:go_default_library",
        "//vendor/google.golang.org/api/urlshortener/v1:go_default_library",
        "//vendor/google.golang.org/api/youtube/v3:go_default_library",
    ],
)

go_binary(
    name = "examples",
    importpath = "google.golang.org/api/examples",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
