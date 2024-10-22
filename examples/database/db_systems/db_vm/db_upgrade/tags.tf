# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      tags.tf
#
#    USAGE
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade
#
#    NOTES
#      Terraform Example: TestDatabaseDatabaseUpgradeResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   11/12/2024 - Created


resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}
