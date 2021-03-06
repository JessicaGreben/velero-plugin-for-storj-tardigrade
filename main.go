package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	veleroplugin "github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func main() {
	veleroplugin.NewServer().
		BindFlags(pflag.CommandLine).
		RegisterObjectStore("velero.io/gcp", newStorjObjectStore).
		Serve()
}

func newStorjObjectStore(logger logrus.FieldLogger) (interface{}, error) {
	return newObjectStore(logger), nil
}
