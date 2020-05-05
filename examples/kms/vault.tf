resource "oci_kms_vault" "private-vault-kms" {
  //If restoring using a pre-authenticated-uri use the config below

  /*restore_from_object_store {
    bucket      = "${data.oci_objectstorage_object.vault_backup_object.bucket}"
    destination = "${var.destination.bucket}"
    namespace   = "${data.oci_objectstorage_object.vault_backup_object.namespace}"
    object      = "${data.oci_objectstorage_object.vault_backup_object.object}"
  }*/

  //If restoring using a pre-authenticated-uri use the config below

  /*restore_from_object_store {
    destination = "${var.destination[1]}"
    uri         = "${data.oci_objectstorage_preauthrequest.vault_backup_preauthenticated_request}"
  }*/

  compartment_id  = "${var.compartment_id}"
  display_name    = "${var.vault_display_name}"
  vault_type      = "${var.vault_type[1]}"
  restore_trigger = "${var.vault_restore_trigger}"
}
