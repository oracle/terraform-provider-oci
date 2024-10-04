// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_id" {}

variable "log_group_defined_tags_value" {
  default = "value2"
}

variable "log_group_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}