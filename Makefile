VERSION?=$(shell git describe --tags --always --dirty)

.PHONY: version gen

PROTOC=/usr/local/bin/protoc
PROTO_INCLUDE=-I=/usr/local/include -I=/usr/include -I=.

gen:
	./.circleci/scripts/generate.sh golang
	./.circleci/scripts/generate.sh swift
	@ls -al | grep "pb.go"
	@ls -al | grep "pb.swift"

version:
	@echo ${VERSION} > VERSION.txt
