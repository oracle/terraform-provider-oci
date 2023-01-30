// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "volume_group_defined_tags_value" {
  default = "value"
}

variable "volume_group_display_name" {
  default = "displayName"
}

variable "volume_group_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "volume_group_source_details_type" {
  default = "volumeIds"
}

variable "volume_group_source_details_volume_ids" {
  default = []
}

variable "volume_group_state" {
  default = "AVAILABLE"
}

variable "volume_group_volume_group_replicas_availability_domain" {
  default = "availabilityDomain"
}

variable "volume_group_volume_group_replicas_display_name" {
  default = "displayName"
}

data "oci_identity_availability_domains" "test_availability_domains" {
    compartment_id = var.tenancy_ocid
}

resource "oci_core_volume" "source_volume_list" {
	count = 2
	display_name = format("source-volume-%d", count.index + 1)

	#Required
	availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
	compartment_id = var.compartment_id
}

resource "oci_core_volume_group" "test_volume_group" {
  #Required
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id      = var.compartment_id
  source_details {
    #Required
    type = var.volume_group_source_details_type

    #Optional
    volume_ids              = oci_core_volume.source_volume_list.*.id
  }

  #Optional
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_group_defined_tags_value}")
  display_name     = var.volume_group_display_name
  freeform_tags    = var.volume_group_freeform_tags
  volume_group_replicas {
    #Required
    availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name

    #Optional
    display_name = var.volume_group_volume_group_replicas_display_name
  }
}

data "oci_core_volume_groups" "test_volume_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
  display_name        = var.volume_group_display_name
  state               = var.volume_group_state
}

