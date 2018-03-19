resource "oci_file_storage_file_system" "my_fs_1" {
  #Required
  availability_domain = "${var.availability_domain}"
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.file_system_1_display_name}"
}

resource "oci_file_storage_file_system" "my_fs_2" {
  #Required
  availability_domain = "${var.availability_domain}"
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.file_system_2_display_name}"
}