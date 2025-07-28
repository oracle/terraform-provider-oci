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
  default = "AMAZON_KINESIS"
}
variable "display_name" {
  default = "Amazon_Kinesis_TerraformTest"
}
variable "technology_type" {
  default = "AMAZON_KINESIS"
}
variable "access_key_id" {
  default = "access_key_id_access_key_id_access_key_id"
}
variable "secret_access_key" {
  default = "access-key"
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
  access_key_id = var.access_key_id

  #Optional
  description = var.description
  freeform_tags = var.freeform_tags
  secret_access_key_secret_id = var.password_secret_id
  does_use_secret_ids = var.does_use_secret_ids
}