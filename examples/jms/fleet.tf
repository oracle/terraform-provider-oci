# // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
# // Licensed under the Mozilla Public License v2.0

resource "oci_jms_fleet" "example_fleet" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "Example Fleet"
  inventory_log {
    log_group_id = var.fleet_log_group_ocid
    log_id       = var.fleet_inventory_log_ocid
  }

  #Optional
  description                  = "Example Fleet created by Terraform"
  freeform_tags                = var.fleet_freeform_tags
  operation_log {
    log_group_id = var.fleet_log_group_ocid
    log_id       = var.fleet_operation_log_ocid
  }

  # Create the Tag namespace in OCI before enabling
  # See user guide: https://docs.oracle.com/en-us/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm
  # defined_tags  = var.fleet_defined_tags
}

data "oci_jms_fleets" "example_fleets" {

  #Optional
  compartment_id        = var.compartment_ocid
  id                    = oci_jms_fleet.example_fleet.id
  display_name          = oci_jms_fleet.example_fleet.display_name
  display_name_contains = oci_jms_fleet.example_fleet.display_name
  state                 = oci_jms_fleet.example_fleet.state
}

