
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "target_database_display_name" {
  default = "displayName"
}

variable "target_database_description" {
  default = "description"
}

variable "datasafe_private_endpoint_id" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "autonomous_db_id" {
}
resource "oci_data_safe_target_database" "test_adb_target_database" {
  compartment_id = var.compartment_ocid
  display_name   = var.target_database_display_name
  description    = var.target_database_description
  database_details {
    database_type          = "AUTONOMOUS_DATABASE"
    infrastructure_type    = "ORACLE_CLOUD"
    autonomous_database_id = var.autonomous_db_id
  }
}

data "oci_data_safe_target_database" "test_adb_target_database" {
  target_database_id = oci_data_safe_target_database.test_adb_target_database.id
}

data "oci_data_safe_target_databases" "test_adb_target_databases" {
  compartment_id = var.compartment_ocid
}