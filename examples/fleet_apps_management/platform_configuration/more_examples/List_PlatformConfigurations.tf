
variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "linked_product" { default = "ocid1.famsplatformconfiguration.oc1.." }
variable "compatible_products" { default = "ocid1.famsplatformconfiguration.oc1.." }
variable "credentials" { default = "ocid1.famsplatformconfiguration.oc1.." }
variable "patch_types" { default = "ocid1.famsplatformconfiguration.oc1.." }


data "oci_fleet_apps_management_platform_configurations" "test_platform_configurations" {
  compartment_id            = "${var.compartment_id}"
  compartment_id_in_subtree = "false"
  config_category           = "PRODUCT"
  display_name              = "displayName2"
  filter {
    name   = "id"
    values = ["${oci_fleet_apps_management_platform_configuration.test_platform_configuration.id}"]
  }
  id    = "${oci_fleet_apps_management_platform_configuration.test_platform_configuration.id}"
  state = "ACTIVE"
  type  = "USER_DEFINED"
}


resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = "${var.compartment_id}"
  config_category_details {
    compatible_products {
      id = "${var.compatible_products}"
    }
    components      = ["components2"]
    config_category = "PRODUCT"
    credentials {
      id = "${var.credentials}"
    }
    is_compliance_policy_required_for_softlink = "true"
    is_softlink                                = "true"
    link_product_id                            = "${var.linked_product_name}"
    patch_types {
      id = "${var.patch_types}"
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
