resource "oci_file_storage_export" "my_export_fs1_mt1" {
  #Required
  export_set_id  = "${oci_file_storage_export_set.my_export_set_1.id}"
  file_system_id = "${oci_file_storage_file_system.my_fs_1.id}"
  path           = "${var.export_path_fs1_mt1}"

  export_options = [
    {
      source                         = "${var.export_read_write_access_source}"
      access                         = "READ_WRITE"
      identity_squash                = "NONE"
      require_privileged_source_port = true
    },
    {
      source                         = "${var.export_read_only_access_source}"
      access                         = "READ_ONLY"
      identity_squash                = "ALL"
      require_privileged_source_port = true
    },
  ]
}

resource "oci_file_storage_export" "my_export_fs1_mt2" {
  #Required
  export_set_id  = "${oci_file_storage_export_set.my_export_set_2.id}"
  file_system_id = "${oci_file_storage_file_system.my_fs_1.id}"
  path           = "${var.export_path_fs1_mt2}"
}

resource "oci_file_storage_export" "my_export_fs2_mt1" {
  #Required
  export_set_id  = "${oci_file_storage_export_set.my_export_set_1.id}"
  file_system_id = "${oci_file_storage_file_system.my_fs_2.id}"
  path           = "${var.export_path_fs2_mt1}"
}
