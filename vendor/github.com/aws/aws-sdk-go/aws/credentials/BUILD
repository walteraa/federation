load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "chain_provider.go",
        "credentials.go",
        "env_provider.go",
        "shared_credentials_provider.go",
        "static_provider.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/aws/credentials",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws/awserr:go_default_library",
        "//vendor/github.com/go-ini/ini:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "chain_provider_test.go",
        "credentials_test.go",
        "env_provider_test.go",
        "shared_credentials_provider_test.go",
        "static_provider_test.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/aws/credentials",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws/awserr:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials/endpointcreds:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials/stscreds:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
