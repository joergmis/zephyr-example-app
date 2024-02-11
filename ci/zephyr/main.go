package zephyr

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

type Image struct {
	ctx       context.Context
	dag       *dagger.Client
	container *dagger.Container
}

func New(ctx context.Context, dag *dagger.Client) *Image {
	image := &Image{ctx: ctx, dag: dag}
	image.container = image.NewContainer()
	return image
}

func (image *Image) NewContainer() *dagger.Container {
	return image.dag.Container().
		WithEnvVariable("DEBIAN_FRONTEND", "noninteractive")
}

// OK forces the container to be built in case you don't publish the image to a
// registry.
func (image *Image) OK() *Image {
	fmt.Println(image.container.Stdout(image.ctx))

	return image
}
