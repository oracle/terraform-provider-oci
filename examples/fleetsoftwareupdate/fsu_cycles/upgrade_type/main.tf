// path: fleetsoftwareupdate/fsu_cycles/upgrade_type

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

// Fsu Cycle resource for database version 19 with goal version 23ai
resource "oci_fleet_software_update_fsu_cycle" "test_fsu_cycle_upgrade_23" {
  display_name      = "tf-test-upgrade-cycle"
  compartment_id    = var.compartment_id
  fsu_collection_id = oci_fleet_software_update_fsu_collection.test_fsu_collection_db19.id
  type = "UPGRADE"
  goal_version_details {
		type = "VERSION"
		version = "23.6.0.24.10"
	}
}
