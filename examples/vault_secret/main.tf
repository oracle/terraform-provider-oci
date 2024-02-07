// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "kms_vault_ocid" {
}

variable "kms_key_ocid" {}

variable "adb_ocid" {}

variable "function_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_vault_secrets" "test_secrets" {
  compartment_id = var.compartment_ocid
  state          = "ACTIVE"
  vault_id       = var.kms_vault_ocid
}

resource "oci_vault_secret" "test_secret" {
  #Required
  compartment_id = var.compartment_ocid
  secret_content {
    #Required
    content_type = "BASE64"

    #Optional
    content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    name    = "name"
    stage   = "CURRENT"
  }
  key_id = var.kms_key_ocid
  secret_name = "TFsample1"
  vault_id    = var.kms_vault_ocid
}

resource "oci_vault_secret" "test_secret_with_auto_rotation" {
  compartment_id = var.compartment_ocid
  key_id         = var.kms_key_ocid
  secret_name    = "TFSecretAutoRotation"
  vault_id       = var.kms_vault_ocid

  secret_content {
    #Required
    content_type = "BASE64"

    #Optional
    content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    name    = "name"
    stage   = "CURRENT"
  }

  rotation_config {
    is_scheduled_rotation_enabled = "true"
    rotation_interval = "P90D"
    target_system_details {
      function_id = var.function_ocid
      target_system_type = "FUNCTION"
    }
  }
}

resource "oci_vault_secret" "test_secret_with_target_system" {
  compartment_id = var.compartment_ocid
  key_id         = var.kms_key_ocid
  secret_name    = "TFSecretWithTargetSystem"
  vault_id       = var.kms_vault_ocid

  secret_content {
    #Required
    content_type = "BASE64"

    #Optional
    content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    name    = "name"
    stage   = "CURRENT"
  }

  rotation_config {
    is_scheduled_rotation_enabled = "false"
    target_system_details {
      adb_id = var.adb_ocid
      target_system_type = "ADB"
    }
  }
}


data "oci_vault_secret" "test_secret" {
  secret_id = oci_vault_secret.test_secret.id
}

data "oci_secrets_secretbundle_versions" "test_secretbundle_versions" {
  #Required
  secret_id = oci_vault_secret.test_secret.id
}

// Get Secret content
data "oci_secrets_secretbundle" "test_secretbundles" {
  #Required
  secret_id = oci_vault_secret.test_secret.id
  stage               = "CURRENT"
}