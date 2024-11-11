---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_plugin_error_analytics"
sidebar_current: "docs-oci-datasource-jms-plugin_error_analytics"
description: |-
  Provides the list of Plugin Error Analytics in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_plugin_error_analytics
This data source provides the list of Plugin Error Analytics in Oracle Cloud Infrastructure Jms service.

Returns a high level summary of PluginErrors.

## Example Usage

```hcl
data "oci_jms_plugin_error_analytics" "test_plugin_error_analytics" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.plugin_error_analytic_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `compartment_id_in_subtree` - (Optional) Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments. 


## Attributes Reference

The following attributes are exported:

* `plugin_error_aggregation_collection` - The list of plugin_error_aggregation_collection.

### PluginErrorAnalytic Reference

The following attributes are exported:

* `items` - A list of PluginErrorAggregationSummary.
	* `healthy_plugin_count` - Count of plugins with no problems.
	* `plugin_error_aggregations` - List of plugin aggregation errors.
		* `plugin_error_analytic_count` - Number of FleetErrors encountered for the specific reason.
		* `reason` - Enum that uniquely identifies the plugin error.

