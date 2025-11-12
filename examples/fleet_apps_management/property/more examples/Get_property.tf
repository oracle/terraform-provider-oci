
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}


data "oci_fleet_apps_management_property" "test_property" {
  property_id = "${oci_fleet_apps_management_property.test_property.id}"
}
variable "compartment_id" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_property" "test_property" {
  compartment_id = "${var.compartment_id}"
  defined_tags   = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  display_name   = "displayName2"
  selection      = "MULTI_CHOICE"
  value_type     = "NUMERIC"
  values         = ["values2"]
}
