// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage a log analytics entity
 */

variable "managed_agent_id" {}

# Sample create entity with required parameters.
resource "oci_log_analytics_log_analytics_entity" "entityRequired" {
  compartment_id          = var.compartment_ocid
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  name                    = "tf-entity-example-req"
  entity_type_name        = "Host (Linux)"
}

# Get details of above created entity with required parameters
data "oci_log_analytics_log_analytics_entity" "entityRequiredDetails" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_entity_id = oci_log_analytics_log_analytics_entity.entityRequired.id
}

# Sample create entity with optional parameters
resource "oci_log_analytics_log_analytics_entity" "entityOptional" {
  compartment_id          = var.compartment_ocid
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  name                    = "tf-entity-example-opt"
  entity_type_name        = "Host (Linux)"
  management_agent_id     = var.managed_agent_id
  hostname                = "aa.domainname.com"
  time_last_discovered    = "2023-09-12T20:30:00.006Z"
  metadata {
    items {
      name      = "environment1"
      type     = "infrastructure1"
      value    = "test1"
    } 
    items {
      name      = "environment2"
      type     = "infrastructure2"
      value    = "test2"
    }
  }
  timezone_region         = "PST8PDT"
  properties              = tomap({"JAVA_HOME" = "/usr/java/jdk1.8.0_202-amd64", "version" = "OEL-7uek"})
  freeform_tags           = tomap({"servicegroup" = "test", "Dept" = "Devops"})
}

# Get details of above created entity with optional parameters
data "oci_log_analytics_log_analytics_entity" "entityOptionalDetails" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_entity_id = oci_log_analytics_log_analytics_entity.entityOptional.id
}

# Get a list of entities with some query parameters
data "oci_log_analytics_log_analytics_entities" "entitiesList" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  state                      = "ACTIVE"
  lifecycle_details_contains = "READY"
}

# Get an overall summary of entities
data "oci_log_analytics_log_analytics_entities_summary" "entitiesSummary" {
  compartment_id = var.compartment_ocid
  namespace      = data.oci_objectstorage_namespace.ns.namespace
}
