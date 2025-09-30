variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "description" {
  default = "Created as example for TERSI-7228 Connections R7"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "AMAZON_S3"
}
variable "display_name" {
  default = "S3_TerraformTest"
}
variable "technology_type" {
  default = "AMAZON_S3"
}

variable "access_key_id" {
  default = "AKIAIOSFODNN7EXAMPLE"
}

variable "secret_access_key" {
  default = "mySecret"
}

variable "security_attributes" {
  default = {
    "oracle-zpr.sensitivity.value" = "42"
    "oracle-zpr.sensitivity.mode" = "enforce"
  }
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
  secret_access_key = var.secret_access_key

  description = var.description

  security_attributes = var.security_attributes

}

data "oci_golden_gate_connection" "fetched_connection" {
  connection_id = oci_golden_gate_connection.test_connection.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.fetched_connection.display_name
}
