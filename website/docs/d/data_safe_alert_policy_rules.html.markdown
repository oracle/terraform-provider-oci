---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_policy_rules"
sidebar_current: "docs-oci-datasource-data_safe-alert_policy_rules"
description: |-
  Provides the list of Alert Policy Rules in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_alert_policy_rules
This data source provides the list of Alert Policy Rules in Oracle Cloud Infrastructure Data Safe service.

Lists the rules of the specified alert policy. The alert policy is said to be satisfied when all rules in the policy evaulate to true.
If there are three rules: rule1,rule2 and rule3, the policy is satisfied if rule1 AND rule2 AND rule3 is True.


## Example Usage

```hcl
data "oci_data_safe_alert_policy_rules" "test_alert_policy_rules" {
  #Required
  alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `alert_policy_id` - (Required) The OCID of the alert policy.


## Attributes Reference

The following attributes are exported:

* `alert_policy_rule_collection` - The list of alert_policy_rule_collection.

### AlertPolicyRule Reference

The following attributes are exported:

* `description` - Describes the alert policy rule.
* `display_name` - The display name of the alert policy rule.
* `expression` - The conditional expression of the alert policy rule which evaluates to boolean value.
* `key` - The unique key of the alert policy rule.
* `state` - The current state of the alert policy rule.
* `time_created` - Creation date and time of the alert policy rule, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
