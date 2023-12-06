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

Lists the rules of the specified alert policy. The alert policy is said to be satisfied when all rules in the policy evaulate to true.
If there are three rules: rule1,rule2 and rule3, the policy is satisfied if rule1 AND rule2 AND rule3 is True.


## Example Usage

```hcl
data "oci_data_safe_alert_policy_rule" "test_alert_policy_rule" {
	#Required
	alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `alert_policy_id` - (Required) The OCID of the alert policy.


## Attributes Reference

The following attributes are exported:

* `items` - Array of alert policy rules summary
	* `description` - Describes the alert policy rule.
	* `expression` - The conditional expression of the alert policy rule which evaluates to boolean value.
	* `key` - The unique key of the alert policy rule.

