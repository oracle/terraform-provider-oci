// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_identity_tag_namespace" "tag_namespace_rd" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "Just a test"
  name           = "${var.tag_namespace_name}"

  is_retired = false
}

resource "oci_identity_tag" "tag_rd" {
  #Required
  description      = "example tag rd"
  name             = "tagRD"
  tag_namespace_id = "${oci_identity_tag_namespace.tag_namespace_rd.id}"

  #Optional
  is_cost_tracking = false // default is "false". The value "true" is only permitted if the associated tag namespace is part of the root compartment.
  is_retired       = false

  validator {
    validator_type = "ENUM"
    values         = ["value1", "value2"]
  }
}

resource "oci_identity_tag_default" "tag_default_rd" {
  compartment_id    = "${var.compartment_ocid}"
  tag_definition_id = "${oci_identity_tag.tag_rd.id}"
  value             = "value1"
}
