// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "service_catalog_defined_tags_value" {
  default = "value"
}

variable "service_catalog_display_name" {
  default = "displayName"
}

variable "service_catalog_freeform_tags" {
  default = { "bar-key" = "value" }
}


resource "oci_service_catalog_service_catalog" "test_service_catalog" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.service_catalog_display_name

  #Optional
  freeform_tags = var.service_catalog_freeform_tags
}

data "oci_service_catalog_service_catalogs" "test_service_catalogs" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name       = var.service_catalog_display_name
  service_catalog_id = oci_service_catalog_service_catalog.test_service_catalog.id
}

