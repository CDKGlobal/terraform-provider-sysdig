provider "sysdig" {

  token = "XXXXXXXXXX"

}
resource "sysdig_alert" "foo" {
  name = "foo"
  description = "this is the provider"
  enabled = true
  severity = 7
  condition = "timeAvg(cpu.used.percent) >= 95"
  timespan = 600000000
  type = "MANUAL"
  segmentcondition = "ANY"
  segmentby = ["host.mac"]
  filter = "agent.tag.location='prod1-pdx'"
  notificationchannelids = [8227,8611]
}
