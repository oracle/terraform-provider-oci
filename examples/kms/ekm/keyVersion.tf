resource "oci_kms_key_version" "test_key_version" {
  #Required
  key_id = oci_kms_key.test_key.id
  management_endpoint = "avsnmg6paahhm-management.kms.r1.oracleiaas.com"
  external_key_version_id = var.ext_key_version_id
}