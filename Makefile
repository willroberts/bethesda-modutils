.PHONY: test
test:
	go test -v ./...

.PHONY: coverage
coverage:
	go test -coverprofile cover.out
	go tool cover -html cover.out

.PHONY: lint
lint:
	go vet ./...
	golint ./...
