#!/bin/sh
protoc -I proto --micro_out=proto --go_out=proto proto/hello.proto
