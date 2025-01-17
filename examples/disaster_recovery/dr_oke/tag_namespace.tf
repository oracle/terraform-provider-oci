// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tag_namespace_defined_tags_value" {
  default = "value"
}

variable "tag_namespace_description" {
  default = "This namespace contains tags that will be used in billing."
}

variable "tag_namespace_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "tag_namespace_include_subcompartments" {
  default = false
}

variable "tag_namespace_name" {
  default = "BillingTags"
}

variable "tag_namespace_state" {
  default = "ACTIVE"
}

variable defined_tag_namespace_name { default = "" }

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

resource "oci_identity_tag_namespace" "test_tag_namespace" {
  #Required
  compartment_id = var.compartment_id
  description    = var.tag_namespace_description
  name           = var.tag_namespace_name

  #Optional
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.tag_namespace_defined_tags_value}")
  freeform_tags = var.tag_namespace_freeform_tags
}

data "oci_identity_tag_namespaces" "test_tag_namespaces" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  include_subcompartments = var.tag_namespace_include_subcompartments
  state                   = var.tag_namespace_state
}

