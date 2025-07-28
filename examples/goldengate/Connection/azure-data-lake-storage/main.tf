variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "deployment_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}
variable "password_secret_id" {}

variable "description" {
  default = "Created as example for TERSI-4594 Connections R8"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "AZURE_DATA_LAKE_STORAGE"
}
variable "display_name" {
  default = "Azure_Data_Lake_Storage_TerraformTest"
}
variable "technology_type" {
  default = "AZURE_DATA_LAKE_STORAGE"
}
variable "account_name"{
  default = "testaccount"
}
variable "authentication_type" {
  default = "SHARED_KEY"
}
variable "does_use_secret_ids" {
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
  technology_type = var.technology_type
  display_name = var.display_name
  account_name = var.account_name
  authentication_type = var.authentication_type

  #Required for SHARED_KEY authentication_type  
  account_key_secret_id = var.password_secret_id

  #Optional
  description = var.description
  freeform_tags = var.freeform_tags
  does_use_secret_ids = var.does_use_secret_ids
}