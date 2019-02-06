IMAGE_VERSION = 1

all:
	mkdir -p ./bin/
	cp $(GOPATH)/bin/graphspace ./bin/
	docker build -t graphspace:$(IMAGE_VERSION) .
