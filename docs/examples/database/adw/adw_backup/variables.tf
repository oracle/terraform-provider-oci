variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

# The Autonomous Data Warehouse needs to be pre-configured for manual backups. Refer:
# https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/adwbackingup.htm
variable "autonomous_data_warehouse_id" {
  default = ""
}

variable "autonomous_data_warehouse_backup_display_name" {
  default = "Monthly Backup"
}
