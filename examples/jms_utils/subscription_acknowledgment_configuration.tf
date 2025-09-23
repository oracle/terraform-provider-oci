// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "subscription_acknowledgment_configuration_is_acknowledged" {
  default = true
}


data "oci_jms_utils_subscription_acknowledgment_configuration" "test_subscription_acknowledgment_configuration" {

  #Optional
  compartment_id  = var.tenancy_ocid
  is_acknowledged = var.subscription_acknowledgment_configuration_is_acknowledged
}
