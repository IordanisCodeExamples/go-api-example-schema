# go-api-example-schema

This repository hosts the schema definitions for the `go-api-example` project. The schema definitions are based on Protocol Buffers, a language-neutral, platform-neutral, extensible mechanism for serializing structured data.

The main purpose of this repository is to allow for a consistent and versioned schema definition that can be shared across different components or services of a larger system. By having the schema separate, you ensure that changes are properly tracked and can be used as a source of truth for what the API contract should be.

## Prerequisites

Before generating Go code from the `.proto` files, ensure you have the necessary tools installed:

- **Protocol Buffers Compiler**: Used to compile `.proto` files.
- **protoc-gen-go and protoc-gen-go-grpc**: Go plugins for the Protocol Buffers compiler to generate Go code with support for gRPC.

You can install the required Go plugins with:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


## Usage

To generate Go code from the `.proto` files, you can use the provided script:

```
./generate.sh
```

This will take all the `.proto` files within the `protos` directory and generate Go code with support for gRPC in the specified output directory.

