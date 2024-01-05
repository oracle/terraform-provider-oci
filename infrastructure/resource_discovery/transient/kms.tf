// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_kms_key" "key_rd" {
  #Required
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "KmsRd"
  management_endpoint = "${data.oci_kms_vault.vault_rd.management_endpoint}"

  key_shape {
    #Required
    algorithm = "AES"
    length    = "32"
  }
}

resource "oci_kms_vault" "private-vault-kms-rd" {
  compartment_id = "${var.compartment_ocid}"

  display_name = "VaultRd"
  vault_type   = "DEFAULT"
}

data "oci_kms_vault" "vault_rd" {
  #Required
  vault_id = "${oci_kms_vault.private-vault-kms-rd.id}"
}
