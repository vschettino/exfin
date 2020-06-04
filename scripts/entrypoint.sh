#!/bin/sh
(cd exfin-cli && go install)
(cd migrate && go run *.go init && go run *.go)
air