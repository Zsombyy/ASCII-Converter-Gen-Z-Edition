APP_NAME=brainrot-ascii

.PHONY all build clean test release snapshot install-deps

build:
	go build -o bin/$(APP_NAME) main.go

	clean rm -rf bin dist

install-deps:
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest

release:
	goreleaser release --clean

snapshot:
	goreleaser release --snapshot --clean --skip-publish
