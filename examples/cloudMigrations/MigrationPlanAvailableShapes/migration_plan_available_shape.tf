// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "migration_plan_available_shape_availability_domain" {
  default = "availabilityDomain"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_cloud_migrations_migration_plan_available_shapes" "test_migration_plan_available_shapes" {
  #Required
  migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id

  #Optional
  availability_domain  = var.migration_plan_available_shape_availability_domain
  compartment_id       = var.compartment_id
  dvh_host_id          = oci_cloud_migrations_dvh_host.test_dvh_host.id
  reserved_capacity_id = oci_cloud_migrations_reserved_capacity.test_reserved_capacity.id
}

