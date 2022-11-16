---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_ingest_time_rules_management"
sidebar_current: "docs-oci-resource-log_analytics-namespace_ingest_time_rules_management"
description: |-
  Provides the Namespace Ingest Time Rules Management resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_ingest_time_rules_management
This resource provides the Namespace Ingest Time Rules Management resource in Oracle Cloud Infrastructure Log Analytics service.

Enables the specified ingest time rule.


## Example Usage

```hcl
resource "oci_log_analytics_namespace_ingest_time_rules_management" "test_namespace_ingest_time_rules_management" {
	#Required
	ingest_time_rule_id = oci_events_rule.test_rule.id
	namespace = var.namespace_ingest_time_rules_management_namespace
	enable_ingest_time_rule = var.enable_ingest_time_rule
}
```

## Argument Reference

The following arguments are supported:

* `ingest_time_rule_id` - (Required) Unique ocid of the ingest time rule. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `enable_ingest_time_rule` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Ingest Time Rules Management
	* `update` - (Defaults to 20 minutes), when updating the Namespace Ingest Time Rules Management
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Ingest Time Rules Management
