package main

import (
	"github.com/ForestEckhardt/no-source-cnb/nosource"
	"github.com/cloudfoundry/packit"
)

func main() {
	packit.Build(nosource.Build())
}