package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/paperg/terraform-etcd/etcd"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: etcd.Provider,
	})
}
