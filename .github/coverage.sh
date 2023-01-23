#!/bin/bash
set -eu

cov_dir="./coverage"
mkdir -p "${cov_dir}"
go test ./... -covermode=set -coverprofile="${cov_dir}/coverage.out" -coverpkg=./...
go tool cover -o "${cov_dir}/index.html" -html="${cov_dir}/coverage.out"
open "${cov_dir}/index.html" &