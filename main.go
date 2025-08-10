package main

import (
	"helloworld/router"
	"helloworld/service"

	"github.com/forbearing/golib/bootstrap"
	. "github.com/forbearing/golib/util"
)

func main() {
	RunOrDie(bootstrap.Bootstrap)
	RunOrDie(service.Init)
	RunOrDie(router.Init)
	RunOrDie(bootstrap.Run)
}
