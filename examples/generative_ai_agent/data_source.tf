resource "oci_generative_ai_agent_data_source" "test_data_source" {
  #Required
  compartment_id                 = var.compartment_ocid
  knowledge_base_id              = oci_generative_ai_agent_knowledge_base.test_service_managed_knowledge_base.id
  data_source_config  {
    data_source_config_type = "OCI_OBJECT_STORAGE"
    object_storage_prefixes {
      bucket = data.oci_objectstorage_bucket.bucket.name
      namespace = data.oci_objectstorage_namespace.ns.namespace
      prefix = var.test_data_source_prefix
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

data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_objectstorage_bucket" "bucket" {
  name                         = "oci-docs"
  namespace                    = data.oci_objectstorage_namespace.ns.namespace
}
