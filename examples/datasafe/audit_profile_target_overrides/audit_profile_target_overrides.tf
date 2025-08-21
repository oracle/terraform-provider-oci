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

data "oci_data_safe_audit_profile_target_overrides" "test_audit_profile_target_overrides" {
  // Required 
  audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id

  // Optional
  display_name     = "displayName"
}