resource "oci_generative_ai_agent_data_source" "test_data_source" {
  #Required
  compartment_id                 = var.compartment_ocid
  knowledge_base_id              = var.test_knowledge_base_id
  data_source_config  {
    data_source_config_type      = "OCI_OBJECT_STORAGE"
    object_storage_prefixes {
      bucket = var.test_bucket
      namespace = var.test_namespace
      prefix = var.test_prefix
    }
  }

  #Optional
  display_name                  = var.test_data_source_display_name
  description                   = var.test_data_source_description
  #defined_tags not tested - cannot test in home region
  freeform_tags                 = var.test_freeform_tags
}

data "oci_generative_ai_agent_data_source" "test_data_source" {
  #Required
  data_source_id                   = oci_generative_ai_agent_data_source.test_data_source.id
}

data "oci_generative_ai_agent_data_sources" "test_data_sources" {
  #Required
  compartment_id                = var.compartment_ocid
}
