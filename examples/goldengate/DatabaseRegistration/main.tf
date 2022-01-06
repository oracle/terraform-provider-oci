variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "kms_vault_id" {}
variable "kms_key_id" {}
variable "test_subnet_id" {}
variable "test_db_id" {}

variable "database_registration_alias_name" {
  default = "aliasName"
}

variable "database_registration_connection_string" {
  default = "fqdndb.ggs.com:1521/orcl.us.oracle.com"
}

variable "database_registration_session_mode" {
  default = "DIRECT"
}

variable "database_registration_defined_tags_value" {
  default = "value"
}

variable "database_registration_description" {
  default = "description"
}

variable "database_registration_display_name" {
  default = "displayName"
}

variable "database_registration_fqdn" {
  default = "fqdndb.ggs.com"
}

variable "database_registration_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_registration_ip_address" {
  default = "10.0.0.10"
}

variable "database_registration_password" {
  default = "BEstrO0ng_#11"
}

variable "database_registration_state" {
  default = "ACTIVE"
}

variable "database_registration_username" {
  default = "username"
}

variable "database_registration_wallet" {
  default = "wallet"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_database_registration" "test_database_registration" {
  #Required
  alias_name     = var.database_registration_alias_name
  compartment_id = var.compartment_id
  display_name   = var.database_registration_display_name
  fqdn           = var.database_registration_fqdn
  password       = var.database_registration_password
  username       = var.database_registration_username

  #Optional
  connection_string     = var.database_registration_connection_string
  session_mode          = var.database_registration_session_mode
  database_id           = var.test_db_id
  #defined_tags          = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.database_registration_defined_tags_value)
  description           = var.database_registration_description
  freeform_tags         = var.database_registration_freeform_tags
  ip_address            = var.database_registration_ip_address
  key_id                = var.kms_key_id
  secret_compartment_id = var.compartment_id
  subnet_id             = var.test_subnet_id
  vault_id              = var.kms_vault_id
  wallet                = var.database_registration_wallet
}

data "oci_golden_gate_database_registrations" "test_database_registrations" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.database_registration_display_name
  state        = var.database_registration_state
}
