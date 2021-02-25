
.PHONY:  docker

GOPATH = $(shell go env GOPATH)
GOBINDATA := $(GOPATH)/bin/go-bindata

# BIN := go-bindata-assetfs

all: dep
	$(GOBINDATA) -pkg data -o data/bindata.go static static/images

dev:
	$(GOBINDATA) -debug -pkg data -o data/bindata.go static static/images
	go install github.com/sigmonsays/graphspace/...

dep:
	go get github.com/jteeuwen/go-bindata/go-bindata
	go get github.com/elazarl/go-bindata-assetfs/...

docker:
    # build docker image
	docker build -t graphspace:latest .
docker-push:
	docker tag graphspace:latest sigmonsays/graphspace:latest
	docker push sigmonsays/graphspace:latest

