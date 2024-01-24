// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_export" "my_export_fs1_mt1" {
  #Required
  export_set_id  = oci_file_storage_export_set.my_export_set_1.id
  file_system_id = oci_file_storage_file_system.my_fs_1.id
  path           = var.export_path_fs1_mt1

  export_options {
    source                         = var.export_read_write_access_source
    access                         = "READ_WRITE"
    identity_squash                = "NONE"
    require_privileged_source_port = true
  }

  export_options {
    source                         = var.export_read_only_access_source
    access                         = "READ_ONLY"
    identity_squash                = "ALL"
    require_privileged_source_port = true
  }
}

resource "oci_file_storage_export" "my_export_fs1_mt2" {
  #Required
  export_set_id  = oci_file_storage_export_set.my_export_set_2.id
  file_system_id = oci_file_storage_file_system.my_fs_1.id
  path           = var.export_path_fs1_mt2
}

resource "oci_file_storage_export" "my_export_fs2_mt1" {
  #Required
  export_set_id  = oci_file_storage_export_set.my_export_set_1.id
  file_system_id = oci_file_storage_file_system.my_fs_2.id
  path           = var.export_path_fs2_mt1
}

resource "oci_file_storage_export" "my_krb_export_krbfs_krbmt" {
  #Required
  export_set_id  = oci_file_storage_export_set.my_krb_export_set.id
  file_system_id = oci_file_storage_file_system.my_krb_file_system.id
  path           = var.export_path_kfs_kmt

  #Optional
  export_options {
    #Required
    source = var.export_read_write_access_source
    #Optional
    access                           = "READ_WRITE"
    allowed_auth                     = var.krb_export_export_options_allowed_auth
    # anonymous_gid                  = var.export_export_options_anonymous_gid
    # anonymous_uid                  = var.export_export_options_anonymous_uid
    # identity_squash                = var.export_export_options_identity_squash
    is_anonymous_access_allowed      = var.krb_export_export_options_is_anonymous_access_allowed
    # require_privileged_source_port = var.export_export_options_require_privileged_source_port
  }
  is_idmap_groups_for_sys_auth = var.krb_export_is_idmap_groups_for_sys_auth
}

