// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "auth" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "kms_key_ocid_cross_region" {
}

variable "replica_availability_domain" {
}

variable "source_availability_domain" {
}

variable "config_file_profile" {
}

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_core_volume" "test_volume" {
  availability_domain = var.source_availability_domain
  compartment_id      = var.compartment_ocid
  display_name        = "test-volume-group-replica-kms-volume"
  size_in_gbs         = 50
}

resource "oci_core_volume_group" "test_volume_group" {
  availability_domain = var.source_availability_domain
  compartment_id      = var.compartment_ocid
  display_name        = "test-volume-group-replica-kms"

  source_details {
    type       = "volumeIds"
    volume_ids = [oci_core_volume.test_volume.id]
  }

  volume_group_replicas {
    availability_domain = var.replica_availability_domain
    display_name        = "test-volume-group-replica-kms-replica"
    xrr_kms_key_id      = var.kms_key_ocid_cross_region
  }

  volume_group_replicas_deletion = true
}

output "volume_group" {
  value = {
    id                    = oci_core_volume_group.test_volume_group.id
    volume_group_replicas = oci_core_volume_group.test_volume_group.volume_group_replicas
  }
}
