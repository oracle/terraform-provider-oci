// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "private_application_package_display_name" {
  default = "displayName"
}

variable "private_application_package_package_type" {
  default = []
}

data "oci_service_catalog_private_application_packages" "test_private_application_packages" {
  #Required
  private_application_id = oci_service_catalog_private_application.test_private_application.id

  #Optional
  display_name                   = var.private_application_package_display_name
  package_type                   = var.private_application_package_package_type
}
