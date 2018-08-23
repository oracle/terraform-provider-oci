resource "oci_file_storage_file_system" "my_fs_1" {
  #Required
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.file_system_1_display_name}"
}

resource "oci_file_storage_file_system" "my_fs_2" {
  #Required
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.file_system_2_display_name}"
}
