variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_policy_ocid" {}

variable "description" {
  default = "description"
}

variable "display_name" {
  default = "security_policy_config_updated"
}

variable "exclude_datasafe_user" {
  default = "ENABLED"
}

variable "security_policy_config_access_level" {
  default = "ACCESSIBLE"
}

variable "security_policy_config_compartment_id_in_subtree" {
  default = false
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_security_policy_config" "test_security_policy_config" {
  #Required
  compartment_id = var.compartment_ocid
  security_policy_id = var.security_policy_ocid
  unified_audit_policy_config {
    exclude_datasafe_user = var.exclude_datasafe_user
  }

  #Optional
  description = var.description
  display_name = var.display_name
}

data "oci_data_safe_security_policy_configs" "test_security_policy_configs" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  security_policy_config_id     = oci_data_safe_security_policy_config.test_security_policy_config.id
  access_level                  = var.security_policy_config_access_level
  compartment_id_in_subtree     = var.security_policy_config_compartment_id_in_subtree
}