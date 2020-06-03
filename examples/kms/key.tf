// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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

  //If using bucket to restore a key use below
  /*restore_from_object_store {
    bucket      = "${data.oci_objectstorage_object.key_backup_object.bucket}"
    destination = "${var.destination[0]}"
    namespace   = "${data.oci_objectstorage_object.key_backup_object.namespace}"
    object      = "${data.oci_objectstorage_object.key_backup_object.object}"
  }*/


  //If using Pre-authenticated-request to restore a key use below
  /*restore_from_object_store {
    destination = "${var.destination[1]}"
    uri         = "${data.oci_objectstorage_preauthrequest.key_backup_preauthenticated_request}"
  }*/

  restore_trigger = "${var.key_restore_trigger}"
}

resource "oci_kms_key_version" "test_key_version" {
  #Required
  key_id              = "${oci_kms_key.test_key.id}"
  management_endpoint = "${data.oci_kms_vault.test_vault.management_endpoint}"
}
