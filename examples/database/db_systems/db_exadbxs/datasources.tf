// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

locals {
  ad = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
}

data "oci_database_backups" "test_database_backups_by_exadbxs" {
  compartment_id = var.compartment_ocid
  shape_family   = "EXADB_XS"
}

data "oci_database_gi_versions" "test_gi_versions" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  shape               = "ExaDbXS"
  availability_domain = local.ad
}

data "oci_database_gi_version_minor_versions" "test_gi_minor_versions" {
  #Required
  version = data.oci_database_gi_versions.test_gi_versions.gi_versions[0].version
  #Optional
  compartment_id                 = data.oci_database_gi_versions.test_gi_versions.compartment_id
  availability_domain            = data.oci_database_gi_versions.test_gi_versions.availability_domain
  shape_family                   = "EXADB_XS"
  shape                          = data.oci_database_gi_versions.test_gi_versions.shape
  is_gi_version_for_provisioning = false
}

