variable "tenancy_ocid" {}
variable "region" {}
variable "kms_vault_ocid" {}
variable "kms_key_ocid" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  region           = var.region
  // .. auth options
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
    name    = "XRRSecretSample1"
    stage   = "CURRENT"
  }
  replication_config {
    replication_targets {
      target_key_id   = var.kms_key_ocid
      target_region   = "us-phoenix-1"
      target_vault_id = var.kms_vault_ocid
    }
    #Optional
    is_write_forward_enabled = false
  }

  key_id = var.kms_key_ocid
  secret_name = "XRRSecretSample1105"
  vault_id    = var.kms_vault_ocid
}

resource "oci_vault_secret" "test_secret_wf_enabled" {
  #Required
  compartment_id = var.compartment_ocid
  secret_content {
    #Required
    content_type = "BASE64"

    #Optional
    content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    name    = "XRRSecretSample102"
    stage   = "CURRENT"
  }
  replication_config {
    replication_targets {
      target_key_id   = var.kms_key_ocid
      target_region   = "us-phoenix-1"
      target_vault_id = var.kms_vault_ocid
    }
    #Optional
    is_write_forward_enabled = true
  }

  key_id = var.kms_key_ocid
  secret_name = "XRRSecretSample1104"
  vault_id    = var.kms_vault_ocid
}

resource "oci_vault_secret" "test_secret_multiple_replication_targets" {
  #Required
  compartment_id = var.compartment_ocid
  secret_content {
    #Required
    content_type = "BASE64"

    #Optional
    content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    name    = "XRRSecretSample1"
    stage   = "CURRENT"
  }
  replication_config {
    replication_targets {
      target_key_id   = var.kms_key_ocid
      target_region   = "us-phoenix-1"
      target_vault_id = var.kms_vault_ocid
    }
    replication_targets {
      target_key_id   = var.kms_key_ocid
      target_region   = "us-sanjose-1"
      target_vault_id = var.kms_vault_ocid
    }
    #Optional
    is_write_forward_enabled = true
  }

  key_id = var.kms_key_ocid
  secret_name = "XRRSecretSample1103"
  vault_id    = var.kms_vault_ocid
}

data "oci_vault_secret" "test_secret" {
  secret_id = oci_vault_secret.test_secret.id
}

data "oci_vault_secret" "test_secret_wf_enabled" {
  secret_id = oci_vault_secret.test_secret_wf_enabled.id
}

data "oci_vault_secret" "test_secret_multiple_replication_targets" {
  secret_id = oci_vault_secret.test_secret_multiple_replication_targets.id
}

data "oci_vault_secrets" "test_secret_xrr" {
  compartment_id = var.compartment_ocid
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

output "all_vault_secrets_data" {
  value = data.oci_vault_secrets.test_secret_xrr
}

output "all_vault_secrets_data_for_xrr_secret" {
  value = data.oci_vault_secret.test_secret
}

output "all_vault_secrets_data_for_xrr_secret_wf_enabled" {
  value = data.oci_vault_secret.test_secret_wf_enabled
}

output "all_vault_secrets_data_for_xrr_secret_multiple_replication_targets" {
  value = data.oci_vault_secret.test_secret_multiple_replication_targets
}