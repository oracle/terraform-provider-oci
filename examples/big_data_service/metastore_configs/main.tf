// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "api_key_id" {}
variable "bds_instance_id" {}
variable "metastore_id" {}

variable "bds_instance_metastore_config_bds_api_key_passphrase" {
  default = "V2VsY29tZTE="
}

variable "bds_instance_metastore_config_cluster_admin_password" {
  default = "V2VsY29tZTE="
}

variable "bds_instance_metastore_config_display_name" {
  default = "displayName"
}

variable "bds_instance_metastore_config_metastore_type" {
  default = "EXTERNAL"
}

variable "bds_instance_metastore_config_state" {
  default = "INACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_bds_bds_instance_metastore_config" "test_bds_instance_metastore_config" {
  #Required
  bds_api_key_id         = var.api_key_id
  bds_api_key_passphrase = var.bds_instance_metastore_config_bds_api_key_passphrase
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.bds_instance_metastore_config_cluster_admin_password
  metastore_id           = var.metastore_id

  #Optional
  display_name = var.bds_instance_metastore_config_display_name
}

data "oci_bds_bds_instance_metastore_configs" "test_bds_instance_metastore_configs" {
  #Required
  bds_instance_id = var.bds_instance_id

  #Optional
  display_name   = var.bds_instance_metastore_config_display_name
}

