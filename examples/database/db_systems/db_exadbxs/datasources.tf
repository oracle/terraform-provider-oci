// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

locals {
  ad = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  # the list of supported shape_attributes of ExaDbXS
  exadbxs_shape_name        = data.oci_database_db_system_shapes.exadbxs_shape.db_system_shapes.0.name
  exadbxs_shape_family_name = data.oci_database_db_system_shapes.exadbxs_shape.db_system_shapes.0.shape_family
  exadbxs_shape_attributes  = data.oci_database_db_system_shapes.exadbxs_shape.db_system_shapes.0.shape_attributes
  shape_attribute           = "BLOCK_STORAGE"
}

data "oci_database_backups" "test_database_backups_by_exadbxs" {
  compartment_id = var.compartment_ocid
  shape_family   = "EXADB_XS"
}

data "oci_database_gi_versions" "test_gi_versions" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  shape               = local.exadbxs_shape_name
  shape_attribute     = local.shape_attribute
  availability_domain = local.ad
}

data "oci_database_gi_version_minor_versions" "test_gi_minor_versions" {
  #Required
  version = data.oci_database_gi_versions.test_gi_versions.gi_versions[0].version
  #Optional
  compartment_id                 = data.oci_database_gi_versions.test_gi_versions.compartment_id
  availability_domain            = data.oci_database_gi_versions.test_gi_versions.availability_domain
  shape_family                   = local.exadbxs_shape_family_name
  shape                          = local.exadbxs_shape_name
  is_gi_version_for_provisioning = false
}

data "oci_database_db_system_shapes" "exadbxs_shape" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain = local.ad
  shape_attribute     = local.shape_attribute
  filter {
    name = "shape"
    values = ["ExaDbXS"]
  }
}

data "oci_database_db_versions" "test_db_versions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  db_system_shape = local.exadbxs_shape_name
  shape_attribute = local.shape_attribute
}