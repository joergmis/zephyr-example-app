package image

func (image *Image) ZephyrTest() *Image {
	image.container = image.
		WithRubyFull().
		WithLibasioDev().
		WithLibGtestDev().
		WithLibTclapDev().
		WithCucumberGems().
		container

	return image
}

func (image *Image) WithRubyFull() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "ruby-full",
		})

	return image
}

func (image *Image) WithLibasioDev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libasio-dev",
		})

	return image
}

func (image *Image) WithLibGtestDev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libgtest-dev",
		})

	return image
}

func (image *Image) WithLibTclapDev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libtclap-dev",
		})

	return image
}

func (image *Image) WithNlohmannJson3Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "nlohmann-json3-dev",
		})

	return image
}

func (image *Image) WithCucumberGems() *Image {
	image.container = image.WithRubyFull().container.
		WithExec([]string{
			"gem", "install", "cucumber:7.1.0", "cucumber-wire:6.2.1",
		})

	return image
}
