// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}
variable "compartment_id" {
  default = "OCID"
}
variable "vaultSecretId" {
  default = "OCID"
}

variable "discovery_schedule_execution_recurrences" {
  default = "FREQ=DAILY;BYHOUR=6"
}

variable "discovery_schedule_display_name" {
  default = "displayName"
}

resource "oci_cloud_bridge_discovery_schedule" "test_discovery_schedule" {
  compartment_id        = var.compartment_id
  execution_recurrences = var.discovery_schedule_execution_recurrences
  display_name          = var.discovery_schedule_display_name
}

variable "environment_display_name" {
  default = "displayName"
}

resource "oci_cloud_bridge_environment" "test_environment" {
  compartment_id = var.compartment_id
  display_name   = var.environment_display_name
}

variable "asset_source_is_cost_information_collected" {
  default = false
}

variable "asset_source_are_historical_metrics_collected" {
  default = false
}

variable "asset_source_are_realtime_metrics_collected" {
  default = false
}

variable "asset_source_discovery_credentials_type" {
  default = "BASIC"
}

variable "aws_asset_source_discovery_credentials_type" {
  default = "API_KEY"
}

variable "asset_source_display_name" {
  default = "displayName"
}

variable "asset_source_replication_credentials_type" {
  default = "BASIC"
}

variable "asset_source_state" {
  default = "ACTIVE"
}

variable "asset_source_type" {
  default = "VMWARE"
}

variable "aws_asset_source_type" {
  default = "AWS"
}

variable "aws_asset_source_account_key" {
  default = "000000000000"
}

variable "aws_asset_source_region" {
  default = "eu-central-1"
}

variable "olvm_asset_source_type" {
  default = "OLVM"
}

variable "asset_source_vcenter_endpoint" {
  default = "https://11.0.11.130/sdk"
}

variable "asset_source_olvm_endpoint" {
  default = "https://11.0.11.131:443"
}

variable "inventory_id" {
  default = "OCID"
}


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  # version             = "8.3.0"
}

resource "oci_cloud_bridge_asset_source" "test_asset_source" {
  assets_compartment_id = var.compartment_id
  compartment_id        = var.compartment_id
  discovery_credentials {
    secret_id = var.vaultSecretId
    type      = var.asset_source_discovery_credentials_type
  }
  environment_id                   = oci_cloud_bridge_environment.test_environment.id
  inventory_id                     = var.inventory_id
  type                             = var.asset_source_type
  vcenter_endpoint                 = var.asset_source_vcenter_endpoint
  are_historical_metrics_collected = var.asset_source_are_historical_metrics_collected
  are_realtime_metrics_collected   = var.asset_source_are_realtime_metrics_collected
  discovery_schedule_id            = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name                     = var.asset_source_display_name
  replication_credentials {
    secret_id = var.vaultSecretId
    type      = var.asset_source_replication_credentials_type
  }
}

data "oci_cloud_bridge_asset_sources" "test_asset_sources" {
  compartment_id  = var.compartment_id
  asset_source_id = oci_cloud_bridge_asset_source.test_asset_source.id
  display_name    = var.asset_source_display_name
  state           = var.asset_source_state
}

resource "oci_cloud_bridge_asset_source" "test_aws_asset_source" {
  assets_compartment_id = var.compartment_id
  compartment_id        = var.compartment_id
  discovery_credentials {
    secret_id = var.vaultSecretId
    type      = var.aws_asset_source_discovery_credentials_type
  }
  environment_id                   = oci_cloud_bridge_environment.test_environment.id
  inventory_id                     = var.inventory_id
  type                             = var.aws_asset_source_type
  aws_account_key                  = var.aws_asset_source_account_key
  aws_region                       = var.aws_asset_source_region
  is_cost_information_collected    = var.asset_source_is_cost_information_collected
  are_historical_metrics_collected = var.asset_source_are_historical_metrics_collected
  are_realtime_metrics_collected   = var.asset_source_are_realtime_metrics_collected
  discovery_schedule_id            = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name                     = var.asset_source_display_name
  replication_credentials {
    secret_id = var.vaultSecretId
    type      = var.aws_asset_source_discovery_credentials_type
  }
}

data "oci_cloud_bridge_asset_sources" "test_aws_asset_sources" {
  compartment_id  = var.compartment_id
  asset_source_id = oci_cloud_bridge_asset_source.test_aws_asset_source.id
  display_name    = var.asset_source_display_name
  state           = var.asset_source_state
}

resource "oci_cloud_bridge_asset_source" "test_olvm_asset_source" {
  assets_compartment_id = var.compartment_id
  compartment_id        = var.compartment_id
  discovery_credentials {
    secret_id = var.vaultSecretId
    type      = var.asset_source_discovery_credentials_type
  }
  environment_id                   = oci_cloud_bridge_environment.test_environment.id
  inventory_id                     = var.inventory_id
  type                             = var.olvm_asset_source_type
  olvm_endpoint                    = var.asset_source_olvm_endpoint
  discovery_schedule_id            = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name                     = var.asset_source_display_name
  replication_credentials {
    secret_id = var.vaultSecretId
    type      = var.asset_source_replication_credentials_type
  }
}

data "oci_cloud_bridge_asset_sources" "test_olvm_asset_sources" {
  compartment_id  = var.compartment_id
  asset_source_id = oci_cloud_bridge_asset_source.test_olvm_asset_source.id
  display_name    = var.asset_source_display_name
  state           = var.asset_source_state
}
