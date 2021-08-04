// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_management_agent_management_agent_plugin_count" "test_management_agent_plugin_count" {
  #Required
  compartment_id = var.compartment_ocid
  group_by = "pluginName"
}