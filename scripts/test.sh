#!/bin/sh
go test -cover ./auth ./db ./models ./resources ./router "$@"