
// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "entity_id" {
  type        = string
  description = "ID of the entity"
}

variable "entity_name" {
  type        = string
  description = "Name of the entity"
}

variable "entity_type_name" {
  type        = string
  description = "Type name of the entity"
}

variable "source_name" {
  type        = string
  description = "Name of the source"
}

variable "source_type_name" {
  type        = string
  description = "Type name of the source"
}

variable "pattern_id" {
  type        = number
  description = "ID of the pattern"
}

variable "host" {
  type        = string
  description = "Hostname"
}

variable "agent_id" {
  type        = string
  description = "ID of the agent"
}

variable "log_group_id" {
  type        = string
  description = "ID of the log group"
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "test_namespace" {
  compartment_id   = var.compartment_id
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// Create an association with required parameters
resource "oci_log_analytics_namespace_association" "test_namespace_association_mandatory" {
  compartment_id   = var.compartment_id
  entity_id        = var.entity_id
  log_group_id     = var.log_group_id
  namespace        = data.oci_objectstorage_namespace.test_namespace.namespace
  source_name      = var.source_name
}

// Create an association with optional parameters
resource "oci_log_analytics_namespace_association" "test_namespace_association_optional" {
  compartment_id   = var.compartment_id
  entity_id        = var.entity_id
  source_name      = var.source_name
  log_group_id     = var.log_group_id
  namespace        = data.oci_objectstorage_namespace.test_namespace.namespace

  is_from_republish = "true"
  association_properties {
    name = "management_agent.os_file.timezone"

    patterns {
      id    = var.pattern_id
      value = "PST"
    }
    value = "IST"
  }
}


