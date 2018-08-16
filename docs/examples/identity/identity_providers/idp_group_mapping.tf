resource "oci_identity_idp_group_mapping" "test_idp_group_mapping" {
  #Required
  group_id             = "${oci_identity_group.group1.id}"
  idp_group_name       = "${var.idp_group_mapping_idp_group_name}"
  identity_provider_id = "${oci_identity_identity_provider.test_identity_provider.id}"
}

data "oci_identity_idp_group_mappings" "test_idp_group_mappings" {
  #Required
  identity_provider_id = "${oci_identity_identity_provider.test_identity_provider.id}"
}

output "idp_group_mappings" {
  value = "${data.oci_identity_idp_group_mappings.test_idp_group_mappings.idp_group_mappings}"
}
