variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "autonomous_data_warehouse_cpu_core_count" {
  default = 1
}

variable "autonomous_data_warehouse_data_storage_size_in_tbs" {
  default = 1
}

variable "autonomous_data_warehouse_db_name" {
  default = "adwdb1"
}

variable "autonomous_data_warehouse_display_name" {
  default = "example_autonomous_data_warehouse"
}

variable "autonomous_data_warehouse_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_data_warehouse_state" {
  default = "AVAILABLE"
}

variable "autonomous_data_warehouse_backup_display_name" {
  default = "Monthly Backup"
}

variable "autonomous_data_warehouse_backup_state" {
  default = "AVAILABLE"
}
