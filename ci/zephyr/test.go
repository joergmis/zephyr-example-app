package zephyr

func (image *Image) AddTestDependencies() *Image {
	image.container = image.
		WithRubyFull().
		WithLibasioDev().
		WithLibGtestDev().
		WithLibTclapDev().
		WithCucumberGems().
		WithBashrc().
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
	image.container = image.container.
		WithExec([]string{
			"gem", "install", "cucumber:7.1.0", "cucumber-wire:6.2.1",
		})

	return image
}

func (image *Image) WithBashrc() *Image {
	image.container = image.container.WithExec([]string{
		"sh", "-c", "echo 'source ../zephyr/zephyr-env.sh' >> $HOME/.bashrc",
	}).WithExec([]string{
		"sh", "-c", "echo 'source $HOME/.zephyr' >> $HOME/.bashrc",
	})

	return image
}

func (image *Image) RunTests() *Image {
	image.container = image.container.
		WithWorkdir("/zephyr/src").
		WithExec([]string{
			"bash", "-c", "west twister -p custom_plank -T ./tests",
		})

	return image
}

func (image *Image) ExportTestReports() *Image {
	dir := image.container.Directory("/zephyr/src/twister-out")

	if _, err := dir.Export(image.ctx, "./artifacts"); err != nil {
		panic(err)
	}

	return image
}
