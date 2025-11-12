
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.tenancy.oc1.." }

resource "oci_fleet_apps_management_onboarding" "test_onboarding" {
  compartment_id = "${var.compartment_id}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "value")}"
  freeform_tags = {
    "bar-key" = "value"
  }
  is_cost_tracking_tag_enabled = "false"
  is_fams_tag_enabled          = "false"
}
