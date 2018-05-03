VERSION?=$(shell git describe --tags --always --dirty)
PROTOC=/usr/local/bin/protoc

.PHONY: version gen

PROTO_INCLUDE=-I=/usr/include -I=.

gen:
	${PROTOC} ${PROTO_INCLUDE} --lint_out=. *.proto
	${PROTOC} ${PROTO_INCLUDE} --java_out=./tmp/java --go_out=${GOPATH}/src *.proto
	python -m grpc_tools.protoc ${PROTO_INCLUDE} --python_out=./ntypes *.proto
	ls -al | grep "pb.go"
	ls -al ./ntypes | grep "_pb2"

version:
	@echo ${VERSION} > VERSION.txt

get:
	go get -u github.com/golang/protobuf/...
	go get -u google.golang.org/grpc
	go get -u github.com/ckaznocha/protoc-gen-lint