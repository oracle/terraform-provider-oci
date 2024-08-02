// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_summarize_resource_inventory" "test_summarize_resource_inventories" {

  #Optional
  compartment_id = var.compartment_ocid
  time_end       = var.time_end
  time_start     = var.time_start
}
