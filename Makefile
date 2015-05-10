all:
	go-bindata -pkg data -o data/bindata.go static static/images
dev:
	go-bindata -debug -pkg data -o data/bindata.go static static/images
dep:
	go get github.com/jteeuwen/go-bindata/go-bindata

	
