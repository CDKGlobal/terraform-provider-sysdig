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
						"sysdig_alert.foo", "segmentcondition", "None"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "severity", "4"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "condition", "timeAvg(cpu.used.percent) >= 95"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "timespan", "600000000"),
					resource.TestCheckResourceAttr(
						"sysdig_alert.foo", "segmentby.0", "host.mac"),
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
		  token = "XXXXXXXX"
		}
		resource "sysdig_alert" "foo" {
			name = "foo"
			description = "this is the provider"
			enabled = true
			severity = 4
			condition = "timeAvg(cpu.used.percent) >= 95"
			timespan = 600000000
			type = "MANUAL"
			segmentcondition = "None"
			segmentby = ["host.mac"]
			filter = "agent.tag.location='prod-pdx'"
			notificationchannelids =[8227,8611]
	  }
	  `
func alertDestroyHelper(s *terraform.State, api *swagger.DefaultApiService) error {
	log.Printf("****** Destroy Helper is called ******")

	for _, r := range s.RootModule().Resources {
		id, _:= strconv.Atoi(r.Primary.Attributes["alert_id"])
		log.Printf("ID in destroy helper %v", id)

		alert,response, err := api.GetAlert(context.Background(),int64(id))
		log.Printf("There is an alert %v ", alert)
		log.Printf("There is an alert response %v ", response)


		if err != nil {
			if strings.Contains(err.Error(), "unexpected end of JSON input") {
				continue
			}
			if strings.Contains(err.Error(), "404 Not Found") {
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
