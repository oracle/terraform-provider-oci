// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Create entityType with required parameters
resource "oci_log_analytics_log_analytics_entity_type" "test_log_analytics_entity_type_required" {
  #Required
  name      = "tf-example-entity-type-req"
  namespace = data.oci_objectstorage_namespace.ns.namespace
}

# Get details of above created entityType with required parameters
data "oci_log_analytics_log_analytics_entity_type" "test_log_analytics_entity_type_required_details" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  entity_type_name = oci_log_analytics_log_analytics_entity_type.test_log_analytics_entity_type_required.name
}

# Create entityType with optional parameters
resource "oci_log_analytics_log_analytics_entity_type" "test_log_analytics_entity_type_optional" {
  #Required
  name      = "tf-example-entity-type-opt"
  namespace = data.oci_objectstorage_namespace.ns.namespace

  #Optional
  category = "CUSTOM"
  properties {
    #Required
    name = "propertyName"

    #Optional
    description = "propertyDescription"
  }
}

data "oci_log_analytics_log_analytics_entity_types" "test_log_analytics_entity_type_optional_details" {
  #Required
  namespace = data.oci_objectstorage_namespace.ns.namespace

  #Optional
  cloud_type    = "NON_CLOUD"
  name          = oci_log_analytics_log_analytics_entity_type.test_log_analytics_entity_type_optional.name
  state         = "ACTIVE"
}
