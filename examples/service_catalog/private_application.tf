// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "private_application_defined_tags_value" {
  default = "value"
}

variable "private_application_display_name" {
  default = "displayName"
}

variable "private_application_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "private_application_logo_file_base64encoded" {
}

variable "private_application_long_description" {
  default = "longDescription"
}

variable "private_application_package_details_package_type" {
  default = "STACK"
}

variable "private_application_package_details_version" {
  default = "version"
}

variable "private_application_package_details_zip_file_base64encoded" {
}

variable "private_application_short_description" {
  default = "shortDescription"
}


resource "oci_service_catalog_private_application" "test_private_application" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.private_application_display_name
  package_details {
    #Required
    package_type = var.private_application_package_details_package_type
    version      = var.private_application_package_details_version

    #Optional
    zip_file_base64encoded = var.private_application_package_details_zip_file_base64encoded
  }
  short_description = var.private_application_short_description

  #Optional
  freeform_tags           = var.private_application_freeform_tags
  logo_file_base64encoded = var.private_application_logo_file_base64encoded
  long_description        = var.private_application_long_description
}

data "oci_service_catalog_private_applications" "test_private_applications" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name           = var.private_application_display_name
  private_application_id = oci_service_catalog_private_application.test_private_application.id
}

