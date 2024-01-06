package main

import (
	_ "embed"

	"github.com/coregenesis/coregenesis/command/root"
	"github.com/coregenesis/coregenesis/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
