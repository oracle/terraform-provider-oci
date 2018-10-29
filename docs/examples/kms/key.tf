variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "vault_id" {}

variable "key_display_name" {
  default = "Key C"
}

variable "key_key_shape_algorithm" {
  default = "AES"
}

variable "key_key_shape_length" {
  default = 32
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_kms_vault" "test_vault" {
  #Required
  vault_id = "${var.vault_id}"
}

resource "oci_kms_key" "test_key" {
  #Required
  compartment_id      = "${var.compartment_id}"
  display_name        = "${var.key_display_name}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

  key_shape {
    #Required
    algorithm = "${var.key_key_shape_algorithm}"
    length    = "${var.key_key_shape_length}"
  }
}

resource "oci_kms_key_version" "test_key_version" {
  #Required
  key_id              = "${oci_kms_key.test_key.id}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
}

data "oci_kms_keys" "test_keys" {
  #Required
  compartment_id      = "${var.compartment_id}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"

  filter {
    name   = "display_name"
    values = ["${var.key_display_name}"]
  }
}

output "key_id" {
  value = "${oci_kms_key.test_key.id}"
}
