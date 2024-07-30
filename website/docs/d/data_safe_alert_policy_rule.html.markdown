---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_policy_rule"
sidebar_current: "docs-oci-datasource-data_safe-alert_policy_rule"
description: |-
  Provides details about a specific Alert Policy Rule in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_alert_policy_rule
This data source provides details about a specific Alert Policy Rule resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of a policy rule by its key.

## Example Usage

```hcl
data "oci_data_safe_alert_policy_rule" "test_alert_policy_rule" {
  #Required
  alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
  rule_key = var.alert_policy_rule_rule_key
}
```

## Argument Reference

The following arguments are supported:

* `alert_policy_id` - (Required) The OCID of the alert policy.
* `rule_key` - (Required) The key of the alert policy rule.


## Attributes Reference

The following attributes are exported:

* `description` - Describes the alert policy rule.
* `display_name` - The display name of the alert policy rule.
* `expression` - The conditional expression of the alert policy rule which evaluates to boolean value.
* `key` - The unique key of the alert policy rule.
* `state` - The current state of the alert policy rule.
* `time_created` - Creation date and time of the alert policy rule, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
