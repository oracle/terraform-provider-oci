resource "oci_kms_vault" "private-vault-kms" {
  //If restoring using a pre-authenticated-uri use the config below

  /*restore_from_object_store {
    bucket      = data.oci_objectstorage_object.vault_backup_object.bucket
    destination = var.destination[0]
    namespace   = data.oci_objectstorage_object.vault_backup_object.namespace
    object      = data.oci_objectstorage_object.vault_backup_object.object
  }*/

  //If restoring using a pre-authenticated-uri use the config below

  /*restore_from_object_store {
          destination = var.destination[1]
          uri         = data.oci_objectstorage_preauthrequest.vault_backup_preauthenticated_request
        }*/

  //If restoring using object in object storage use the config below

  /*restore_from_file {
    restore_vault_from_file_details = data.oci_objectstorage_object.vault_backup_object.content
    content_length                = data.oci_objectstorage_object.vault_backup_object.content_length

    //Optional
    content_md5                   = data.oci_objectstorage_object.vault_backup_object.content_md5
  }*/

  compartment_id = var.compartment_id

  display_name = var.vault_display_name
  vault_type   = var.vault_type[0]
  //Flip the trigger when restore operation on vault needs to be performed
  //restore_trigger = var.vault_restore_trigger
}

// Use oci_kms_vault_replication to create, update and delete a replica of a vault.
// Currently only support virtual private vault.

/*resource "oci_kms_vault_replication" "test_replica" {
  # Required
  vault_id = data.oci_kms_vault.test_vault.id
  replica_region = var.destination_region
}*/

