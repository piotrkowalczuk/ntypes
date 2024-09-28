// swift-tools-version:6.0
import PackageDescription

let package = Package(
    name: "ntypes",
    platforms: [
        .macOS(.v10_15), // Specify the minimum macOS version required
        .iOS(.v13),
        .watchOS(.v6),
        .tvOS(.v13),
        .visionOS(.v1)
    ],
    products: [
        .library(
            name: "NTypes",
            targets: ["NTypes"]
        )
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", from: "1.28.1"),
        .package(url: "https://github.com/grpc/grpc-swift.git", from: "1.23.1"),
        .package(url: "https://github.com/swiftlang/swift-testing.git", branch: "main"),
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
                "ntypes_test.pb.swift",
            ],
            sources: [
                "ntypes.pb.swift",
            ],
            publicHeadersPath: ""
        ),
       .testTarget(
            name: "NTypesTests",
            dependencies: [
                .target(name: "NTypes"),
                .product(name: "Testing", package: "swift-testing"),
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
                "ntypes.pb.swift"
            ],
            sources: [
                "ntypes_test.pb.swift"
            ]
        ),
    ]
)
