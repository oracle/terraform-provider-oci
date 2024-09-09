// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "fleet_application_type" {
  default = "applicationType"
}

variable "fleet_description" {
  default = "fleet description"
}

variable "fleet_display_name" {
  default = "fleetDisplayName"
}

variable "fleet_environment_type" {
  default = "environmentType"
}

variable "fleet_fleet_type" {
  default = "GENERIC"
}

variable "fleet_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "fleet_group_type" {
  default = "ENVIRONMENT"
}

variable "fleet_is_target_auto_confirm" {
  default = true
}

variable "fleet_notification_preferences_preferences_on_job_failure" {
  default = false
}

variable "fleet_notification_preferences_preferences_on_topology_modification" {
  default = false
}

variable "fleet_notification_preferences_preferences_on_upcoming_schedule" {
  default = false
}

variable "fleet_product" {
  default = "OS(COMPUTE)"
}

variable "fleet_products" {
  default = ["OS(COMPUTE)"]
}

variable "fleet_resource_selection_type" {
  default = "MANUAL"
}

variable "fleet_rule_selection_criteria_match_condition" {
  default = "MATCH_ALL"
}

variable "fleet_rule_selection_criteria_rules_basis" {
  default = "basis"
}

variable "fleet_rule_selection_criteria_rules_conditions_attr_group" {
  default = "attrGroup"
}

variable "fleet_rule_selection_criteria_rules_conditions_attr_key" {
  default = "attrKey"
}

variable "fleet_rule_selection_criteria_rules_conditions_attr_value" {
  default = "attrValue"
}

variable "fleet_state" {
  default = "NEEDS_ATTENTION"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

######## Supporting resources used with Fleet.

# Topic creation
resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_id
  name           = random_string.topicname.result
}

########

resource "oci_fleet_apps_management_fleet" "test_fleet" {
  #Required
  compartment_id = var.tenancy_ocid
  fleet_type     = var.fleet_fleet_type

  lifecycle {
    ignore_changes = ["defined_tags"]
  }

  #Optional
  application_type       = var.fleet_application_type
  description            = var.fleet_description
  display_name           = var.fleet_display_name
  environment_type       = var.fleet_environment_type
  freeform_tags          = var.fleet_freeform_tags
  group_type             = var.fleet_group_type
  is_target_auto_confirm = var.fleet_is_target_auto_confirm
  notification_preferences {
    #Required
    compartment_id = var.compartment_id
    topic_id       = oci_ons_notification_topic.test_notification_topic.id

    #Optional
    preferences {

      #Optional
      on_job_failure           = var.fleet_notification_preferences_preferences_on_job_failure
      on_topology_modification = var.fleet_notification_preferences_preferences_on_topology_modification
      on_upcoming_schedule     = var.fleet_notification_preferences_preferences_on_upcoming_schedule
    }
  }
  products                = var.fleet_products
  resource_selection_type = var.fleet_resource_selection_type
  rule_selection_criteria {
    #Required
    match_condition = var.fleet_rule_selection_criteria_match_condition
    rules {
      #Required
      compartment_id = var.tenancy_ocid
      conditions {
        #Required
        attr_group = var.fleet_rule_selection_criteria_rules_conditions_attr_group
        attr_key   = var.fleet_rule_selection_criteria_rules_conditions_attr_key
        attr_value = var.fleet_rule_selection_criteria_rules_conditions_attr_value
      }
      resource_compartment_id = var.compartment_id
      #Optional
      basis = var.fleet_rule_selection_criteria_rules_basis
    }
  }
}

data "oci_fleet_apps_management_fleets" "test_fleets" {

  #Optional
  application_type = var.fleet_application_type
  compartment_id   = var.tenancy_ocid
  display_name     = var.fleet_display_name
  environment_type = var.fleet_environment_type
  fleet_type       = var.fleet_fleet_type
  product          = var.fleet_product
  state            = var.fleet_state
}

