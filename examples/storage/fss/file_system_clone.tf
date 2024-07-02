resource "oci_file_storage_file_system" "my_fs_clone" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_clone_display_name
  source_snapshot_id = oci_file_storage_snapshot.my_snapshot_clone.id
}
resource "oci_file_storage_file_system" "my_fs_simple" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_simple_display_name
}
resource "oci_file_storage_snapshot" "my_snapshot_clone" {
  #Required
  file_system_id = oci_file_storage_file_system.my_fs_simple.id
  name           = var.snapshot_name_clone
}
resource "oci_file_storage_file_system" "my_fs_clone_with_detach" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_clone_with_detach_display_name
  source_snapshot_id = oci_file_storage_snapshot.my_snapshot_clone_1.id
  clone_attach_status = var.clone_attach_status_value
}
resource "oci_file_storage_file_system" "my_fs_simple_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  display_name = var.file_system_simple_1_display_name
}
resource "oci_file_storage_snapshot" "my_snapshot_clone_1" {
  #Required
  file_system_id = oci_file_storage_file_system.my_fs_simple_1.id
  name           = var.snapshot_name_clone_1
}