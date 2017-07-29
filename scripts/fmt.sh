#!/usr/bin/env bash

echo "Running goimports to format go files..."
goimports -w $(go list -f {{.Dir}} ./... | grep -v /vendor/)
