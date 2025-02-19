---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_plugin_errors"
sidebar_current: "docs-oci-datasource-jms-plugin_errors"
description: |-
  Provides the list of Plugin Errors in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_plugin_errors
This data source provides the list of Plugin Errors in Oracle Cloud Infrastructure Jms service.

Returns a list of plugin errors that describe all detected errors.

## Example Usage

```hcl
data "oci_jms_plugin_errors" "test_plugin_errors" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.plugin_error_compartment_id_in_subtree
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	time_first_seen_greater_than_or_equal_to = var.plugin_error_time_first_seen_greater_than_or_equal_to
	time_first_seen_less_than_or_equal_to = var.plugin_error_time_first_seen_less_than_or_equal_to
	time_last_seen_greater_than_or_equal_to = var.plugin_error_time_last_seen_greater_than_or_equal_to
	time_last_seen_less_than_or_equal_to = var.plugin_error_time_last_seen_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `compartment_id_in_subtree` - (Optional) Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments. 
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the managed instance.
* `time_first_seen_greater_than_or_equal_to` - (Optional) If specified, only errors with a first seen time later than this parameter will be included in the search (formatted according to RFC3339).
* `time_first_seen_less_than_or_equal_to` - (Optional) If specified, only errors with a first seen time earlier than this parameter will be included in the search (formatted according to RFC3339).
* `time_last_seen_greater_than_or_equal_to` - (Optional) If specified, only errors with a last seen time later than this parameter will be included in the search (formatted according to RFC3339).
* `time_last_seen_less_than_or_equal_to` - (Optional) If specified, only errors with a last seen time earlier than this parameter will be included in the search (formatted according to RFC3339).


## Attributes Reference

The following attributes are exported:

* `plugin_error_collection` - The list of plugin_error_collection.

### PluginError Reference

The following attributes are exported:

* `items` - A list of PluginErrorSummary.
	* `agent_type` - The agent type.
	* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet. 
	* `errors` - List of plugin error details.
		* `details` - Optional string containing additional details.
		* `reason` - The plugin error reason.
		* `time_last_seen` - The date and time the resource was _last_ reported to JMS. This is potentially _after_ the specified time period provided by the filters. For example, a resource can be last reported to JMS before the start of a specified time period, if it is also reported during the time period. 
	* `host_name` - The HostName or Compute Instance name of the Managed Instance running the plugin.
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Instance running the plugin.
	* `time_first_seen` - The timestamp of the first time an error was detected. 
	* `time_last_seen` - The timestamp of the last time an error was detected. 

