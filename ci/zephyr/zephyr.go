package zephyr

import (
	"fmt"

	"dagger.io/dagger"
)

func (image *Image) AddZephyrDependencies(zephyrSDKVersion string) *Image {
	image.container = image.Ubuntu().
		WithAptUpdate().
		WithSoftwarePropertiesCommon().
		WithLsbRelease().
		WithAutoconf().
		WithAutomake().
		WithBison().
		WithBuildEssential().
		WithCaCertificates().
		WithCcache().
		WithChrPath().
		WithCmake().
		WithCpio().
		WithDeviceTreeCompiler().
		WithDfuUtil().
		WithDiffstat().
		WithDos2Unix().
		WithDoxygen().
		WithFile().
		WithFlex().
		WithGpp().
		WithGcovr().
		WithGdb().
		WithGit().
		WithGitCore().
		WithGnupg().
		WithGperf().
		WithGtkSharp2().
		WithHelp2man().
		WithIproute2().
		WithLcov().
		WithLibcairo2Dev().
		WithLibGlib20Dev().
		WithLibGtk200().
		WithLibLocaleGettextPerl().
		WithLibNcurses5Dev().
		WithLibPcapDev().
		WithLibPopt0().
		WithLibSdl12Dev().
		WithLibSdl2Dev().
		WithLibSslDev().
		WithLibTool().
		WithLibtoolBin().
		WithLocales().
		WithMake().
		WithNetTools().
		WithNinjaBuild().
		WithOpensshClient().
		WithParallel().
		WithPkgConfig().
		WithPython3Dev().
		WithPython3Pip().
		WithPython3Ply().
		WithPython3Setuptools().
		WithPythonIsPython3().
		WithQemu().
		WithRsync().
		WithSocat().
		WithSrecord().
		WithSudo().
		WithTexinfo().
		WithUnzip().
		WithValgrind().
		WithWget().
		WithOvmf().
		WithXzUtils().
		WithThriftCompiler().
		WithGccMultilib().
		WithGppMultilib().
		WithLocaleGen().
		WithZephyrSDK(zephyrSDKVersion).
		WithPythonDepdencies().
		WithAptClean().
		container

	return image
}

func (image *Image) Ubuntu() *Image {
	image.container = image.NewContainer().
		From("ubuntu:22.04")

	return image
}

func (image *Image) WithAptUpdate() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "update", "--yes",
		})

	return image
}

func (image *Image) WithAptClean() *Image {
	image.container = image.container.WithExec([]string{
		"apt-get", "clean", "--yes",
	}).WithExec([]string{
		"apt-get", "autoremove", "--purge", "--yes",
	}).WithExec([]string{
		"rm", "-rf", "/var/lib/apt/lists/*",
	})

	return image
}

func (image *Image) WithSoftwarePropertiesCommon() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "software-properties-common",
		})

	return image
}

func (image *Image) WithGitPPA() *Image {
	image.container = image.container.
		WithExec([]string{
			"add-apt-repository", "ppa:git-core/ppa",
		})

	return image
}

func (image *Image) WithLsbRelease() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "lsb-release",
		})

	return image
}

func (image *Image) WithAutoconf() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "autoconf",
		})

	return image
}

func (image *Image) WithAutomake() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "automake",
		})

	return image
}

func (image *Image) WithBison() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "bison",
		})

	return image
}

func (image *Image) WithBuildEssential() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "build-essential",
		})

	return image
}

func (image *Image) WithCaCertificates() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "ca-certificates",
		})

	return image
}

func (image *Image) WithCcache() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "ccache",
		})

	return image
}

func (image *Image) WithChrPath() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "chrpath",
		})

	return image
}

func (image *Image) WithCmake() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "cmake",
		})

	return image
}

func (image *Image) WithCpio() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "cpio",
		})

	return image
}

func (image *Image) WithDeviceTreeCompiler() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "device-tree-compiler",
		})

	return image
}

func (image *Image) WithDfuUtil() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "dfu-util",
		})

	return image
}

func (image *Image) WithDiffstat() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "diffstat",
		})

	return image
}

func (image *Image) WithDos2Unix() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "dos2unix",
		})

	return image
}

func (image *Image) WithDoxygen() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "doxygen",
		})

	return image
}

func (image *Image) WithFile() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "file",
		})

	return image
}

func (image *Image) WithFlex() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "flex",
		})

	return image
}

func (image *Image) WithGpp() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "g++",
		})

	return image
}

func (image *Image) WithGawk() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gawk",
		})

	return image
}

func (image *Image) WithGcc() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gcc",
		})

	return image
}

func (image *Image) WithGcovr() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gcovr",
		})

	return image
}

func (image *Image) WithGdb() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gdb",
		})

	return image
}

func (image *Image) WithGit() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "git",
		})

	return image
}

func (image *Image) WithGitCore() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "git-core",
		})

	return image
}

func (image *Image) WithGnupg() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gnupg",
		})

	return image
}

func (image *Image) WithGperf() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gperf",
		})

	return image
}

func (image *Image) WithGtkSharp2() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gtk-sharp2",
		})

	return image
}

func (image *Image) WithHelp2man() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "help2man",
		})

	return image
}

func (image *Image) WithIproute2() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "iproute2",
		})

	return image
}

func (image *Image) WithLcov() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "lcov",
		})

	return image
}

func (image *Image) WithLibcairo2Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libcairo2-dev",
		})

	return image
}

func (image *Image) WithLibGlib20Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libglib2.0-dev",
		})

	return image
}

func (image *Image) WithLibGtk200() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libgtk2.0-0",
		})

	return image
}

func (image *Image) WithLibLocaleGettextPerl() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "liblocale-gettext-perl",
		})

	return image
}

func (image *Image) WithLibNcurses5Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libncurses5-dev",
		})

	return image
}

func (image *Image) WithLibPcapDev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libpcap-dev",
		})

	return image
}

func (image *Image) WithLibPopt0() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libpopt0",
		})

	return image
}

func (image *Image) WithLibSdl12Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libsdl1.2-dev",
		})

	return image
}

func (image *Image) WithLibSdl2Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libsdl2-dev",
		})

	return image
}

func (image *Image) WithLibSslDev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libssl-dev",
		})

	return image
}

func (image *Image) WithLibTool() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libtool",
		})

	return image
}

func (image *Image) WithLibtoolBin() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "libtool-bin",
		})

	return image
}

func (image *Image) WithLocales() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "locales",
		})

	return image
}

func (image *Image) WithMake() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "make",
		})

	return image
}

func (image *Image) WithNetTools() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "net-tools",
		})

	return image
}

func (image *Image) WithNinjaBuild() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "ninja-build",
		})

	return image
}

func (image *Image) WithOpensshClient() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "openssh-client",
		})

	return image
}

func (image *Image) WithParallel() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "parallel",
		})

	return image
}

func (image *Image) WithPkgConfig() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "pkg-config",
		})

	return image
}

func (image *Image) WithPython3Dev() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "python3-dev",
		})

	return image
}

func (image *Image) WithPython3Pip() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "python3-pip",
		})

	return image
}

func (image *Image) WithPython3Ply() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "python3-ply",
		})

	return image
}

func (image *Image) WithPython3Setuptools() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "python3-setuptools",
		})

	return image
}

func (image *Image) WithPythonIsPython3() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "python-is-python3",
		})

	return image
}

func (image *Image) WithQemu() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "qemu",
		})

	return image
}

func (image *Image) WithRsync() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "rsync",
		})

	return image
}

func (image *Image) WithSocat() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "socat",
		})

	return image
}

func (image *Image) WithSrecord() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "srecord",
		})

	return image
}

func (image *Image) WithSudo() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "sudo",
		})

	return image
}

func (image *Image) WithTexinfo() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "texinfo",
		})

	return image
}

func (image *Image) WithUnzip() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "unzip",
		})

	return image
}

func (image *Image) WithValgrind() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "valgrind",
		})

	return image
}

func (image *Image) WithWget() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "wget",
		})

	return image
}

func (image *Image) WithOvmf() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "ovmf",
		})

	return image
}

func (image *Image) WithXzUtils() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "xz-utils",
		})

	return image
}

func (image *Image) WithThriftCompiler() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "thrift-compiler",
		})

	return image
}

func (image *Image) WithGccMultilib() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "gcc-multilib",
		})

	return image
}

func (image *Image) WithGppMultilib() *Image {
	image.container = image.container.
		WithExec([]string{
			"apt-get", "install", "--yes", "g++-multilib",
		})

	return image
}

func (image *Image) WithLocaleGen() *Image {
	image.container = image.container.WithExec([]string{
		"sh", "-c", "locale-gen en_US.UTF-8",
	}).WithEnvVariable("LANG", "en_US.UTF-8").
		WithEnvVariable("LANGUAGE", "en_US:en").
		WithEnvVariable("LC_ALL", "en_US.UTF-8")

	return image
}

func (image *Image) WithZephyrSDK(zephyrSDKVersion string) *Image {
	image.container = image.container.WithExec([]string{
		"wget", "--quiet", fmt.Sprintf("https://github.com/zephyrproject-rtos/sdk-ng/releases/download/v%s/zephyr-sdk-%s_linux-x86_64_minimal.tar.xz", zephyrSDKVersion, zephyrSDKVersion),
	}).WithExec([]string{
		"wget", "--quiet", "-O", "shasum.txt", fmt.Sprintf("https://github.com/zephyrproject-rtos/sdk-ng/releases/download/v%s/sha256.sum", zephyrSDKVersion),
	}).WithExec([]string{
		"shasum", "--check", "--ignore-missing", "shasum.txt",
	}).WithExec([]string{
		"tar", "xf", fmt.Sprintf("zephyr-sdk-%s_linux-x86_64_minimal.tar.xz", zephyrSDKVersion), "-C", "/opt/",
	}).WithExec([]string{
		"rm", fmt.Sprintf("zephyr-sdk-%s_linux-x86_64_minimal.tar.xz", zephyrSDKVersion),
	}).WithExec([]string{
		fmt.Sprintf("/opt/zephyr-sdk-%s/setup.sh", zephyrSDKVersion), "-t", "arm-zephyr-eabi",
	}).WithExec([]string{
		"sh", "-c", "echo export ZEPHYR_TOOLCHAIN_VARIANT=zephyr > $HOME/.zephyrrc",
	}).WithExec([]string{
		"sh", "-c", "echo export ZEPHYR_SDK_INSTALL_DIR=/opt/zephyr-sdk-${ZEPHYR_SDK_VERSION} >> $HOME/.zephyrrc",
	})

	return image
}

func (image *Image) WithPythonDepdencies() *Image {
	image.container = image.container.WithExec([]string{
		"python3", "-m", "pip", "install", "-U", "--no-cache-dir", "pip",
	}).WithExec([]string{
		"pip3", "install", "-U", "--no-cache-dir", "wheel", "setuptools",
	}).WithExec([]string{
		"pip3", "install", "--no-cache-dir", "pygobject",
	}).WithExec([]string{
		"pip3", "install", "--no-cache-dir",
		"-r", "https://raw.githubusercontent.com/zephyrproject-rtos/zephyr/main/scripts/requirements.txt",
		"-r", "https://raw.githubusercontent.com/zephyrproject-rtos/mcuboot/main/scripts/requirements.txt",
		"GitPython", "imgtool", "junitparser", "numpy", "protobuf", "PyGithub",
		"pylint", "sh", "statistics", "west",
	}).WithExec([]string{
		"pip3", "check",
	})

	return image
}

func (image *Image) AddSourceCode() *Image {
	source := image.dag.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{
			"ci",
			"build*",
			"binaries",
			"twister*",
		},
	})

	image.container = image.container.
		WithDirectory("/zephyr/src", source).
		WithWorkdir("/zephyr")

	return image
}

func (image *Image) SetupZephyrModules() *Image {
	image.container = image.container.
		WithMountedCache("/zephyr/modules", image.dag.CacheVolume("zephyr-modules")).
		WithMountedCache("/zephyr/zephyr", image.dag.CacheVolume("zephyr")).
		WithWorkdir("/zephyr").
		WithExec([]string{
			"west", "init", "-l", "src",
		}).WithExec([]string{
		"west", "update",
	})

	return image
}

func (image *Image) BuildApp() *Image {
	image.container = image.container.
		WithWorkdir("/zephyr/src").
		WithExec([]string{
			"west", "build", "-b", "custom_plank", "app",
		})

	return image
}

func (image *Image) ExportBinaries() *Image {
	dir := image.container.Directory("/zephyr/src/build/zephyr")

	if _, err := dir.Export(image.ctx, "./artifacts"); err != nil {
		panic(err)
	}

	return image
}
