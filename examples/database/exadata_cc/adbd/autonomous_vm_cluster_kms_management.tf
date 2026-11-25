resource "oci_database_key_store" "thales_key_store" {
  compartment_id = "${var.compartment_ocid}"
  defined_tags = "${tomap({ "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value" })}"
  display_name = "Thales Key Store1"
  freeform_tags = {
    "Department" = "Finance"
  }
  type_details {
    admin_username = "username1"
    connection_ips = ["192.1.1.1"]
    secret_id = "${var.okv_secret}"
    type = "THALES"
    vault_id = "${var.kms_vault_ocid}"
  }
}
resource "oci_database_autonomous_vm_cluster_kms_management" "test_autonomous_vm_cluster_kms_management" {
  autonomous_vm_cluster_id = "${oci_database_autonomous_vm_cluster.autonomous_vm_cluster_2.id}"
  is_exclusive = "true"
  key_management_system = "THALES"
  key_store_id = "${oci_database_key_store.thales_key_store.id}"
  thales_hsm_port = "9000"
}