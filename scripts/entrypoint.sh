#!/bin/sh
cd migrate && go run *.go init && go run *.go
cd ../ && air