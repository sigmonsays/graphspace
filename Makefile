all:
	go-bindata -pkg data -o data/bindata.go static
dev:
	go-bindata -debug -pkg data -o data/bindata.go static
dep:
	go get github.com/jteeuwen/go-bindata/go-bindata

	
