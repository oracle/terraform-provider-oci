variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "source_connection_id" { }
variable "target_connection_id" { }
variable "test_subnet_id" {}
variable "display_name" {
  default = "Data fabric pipeline display name"
}
variable "license_model" {
  default = "LICENSE_INCLUDED"
}
variable "recipe_type" {
  default = "ZERO_ETL"
}
variable "pipeline_cpu_core_count" {
  default = 1
}
variable "pipeline_is_auto_scaling_enabled" {
  default = false
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

  # Optional
  subnet_id = var.test_subnet_id
  is_auto_scaling_enabled = var.pipeline_is_auto_scaling_enabled
  cpu_core_count          = var.pipeline_cpu_core_count
}

data "oci_golden_gate_pipeline" "test_pipelines" {
  pipeline_id = oci_golden_gate_pipeline.test_pipeline.id
}

# Output the subnet ID
output "subnet_id" {
  value = data.oci_golden_gate_pipeline.test_pipelines.subnet_id
}