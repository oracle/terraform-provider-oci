// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# You need actual OCID value of compute instance or managed instance to create this resource.
# variable "agent_id" {}
# resource "oci_jms_jms_plugin" "example_jms_plugin" {
#   #Required
#   agent_id = var.agent_id
#   compartment_id = var.compartment_ocid

#   #Optional
#   fleet_id = var.fleet_id
# }

data "oci_jms_jms_plugins" "example_jms_plugins" {
  compartment_id = var.compartment_ocid
  availability_status = "ACTIVE"
}