package sysdig

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/CDKGlobal/go-sysdig/generated"
	"log"
	"fmt"
	"context"
	"strconv"
	"strings"

)

func resourceAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertCreate,
		Read:    resourceAlertRead,
		Update: resourceAlertUpdate,
		Delete: resourceAlertDelete,
		Schema: map[string]*schema.Schema{
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
			"autocreated": {
				Type:     schema.TypeBool,
				Computed: true,
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
			"version" : {
				Type: schema.TypeInt,
				Computed: true,
			},
			"teamid" : {
				Type: schema.TypeInt,
				Computed: true,
			},
			"createdon" : {
				Type: schema.TypeInt,
				Computed: true,
			},
			"modifiedon" : {
				Type: schema.TypeInt,
				Computed: true,
			},
			"alert_id": {
				Type:     schema.TypeInt,
				Computed: true,
				ForceNew: true,
			},

		},
	}
}

//building the alert construct
func buildAlertStruct(d *schema.ResourceData) *swagger.Alert {

	var alert swagger.Alert

	if attr, ok := d.GetOk("alert_id"); ok {
		alert.Id=(int64)(attr.(int))
	}
	if attr, ok := d.GetOk("name"); ok {
		alert.Name=attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		alert.Type_=attr.(string)
	}
	if attr, ok := d.GetOk("severity"); ok {
		alert.Severity= (int32)(attr.(int))
	}
	if attr, ok := d.GetOk("version"); ok {
		alert.Version= (int32)(attr.(int))
	}
	if attr, ok := d.GetOk("description"); ok {
		alert.Description=attr.(string)
	}
	if attr, ok := d.GetOk("enabled"); ok {
		alert.Enabled=attr.(bool)
	}
	if attr, ok := d.GetOk("autocreated"); ok {
		alert.AutoCreated=attr.(bool)
	}
	if attr, ok := d.GetOk("timespan"); ok {
		alert.Timespan=(int64)(attr.(int))
	}
	if attr, ok := d.GetOk("teamid"); ok {
		alert.TeamId=(int64)(attr.(int))
	}
	if attr, ok := d.GetOk("createdon"); ok {
		alert.CreatedOn=(int64)(attr.(int))
	}
	if attr, ok := d.GetOk("modifiedon"); ok {
		alert.ModifiedOn=(int64)(attr.(int))
	}
	if attr, ok := d.GetOk("condition"); ok {
		alert.Condition= attr.(string)
	}

	if attr, ok := d.GetOk("segmentcondition"); ok {
		alert.SegmentCondition = &swagger.SegmentCondition{
			Type_: attr.(string),
		}
	}

	if attr, ok := d.GetOk("filter"); ok {
		alert.Filter=attr.(string)
	}
	segmentby := []string{}
	for _, s := range d.Get("segmentby").([]interface{}) {
		segmentby = append(segmentby, s.(string))
	}
	alert.SegmentBy = segmentby

	notificationchannelids := []int64{}
	for _, s := range d.Get("notificationchannelids").([]interface{}) {
		notificationchannelids = append(notificationchannelids,(int64)(s.(int)))
	}
	alert.NotificationChannelIds = notificationchannelids

	return &alert
}

//building resource data to maintain consistency between terraform state and sysdig
func updateResourceData(d *schema.ResourceData, alert *swagger.Alert) error{

	log.Printf("The alert received in the buildResourceData function: %v",alert)

	d.SetId(strconv.Itoa(int(alert.Id)))
	d.Set("alert_id",alert.Id)
	d.Set("name",alert.Name)
	d.Set("type",alert.Type_)
	d.Set("description",alert.Description)
	d.Set("segmentby",alert.SegmentBy)
	d.Set("segmentcondition",alert.SegmentCondition)
	d.Set("condition",alert.Condition)
	d.Set("enabled",alert.Enabled)
	d.Set("severity",alert.Severity)
	d.Set("version", alert.Version)
	d.Set("filter", alert.Filter)
	d.Set("teamid",alert.TeamId)
	d.Set("createdon",alert.CreatedOn)
	d.Set("modifiedon",alert.ModifiedOn)
	d.Set("notificationchannelids", alert.NotificationChannelIds)

	log.Printf("Resource Data object after sets: %v", d)
	return nil
}


// Define the CRUD functions - Create, Read, Update and Delete

func resourceAlertCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("****** We are in the create function ******")
	api := meta.(*swagger.APIClient).DefaultApi

	alertInput := swagger.AlertInput{buildAlertStruct(d)}

	log.Printf("This is what the generated alert is %v", alertInput.Alert)

	alert,alertresponse, err := api.CreateAlert(context.Background(),alertInput)
	if err != nil {
		if strings.Contains(err.Error(), "422 Unprocessable Entity") {
			return fmt.Errorf("****** Looks like an alert with the same name already exists in this sysdig account. Names need to be unique ******")
		}
		return fmt.Errorf("error creating alert: %s", err.Error())
	}
	log.Printf("This is what the alertresponse is %v", alertresponse)

	log.Printf("The alert created: %v",alert.Alert)
	updateResourceData(d,alert.Alert)
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

	updateResourceData(d,alert.Alert)

	return nil
}


func resourceAlertUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("****** We are in the update function ******")


	api := meta.(*swagger.APIClient).DefaultApi

	alertInput := swagger.AlertInput{buildAlertStruct(d)}
	log.Printf("Alertinput in update : %v",alertInput.Alert)
	log.Printf("Updating alert with id: %v",alertInput.Alert.Id)

	alertinput,alertresponse, err := api.UpdateAlert(context.Background(),alertInput.Alert.Id,alertInput)
	log.Printf("%v",alertinput)
	log.Printf("%v",alertresponse)
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
