devtool := "go tool devtool"

gen:
    {{ devtool }} gen --all ./internal/cmd/postgres-operator

fmt:
    {{ devtool }} fmt -w -l .

update:
    go get -u ./pkg/... ./internal/...

dep:
    go mod tidy

test:
    CGO_ENABLED=0 \
      go test -failfast -count=1 ./internal/... ./pkg/...
