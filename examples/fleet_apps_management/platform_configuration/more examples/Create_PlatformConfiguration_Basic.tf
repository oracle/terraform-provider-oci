
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = "${var.compartment_id}"
  config_category_details {
    config_category = "PRODUCT"
    versions        = ["1"]
  }
  display_name = "displayName"
}
