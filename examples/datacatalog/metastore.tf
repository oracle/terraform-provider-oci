resource "oci_datacatalog_metastore" "test_metastore" {
  #Required
  compartment_id = var.compartment_id
  default_external_table_location = var.metastore_default_external_table_location
  default_managed_table_location = var.metastore_default_managed_table_location
  lifecycle {
    ignore_changes = [
      system_tags,
    ]
  }
}

data "oci_datacatalog_metastores" "test_metastores" {
  #Required
  compartment_id = var.compartment_id
}