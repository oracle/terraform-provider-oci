// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "scheduling_plan_defined_tags_value" {
  default = "definedTags"
}

variable "scheduling_plan_display_name" {
  default = "displayName"
}

variable "scheduling_plan_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "scheduling_plan_id" {
  default = "id"
}

variable "scheduling_plan_is_using_recommended_scheduled_actions" {
  default = false
}

variable "scheduling_plan_service_type" {
  default = "EXACC"
}

variable "scheduling_plan_state" {
  default = "AVAILABLE"
}

resource "oci_database_scheduling_plan" "test_scheduling_plan" {
  #Required
  compartment_id       = var.compartment_id
  resource_id          = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
  service_type         = var.scheduling_plan_service_type

  #Optional
  defined_tags                           = var.scheduling_plan_defined_tags_value
  freeform_tags                          = var.scheduling_plan_freeform_tags
  is_using_recommended_scheduled_actions = var.scheduling_plan_is_using_recommended_scheduled_actions
}

data "oci_database_scheduling_plans" "test_scheduling_plans" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name         = var.scheduling_plan_display_name
  id                   = var.scheduling_plan_id
  resource_id          = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
  state                = var.scheduling_plan_state
}