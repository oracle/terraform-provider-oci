// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "service_catalog_association_entity_type" {
  default = "privateapplication"
}

resource "oci_service_catalog_service_catalog_association" "test_service_catalog_association" {
  #Required
  entity_id          = oci_service_catalog_private_application.test_private_application.id
  service_catalog_id = oci_service_catalog_service_catalog.test_service_catalog.id

  #Optional
  entity_type = var.service_catalog_association_entity_type
}

data "oci_service_catalog_service_catalog_associations" "test_service_catalog_associations" {

  #Optional
  entity_id                      = oci_service_catalog_private_application.test_private_application.id
  entity_type                    = var.service_catalog_association_entity_type
  service_catalog_association_id = oci_service_catalog_service_catalog_association.test_service_catalog_association.id
  service_catalog_id             = oci_service_catalog_service_catalog.test_service_catalog.id
}

