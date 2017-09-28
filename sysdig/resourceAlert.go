package sysdig

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/CDKGlobal/go-sysdig"
	"log"
	"fmt"
	"github.com/tidwall/gjson"
)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertCreate,
		Read:    resourceAlertRead,
		Update: resourceAlertUpdate,
		Delete: resourceAlertDelete,
		Schema: map[string]*schema.Schema{
			"alert_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"severity": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"timespan": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"segmentcondition": {
				Type:     schema.TypeString,
				Required: true,

			},
			"segmentby" : {
				Type: schema.TypeList,
				Required: true,
				Elem: &schema.Schema {Type : schema.TypeString},
			},


		},
	}
}

// Define the CRUD functions - Create, Read, Update and Delete

func resourceAlertCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("****** We are in the create function ******")
	api := meta.(*swagger.DefaultApi)

	alertInput := swagger.AlertInput{
		Alert : swagger.Alert{
			Name : d.Get("name").(string),
			Type_ : d.Get("type").(string),
			Severity : int32(d.Get("severity").(int)),
			Description : d.Get("description").(string),
			Enabled : d.Get("enabled").(bool),
			Condition : d.Get("condition").(string),
			Timespan : int64(d.Get("timespan").(int)),
			SegmentCondition :  swagger.SegmentCondition{Type_: "ANY"},
			SegmentBy:  []string{ "host.mac"},

		},
	}

	alert,alertresponse, err := api.CreateAlert(alertInput)

	log.Printf("The alert created: %s",alert)
	if err != nil {
		return fmt.Errorf("error creating alert: %s", err.Error())
	}

	d.Set("alert_id",alert.Alert.Id)
	d.SetId(gjson.GetBytes(alertresponse.Payload,"alert.name").String());

	log.Printf("This is what the ID is %v", d.Get("alert_id"))

	return  nil

}
func resourceAlertRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the read function ******")

	api := meta.(*swagger.DefaultApi)

	alert_id:= int64(d.Get("alert_id").(int))

	alert,response, err := api.GetAlert(alert_id)
	log.Print(alert.Alert)

	if err != nil {
		return fmt.Errorf("error retrieving alert using the id given : %s", err.Error())
	}
	log.Print("The response generated while fetching the alert using `id` in the read function %s",response)

	d.Set("name",alert.Alert.Name)
	d.Set("type",alert.Alert.Type_)
	d.Set("description",alert.Alert.Description)
	d.Set("segmentby",alert.Alert.SegmentBy)
	d.Set("segmentcondition",alert.Alert.SegmentCondition)
	d.Set("condition",alert.Alert.Condition)
	d.Set("enabled",alert.Alert.Enabled)
	d.Set("severity",alert.Alert.Severity)


	return nil
}


func resourceAlertUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the update function ******")

	api := meta.(*swagger.DefaultApi)
	alert_id:= int64(d.Get("alert_id").(int))

	alertInput := swagger.AlertInput{
		Alert : swagger.Alert{
			Name : d.Get("name").(string),
			Type_ : d.Get("type").(string),
			Severity : int32(d.Get("severity").(int)),
			Description : d.Get("description").(string),
			Enabled : d.Get("enabled").(bool),
			Condition : d.Get("condition").(string),
			Timespan : int64(d.Get("timespan").(int)),
			SegmentCondition :  swagger.SegmentCondition{Type_: "ANY"},
			SegmentBy:  []string{ "host.mac"},

		},
	}
	alertinput,alertresponse, err := api.UpdateAlert(alert_id,alertInput)
	log.Printf("%s",alertinput)
	log.Printf("%s",alertresponse.Payload)
	if err != nil {
		return fmt.Errorf("error updating alert: %s", err.Error())
	}
	return  nil
}


func resourceAlertDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the delete function ******")

	api := meta.(*swagger.DefaultApi)
		log.Printf("Before Deletion %v", d.State())
		alert_id:= int64(d.Get("alert_id").(int))

		_, err := api.DeleteAlert(int64(alert_id))

		log.Printf("After Deletion %v", d.State())


	if err != nil {
			return err
		}
	d.SetId("")
	return nil
}
