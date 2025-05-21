resource "oci_database_key_store" "test_key_store" {
  compartment_id           = var.compartment_ocid
  display_name             = "example-key-store"
  type_details {
    admin_username = "example-username"
    connection_ips = ["192.1.1.1"]
    secret_id      = var.okv_secret
    type           = "ORACLE_KEY_VAULT"
    vault_id       = var.kms_vault_ocid
  }
}