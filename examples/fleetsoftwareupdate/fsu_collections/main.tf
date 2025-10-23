// path: fleetsoftwareupdate/fsu_collections

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

// Fsu Collection resource for database version 23
resource "oci_fleet_software_update_fsu_collection" "test_fsu_collection_db23" {
  display_name   = "tf-test-db23-collection"
  compartment_id = var.compartment_id
  fleet_discovery {
    strategy = "TARGET_LIST"
    targets  = [var.fsu_db_23_target_1]
  }
  service_type         = "EXACS"
  source_major_version = "DB_23"
  type                 = "DB"
}


// Fsu Collection resource for database version 26
resource "oci_fleet_software_update_fsu_collection" "test_fsu_collection_db26" {
  display_name   = "tf-test-db23-collection"
  compartment_id = var.compartment_id
  fleet_discovery {
    strategy = "TARGET_LIST"
    targets  = [var.fsu_db_26_target_1]
  }
  service_type         = "EXACS"
  source_major_version = "DB_26"
  type                 = "DB"
}

