provider "sysdig" {

  token = "Value of your token here"

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
  filter = "agent.tag.location='prod1-pdx'"
}
