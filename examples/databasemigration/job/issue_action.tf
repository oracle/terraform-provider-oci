variable "job_ocid" {
  default = ""
}

variable "job_id" {
  default = ""
}

variable "migration_id" {
  default = ""
}

variable "tenancy_ocid" {
  default = ""
}
variable "user_ocid" {
  default = ""
}
variable "fingerprint" {
  default = ""
}
variable "private_key_path" {
  default = ""
}
variable "region" {
  default = ""
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  auth             = "SecurityToken"
  region           = var.region

}
data "oci_database_migration_job" "test_job" {
  job_id = var.job_id
}

output "phase_list" {
  value = data.oci_database_migration_job.test_job.progress
}