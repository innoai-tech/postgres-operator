gen:
    go tool devtool gen --all ./internal/cmd/postgres-operator

update:
    go get -u ./...

dep:
    go mod tidy

test:
    CGO_ENABLED=0 \
      go test -failfast -count=1 ./internal/... ./pkg/...

fmt:
    go tool gofumpt -w -l .
