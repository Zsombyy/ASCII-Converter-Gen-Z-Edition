APP = brainrot-ascii
DIST = dist

.PHONY:  all build clean snapshot release install-deps appimage flatpak

all: clean build
	mkdir -p $(DIST)
	go build -o $(DIST)/$(APP) main.go

clean: 
	rm -rf $(DIST)

install-deps:
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest

release: install-deps
	goreleaser release --clean

appimage: 
	mkdir -p $(DIST)
	bash /scripts/build-appimage.sh

flatpak:
	mkdir -p $(DIST)
	cp packing/$(APP).flatpakref $(DIST)/$(APP).flatpakref
