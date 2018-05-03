VERSION?=$(shell git describe --tags --always --dirty)

.PHONY: version gen

PROTOC=/usr/local/bin/protoc
PROTO_INCLUDE=-I=/usr/include -I=.

gen:
	./.circleci/scripts/generate.sh golang
	@ls -al | grep "pb.go"

version:
	@echo ${VERSION} > VERSION.txt