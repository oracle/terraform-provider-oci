variable "job_ocid" {}

data "oci_database_migration_job" "test_job" {
  job_id = var.job_ocid
}