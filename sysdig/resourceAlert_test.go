package sysdig

import (

	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/CDKGlobal/go-sysdig/generated"
	"log"
	"strconv"
	"fmt"
	"strings"
	"context"
)

func TestAccAlert_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAlertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckAlertConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "name", "foo"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "description", "this is the provider"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "type", "MANUAL"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "segmentcondition", "ANY"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "severity", "4"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "condition", "timeAvg(cpu.used.percent) >= 95"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "timespan", "600000000"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "segmentby.0", "host.mac"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "notificationchannelids.0", "8227"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "notificationchannelids.1", "8611"),
				),
			},
		},
	})
}
func testAccCheckAlertDestroy(s *terraform.State) error {
	log.Printf("****** We are in the delete test function ******")

	api := testAccProvider.Meta().(*swagger.APIClient).DefaultApi

	if err := alertDestroyHelper(s, api); err != nil {
		return err
	}
	return nil
}


var testAccCheckAlertConfig = `
			provider "sysdig" {

			  token = "XXXXXXXXXX"

			}
			resource "sysdig_alert" "foo" {
			  name = "foo"
			  description = "this is the provider"
			  enabled = true
			  severity = 4
			  condition = "timeAvg(cpu.used.percent) >= 95"
			  timespan = 600000000
			  type = "MANUAL"
			  segmentcondition = "ANY"
			  segmentby = ["host.mac"]
			  filter = "agent.tag.location='prod1-pdx'"
			  notificationchannelids = [8227,8611]
			}
	  `
func alertDestroyHelper(s *terraform.State, api *swagger.DefaultApiService) error {
	log.Printf("****** Destroy Helper is called ******")

	for _, r := range s.RootModule().Resources {
		id, iderr:= strconv.Atoi(r.Primary.Attributes["alert_id"])
		if iderr != nil {
			log.Printf("There was an error fetching the alert_id %v", iderr.Error())
		}
		log.Printf("ID in destroy helper %v", id)

		alert,response, err := api.GetAlert(context.Background(),int64(id))
		log.Printf("There is an alert %v ", alert)
		log.Printf("There is an alert response %v ", response)


		if err != nil {
			if strings.Contains(err.Error(), "unexpected end of JSON input") {
				log.Printf("******The alert has been deleted succesfully******")
				continue
			}
			if strings.Contains(err.Error(), "404 Not Found") {
				log.Printf("******The alert has been deleted succesfully******")
				continue
			}
		}

		log.Printf("Alert in destroy helper %v", alert.Alert.Description)
		log.Printf("Alert response in destroy helper %v", response)
		log.Printf("Error in destroy helper %v", err)


		 if alert.Alert.Name != "" {
			 return fmt.Errorf("The alert still exists or this might be some lingering alert")

		 }
	}
	return nil
}
