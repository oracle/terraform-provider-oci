// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.tag_namespace_description
  name           = var.tag_namespace_name
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

resource "oci_identity_tag" "tag2" {
  #Required
  description      = "tf example tag 2"
  name             = "tf-example-tag-2"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

