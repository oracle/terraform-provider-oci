// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "auth" {
}

variable "config_file_profile" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "destination_region" {
}

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

// Example 1: Using predefined policy

data "oci_core_volume_backup_policies" "test_predefined_volume_backup_policies" {
  filter {
    name   = "display_name"
    values = ["silver"]
  }
}

output "silver_policy_id" {
  value = data.oci_core_volume_backup_policies.test_predefined_volume_backup_policies.volume_backup_policies[0].id
}

resource "oci_core_volume" "test_volume_1" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "-tf-volume-1"
  size_in_gbs         = "50"
}

resource "oci_core_volume_backup_policy_assignment" "test_backup_policy_assignment" {
  asset_id  = oci_core_volume.test_volume_1.id
  policy_id = data.oci_core_volume_backup_policies.test_predefined_volume_backup_policies.volume_backup_policies[0].id
}

data "oci_core_volume_backup_policy_assignments" "test_backup_policy_assignments" {
  asset_id = oci_core_volume.test_volume_1.id

  filter {
    name   = "id"
    values = [oci_core_volume_backup_policy_assignment.test_backup_policy_assignment.id]
  }
}

output "test_backup_policy_assignments" {
  value = data.oci_core_volume_backup_policy_assignments.test_backup_policy_assignments.volume_backup_policy_assignments
}

// Example 2: Using custom scheduled policy

resource "oci_core_volume_backup_policy" "test_volume_backup_policy_custom" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  defined_tags = {
    "example-tag-namespace-all.example-tag" = "originalValue"
  }
  display_name = "BackupPolicy1"

  freeform_tags = {
    "Department" = "Finance"
  }

  schedules {
    #Required
    backup_type       = "INCREMENTAL"
    period            = "ONE_YEAR"
    retention_seconds = "604800"

    #Optional
    day_of_month   = "10"
    day_of_week    = "TUESDAY"
    hour_of_day    = "10"
    month          = "FEBRUARY"
    offset_seconds = "0"
    offset_type    = "STRUCTURED"
    time_zone      = "UTC"
  }

  destination_region = var.destination_region
}

resource "oci_core_volume" "test_volume_2" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "-tf-volume-2"
  size_in_gbs         = "50"
}

resource "oci_core_volume_backup_policy_assignment" "test_backup_policy_assignment_custom" {
  asset_id  = oci_core_volume.test_volume_2.id
  policy_id = oci_core_volume_backup_policy.test_volume_backup_policy_custom.id
}

data "oci_core_volume_backup_policy_assignments" "test_backup_policy_assignments_custom" {
  asset_id = oci_core_volume.test_volume_2.id

  filter {
    name   = "id"
    values = [oci_core_volume_backup_policy_assignment.test_backup_policy_assignment_custom.id]
  }
}

output "test_backup_policy_assignments_custom" {
  value = data.oci_core_volume_backup_policy_assignments.test_backup_policy_assignments_custom.volume_backup_policy_assignments
}

// Example 3: Assign a policy to a volume group

resource "oci_core_volume" "test_volume" {
  count               = 2
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = format("-tf-volume-%d", count.index + 1)
  size_in_gbs         = "50"
}

resource "oci_core_volume_group" "test_volume_group" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  source_details {
    #Required
    type = "volumeIds"

    volume_ids = oci_core_volume.test_volume.*.id
  }

  #Optional
  display_name = "test-volume-group-with-backup-policy"
}

resource "oci_core_volume_backup_policy_assignment" "test_volume_group_backup_policy_assignment" {
  asset_id  = oci_core_volume_group.test_volume_group.id
  policy_id = oci_core_volume_backup_policy.test_volume_backup_policy_custom.id
}

data "oci_core_volume_backup_policy_assignments" "test_volume_group_backup_policy_assignments" {
  asset_id = oci_core_volume_group.test_volume_group.id

  filter {
    name   = "id"
    values = [oci_core_volume_backup_policy_assignment.test_volume_group_backup_policy_assignment.id]
  }
}

output "volume_group_backup_policy_assignments" {
  value = data.oci_core_volume_backup_policy_assignments.test_volume_group_backup_policy_assignments
}