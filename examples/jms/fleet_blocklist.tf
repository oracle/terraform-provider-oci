// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "fleet_blocklist_operation" {
  default = "DELETE_JAVA_INSTALLATION"
}

data "oci_jms_fleet_blocklists" "test_fleet_blocklists" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  managed_instance_id = var.managed_instance_ocid
  operation           = var.fleet_blocklist_operation
}