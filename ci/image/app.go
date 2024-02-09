package image

import "dagger.io/dagger"

func (image *Image) WithSourceCode() *Image {
	source := image.dag.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"ci", "build", "binaries"},
	})

	image.container = image.container.
		WithDirectory("/zephyr/src", source).
		WithWorkdir("/zephyr")

	return image
}

func (image *Image) SetupZephyrModules() *Image {
	image.container = image.container.
		WithWorkdir("/zephyr").
		WithExec([]string{
			"west", "init", "-l", "src",
		}).WithExec([]string{
		"west", "update",
	})

	return image
}

func (image *Image) Build() *Image {
	image.container = image.container.
		WithWorkdir("/zephyr/src").
		WithExec([]string{
			"west", "build", "-b", "custom_plank", "app",
		})

	return image
}

func (image *Image) ExportBinaries() *Image {
	dir := image.container.Directory("/zephyr/src/build/zephyr")

	if _, err := dir.Export(image.ctx, "./binaries"); err != nil {
		panic(err)
	}

	return image
}
