package sysdig

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/CDKGlobal/go-sysdig/generated"
	"log"
	"fmt"
	"context"
	"strconv"
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
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
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
			"notificationchannelids" : {
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Schema {Type : schema.TypeInt},
			},
			"filter" : {
				Type: schema.TypeString,
				Required: true,
			},


		},
	}
}

// Define the CRUD functions - Create, Read, Update and Delete

func resourceAlertCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("****** We are in the create function ******")
	api := meta.(*swagger.APIClient).DefaultApi


	alertInput := swagger.AlertInput{
		Alert : &swagger.Alert{
			Name : d.Get("name").(string),
			Type_ : d.Get("type").(string),
			Severity : int32(d.Get("severity").(int)),
			Description : d.Get("description").(string),
			Enabled : d.Get("enabled").(bool),
			Condition : d.Get("condition").(string),
			Timespan : int64(d.Get("timespan").(int)),
			SegmentCondition :  &swagger.SegmentCondition{Type_: "ANY"},
			SegmentBy:  []string{ "host.mac"},
			NotificationChannelIds: []int64 {8227, 8611},
			Filter: d.Get("filter").(string),
		},
	}

	alert,alertresponse, err := api.CreateAlert(context.Background(),alertInput)
	log.Printf("This is what the alertresponse is %v", alertresponse)

	log.Printf("The alert created: %s",alert.Alert)
	if err != nil {
		return fmt.Errorf("error creating alert: %s", err.Error())
	}
	d.SetId(strconv.Itoa(int(alert.Alert.Id)))
	d.Set("alert_id",alert.Alert.Id)
	log.Printf("This is what the ID is %v", d.Get("alert_id"))
	log.Printf("This is what the name is %s", d.Get("name"))


	return  nil

}
func resourceAlertRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the read function ******")

	api := meta.(*swagger.APIClient).DefaultApi

	alert_id:= int64(d.Get("alert_id").(int))

	alert,response, err := api.GetAlert(context.Background(),alert_id)
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
	d.Set("filter", alert.Alert.Filter)
	d.Set("notificationchannelids", alert.Alert.NotificationChannelIds)


	return nil
}


func resourceAlertUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the update function ******")

	api := meta.(*swagger.APIClient).DefaultApi
	alert_id:= int64(d.Get("alert_id").(int))

	alertInput := swagger.AlertInput{
		Alert : &swagger.Alert{
			Name : d.Get("name").(string),
			Type_ : d.Get("type").(string),
			Severity : int32(d.Get("severity").(int)),
			Description : d.Get("description").(string),
			Enabled : d.Get("enabled").(bool),
			Condition : d.Get("condition").(string),
			Timespan : int64(d.Get("timespan").(int)),
			SegmentCondition :  &swagger.SegmentCondition{Type_: "ANY"},
			SegmentBy:  []string{ "host.mac"},
			NotificationChannelIds: []int64 {8227, 8611},
			Filter: d.Get("filter").(string),

		},
	}
	alertinput,alertresponse, err := api.UpdateAlert(context.Background(),alert_id,alertInput)
	log.Printf("%s",alertinput)
	log.Printf("%s",alertresponse)
	if err != nil {
		return fmt.Errorf("error updating alert: %s", err.Error())
	}
	return  nil
}


func resourceAlertDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the delete function ******")

		api := meta.(*swagger.APIClient).DefaultApi
		log.Printf("Before Deletion %v", d.State())
		alert_id:= int64(d.Get("alert_id").(int))

		_, err := api.DeleteAlert(context.Background(),int64(alert_id))

		log.Printf("After Deletion %v", d.State())


	if err != nil {
			return err
		}
	d.SetId("")
	return nil
}
