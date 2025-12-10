resource "oci_generative_ai_imported_model" "test_imported_model" {
  #Required
  compartment_id = var.compartment_ocid
  data_source {
    #Optional
    access_token = var.hf_access_token
    model_id     = var.imported_model_model_id
    source_type  = var.imported_model_data_source_source_type
  }

  #Optional
  capabilities  = var.imported_model_capabilities
  description   = var.imported_model_description
  display_name  = var.imported_model_display_name
  freeform_tags = var.imported_model_freeform_tags
  vendor        = var.imported_model_vendor
  version       = var.imported_model_version
}

data "oci_generative_ai_imported_models" "test_imported_models" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  capability   = var.imported_model_capability
}
