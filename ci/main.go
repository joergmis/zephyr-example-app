package main

import (
	"context"
	"flag"
	"os"

	"dagger.io/dagger"
	"github.com/joergmis/zephyr-example-app/ci/zephyr"
)

const (
	zephyrSDKVersion = "0.16.1"
)

func main() {
	build := flag.Bool("build", false, "Trigger a build of the pap")
	test := flag.Bool("test", false, "Trigger running the tests")

	flag.Parse()

	ctx := context.Background()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer dag.Close()

	if *build {
		buildApp(ctx, dag)
	}

	if *test {
		testApp(ctx, dag)
	}
}

func buildApp(ctx context.Context, dag *dagger.Client) {
	zephyr.New(ctx, dag).
		AddZephyrDependencies(zephyrSDKVersion).
		AddSourceCode().
		SetupZephyrModules().
		BuildApp().
		ExportBinaries().
		OK()
}

func testApp(ctx context.Context, dag *dagger.Client) {
	zephyr.New(ctx, dag).
		AddZephyrDependencies(zephyrSDKVersion).
		AddTestDependencies().
		AddSourceCode().
		SetupZephyrModules().
		RunTests().
		ExportTestReports().
		OK()
}
