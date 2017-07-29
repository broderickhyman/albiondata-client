
fmt:
	goimports -w $(go list -f {{.Dir}} ./... | grep -v /vendor/)
