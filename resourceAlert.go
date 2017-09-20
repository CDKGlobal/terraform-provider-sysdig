package main

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/jgensler8/go-sysdig"
  "fmt"
  "log"
)

func resourceAlert() *schema.Resource {
  return &schema.Resource{
    Create: resourceAlertCreate,
    Read:   resourceAlertRead,
    Exists: resourceAlertExists,
    Delete: resourceAlertDelete,
    Schema: map[string]*schema.Schema{
      "name": {
				Type:     schema.TypeString,
				Required: true,
        ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
        ForceNew: true,

			},
      "alert_id": {
				Type:     schema.TypeInt,
				Required: true,
        ForceNew: true,

			},
      "enabled": {
				Type:     schema.TypeBool,
				Required: true,
        ForceNew: true,

			},
    },
  }
}

// Define the CRUD functions - Create, Read, Update and Delete


func resourceAlertCreate(d *schema.ResourceData, meta interface{}) error {
  api := meta.(*swagger.DefaultApi)

  alertInput := swagger.AlertInput{
    Alert: swagger.Alert{

      Name : d.Get("name").(string),
      Description : d.Get("description").(string),
      Id : d.Get("alert_id").(int64),
      Enabled : d.Get("enabled").(bool),

    },
  }
  log.Println(alertInput)
  alertinput,alertresponse, err := api.CreateAlert(alertInput)
  log.Println(alertinput)
  log.Println(alertresponse)
  log.Println(err)
  if err != nil {
		return fmt.Errorf("error creating alert: %s", err.Error())
	}
  return  nil

}



func resourceAlertExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
  return true, nil
}


func resourceAlertRead(d *schema.ResourceData, meta interface{}) error {
  return nil
}

func resourceAlertDelete(d *schema.ResourceData, meta interface{}) error {
  return nil
}
