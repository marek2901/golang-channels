#!/usr/bin/env bash

if !(which go &> /dev/null); then
    echo missing go binary
    exit 1
fi

(cd .. && go run main.go)
