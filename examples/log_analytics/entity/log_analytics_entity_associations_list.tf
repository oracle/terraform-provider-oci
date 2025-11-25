// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "log_analytics_entity_associations_list_direct_or_all_associations" {
  default = "DIRECT"
}

# Get details of created entity associations with required and optional parameters
data "oci_log_analytics_log_analytics_entity_associations_list" "test_log_analytics_entity_associations_list" {
  #Required
  log_analytics_entity_id = oci_log_analytics_log_analytics_entity.entityRequired.id
  namespace               = data.oci_objectstorage_namespace.ns.namespace

  #Optional
  direct_or_all_associations = var.log_analytics_entity_associations_list_direct_or_all_associations
}