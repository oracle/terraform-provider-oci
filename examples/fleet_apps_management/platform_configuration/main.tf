
variable "tenancy_ocid" {}

variable "private_key_path" {}
variable "user_ocid" {}
variable "fingerprint" {}

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {}

variable "compatible_product_id" {}

variable "credential_id" {}

variable "patch_type_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = var.compartment_id
  config_category_details {
    compatible_products {
      id = var.compatible_product_id
    }
    components      = ["components2"]
    config_category = "PRODUCT"
    credentials {
      id = var.credential_id
    }
    patch_types {
      id = var.patch_type_id
    }
    versions = ["1"]
  }
  description  = "description2"
  display_name = "displayName2"
}
