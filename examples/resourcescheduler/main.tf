// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// To Test these examole first export these variables
// present here - https://confluence.oraclecorp.com/confluence/display/TERSI/Quick+start#Quickstart-Localenvironmentsetup
// Also export this variable to test body parameter resource test below - rs_body_param_test
// export TF_VAR_function_ocid and resource_id = <any current functions, instance or adbd ocid, if not preset create it in the same region which you exported>

variable "tenancy_ocid" {}
# variable "user_ocid" {}
# variable "fingerprint" {}
# variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "resource_id" {}

variable "schedule_action" {
  default = "START_RESOURCE"
}

variable "schedule_defined_tags_value" {
  default = "value"
}

variable "schedule_description" {
  default = "description"
}

variable "schedule_display_name" {
  default = "displayName"
}

variable "schedule_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "schedule_recurrence_details" {
  default = "FREQ=DAILY;INTERVAL=1"
}

variable "schedule_recurrence_type" {
  default = "ICAL"
}

variable "schedule_resource_filters_attribute" {
  default = "COMPARTMENT_ID"
}

variable "schedule_resource_filters_condition" {
  default = "EQUAL"
}

variable "schedule_resource_filters_should_include_child_compartments" {
  default = false
}

variable "schedule_resource_filters_value_namespace" {
  default = "namespace"
}

variable "schedule_resource_filters_value_tag_key" {
  default = "tagKey"
}

variable "schedule_resource_filters_value_value" {
  default = "value"
}

variable "schedule_resources_id" {
  default = "id"
}

variable "schedule_resources_metadata" {
  default = "metadata"
}

variable "schedule_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  # user_ocid        = var.user_ocid
  # fingerprint      = var.fingerprint
  # private_key_path = var.private_key_path
  region           = var.region
  auth             = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_resource_scheduler_schedule" "test_schedule" {
  #Required
  action             = var.schedule_action
  compartment_id     = var.compartment_id
  recurrence_details = var.schedule_recurrence_details
  recurrence_type    = var.schedule_recurrence_type

  resource_filters {
    # Required
    attribute = "DEFINED_TAGS"
    value {
      namespace="ResourceSchedulerCanary"
      tag_key="ScheduleTagFilterTestKey"
      value="foo"
    }
  }
  resource_filters {
    # Required
    attribute = "LIFECYCLE_STATE"
    value {
      value="running"
    }
    value {
      value="stopped"
    }
  }
  resource_filters {
    # Required
    attribute = "COMPARTMENT_ID"
    value {
      value=var.compartment_id
    }
  }

  #Optional
  #   defined_tags  = tomap({oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.schedule_defined_tags_value})
  description   = var.schedule_description
  display_name  = var.schedule_display_name
  freeform_tags = var.schedule_freeform_tags
  time_ends   = local.schedule_time_ends
  time_starts = local.start_time
}

data "oci_resource_scheduler_schedules" "test_schedules" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.schedule_display_name
  schedule_id    = oci_resource_scheduler_schedule.test_schedule.id
  state          = var.schedule_state
}

locals {
  start_time = "2035-05-01T00:00:00Z"
  schedule_time_ends = "2035-12-31T00:00:00Z"
  schedule_body1 = [
    {
      "scan_tool" = "bandit",
      "scan_tool_display_name" = "OHAI Bandit",
      "scan_tool_type" = "static_code_analysis",
      "scanner_bucket_name" = "ohai-bandit-data-unstable",
      "scan_tool_plugin_mnemonic" = "bandit",
      "plugin_family" = "static_code_analysis"
    }
  ]
  schedule_body2 = [
    {
      "scan_tool" = "bandit",
      "scan_tool_display_name" = "OHAI Bandit",
      "scan_tool_type" = "static_code_analysis",
      "scanner_bucket_name" = "ohai-bandit-data-unstable",
      "scan_tool_plugin_mnemonic" = "bandit",
      "plugin_family" = "static_code_analysis"
    },
    {
      "scan_tool" = "dependency_check",
      "scan_tool_display_name" = "OHAI OWASP Dependency Check",
      "scan_tool_type" = "static_dependency_analysis",
      "scanner_bucket_name" = "ohai-dependency-check-data-unstable",
      "scan_tool_plugin_mnemonic" = "dependency_check",
      "plugin_family" = "static_dependency_analysis"
    }
  ]
}

resource "oci_resource_scheduler_schedule" "rs_body_param_test" {
  #Required
  action             = var.schedule_action
  compartment_id     = var.compartment_id
  description        = var.schedule_description
  display_name       = "test-fn-body-1-tf-created"
  recurrence_details = "FREQ=DAILY;COUNT=1"
  recurrence_type    = var.schedule_recurrence_type
  time_starts        = local.start_time

  resources {
    id = var.resource_id

    parameters {
      parameter_type = "BODY"
      value = [jsonencode(local.schedule_body1)]
    }
  }

  lifecycle {
    ignore_changes = [
      time_starts
    ]
  }
}

output "schedule_id" {
  value = oci_resource_scheduler_schedule.rs_body_param_test.id
}
