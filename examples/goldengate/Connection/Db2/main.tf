variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "deployment_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}
variable "subnet_id" {}
variable "vault_id" {}
variable "kms_key_id" {}

variable "username" {
  default = "admin"
}
variable "password" {
  default = "bEStrO0nG_1"
}
variable "database_name" {
  default = "TF_DB2"
}
variable "security_protocol" {
  default = "PLAIN"
}
variable "description" {
  default = "Created as example for TERSI-3444 Connections R5"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "DB2"
}
variable "display_name" {
  default = "Db2_TerraformTest"
}
variable "technology_type" {
  default = "DB2_ZOS"
}
variable "host"{
  default = "10.0.0.129"
}
variable "port" {
  default = "10000"
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

  #Required for Postgresql connection_type
  technology_type = var.technology_type
  host = var.host
  port = var.port
  username = var.username
  password = var.password
  database_name = var.database_name
  security_protocol = var.security_protocol

  #Optional
  description = var.description
  freeform_tags = var.freeform_tags
  subnet_id = var.subnet_id
  vault_id = var.vault_id
  key_id = var.kms_key_id
}