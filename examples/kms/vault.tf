resource "oci_kms_vault" "private-vault-kms" {

  compartment_id = var.compartment_id

  display_name = var.vault_display_name
  vault_type   = var.vault_type[0]
}


