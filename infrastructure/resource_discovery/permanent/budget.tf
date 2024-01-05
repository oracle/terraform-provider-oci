// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_budget_budget" "budget_rd" {
  #Required
  amount         = "1"
  compartment_id = "${var.tenancy_ocid}"
  reset_period   = "MONTHLY"
  target_type    = "COMPARTMENT"

  targets = [
    "${var.compartment_ocid}",
  ]

  #Optional
  description  = "budget1 description"
  display_name = "budgetRD"
}

resource "oci_budget_alert_rule" "alert_rule_rd" {
  #Required
  budget_id      = "${oci_budget_budget.budget_rd.id}"
  threshold      = "100"
  threshold_type = "ABSOLUTE"
  type           = "ACTUAL"

  #Optional
  description  = "alertRuleDescription"
  display_name = "alertRuleRD"
  message      = "possible overspend"
  recipients   = "JohnSmith@example.com"
}
