git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "551e1f46cbeb470fe9d085e873c3bb83f075e18c",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")

go_rules_dependencies()
go_register_toolchains()
proto_register_toolchains()

git_repository(
    name = "com_github_trystanj_proto",
    commit = "72b5fee4fcb62697cdb29657f550110480a7eb39",
    remote = "https://github.com/trystanj/proto",
)
