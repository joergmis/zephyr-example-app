package image

func (image *Image) WithReporting() *Image {
	image.container = image.
		WithLibreofficeWriter().
		WithZip().
		container

	return image
}

func (image *Image) WithLibreofficeWriter() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libreoffice-writer",
		})

	return image
}

func (image *Image) WithZip() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "zip",
		})

	return image
}
