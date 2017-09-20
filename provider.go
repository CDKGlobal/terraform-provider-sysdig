package main

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  sysdig "github.com/jgensler8/go-sysdig"
	"fmt"
  "log"

)

func Provider() terraform.ResourceProvider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "token": &schema.Schema{
        Type: schema.TypeString,
        Required: true,
        DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
        Description: "API Token",
      },
    },
    ResourcesMap: map[string]*schema.Resource{
      "sysdig_alert": resourceAlert(),
    },
    ConfigureFunc: configureProvider,
  }
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
    token := d.Get("token").(string)
    configuration := sysdig.NewConfiguration()

  	configuration.APIKeyPrefix["Authorization"] = fmt.Sprintf("Bearer %s", token)

  	api := sysdig.NewDefaultApi()
  	api.Configuration = configuration

    log.Println("[INFO] Initializing sysdig client")
    return api,nil
}
