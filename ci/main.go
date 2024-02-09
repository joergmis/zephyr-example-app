package main

import (
	"context"
	"os"

	"dagger.io/dagger"
	"github.com/joergmis/zephyr-example-app/ci/image"
)

const (
	zephyrSDKVersion = "0.16.1"
)

func main() {
	ctx := context.Background()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer dag.Close()

	image.New(ctx, dag).
		ZephyrBase(zephyrSDKVersion).
		WithSourceCode().
		SetupZephyrModules().
		Build().
		ExportBinaries().
		OK()
}
