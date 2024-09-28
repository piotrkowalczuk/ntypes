// swift-tools-version:6.0
import PackageDescription

let package = Package(
    name: "ntypes",
    products: [
        .library(
            name: "ntypes",
            targets: ["NTypes"]
        )
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", from: "1.28.1"),
        .package(url: "https://github.com/grpc/grpc-swift.git", from: "1.23.1"),
    ],
    targets: [
        // Główny target dla kodu źródłowego
        .target(
            name: "NTypes",
            dependencies: [
                .product(name: "SwiftProtobuf", package: "swift-protobuf"),
                .product(name: "GRPC", package: "grpc-swift"),
            ],
            path: ".",
            exclude: [
                "go.mod",
                "go.sum",
                "ntypes.go",
                "ntypes.pb.go",
                "ntypes_test.go",
                "ntypespq",
                "ntypes",
                "Makefile",
                "setup.py",
                "setup.cfg",
                "MANIFEST.in",
                "LICENSE.txt",
                "README.md",
                "ntypes.proto",
            ],
            sources: [
                "ntypes.pb.swift",
            ],
            publicHeadersPath: ""
        ),
    ]
)
