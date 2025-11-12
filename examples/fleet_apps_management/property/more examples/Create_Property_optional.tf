
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}


variable "compartment_id" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_property" "test_property" {
  compartment_id = "${var.compartment_id}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "value")}"
  display_name   = "displayName"
  selection      = "SINGLE_CHOICE"
  value_type     = "STRING"
  values         = ["values"]
}
