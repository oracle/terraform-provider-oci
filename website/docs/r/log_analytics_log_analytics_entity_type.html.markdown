---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity_type"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_entity_type"
description: |-
  Provides the Log Analytics Entity Type resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_entity_type
This resource provides the Log Analytics Entity Type resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/LogAnalyticsEntityType

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics

Add custom log analytics entity type.

## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_entity_type" "test_log_analytics_entity_type" {
	#Required
	name = var.log_analytics_entity_type_name
	namespace = var.log_analytics_entity_type_namespace

	#Optional
	category = var.log_analytics_entity_type_category
	properties {
		#Required
		name = var.log_analytics_entity_type_properties_name

		#Optional
		description = var.log_analytics_entity_type_properties_description
	}
}
```

## Argument Reference

The following arguments are supported:

* `category` - (Optional) Log analytics entity type category. Category will be used for grouping and filtering. 
* `name` - (Required) Log analytics entity type name. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `properties` - (Optional) Log analytics entity type property definition.
	* `description` - (Optional) Description for the log analytics entity type property. 
	* `name` - (Required) Log analytics entity type property name. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - Array of log analytics entity type summary.
	* `category` - Log analytics entity type category. Category will be used for grouping and filtering. 
	* `cloud_type` - Log analytics entity type group. This can be CLOUD (OCI) or NON_CLOUD otherwise. 
	* `internal_name` - Internal name for the log analytics entity type. 
	* `name` - Log analytics entity type name. 
	* `state` - The current lifecycle state of the log analytics entity type. 
	* `time_created` - Time the log analytics entity type was created. An RFC3339 formatted datetime string. 
	* `time_updated` - Time the log analytics entity type was updated. An RFC3339 formatted datetime string. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Entity Type
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Entity Type
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Entity Type


## Import

LogAnalyticsEntityTypes can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_log_analytics_entity_type.test_log_analytics_entity_type "namespaces/{namespaceName}/logAnalyticsEntityTypes" 
```

