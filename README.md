# terraform-provider-sysdig_v1.0.0
Terraform provider for sysdig. Uses the package: https://github.com/CDKGlobal/go-sysdig

**Create an alert in Sysdig**

```
provider "sysdig" {

  token = "XXXXXXXXXXXXXXXXXXXXXX"

}
resource "sysdig_alert" "foo" {
  name = "foo"
  description = "this is the provider"
  enabled = false
  severity = 7
  condition = "timeAvg(cpu.used.percent) >= 95"
  timespan = 600000000
  type = "MANUAL"
  segmentcondition = "None"
  segmentby = ["host.hostname", "host.mac"]
}

```

Note :: Make Sure you do not add an alert ID during the creation of the alert in your terraform plan since that is a 'computed' field and sysdig will
auto-generate it for you.



**Run tests**

Step 1: You will need to add the sysdig token to the config file in the resourceAlert_test.go file

```
var testAccCheckAlertConfig = `
		provider "sysdig" {
		  token = "XXXXXXXXXXXXXXXXXXXXXX"
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
	  }
	  `

```

Step 2: From within the sysdig directory, use the following command:


```
go test -v

```


Note: Acceptance tests create real resources, and often cost money to run.

