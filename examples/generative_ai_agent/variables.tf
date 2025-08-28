// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "agent_display_name" {
  default = "agent"
}

variable "agent_description" {
  default = "this is an agent"
}

variable "agent_welcome_message" {
  default = "welcomeMessage"
}

variable "should_enable_citation" {
  default = "false"
}

variable "should_enable_session" {
  default = "true"
}

variable "should_enable_trace" {
  default = "false"
}

variable "should_enable_on_input" {
  default = "false"
}

variable "should_enable_on_output" {
  default = "false"
}

variable "idle_timeout_in_seconds" {
  default = 3600
}

variable "index_config_type_default" {
  default = "DEFAULT_INDEX_CONFIG"
}

variable "should_enable_hybrid_search" {
  default = "false"
}

variable "test_agent_endpoint_display_name" {
  default = "test_agent_endpoint"
}

variable "test_agent_endpoint_description" {
  default = "test agent endpoint"
}

variable "test_agent_tool_display_name" {
  default = "test_agent_tool"
}

variable "test_agent_tool_description" {
  default = "test agent tool"
}

variable "test_knowledge_base_display_name" {
  default = "test_knowledge_base"
}

variable "test_knowledge_base_description" {
  default = "test knowledgeBase"
}

variable "test_data_source_display_name" {
  default = "test_data_source"
}

variable "test_data_source_description" {
  default = "test dataSource"
}

variable "test_data_ingestion_job_display_name" {
  default = "test_data_ingestion_job"
}

variable "test_data_ingestion_job_description" {
  default = "test dataIngestionJob"
}

variable "test_data_source_prefix" {
  default = "sample text.pdf"
}

variable "test_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "should_enable_multi_modality"{
  default = "true"
}

variable "test_agent_id" {}

variable "test_knowledge_base_id" {}

variable "test_bucket" {}

variable "test_namespace" {}

variable "test_prefix" {}

variable "test_subnet_id" {}

variable "test_agent_endpoint_id" {}

variable "agent_endpoint_guardrail_config_content_moderation_config_input_guardrail_mode" {
  default = "DISABLE"
}

variable "agent_endpoint_guardrail_config_content_moderation_config_output_guardrail_mode" {
  default = "DISABLE"
}

variable "agent_endpoint_guardrail_config_personally_identifiable_information_config_input_guardrail_mode" {
  default = "DISABLE"
}

variable "agent_endpoint_guardrail_config_personally_identifiable_information_config_output_guardrail_mode" {
  default = "DISABLE"
}

variable "agent_endpoint_guardrail_config_prompt_injection_config_input_guardrail_mode" {
  default = "DISABLE"
}

variable "agent_endpoint_human_input_config_should_enable_human_input" {
  default = false
}

variable "agent_endpoint_metadata" {
  default = "metadata"
}

variable "agent_endpoint_output_config_output_location_bucket" {
  default = "bucket"
}

variable "agent_endpoint_output_config_output_location_namespace" {
  default = "namespace"
}

variable "agent_endpoint_output_config_output_location_output_location_type" {
  default = "OBJECT_STORAGE_PREFIX"
}

variable "agent_endpoint_output_config_output_location_prefix" {
  default = "prefix"
}

variable "agent_endpoint_output_config_retention_period_in_minutes" {
  default = 10
}

variable "agent_endpoint_should_enable_multi_language" {
  default = false
}

variable "agent_llm_config_routing_llm_customization_instruction" {
  default = "instruction"
}