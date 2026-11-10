variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "warehouse_name" {
  default = "terraform-oadc-warehouse"
}

variable "connection_type" {
  default = "ORACLE_AI_DATA_CATALOG_STORAGE"
}

variable "technology_type" {
  default = "ORACLE_AI_DATA_CATALOG_STORAGE"
}

variable "does_use_secret_ids" {
  default = true
}

variable "freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "display_name" {
  default = "OracleAiDataCatalogStorage_TerraformTest"
}

variable "warehouse_storage_access_key_id" {
  default = "warehouse_storage_access_key_id"
}

variable "warehouse_storage_storage_type" {
  default = "OCI_OBJECT_STORAGE_ACCESS_KEY"
}

variable "warehouse_storage_region" {
  default = "us-ashburn-1"
}

variable "warehouse_storage_namespace" {
  type        = string
  description = "Object Storage namespace containing the warehouse bucket."
  default     = "axm8ybx0pwmf"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "iceberg_service_account_key_file_secret_id" {}

resource "oci_golden_gate_connection" "test_connection" {
  # Required
  compartment_id  = var.compartment_id
  connection_type = var.connection_type
  technology_type = var.technology_type
  display_name    = var.display_name

  # Optional
  does_use_secret_ids = var.does_use_secret_ids
  freeform_tags       = var.freeform_tags

  # Required
  warehouse_name = var.warehouse_name

  # Required
  warehouse_storage {
    # Required
    storage_type = var.warehouse_storage_storage_type

    # Required when storage_type is OCI_OBJECT_STORAGE_ACCESS_KEY
    namespace                   = var.warehouse_storage_namespace
    access_key_id               = var.warehouse_storage_access_key_id
    secret_access_key_secret_id = var.iceberg_service_account_key_file_secret_id

    # Optional
    region = var.warehouse_storage_region
  }

  lifecycle {
    ignore_changes = [defined_tags, freeform_tags]
  }
}
