// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "scheduling_policy_recommended_scheduled_action_plan_intent" {
  default = "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"
}


data "oci_database_scheduling_policy_recommended_scheduled_actions" "test_scheduling_policy_recommended_scheduled_actions" {
  #Required
  plan_intent                          = var.scheduling_policy_recommended_scheduled_action_plan_intent
  scheduling_policy_id                 = oci_database_scheduling_policy.test_scheduling_policy.id
  scheduling_policy_target_resource_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}