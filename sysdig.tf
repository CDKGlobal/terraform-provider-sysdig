provider "sysdig" {

  token = "ebf1f040-c8d7-4b0e-98c8-5c3e9aa23015"

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
