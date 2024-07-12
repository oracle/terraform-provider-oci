// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "scheduled_action_param_type" {
  default = "DB_SERVER_FULL_SOFTWARE_UPDATE"
}


data "oci_database_scheduled_action_params" "test_scheduled_action_params" {
  #Required
  type = var.scheduled_action_param_type
}