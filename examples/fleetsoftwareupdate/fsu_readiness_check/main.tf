// path: fleetsoftwareupdate/fsu_readiness_check

// OCI Provider configuration
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// Fsu Readiness Check resource for database version 23
resource "oci_fleet_software_update_fsu_readiness_check" "test_fsu_readiness_check_db23" {
  display_name   = "tf-test-db23-readiness-check"
  compartment_id = var.compartment_id
  type           = "TARGET"

  targets {
    entity_id   = var.fsu_db_23_target_1
    entity_type = "DATABASE"
  }

  freeform_tags = {
    "Environment" = "Test"
    "ManagedBy"   = "Terraform"
  }
}
