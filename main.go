package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
    "github.com/anuiq/terraform-provider-sysdig/sysdig"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{

        ProviderFunc: func() terraform.ResourceProvider {
            return sysdig.Provider()
        },
    })
}
