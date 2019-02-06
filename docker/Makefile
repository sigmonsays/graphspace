IMAGE_VERSION = 2

all:
	mkdir -p ./bin/
	cp $(GOPATH)/bin/graphspace ./bin/
	docker build -t graphspace:$(IMAGE_VERSION) .
