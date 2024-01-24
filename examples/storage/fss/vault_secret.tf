// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_kms_vault" "krb_test_vault" {
  compartment_id = var.compartment_ocid
  display_name = var.krb_vault_display_name
  vault_type   = var.krb_vault_type
}

resource "oci_kms_key" "krb_test_key" {
  #Required
  compartment_id      = var.compartment_ocid
  display_name        = var.krb_key_display_name
  management_endpoint = oci_kms_vault.krb_test_vault.management_endpoint
  key_shape {
    #Required
    algorithm = var.krb_key_shape_algorithm
    length    = var.krb_key_shape_length
  }
}

resource "random_string" "random_keytab_name" {
  length  = 10
  special = false
  upper   = false
  keepers = {
    kms_key_id = oci_kms_key.krb_test_key.id
  }
}

resource "random_string" "random_ldap_pwd_name" {
  length  = 10
  special = false
  upper   = false
  keepers = {
    kms_key_id = oci_kms_key.krb_test_key.id
  }
}

resource "oci_vault_secret" "krb_keytab_secret" {
  #Required
  compartment_id = var.compartment_ocid
  secret_content {
    #Required
    content_type = "BASE64"
    #Optional
    content = var.krb_keytab_content
    stage   = "CURRENT"
  }
  key_id = oci_kms_key.krb_test_key.id
  secret_name = "my_keytab_${random_string.random_keytab_name.result}"
  vault_id    = oci_kms_vault.krb_test_vault.id
}

resource "oci_vault_secret" "krb_ldap_pwd_secret" {
  #Required
  compartment_id = var.compartment_ocid
  secret_content {
    #Required
    content_type = "BASE64"
    #Optional
    content = var.krb_ldap_pwd_content
    stage   = "CURRENT"
  }
  key_id = oci_kms_key.krb_test_key.id
  secret_name = "my_ldap_pwd_${random_string.random_ldap_pwd_name.result}"
  vault_id    = oci_kms_vault.krb_test_vault.id
}