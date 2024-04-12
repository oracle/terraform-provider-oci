resource "oci_database_autonomous_database_software_image" "autonomous_database_software_image" {
  compartment_id = var.compartment_ocid
  display_name = "ADSI-TFTest"
  image_shape_family = "EXACC_SHAPE"
  source_cdb_id = oci_database_autonomous_container_database.autonomous_container_database.id
  freeform_tags = {
      "Department" = "Finance"
    }
}