go test -v -tags integration ./...
go test -v -tags unit ./...

#this does not work bcz of the tags present on the first lines in _test.go files 
go test -v ./... 