// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_generative_ai_agent_agent" "test_agent" {
  #Required
  compartment_id                 = var.compartment_ocid

  #Optional
  display_name                  = var.agent_display_name
  description                   = var.agent_description
  welcome_message               = var.agent_welcome_message
  llm_config {
    routing_llm_customization {
          #Optional
      instruction = var.agent_llm_config_routing_llm_customization_instruction
    }
  }
  #defined_tags not tested - cannot test in home region
  freeform_tags                = var.test_freeform_tags
}

data "oci_generative_ai_agent_agent" "test_agent" {
  #Required
  agent_id       = oci_generative_ai_agent_agent.test_agent.id
}

data "oci_generative_ai_agent_agents" "test_agents" {
  #Required
  compartment_id                = var.compartment_ocid
}