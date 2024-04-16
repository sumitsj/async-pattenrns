build:
	go build -o main .

run:
	go run ./main.go

run-tests:
	go clean -testcache
	go test `go list ./... | grep -v "mocks"`

run-tests-with-coverage:
	go clean -testcache
	go test -coverprofile cover.out `go list ./... | grep -v "mocks\|constants\|contracts\|models"`
	go tool cover -html cover.out -o cover.html
	open cover.html
