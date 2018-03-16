variable "compartment_ocid" {}

variable "label_prefix" {}
variable "ad_count" {}
variable "tenancy_ocid" {}

variable "ad_deployment" {
  default = "0"
}

variable "sql_db_size" {}
variable "sql_backup_size" {}
variable "sql_log_size" {}
variable "witness_volume_size" {}
