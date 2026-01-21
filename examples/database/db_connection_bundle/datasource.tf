provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_database_db_connection_bundles" "db_connection_bundles" {
  compartment_id            = var.compartment_ocid
  associated_resource_id    = var.associated_resource_ocid
  db_connection_bundle_type = "TLS"
  state                     = "ACTIVE"
}

data "oci_database_db_connection_bundle" "db_connection_bundle" {
  count                   = var.db_connection_bundle_ocid != "" ? 1 : 0
  db_connection_bundle_id = var.db_connection_bundle_ocid
}