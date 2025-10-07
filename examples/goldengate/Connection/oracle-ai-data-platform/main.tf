variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}
variable "subnet_id" {}

variable "locks_type" {
  default = "FULL"
}

variable "locks_message" {
  default = "message"
}

variable "is_lock_override" {
  default = true
}

variable "description" {
  default = "Created as example for TERSI-4887 Connections R9"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "ORACLE_AI_DATA_PLATFORM"
}
variable "display_name" {
  default = "OracleAiDataPlatform_TerraformTest"
}
variable "technology_type" {
  default = "ORACLE_AI_DATA_PLATFORM"
}
variable "connection_url" {
  default = "jdbc:spark://gateway.datalake.us-phoenix-1.oci.oraclecloud.com/default"
}
variable "should_use_resource_principal" {
  default = true
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}
resource "oci_golden_gate_connection" "test_connection"{
  #Required
  compartment_id = var.compartment_id
  connection_type = var.connection_type
  display_name = var.display_name

  #Required for OracleAiDataPlatform connection_type
  technology_type = var.technology_type
  connection_url = var.connection_url

  #Optional
  description = var.description

  #Optional for OracleAiDataPlatform connection_type
  should_use_resource_principal = var.should_use_resource_principal
  locks {
    type = var.locks_type
    message = var.locks_message
  }
  is_lock_override = var.is_lock_override
  lifecycle {
    ignore_changes = [defined_tags, locks, freeform_tags, is_lock_override]
  }
}