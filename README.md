# Zephyr Example Application

Proof of concept for building zephyr applications in github actions with 
[dagger](https://dagger.io).

The application is from the [zephyr documentation](https://docs.zephyrproject.org/latest/develop/application/index.html#basic-example-application-usage).

* `dagger run go run main.go`
* check the binaries at `./artifacts`

Inspired by [thechangelog setup](https://github.com/thechangelog/changelog.com/blob/master/magefiles/image/main.go)
and their episodes:

* [launching dagger](https://changelog.com/shipit/48)
* [from docker to dagger](https://changelog.com/podcast/550)
