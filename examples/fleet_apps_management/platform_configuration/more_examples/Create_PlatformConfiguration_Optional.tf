
variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "linked_product" { default = "ocid1.famsplatformconfiguration.oc1.phx." }
variable "compatible_products" { default = "ocid1.famsplatformconfiguration.oc1.phx." }
variable "credentials" { default = "ocid1.famsplatformconfiguration.oc1.phx." }
variable "patch_types" { default = "ocid1.famsplatformconfiguration.oc1.phx." }

resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = "${var.compartment_id}"
  config_category_details {
    compatible_products {
      id = "${var.compatible_products}"
    }
    components      = ["components"]
    config_category = "PRODUCT"
    credentials {
      id = "${var.credentials}"
    }
    is_compliance_policy_required_for_softlink = "false"
    is_softlink                                = "false"
    link_product_id                            = "${var.linked_product}"
    patch_types {
      id = "${var.patch_types}"
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
