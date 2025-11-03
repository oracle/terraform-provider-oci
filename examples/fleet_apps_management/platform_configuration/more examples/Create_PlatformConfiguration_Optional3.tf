
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "compartment_id_for_update" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = "${var.compartment_id}"
  config_category_details {
    compatible_products {
      id = "ocid1.famsplatformconfiguration.oc1."
    }
    components      = ["components2"]
    config_category = "PRODUCT"
    credentials {
      id = "ocid1.famsplatformconfiguration.oc1."
    }
    patch_types {
      id = "ocid1.famsplatformconfiguration.oc1."
    }
    versions = ["2"]
  }
  defined_tags = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  description  = "description2"
  display_name = "displayName2"
  freeform_tags = {
    "bar-key" = "value"
  }
}
