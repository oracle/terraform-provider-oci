resource "oci_generative_ai_agent_agent_endpoint" "test_agent_endpoint" {
  #Required
  compartment_id                 = var.compartment_ocid
  agent_id                       = data.oci_generative_ai_agent_agent.test_agent.id

  #Optional
  display_name                  = var.test_agent_endpoint_display_name
  description                   = var.test_agent_endpoint_description
  #defined_tags not tested - cannot test in home region        
  freeform_tags                 = var.test_freeform_tags
  should_enable_citation        = var.should_enable_citation
  should_enable_session         = var.should_enable_session
  should_enable_trace           = var.should_enable_trace
  content_moderation_config  {
    should_enable_on_input = var.should_enable_on_input
    should_enable_on_output = var.should_enable_on_output
  }
  session_config              {
    idle_timeout_in_seconds = var.idle_timeout_in_seconds
  }
}

data "oci_generative_ai_agent_agent_endpoint" "test_agent_endpoint" {
  #Required
  agent_endpoint_id                   = oci_generative_ai_agent_agent_endpoint.test_agent_endpoint.id
}

data "oci_generative_ai_agent_agent_endpoints" "test_agent_endpoints" {
  #Required
  compartment_id                = var.compartment_ocid
}
