.PHONY: coverage test

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test
