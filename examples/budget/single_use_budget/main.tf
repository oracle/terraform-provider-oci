// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to use the single use budget and alert rule resources.
 */

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "subscription_id" {
} 

provider "oci" {
  # version         = "4.67.0"
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_budget_budget" "test_single_use_budget" {
  # Required 
  amount         = "1"
  compartment_id = var.tenancy_ocid
  reset_period   = "MONTHLY"
  target_type    = "COMPARTMENT"

  targets = [
    var.compartment_ocid,
  ]

  # Optional
  description  = "budget single use"
  display_name = "budget2"
  processing_period_type = "SINGLE_USE"
  start_date   = "2023-07-12T16:01:19.847222+05:30"
  end_date     = "2023-08-12T16:01:19.847222+05:30"
}


data "oci_budget_budgets" "test_budgets" {
  #Required
  compartment_id = var.tenancy_ocid
  #Optional
  //  display_name = oci_budget_budget.test_budget.display_name
  //  state        = "ACTIVE"
}

output "budget_single_use" {
  value = {
    amount           = oci_budget_budget.test_single_use_budget.amount
    compartment_id   = oci_budget_budget.test_single_use_budget.compartment_id
    reset_period     = oci_budget_budget.test_single_use_budget.reset_period
    targets          = oci_budget_budget.test_single_use_budget.targets[0]
    description      = oci_budget_budget.test_single_use_budget.description
    display_name     = oci_budget_budget.test_single_use_budget.display_name
    alert_rule_count = oci_budget_budget.test_single_use_budget.alert_rule_count
    state            = oci_budget_budget.test_single_use_budget.state
    time_created     = oci_budget_budget.test_single_use_budget.time_created
    time_updated     = oci_budget_budget.test_single_use_budget.time_updated
    version          = oci_budget_budget.test_single_use_budget.version
    processing_period_type = oci_budget_budget.test_single_use_budget.processing_period_type
  }
}

resource "oci_budget_alert_rule" "test_alert_rule" {
  #Required
  budget_id      = oci_budget_budget.test_single_use_budget.id
  threshold      = "100"
  threshold_type = "ABSOLUTE"
  type           = "ACTUAL"

  #Optional
  description  = "alertRuleDescription"
  display_name = "alertRule"
  message      = "possible overspend"
  recipients   = "JohnSmith@example.com"
}

output "alert_rule" {
  value = {
    budget_id      = data.oci_budget_alert_rule.test_alert_rule.budget_id
    recipients     = data.oci_budget_alert_rule.test_alert_rule.recipients
    description    = data.oci_budget_alert_rule.test_alert_rule.description
    display_name   = data.oci_budget_alert_rule.test_alert_rule.display_name
    message        = data.oci_budget_alert_rule.test_alert_rule.message
    recipients     = data.oci_budget_alert_rule.test_alert_rule.recipients
    state          = data.oci_budget_alert_rule.test_alert_rule.state
    threshold      = data.oci_budget_alert_rule.test_alert_rule.threshold
    threshold_type = data.oci_budget_alert_rule.test_alert_rule.threshold_type
    time_created   = data.oci_budget_alert_rule.test_alert_rule.time_created
    time_updated   = data.oci_budget_alert_rule.test_alert_rule.time_updated
    type           = data.oci_budget_alert_rule.test_alert_rule.type
    version        = data.oci_budget_alert_rule.test_alert_rule.version
  }
}

data "oci_budget_alert_rule" "test_alert_rule" {
  #Required
  budget_id     = oci_budget_budget.test_single_use_budget.id
  alert_rule_id = oci_budget_alert_rule.test_alert_rule.id
}

data "oci_budget_alert_rules" "test_alert_rules" {
  #Required
  budget_id = oci_budget_budget.test_single_use_budget.id

  #Optional
  //  display_name = oci_budget_alert_rule.test_alert_rule.display_name
  state = "ACTIVE"
}