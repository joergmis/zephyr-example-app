package image

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

func (image *Image) OK() *Image {
	fmt.Println(image.container.Stdout(image.ctx))

	return image
}
