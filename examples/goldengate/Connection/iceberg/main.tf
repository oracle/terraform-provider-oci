variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}
variable "iceberg_service_account_key_file_secret_id" {}

variable "does_use_secret_ids" {
  default = true
}

variable "description" {
  default = "Created as example for TERSI-7228 Connections R7"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "ICEBERG"
}
variable "display_name" {
  default = "Iceberg_TerraformTest"
}
variable "technology_type" {
  default = "APACHE_ICEBERG"
}

variable "catalog_type" {
  default = "HADOOP"
}

variable "storage_type" {
  default = "GOOGLE_CLOUD_STORAGE"
}

variable "bucket" {
  default = "bucket"
}

variable "project_id" {
  default = "projectId"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}
resource "oci_golden_gate_connection" "test_connection" {
  #Required
  compartment_id  = var.compartment_id
  connection_type = var.connection_type
  display_name    = var.display_name

  #Required for Iceberg connection_type
  technology_type = var.technology_type

  #Optional
  description         = var.description
  freeform_tags       = var.freeform_tags
  does_use_secret_ids = var.does_use_secret_ids

  catalog {
    catalog_type = var.catalog_type
  }

  storage {
    storage_type                       = var.storage_type
    bucket                             = var.bucket
    project_id                         = var.project_id
    service_account_key_file_secret_id = var.iceberg_service_account_key_file_secret_id
  }
  lifecycle {
    ignore_changes = [defined_tags, freeform_tags]
  }
}

# Retrieve the created Iceberg connection.
data "oci_golden_gate_connection" "fetched_connection" {
  connection_id = oci_golden_gate_connection.test_connection.id
}
