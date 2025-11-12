
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }


resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = "${var.compartment_id}"
  config_category_details {
    compatible_products {
      id = "ocid1.famsplatformconfiguration.oc1."
    }
    components      = ["components"]
    config_category = "PRODUCT"
    credentials {
      id = "ocid1.famsplatformconfiguration.oc1."
    }
    patch_types {
      id = "ocid1.famsplatformconfiguration.oc1."
    }
    versions = ["1"]
  }
  defined_tags = "${map("Oracle-Tags.CreatedBy", "value")}"
  description  = "description"
  display_name = "displayName"
  freeform_tags = {
    "bar-key" = "value"
  }
}
