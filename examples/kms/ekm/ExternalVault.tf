variable "vault_display_name" {
  default = "Tf-ext-vault-ex"
}

variable "vault_type" {
  default = "EXTERNAL"
}

variable "external_vault_endpoint_url" {
  default = "https://10.0.0.31/api/v1/cckm/oci/ekm/v1/vaults/af872d6e-52f2-4c6b-9694-5b4821d1b5b6"
}

variable "client_app_id" {
  default = "3977f2b65fca4c569f31142959867127"
}

variable "client_app_secret" {
  default = "d82452e5-f5e3-4363-b7a9-0d74052d1236"
}

variable "idcs_account_name_url" {
  default = "https://idcs-87920edcd339458790351b0e4d415385.identity.oraclecloud.com"
}

variable "private_endpoint_id" {
}

resource "oci_kms_vault" "external-vault-kms" {
  compartment_id = var.compartment_ocid
  display_name = var.vault_display_name
  vault_type   = var.vault_type

  external_key_manager_metadata {
    external_vault_endpoint_url = var.external_vault_endpoint_url
    private_endpoint_id = var.private_endpoint_id
    oauth_metadata {
      client_app_id = var.client_app_id
      client_app_secret = var.client_app_secret
      idcs_account_name_url = var.idcs_account_name_url
    }
  }
}