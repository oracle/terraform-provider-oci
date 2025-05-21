# $Header$
#
# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      tags.tf - tags file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/pluggable_databases/local_clone
#    NOTES
#      Terraform Example:
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   05/08/2025 - Created


resource "oci_identity_tag_namespace" "tag_namespace" {
  #Required
  compartment_id = var.compartment_id
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace.id

  is_retired = false
}