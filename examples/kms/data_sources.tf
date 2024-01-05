// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

// Gets the detail of the vault.
data "oci_kms_vault" "test_vault" {
  #Required
  vault_id = var.vault_id
}

/*
//create a new vault
resource "oci_kms_vault" "test_vault" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.vault_display_name
	vault_type = var.vault_vault_type
}
*/

// Gets the list of keys in the compartment and vault.
data "oci_kms_keys" "test_keys" {
  #Required
  compartment_id      = var.compartment_id
  management_endpoint = data.oci_kms_vault.test_vault.management_endpoint

  filter {
    name   = "display_name"
    values = [var.key_display_name]
  }
}

data "oci_core_volumes" "test_volumes" {
  compartment_id = var.compartment_id

  filter {
    name   = "id"
    values = [oci_core_volume.my_volume.id]
  }
}

// Get replication status of a vault.
// Currently only support virtual private vault.
/*data "oci_kms_replication_status" "test_replication_status" {
  # Required
  management_endpoint = data.oci_kms_vault.test_vault.management_endpoint
  replication_id = data.oci_kms_vault.test_vault.replica_details[0].replication_id
}*/

// List replicas of a vault.
// Currently only support virtual private vault.
/*data "oci_kms_vault_replicas" "test_vault_replicas" {
  # Required
  vault_id = data.oci_kms_vault.test_vault.id
}*/

//bucket object details where key was backed up
/*data "oci_objectstorage_object" "key_backup_object" {
  #Required
  bucket    = "bucket-name"
  namespace = "namespace"
  object    = "object"
}*/
//bucket object details where vault was backed up
/*data "oci_objectstorage_object" "vault_backup_object" {
  #Required
  bucket    = "bucket-name"
  namespace = "namespace"
  object    = "object"
}*/
//Pre-authenticated-request details for key backup
/*data "oci_objectstorage_preauthrequest" "key_backup_preauthenticated_request" {
  #Required
  bucket    = "bucket-name"
  namespace = "namespace"
  par_id    = "par_id"
}*/
//Pre-authenticated-request for vault backup
/*data "oci_objectstorage_preauthrequest" "vault_backup_preauthenticated_request" {
  #Required
  bucket    = "bucket-name"
  namespace = "namespace"
  par_id    = "par_id"
}*/
