provider "sysdig" {

  token = "ebf1f040-c8d7-4b0e-98c8-5c3e9aa23015"

}
resource "sysdig_alert" "testalert" {
  name = "testalert"
  description = "this is the provider"
  alert_id = 1111
  enabled = true
}
