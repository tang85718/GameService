#!/usr/bin/env bash
mkdir $GOPATH/src/proto/gm
protoc --go_out=plugins=micro:$GOPATH/src/proto/gm ./gm.api.proto