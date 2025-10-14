// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * Example: Automatic DR Configuration
 * 
 * This example demonstrates creating an Automatic DR Configuration on an existing
 * DR Protection Group with existing resources (DRPG, DR Plans, and ADB).
 *
 * See the comments above the resource definition for more details.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {
  # Update this with the compartment OCID where your DR resources exist
  default = "compartment ocid"
}

variable "automatic_dr_configuration_defined_tags_value" {
  default = "value"
}

variable "automatic_dr_configuration_display_name" {
  default = "My Automatic DR Configuration"
}

variable "automatic_dr_configuration_freeform_tags" {
  default = { "Department" = "Finance" }
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# ============================================================================
# NOTE: This example uses pre-existing resources in Ashburn (IAD) region:
# - DR Protection Group: "terraform-standby-drpg-iad"
# - Default Switchover Plan: "terraform-default-switchover"
# - Default Failover Plan: "terraform-default-failover"
# - Autonomous Database: "fsdradbs01"
# 
# Make sure these resources exist in your environment before running this example.
# ============================================================================

# Data source to get the DR Protection Group
data "oci_disaster_recovery_dr_protection_groups" "test_drpg" {
  compartment_id = var.compartment_id
  display_name   = "terraform-standby-drpg-iad"
}

# Data source to get the default switchover plan
data "oci_disaster_recovery_dr_plans" "test_switchover_plan" {
  dr_protection_group_id = data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id
  display_name           = "terraform-default-switchover"
}

# Data source to get the default failover plan
data "oci_disaster_recovery_dr_plans" "test_failover_plan" {
  dr_protection_group_id = data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id
  display_name           = "terraform-default-failover"
}

# Data source to get the Autonomous Database
data "oci_database_autonomous_databases" "test_adb" {
  compartment_id = var.compartment_id
  display_name   = "fsdradbs01"
}

# Automatic DR Configuration Resource
# 
# ⚠️  IMPORTANT - KNOWN ISSUE ⚠️ 
# This resource creation will fail with a "Work Request error" due to a known bug.
# The error is EXPECTED and the behavior is:
#   1. The resource WILL be created (you'll see it in the OCI Console)
#   2. The work request will fail, causing Terraform to error
#   3. The resource will move to FAILED or DELETING state
#   4. The service will automatically clean up the resource
# 
# When you run 'terraform apply', you will see an error like:
#   "Error: Work Request error... work request did not succeed..."
# This is expected behavior and not a problem with your configuration.
#
# To clean up after the error, you can:
#   - Wait for automatic cleanup by the service, OR
#   - Manually delete the resource from the OCI Console if it gets stuck
#
resource "oci_disaster_recovery_automatic_dr_configuration" "example_automatic_dr_configuration" {
  #Required
  display_name           = var.automatic_dr_configuration_display_name
  dr_protection_group_id = data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id
  
  # Members - Database resources for automatic DR
  # NOTE: Only DATABASE, AUTONOMOUS_DATABASE, or AUTONOMOUS_CONTAINER_DATABASE 
  # member types are supported for Automatic DR Configuration
  members {
    #Required
    member_id   = data.oci_database_autonomous_databases.test_adb.autonomous_databases[0].id
    member_type = "AUTONOMOUS_DATABASE"
    
    #Required - Enable automatic failover and switchover
    is_auto_failover_enabled   = true
    is_auto_switchover_enabled = true
  }

  #Required - Default DR Plans for automatic operations
  default_failover_dr_plan_id   = data.oci_disaster_recovery_dr_plans.test_failover_plan.dr_plan_collection.0.items.0.id
  default_switchover_dr_plan_id = data.oci_disaster_recovery_dr_plans.test_switchover_plan.dr_plan_collection.0.items.0.id
  
  #Optional - Tags
  freeform_tags = var.automatic_dr_configuration_freeform_tags

  lifecycle {
    ignore_changes = [defined_tags]
  }
}

# Data source to read Automatic DR Configurations
data "oci_disaster_recovery_automatic_dr_configurations" "example_automatic_dr_configurations" {
  #Required
  dr_protection_group_id = data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id

  #Optional
  display_name                  = var.automatic_dr_configuration_display_name
  automatic_dr_configuration_id = oci_disaster_recovery_automatic_dr_configuration.example_automatic_dr_configuration.id
}

# Data source to read a specific Automatic DR Configuration
data "oci_disaster_recovery_automatic_dr_configuration" "example_automatic_dr_configuration" {
  #Required
  automatic_dr_configuration_id = oci_disaster_recovery_automatic_dr_configuration.example_automatic_dr_configuration.id
}

