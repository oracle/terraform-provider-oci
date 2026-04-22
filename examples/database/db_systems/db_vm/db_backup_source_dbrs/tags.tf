# $Header$
#
# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      tags.tf - Tag resources file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup_source_dbrs

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-dbvm-dbrs"
  is_retired     = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_retired       = false
}
