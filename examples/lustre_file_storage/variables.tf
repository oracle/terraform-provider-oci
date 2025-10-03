// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "subnet_id" {}


variable "lustre_file_system_availability_domain" {
  default = "availabilityDomain"
}

variable "lustre_file_system_capacity_in_gbs" {
  default = 31200
}

variable "lustre_file_system_defined_tags_value" {
  default = "value"
}

variable "lustre_file_system_display_name" {
  default = "testDisplayName"
}

variable "lustre_file_system_display_name_import_job" {
  default = "testDisplayName-import-job"
}

variable "lustre_file_system_file_system_description" {
  default = "testFileSystemDescription"
}

variable "lustre_file_system_import_job_file_system_description" {
  default = "Lustre file system for starting and stopping an import job"
}


variable "lustre_file_system_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "lustre_file_system_id" {
  default = "id"
}

variable "lustre_file_system_nsg_ids" {
  default = []
}

variable "lustre_file_system_performance_tier" {
  default = "MBPS_PER_TB_125"
}

variable "lustre_file_system_root_squash_configuration_client_exceptions" {
  default = []
}

variable "lustre_file_system_root_squash_configuration_identity_squash" {
  default = "NONE"
}

variable "lustre_file_system_root_squash_configuration_squash_gid" {
  default = null
}

variable "lustre_file_system_root_squash_configuration_squash_uid" {
  default = null
}

variable "lustre_file_system_state" {
  default = "AVAILABLE"
}

variable "cluster_placement_group_id" {
  default = null
}

variable "lustre_file_system_name" {
  default = "lustre"
}

variable "my_vcn-cidr" {
  default = "10.0.0.0/16"
}

variable "my_subnet_cidr" {
  default = "10.0.0.0/24"
}

variable "object_storage_link_display_name" {
  default = "link1"
}

variable "object_storage_link_display_name_export_job" {
  default = "link-export-job"
}

variable "object_storage_link_display_name_import_job" {
  default = "link-import-job"
}

variable "file_system_path" {
  default = "/test"
}

variable "object_storage_prefix" {
  default = "example/tf-test"
}

variable "object_storage_prefix_export_job" {
  default = "example/tf-test/export"
}

variable "object_storage_prefix_import_job" {
  default = "example/tf-test/import"
}

variable "sync_job_id" {
  default = "syncJobId"
}