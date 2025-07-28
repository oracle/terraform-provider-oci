variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "description" {
  default = "Created as example for TERSI-4594 Connections R8"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "HDFS"
}
variable "display_name" {
  default = "Hdfs_TerraformTest"
}
variable "technology_type" {
  default = "HDFS"
}
variable "core_site_xml" {
  default = "PD94bWwgdmVyc2lvbj0iMS4wIj8+Cjw/eG1sLXN0eWxlc2hlZXQgdHlwZT0idGV4dC94c2wiIGhyZWY9ImNvbmZpZ3VyYXRpb24ueHNsIj8+Cgo8IS0tIFB1dCBzaXRlLXNwZWNpZmljIHByb3BlcnR5IG92ZXJyaWRlcyBpbiB0aGlzIGZpbGUuIC0tPgoKPGNvbmZpZ3VyYXRpb24+CiAgPCEtLSBmaWxlIHN5c3RlbSBwcm9wZXJ0aWVzIC0tPgogIDxwcm9wZXJ0eT4KICAgIDxuYW1lPmZzLmRlZmF1bHRGUzwvbmFtZT4KICAgIDx2YWx1ZT5oZGZzOi8vZm9vLmJhci5jb206ODAyMDwvdmFsdWU+CiAgICA8ZGVzY3JpcHRpb24+VGhlIG5hbWUgb2YgdGhlIGRlZmF1bHQgZmlsZSBzeXN0ZW0uICBFaXRoZXIgdGhlCiAgICAgIGxpdGVyYWwgc3RyaW5nICJsb2NhbCIgb3IgYSBob3N0OnBvcnQgZm9yIE5ERlMuCiAgICA8L2Rlc2NyaXB0aW9uPgogICAgPGZpbmFsPnRydWU8L2ZpbmFsPgogIDwvcHJvcGVydHk+CjwvY29uZmlndXJhdGlvbj4="
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
  core_site_xml = var.core_site_xml

  #Optional
  description = var.description
  freeform_tags = var.freeform_tags
}

data "oci_golden_gate_connection" "fetched_connection" {
  connection_id = oci_golden_gate_connection.test_connection.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.fetched_connection.display_name
}
