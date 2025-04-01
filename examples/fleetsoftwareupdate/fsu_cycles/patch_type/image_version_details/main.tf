// path: fleetsoftwareupdate/fsu_cycles/patch_type/image_version_details

// OCI Provider configuration
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// Fsu Collection resource for database version 19
resource "oci_fleet_software_update_fsu_collection" "test_fsu_collection_db19" {
  display_name   = "tf-test-db19-collection"
  compartment_id = var.compartment_id
  fleet_discovery {
    strategy = "TARGET_LIST"
    targets  = [var.fsu_db_19_target_1]
  }
  service_type         = "EXACS"
  source_major_version = "DB_19"
  type                 = "DB"
}

// Fsu Cycle resource for database version 19 with software image details
resource "oci_fleet_software_update_fsu_cycle" "test_fsu_cycle_db19_software_image" {
  display_name      = "tf-test-db19-cycle"
  compartment_id    = var.compartment_id
  fsu_collection_id = oci_fleet_software_update_fsu_collection.test_fsu_collection_db19.id
  type = "PATCH"
  software_image_details {
    type = "IMAGE_ID"
    software_image_id = var.fsu_db_19_software_image_1
  }
}
