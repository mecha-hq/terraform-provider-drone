package main

import (
	"context"
	"flag"
	"log"

	"github.com/FriendsOfDrone/terraform-provider-drone/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	version string = "dev"
	commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: provider.New(version),
	}

	if debugMode {
		if err := plugin.Debug(context.Background(), "registry.terraform.io/FriendsOfDrone/terraform-provider-drone", opts); err != nil {
			log.Fatal(err.Error())
		}

		return
	}

	plugin.Serve(opts)
}
