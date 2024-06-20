// It needs to be run against the destination region, for example,
// if mysql/main.tf created the backup in IAD, and the copy should go to PHX,
// then this needs to be run against the PHX endpoint.

// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
  // Define the region where destination backup will be created.
  // Example: region = "us-phoenix-1"
}

variable "source_region" {
  // Define the region where the source backup is created.
  // Example: source_region = "us-ashburn-1"
}

variable "source_backup_id" {
  // Define the source backup ID to be copied when using resource
  // oci_mysql_mysql_backup.test_mysql_backup_cross_region_backup_copy
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_mysql_mysql_backup" "test_mysql_backup_cross_region_backup_copy" {
  # Required
  source_details {
    region = var.source_region
    backup_id = var.source_backup_id
    compartment_id = var.compartment_ocid
  }

  # Optional
  display_name = "CrossRegionBackupCopy"
  description = "test backup copy created by terraform"
}