
variable "tenancy_ocid" {
#   default = ""
}

variable "ssh_public_key" {
#   default = ""
}

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" { 
    # default = "" 
}

variable "compatible_product_id" { 
    # default = "" 
}

variable "credential_id" { 
    # default = "" 
}

variable "patch_type_id" { 
    # default = "" 
}

resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = "${var.compartment_id}"
  config_category_details {
    compatible_products {
      id = "${var.compatible_product_id}"
    }
    components      = ["components2"]
    config_category = "PRODUCT"
    credentials {
      id = "${var.credential_id}"
    }
    patch_types {
      id = "${var.patch_type_id}"
    }
    versions = ["versions2"]
  }
  description  = "description2"
  display_name = "displayName2"
}
