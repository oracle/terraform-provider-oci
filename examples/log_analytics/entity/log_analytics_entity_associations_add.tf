// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "log_analytics_entity_associations_add_association_entities" {
  default = []
}


# Sample create entity associations for entities created using oci_log_analytics_log_analytics_entity resource
resource "oci_log_analytics_log_analytics_entity_associations_add" "test_log_analytics_entity_associations_add" {
  #Required
  association_entities    = tolist([oci_log_analytics_log_analytics_entity.entityOptional.id])
  log_analytics_entity_id = oci_log_analytics_log_analytics_entity.entityRequired.id
  namespace               = data.oci_objectstorage_namespace.ns.namespace
}