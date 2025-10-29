
// Create a simple object storage link
resource "oci_lustre_file_storage_object_storage_link" "test_object_storage_link" {
    availability_domain = data.oci_identity_availability_domain.ad.name
    compartment_id      = var.compartment_ocid
    display_name        = var.object_storage_link_display_name
    file_system_path    = var.file_system_path
	  lustre_file_system_id =  oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
	  object_storage_prefix =  "${data.oci_objectstorage_namespace.ns.namespace}:/${oci_objectstorage_bucket.bucket1.name}/${var.object_storage_prefix}"
    is_overwrite = false

    # To start an export job
    # start_export_to_object_trigger = 1

    # To stop an export job
    # stop_export_to_object_trigger = 1

    # To start an import job
    # start_import_from_object_trigger = 1

    # To stop an import job
    # stop_import_from_object_trigger = 1
}

// Create an object storage link, start an export job and stop it
resource "oci_lustre_file_storage_object_storage_link" "test_object_storage_link_start_stop_export_job" {
    availability_domain = data.oci_identity_availability_domain.ad.name
    compartment_id      = var.compartment_ocid
    display_name        = var.object_storage_link_display_name_export_job
    file_system_path    = var.file_system_path
	  lustre_file_system_id =  oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
	  object_storage_prefix =  "${data.oci_objectstorage_namespace.ns.namespace}:/${oci_objectstorage_bucket.bucket1.name}/${var.object_storage_prefix_export_job}"
    is_overwrite = false

    # To start an export job
    start_export_to_object_trigger = 1

    # To stop an export job
    stop_export_to_object_trigger = 1
}

// Create an object storage link, start an import job and stop it
resource "oci_lustre_file_storage_object_storage_link" "test_object_storage_link_start_stop_import_job" {
    availability_domain = data.oci_identity_availability_domain.ad.name
    compartment_id      = var.compartment_ocid
    display_name        = var.object_storage_link_display_name_import_job
    file_system_path    = var.file_system_path
	  lustre_file_system_id =  oci_lustre_file_storage_lustre_file_system.test_lustre_file_system_import_job.id
	  object_storage_prefix =  "${data.oci_objectstorage_namespace.ns.namespace}:/${oci_objectstorage_bucket.bucket1.name}/${var.object_storage_prefix_import_job}"
    is_overwrite = true

    # To start an import job
    start_import_from_object_trigger = 1

    # To stop an import job
    stop_import_from_object_trigger = 1
}


data "oci_lustre_file_storage_object_storage_link" "object_storage_link" {
  object_storage_link_id = oci_lustre_file_storage_object_storage_link.test_object_storage_link.id
#   Optional
#   availability_domain   = var.object_storage_link_availability_domain
#   compartment_id        = var.compartment_id
#   display_name          = var.object_storage_link_display_name
#   state                 = var.object_storage_link_state
#   lustre_file_system_id =  oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
}

data "oci_lustre_file_storage_object_storage_links" "object_storage_links" {
  compartment_id      = var.compartment_ocid
  #Optional
#   availability_domain   = var.object_storage_link_availability_domain
#   compartment_id        = var.compartment_id
#   display_name          = var.object_storage_link_display_name
#   state                 = var.object_storage_link_state
#   lustre_file_system_id =  oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
}
data "oci_lustre_file_storage_object_storage_link_sync_job" "sync_job" {
  sync_job_id = var.sync_job_id
  object_storage_link_id = oci_lustre_file_storage_object_storage_link.test_object_storage_link.id
  #Optional
#   state                 = var.sync_job_state
}

data "oci_lustre_file_storage_object_storage_link_sync_jobs" "sync_jobs" {
  object_storage_link_id = oci_lustre_file_storage_object_storage_link.test_object_storage_link.id
  #Optional
#   state                 = var.object_storage_link_state
#   lustre_file_system_id =  oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
}
