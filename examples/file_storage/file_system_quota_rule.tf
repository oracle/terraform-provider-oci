resource "oci_file_storage_file_system" "my_fs_simple_quota_rule" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_simple_qr_display_name
  are_quota_rules_enabled = var.quota_rule_enabled

}

resource "oci_file_storage_file_system_quota_rule" "my_simple_quota_rule" {
  #Required
  file_system_id           = oci_file_storage_file_system.my_fs_simple_quota_rule.id
  is_hard_quota            = var.my_simple_quota_rule_is_hard_quota
  principal_type           = var.my_simple_quota_rule_principal_type
  quota_limit_in_gigabytes = var.my_simple_quota_rule_quota_limit_in_gigabytes

  #Optional
  display_name = var.my_simple_quota_rule_display_name
}