// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired = false
}


resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "chargeback_plan_name" {
  default = "planNameTF1"
}

variable "chargeback_plan_name_updated" {
  default = "planNameTF1u"
}

variable "chargeback_plan_description" {
  default = "planDescription"
}

variable "chargeback_plan_description_updated" {
  default = "planDescription2"
}

variable "defined_tag_value" {
  default = "value"
}

variable "freeform_tags" {
  default = { "bar-key" = "value" }
}

// Create an Operations Insights chargeback plan
resource "oci_opsi_chargeback_plan" "test_chargeback_plan" {
  compartment_id = var.compartment_ocid
  entity_source  = "CHARGEBACK_EXADATA"
  plan_name      = var.chargeback_plan_name
  plan_type      = "WEIGHTED_ALLOCATION"
  plan_description = var.chargeback_plan_description
  defined_tags     = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tag_value}")}"
  freeform_tags    = var.freeform_tags

  plan_custom_items {
    is_customizable = false
    name            = "statistic"
    value           = "AVG"
  }

  plan_custom_items {
    name  = "metricCostSplit"
    value = "CPU:100"
  }
}

// Update the chargeback plan
resource "oci_opsi_chargeback_plan" "test_chargeback_plan_updated" {
  compartment_id   = var.compartment_ocid
  entity_source    = "CHARGEBACK_EXADATA"
  plan_name        = var.chargeback_plan_name_updated
  plan_type        = "WEIGHTED_ALLOCATION"
  plan_description = var.chargeback_plan_description_updated
  defined_tags     = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tag_value}")}"
  freeform_tags    = var.freeform_tags

  plan_custom_items {
    is_customizable = false
    name            = "statistic"
    value           = "MAX"
  }

  plan_custom_items {
    name  = "metricCostSplit"
    value = "CPU:100"
  }

  depends_on = [oci_opsi_chargeback_plan.test_chargeback_plan]
}

// List Operations Insights chargeback plans
data "oci_opsi_chargeback_plans" "test_chargeback_plans" {
  compartment_id = var.compartment_ocid
  chargebackplan_id = oci_opsi_chargeback_plan.test_chargeback_plan_updated.id

  filter {
    name   = "id"
    values = [oci_opsi_chargeback_plan.test_chargeback_plan_updated.id]
  }
}

// Get an Operations Insights chargeback plan
data "oci_opsi_chargeback_plan" "test_chargeback_plan" {
  chargebackplan_id = oci_opsi_chargeback_plan.test_chargeback_plan_updated.id
}