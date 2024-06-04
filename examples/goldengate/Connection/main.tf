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

variable "locks_type" {
  default = "FULL"
}

variable "locks_message" {
  default = "message"
}

variable "is_lock_override" {
  default = true
}

variable "username" {
  default = "admin"
}
variable "password" {
  default = "bEStrO0nG_1"
}
variable "database_name" {
  default = "TF_PostgresqlDB"
}
variable "security_protocol" {
  default = "PLAIN"
}
variable "description" {
  default = "Created as example for TERSI-2066 Connections R2"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "private_ip" {
  default = "10.0.1.78"
}
variable "connection_type" {
  default = "POSTGRESQL"
}
variable "display_name" {
  default = "Postgresql_TerraformTest"
}
variable "technology_type" {
  default = "POSTGRESQL_SERVER"
}
variable "host"{
  default = "10.0.0.129"
}
variable "port" {
  default = "14"
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

  #Optional for Postgresql connection_type
  private_ip = var.private_ip
  locks {
    type = var.locks_type
    message = var.locks_message
  }
  is_lock_override = var.is_lock_override
  lifecycle {
    ignore_changes = [defined_tags, locks, freeform_tags, is_lock_override]
  }
}