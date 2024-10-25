// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "asset_source_are_historical_metrics_collected" {
  default = false
}

variable "asset_source_are_realtime_metrics_collected" {
  default = false
}

variable "asset_source_defined_tags_value" {
  default = "value"
}

variable "asset_source_discovery_credentials_type" {
  default = "BASIC"
}

variable "asset_source_display_name" {
  default = "displayName"
}

variable "asset_source_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "asset_source_replication_credentials_type" {
  default = "BASIC"
}

variable "asset_source_state" {
  default = "AVAILABLE"
}

variable "asset_source_system_tags" {
  default = "value"
}

variable "asset_source_type" {
  default = "VMWARE"
}

variable "asset_source_vcenter_endpoint" {
  default = "vcenterEndpoint"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_asset_source" "test_asset_source" {
  #Required
  assets_compartment_id = oci_identity_compartment.test_compartment.id
  compartment_id        = var.compartment_id
  discovery_credentials {
    #Required
    secret_id = oci_vault_secret.test_secret.id
    type      = var.asset_source_discovery_credentials_type
  }
  environment_id   = oci_cloud_bridge_environment.test_environment.id
  inventory_id     = oci_cloud_bridge_inventory.test_inventory.id
  type             = var.asset_source_type
  vcenter_endpoint = var.asset_source_vcenter_endpoint

  #Optional
  are_historical_metrics_collected = var.asset_source_are_historical_metrics_collected
  are_realtime_metrics_collected   = var.asset_source_are_realtime_metrics_collected
  defined_tags                     = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.asset_source_defined_tags_value)
  discovery_schedule_id            = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name                     = var.asset_source_display_name
  freeform_tags                    = var.asset_source_freeform_tags
  replication_credentials {
    #Required
    secret_id = oci_vault_secret.test_secret.id
    type      = var.asset_source_replication_credentials_type
  }
  system_tags = var.asset_source_system_tags
}

data "oci_cloud_bridge_asset_sources" "test_asset_sources" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  asset_source_id = oci_cloud_bridge_asset_source.test_asset_source.id
  display_name    = var.asset_source_display_name
  state           = var.asset_source_state
}

