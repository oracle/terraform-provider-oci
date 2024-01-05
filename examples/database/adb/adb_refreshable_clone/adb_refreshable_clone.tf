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
}

variable "compartment_id" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

variable "defined_tag_namespace_name" {
  default = ""
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}

data "oci_database_autonomous_db_versions" "test_autonomous_db_versions" {
  compartment_id = var.compartment_id
}

data "oci_database_autonomous_db_versions" "test_autonomous_dw_versions" {
  compartment_id = var.compartment_id
  db_workload    = "DW"
}

resource "oci_database_autonomous_database" "test_autonomous_database_source" {
  admin_password           = "BEstrO0ng_#11"
  compartment_id           = var.compartment_id
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "rcB8w9HgKux1t1"
  db_version               = "19c"
  db_workload              = "OLTP"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }
  display_name = "regular_source"

  freeform_tags = {
    "Department" = "Finance"
  }

  is_auto_scaling_enabled                        = "false"
  is_dedicated                                   = "false"
  is_preview_version_with_service_terms_accepted = "false"
  license_model                                  = "LICENSE_INCLUDED"

  whitelisted_ips = ["1.1.1.1/28"]
}

resource "oci_database_autonomous_database" "test_autonomous_database_refreshable_clone" {
  admin_password           = ""
  compartment_id           = var.compartment_id
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "bjfjkXw4ZutTt2"
  db_version               = "19c"
  db_workload              = "OLTP"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }
  display_name = "refreshable_clone"

  freeform_tags = {
    "Department" = "Finance"
  }

  is_auto_scaling_enabled                        = "false"
  is_dedicated                                   = "false"
  is_preview_version_with_service_terms_accepted = "false"
  is_refreshable_clone                           = "true"
  license_model                                  = "LICENSE_INCLUDED"
  refreshable_mode                               = "MANUAL"
  source                                         = "CLONE_TO_REFRESHABLE"
  source_id                                      = oci_database_autonomous_database.test_autonomous_database_source.id

  whitelisted_ips = ["1.1.1.1/28"]
}

data "oci_database_autonomous_database" "oci_database_autonomous_database" {
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database_refreshable_clone.id
}

data "oci_database_autonomous_databases" "oci_database_autonomous_databases" {
  compartment_id = var.compartment_id

  filter {
    name   = "id"
    values = [oci_database_autonomous_database.test_autonomous_database_refreshable_clone.id]
  }
}

output "autonomous_database_refreshable_clone" {
  value = data.oci_database_autonomous_databases.oci_database_autonomous_databases.autonomous_databases
}

