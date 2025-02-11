variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "deployment_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "description" {
  default = "Created as example for TERSI-6831 Connections R6"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "DATABRICKS"
}
variable "display_name" {
  default = "Databricks_TerraformTest"
}
variable "technology_type" {
  default = "DATABRICKS"
}
variable "connection_url"{
  default = "jdbc:databricks://adb-33934.4.azuredatabricks.net:443/default;transportMode=http;ssl=1;httpPath=sql/protocolv1/o/3393########44/0##3-7-hlrb"
}
variable "authentication_type" {
  default = "PERSONAL_ACCESS_TOKEN"
}

variable "password" {
  default = "bEStrO0nG_1"
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
  connection_url = var.connection_url
  authentication_type = var.authentication_type

  #Required for Postgresql connection_type
  technology_type = var.technology_type
  password = var.password

  #Optional
  description = var.description
  freeform_tags = var.freeform_tags
}