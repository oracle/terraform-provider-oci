variable "secondary_idcs_account_name_url" {
  default = "secondary_idcs_account_name_url"
}

variable "secondary_private_endpoint_id" {
  default = "secondary_private_endpoint_id"
}

variable "replica_region" {
  default = "us-dcc-phoenix-4"
}

variable vault_id {
  default  = "vault_id"
}

variable virtual_vault_id {
  default = "virtual_vault_id"
}

variable "vault_type" {
  default = "EXTERNAL"
}

resource "oci_kms_vault_replication" "test_replication" {
  replica_region = var.replica_region
  vault_id = var.virtual_vault_id
}