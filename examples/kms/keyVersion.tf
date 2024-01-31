resource "oci_kms_key_version" "test_key_version" {
  #Required
  key_id = oci_kms_key.test_key.id
  management_endpoint = data.oci_kms_vault.test_vault.management_endpoint
  schedule_deletion_days = var.schedule_deletion_days
}