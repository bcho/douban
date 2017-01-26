.PHONY: test get-deps

test:
	@go test -v `go list ./... | grep -v /vendor/`

get-deps:
	gvt restore

