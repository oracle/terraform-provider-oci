
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

data "oci_fleet_apps_management_onboardings" "test_onboardings" {
  compartment_id = "${var.tenancy_ocid}"
  filter {
    name   = "id"
    values = ["${oci_fleet_apps_management_onboarding.test_onboarding.id}"]
  }
  state = "ACTIVE"
}
variable "compartment_id" { default = "ocid1.tenancy.oc1.." }

resource "oci_fleet_apps_management_onboarding" "test_onboarding" {
  compartment_id = "${var.compartment_id}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  freeform_tags = {
    "Department" = "Accounting"
  }
  is_cost_tracking_tag_enabled = "false"
  is_fams_tag_enabled          = "false"
}
