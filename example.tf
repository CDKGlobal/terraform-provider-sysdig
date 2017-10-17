provider "sysdig" {

  token = "ebf1f040-439855389583-2924240-3015"

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
