resource "oci_generative_ai_agent_agent_endpoint" "test_agent_endpoint" {
  #Required
  compartment_id                 = var.compartment_ocid
  agent_id                       = var.test_agent_id

  #Optional
  display_name                  = var.test_agent_endpoint_display_name
  description                   = var.test_agent_endpoint_description
  #defined_tags not tested - cannot test in home region
  freeform_tags                 = var.test_freeform_tags
  should_enable_citation        = var.should_enable_citation
  should_enable_session         = var.should_enable_session
  should_enable_trace           = var.should_enable_trace
  should_enable_multi_language  = var.agent_endpoint_should_enable_multi_language
  content_moderation_config  {
    should_enable_on_input = var.should_enable_on_input
    should_enable_on_output = var.should_enable_on_output
  }
  guardrail_config {

    #Optional
    content_moderation_config {

      #Optional
      input_guardrail_mode  = var.agent_endpoint_guardrail_config_content_moderation_config_input_guardrail_mode
      output_guardrail_mode = var.agent_endpoint_guardrail_config_content_moderation_config_output_guardrail_mode
    }
    personally_identifiable_information_config {

      #Optional
      input_guardrail_mode  = var.agent_endpoint_guardrail_config_personally_identifiable_information_config_input_guardrail_mode
      output_guardrail_mode = var.agent_endpoint_guardrail_config_personally_identifiable_information_config_output_guardrail_mode
    }
    prompt_injection_config {

      #Optional
      input_guardrail_mode = var.agent_endpoint_guardrail_config_prompt_injection_config_input_guardrail_mode
    }
  }
  human_input_config {
    #Required
    should_enable_human_input = var.agent_endpoint_human_input_config_should_enable_human_input
  }
  output_config {
    #Required
    output_location {
      #Required
      bucket               = var.agent_endpoint_output_config_output_location_bucket
      namespace            = var.agent_endpoint_output_config_output_location_namespace
      output_location_type = var.agent_endpoint_output_config_output_location_output_location_type

      #Optional
      prefix = var.agent_endpoint_output_config_output_location_prefix
    }

    #Optional
    retention_period_in_minutes = var.agent_endpoint_output_config_retention_period_in_minutes
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
