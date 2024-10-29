resource "oci_generative_ai_agent_knowledge_base" "test_service_managed_knowledge_base" {
  #Required
  compartment_id                 = var.compartment_ocid
  index_config  {
    index_config_type = var.index_config_type_default
    should_enable_hybrid_search   = var.should_enable_hybrid_search
  }

  #Optional
  display_name                  = var.test_knowledge_base_display_name
  description                   = var.test_knowledge_base_description
  #defined_tags not tested - cannot test in home region        
  freeform_tags                 = var.test_freeform_tags
}

data "oci_generative_ai_agent_knowledge_base" "test_service_managed_knowledge_base" {
  #Required
  knowledge_base_id                   = oci_generative_ai_agent_knowledge_base.test_service_managed_knowledge_base.id
}

data "oci_generative_ai_agent_knowledge_bases" "test_knowledge_bases" {
  #Required
  compartment_id                = var.compartment_ocid
}
