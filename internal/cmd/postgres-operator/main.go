package main

import (
	"context"
	"fmt"
	"os"

	"github.com/innoai-tech/infra/pkg/cli"

	"github.com/innoai-tech/postgres-operator/internal/version"
)

var Serve = cli.AddTo(App, &struct {
	cli.C `name:"serve"`
}{})

var App = cli.NewApp(
	"postgres-operator",
	version.Version(),
	cli.WithImageNamespace("ghcr.io/innoai-tech"),
)

func main() {
	if err := cli.Execute(context.Background(), App, os.Args[1:]); err != nil {
		fmt.Println(err)
	}
}
