resource "oci_database_autonomous_database" "undelete_autonomous_database" {
  compartment_id           = var.compartment_id
  db_name                  = var.undelete_db_name
  source                   = "UNDELETE_ADB"
  source_id                = var.source_id
}