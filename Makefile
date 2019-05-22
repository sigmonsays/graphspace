
.PHONY:  docker

BIN := go-bindata
# BIN := go-bindata-assetfs

all:
	$(BIN) -pkg data -o data/bindata.go static static/images

dev:
	$(BIN) -debug -pkg data -o data/bindata.go static static/images
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

