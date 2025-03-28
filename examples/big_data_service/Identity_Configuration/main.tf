variable "bds_instance_id" {
}

variable "cluster_admin_password" {
}

variable "confidential_application_id" {
}

variable "display_name" {
    default = "identityDomainConfig"
}

variable "identity_domain_id" {
}

variable "activate_iam_user_sync_configuration_trigger" {
    default = "false"
}

variable "activate_upst_configuration_trigger" {
    default = "false"
}

variable "refresh_confidential_application_trigger" {
    default = "false"
}

variable "refresh_upst_token_exchange_keytab_trigger" {
    default = "false"
}

variable "is_posix_attributes_addition_required" {
default = "false"
}

variable "confidential_application_id" {
}

variable "vault_id" {
}

resource "oci_bds_bds_instance_identity_configuration" "test_bds_instance_identity_configuration" {
bds_instance_id = var.bds_instance_id
cluster_admin_password = var.cluster_admin_password
confidential_application_id = var.confidential_application_id
display_name = var.display_name
identity_domain_id = var.identity_domain_id
activate_iam_user_sync_configuration_trigger = var.activate_iam_user_sync_configuration_trigger
activate_upst_configuration_trigger = var.activate_upst_configuration_trigger
refresh_confidential_application_trigger = var.refresh_confidential_application_trigger
refresh_upst_token_exchange_keytab_trigger = var.refresh_upst_token_exchange_keytab_trigger

iam_user_sync_configuration_details {
    is_posix_attributes_addition_required = var.is_posix_attributes_addition_required
}

upst_configuration_details {
    master_encryption_key_id = var.confidential_application_id
    vault_id = var.vault_id
}

}
