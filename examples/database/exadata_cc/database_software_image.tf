resource "oci_database_database_software_image" "test_database_software_image" {
  compartment_id = var.compartment_ocid
  database_version                        = "19.0.0.0"
  display_name                            = "image1"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Finance"
  }

  image_shape_family = "EXACC_SHAPE"
  image_type         = "DATABASE_IMAGE"
  patch_set = "19.26.0.0"
}