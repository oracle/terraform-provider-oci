variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "source_connection_id" { }
variable "target_connection_id" { }
variable "display_name" {
  default = "Data fabric pipeline display name"
}
variable "license_model" {
  default = "LICENSE_INCLUDED"
}
variable "recipe_type" {
  default = "ZERO_ETL"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_pipeline" "test_pipeline" {
  # Required
  compartment_id = var.compartment_id
  display_name = var.display_name
  license_model = var.license_model
  recipe_type = var.recipe_type
  source_connection_details {
    connection_id = var.source_connection_id
  }
  target_connection_details {
    connection_id = var.target_connection_id
  }
}