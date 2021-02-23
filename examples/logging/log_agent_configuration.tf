// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_id" {}

variable "unified_agent_configuration_defined_tags_value" {
  default = "value"
}

variable "unified_agent_configuration_description" {
  default = "description2"
}

variable "unified_agent_configuration_display_name" {
  default = "displayName2"
}

variable "unified_agent_configuration_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "unified_agent_configuration_group_association_group_list" {
  default = [""]
}

variable "unified_agent_configuration_is_compartment_id_in_subtree" {
  default = false
}

variable "unified_agent_configuration_is_enabled" {
  default = true
}

variable "unified_agent_configuration_service_configuration_configuration_type" {
  default = "LOGGING"
}

variable "unified_agent_configuration_service_configuration_sources_channels" {
  default = ["Security"]
}

variable "unified_agent_configuration_service_configuration_sources_name" {
  default = "name"
}

variable "unified_agent_configuration_service_configuration_sources_parser_delimiter" {
  default = "delimiter"
}

variable "unified_agent_configuration_service_configuration_sources_parser_expression" {
  default = "expression"
}

variable "unified_agent_configuration_service_configuration_sources_parser_field_time_key" {
  default = "fieldTimeKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_format" {
  default = []
}

variable "unified_agent_configuration_service_configuration_sources_parser_format_firstline" {
  default = "formatFirstline"
}

variable "unified_agent_configuration_service_configuration_sources_parser_grok_failure_key" {
  default = "grokFailureKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_grok_name_key" {
  default = "grokNameKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_estimate_current_event" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_keep_time_key" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_null_empty_string" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_support_colonless_ident" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_with_priority" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_keys" {
  default = []
}

variable "unified_agent_configuration_service_configuration_sources_parser_message_format" {
  default = "RFC3164"
}

variable "unified_agent_configuration_service_configuration_sources_parser_message_key" {
  default = "messageKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_multi_line_start_regexp" {
  default = "multiLineStartRegexp"
}

variable "unified_agent_configuration_service_configuration_sources_parser_null_value_pattern" {
  default = "nullValuePattern"
}

variable "unified_agent_configuration_service_configuration_sources_parser_parser_type" {
  default = "AUDITD"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_format" {
  default = "fieldTimeFormat"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_key" {
  default = "fieldTimeKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_zone" {
  default = "fieldTimeZone"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_name" {
  default = "name"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_pattern" {
  default = "pattern"
}

variable "unified_agent_configuration_service_configuration_sources_parser_rfc5424time_format" {
  default = "rfc5424TimeFormat"
}

variable "unified_agent_configuration_service_configuration_sources_parser_syslog_parser_type" {
  default = "STRING"
}

variable "unified_agent_configuration_service_configuration_sources_parser_time_format" {
  default = "timeFormat"
}

variable "unified_agent_configuration_service_configuration_sources_parser_time_type" {
  default = "FLOAT"
}

variable "unified_agent_configuration_service_configuration_sources_parser_timeout_in_milliseconds" {
  default = 10
}

variable "unified_agent_configuration_service_configuration_sources_parser_types" {
  default = "types"
}

variable "unified_agent_configuration_service_configuration_sources_paths" {
  default = []
}

variable "unified_agent_configuration_service_configuration_sources_source_type" {
  default = "WINDOWS_EVENT_LOG"
}

variable "unified_agent_configuration_state" {
  default = "AVAILABLE"
}

resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration" {
  #Required
  compartment_id = var.compartment_id
  is_enabled     = var.unified_agent_configuration_is_enabled
  service_configuration {
    #Required
    configuration_type = var.unified_agent_configuration_service_configuration_configuration_type

    #Optional
    destination {
      #Required
      log_object_id = "${oci_logging_log.test_log.id}"
    }
    sources {
      #Required
      source_type = var.unified_agent_configuration_service_configuration_sources_source_type

      #Optional
      channels = var.unified_agent_configuration_service_configuration_sources_channels
      name     = var.unified_agent_configuration_service_configuration_sources_name
    }
  }

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = var.log_group_defined_tags_value
  }
  description   = var.unified_agent_configuration_description
  display_name  = var.unified_agent_configuration_display_name
  freeform_tags = var.unified_agent_configuration_freeform_tags
  group_association {

    #Optional
    group_list = ["${oci_logging_log_group.test_log_group.id}"]
  }
}

data "oci_logging_unified_agent_configurations" "test_unified_agent_configurations" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                 = var.unified_agent_configuration_display_name
  group_id                     = "${oci_logging_log_group.test_log_group.id}"
  is_compartment_id_in_subtree = var.unified_agent_configuration_is_compartment_id_in_subtree
  log_id                       = "${oci_logging_log.test_log.id}"
  state                        = var.unified_agent_configuration_state
}

