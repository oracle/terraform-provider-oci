variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_devops_repository_file_line" "test_repo_file_lines" {
  #Required
  repository_id = "repository_id"
  revision      = "main"
  file_path     = "testfile"

  #Optional
  start_line_number = 10
}
