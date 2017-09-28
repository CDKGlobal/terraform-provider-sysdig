package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
    "github.com/anuiq/terraform-provider-sysdig_v1.0.0/sysdig"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{

        ProviderFunc: func() terraform.ResourceProvider {
            return sysdig.Provider()
        },
    })
}
