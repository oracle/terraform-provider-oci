---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_metric_extension"
sidebar_current: "docs-oci-resource-stack_monitoring-metric_extension"
description: |-
  Provides the Metric Extension resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_metric_extension
This resource provides the Metric Extension resource in Oracle Cloud Infrastructure Stack Monitoring service.

Creates a new metric extension resource for a given compartment


## Example Usage

```hcl
resource "oci_stack_monitoring_metric_extension" "test_metric_extension" {
	#Required
	collection_recurrences = var.metric_extension_collection_recurrences
	compartment_id = var.compartment_id
	display_name = var.metric_extension_display_name
	metric_list {
		#Required
		data_type = var.metric_extension_metric_list_data_type
		name = var.metric_extension_metric_list_name

		#Optional
		compute_expression = var.metric_extension_metric_list_compute_expression
		display_name = var.metric_extension_metric_list_display_name
		is_dimension = var.metric_extension_metric_list_is_dimension
		is_hidden = var.metric_extension_metric_list_is_hidden
		metric_category = var.metric_extension_metric_list_metric_category
		unit = var.metric_extension_metric_list_unit
	}
	name = var.metric_extension_name
	query_properties {
		#Required
		collection_method = var.metric_extension_query_properties_collection_method

		#Optional
		arguments = var.metric_extension_query_properties_arguments
		auto_row_prefix = var.metric_extension_query_properties_auto_row_prefix
		command = var.metric_extension_query_properties_command
		delimiter = var.metric_extension_query_properties_delimiter
		identity_metric = var.metric_extension_query_properties_identity_metric
		in_param_details {

			#Optional
			in_param_position = var.metric_extension_query_properties_in_param_details_in_param_position
			in_param_value = var.metric_extension_query_properties_in_param_details_in_param_value
		}
		is_metric_service_enabled = var.metric_extension_query_properties_is_metric_service_enabled
		jmx_attributes = var.metric_extension_query_properties_jmx_attributes
		managed_bean_query = var.metric_extension_query_properties_managed_bean_query
		out_param_details {

			#Optional
			out_param_position = var.metric_extension_query_properties_out_param_details_out_param_position
			out_param_type = var.metric_extension_query_properties_out_param_details_out_param_type
		}
		script_details {

			#Optional
			content = var.metric_extension_query_properties_script_details_content
			name = var.metric_extension_query_properties_script_details_name
		}
		sql_details {

			#Optional
			content = var.metric_extension_query_properties_sql_details_content
			script_file_name = var.metric_extension_query_properties_sql_details_script_file_name
		}
		sql_type = var.metric_extension_query_properties_sql_type
		starts_with = var.metric_extension_query_properties_starts_with
	}
	resource_type = var.metric_extension_resource_type

	#Optional
	description = var.metric_extension_description
}
```

## Argument Reference

The following arguments are supported:

* `collection_recurrences` - (Required) (Updatable) Schedule of metric extension should use RFC 5545 format i.e. recur-rule-part = "FREQ";INTERVAL where FREQ rule part identifies the type of recurrence rule. Valid values are "MINUTELY","HOURLY","DAILY" to specify repeating events based on an interval of a minute, an hour and a day or more. Example- FREQ=DAILY;INTERVAL=1
* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `description` - (Optional) (Updatable) Description of the metric extension.
* `display_name` - (Required) (Updatable) Metric Extension display name.
* `metric_list` - (Required) (Updatable) List of metrics which are part of this metric extension
	* `compute_expression` - (Optional) (Updatable) Compute Expression to calculate the value of this metric
	* `data_type` - (Required) (Updatable) Data type of value of this metric
	* `display_name` - (Optional) (Updatable) Display name of the metric.
	* `is_dimension` - (Optional) (Updatable) Current metric need to be included as dimension or not
	* `is_hidden` - (Optional) (Updatable) Flag to marks whether a metric has to be uploaded or not. When isHidden = false -> Metric is uploaded, isHidden = true -> Metric is NOT uploaded
	* `metric_category` - (Optional) (Updatable) Metric category
	* `name` - (Required) (Updatable) Name of the metric.
	* `unit` - (Optional) (Updatable) Unit of metric value
* `name` - (Required) Metric Extension Resource name.
* `query_properties` - (Required) (Updatable) Collection method and query properties details of metric extension
	* `arguments` - (Applicable when collection_method=OS_COMMAND) (Updatable) Arguments required by either command or script
	* `auto_row_prefix` - (Applicable when collection_method=JMX) (Updatable) Prefix for an auto generated metric, in case multiple rows with non unique key values are returned
	* `collection_method` - (Required) (Updatable) Type of possible collection methods.
	* `command` - (Required when collection_method=OS_COMMAND) (Updatable) OS command to execute without arguments
	* `delimiter` - (Required when collection_method=OS_COMMAND) (Updatable) Character used to delimit multiple metric values in single line of output
	* `identity_metric` - (Applicable when collection_method=JMX) (Updatable) Semi-colon separated list of key properties from Managed Bean ObjectName to be used as key metrics
	* `in_param_details` - (Applicable when collection_method=SQL) (Updatable) List of values and position of PL/SQL procedure IN parameters
		* `in_param_position` - (Required when collection_method=SQL) (Updatable) Position of IN parameter
		* `in_param_value` - (Required when collection_method=SQL) (Updatable) Value of IN parameter
	* `is_metric_service_enabled` - (Applicable when collection_method=JMX) (Updatable) Indicates if Metric Service is enabled on server domain
	* `jmx_attributes` - (Required when collection_method=JMX) (Updatable) List of JMX attributes or Metric Service Table columns separated by semi-colon
	* `managed_bean_query` - (Required when collection_method=JMX) (Updatable) JMX Managed Bean Query or Metric Service Table name
	* `out_param_details` - (Applicable when collection_method=SQL) (Updatable) Position and SQL Type of PL/SQL OUT parameter
		* `out_param_position` - (Required when collection_method=SQL) (Updatable) Position of PL/SQL procedure OUT parameter
		* `out_param_type` - (Required when collection_method=SQL) (Updatable) SQL Type of PL/SQL procedure OUT parameter
	* `script_details` - (Applicable when collection_method=OS_COMMAND) (Updatable) Script details applicable to any OS Command based Metric Extension which needs to run a script to collect data
		* `content` - (Required when collection_method=OS_COMMAND) (Updatable) Content of the script file as base64 encoded string
		* `name` - (Required when collection_method=OS_COMMAND) (Updatable) Name of the script file
	* `sql_details` - (Required when collection_method=SQL) (Updatable) Details of Sql content which needs to execute to collect Metric Extension data
		* `content` - (Required when collection_method=SQL) (Updatable) Sql statement or script file content as base64 encoded string
		* `script_file_name` - (Applicable when collection_method=SQL) (Updatable) If a script needs to be executed, then provide file name of the script
	* `sql_type` - (Required when collection_method=SQL) (Updatable) Type of SQL data collection method i.e. either a Statement or SQL Script File
	* `starts_with` - (Applicable when collection_method=OS_COMMAND) (Updatable) String prefix used to identify metric output of the OS Command
* `resource_type` - (Required) Resource type to which Metric Extension applies
* `publish_trigger` - (Optional) (Updatable) An optional property when set to `true` triggers Publish of a metric extension. Once set to `true`, it cannot be changed back to `false`. Update of publish_trigger cannot be combined with other updates in the same request. A metric extension cannot be tested and its definition cannot be updated once it is marked published or publish_trigger is updated to `true`.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `collection_method` - Collection Method  Metric Extension applies
* `collection_recurrences` - Schedule of metric extension should use RFC 5545 format -> recur-rule-part = "FREQ";"INTERVAL" where FREQ rule part identifies the type of recurrence rule. Valid values are "MINUTELY","HOURLY","DAILY" to specify repeating events based on an interval of a minute, an hour and a day or more. Example- FREQ=DAILY;INTERVAL=1
* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `created_by` - Created by user
* `description` - Description of the metric extension.
* `display_name` - Metric Extension resource display name
* `enabled_on_resources` - List of resource objects on which this metric extension is enabled.
	* `resource_id` - The OCID of the resource on which Metric Extension is enabled
* `enabled_on_resources_count` - Count of resources on which this metric extension is enabled.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Metric Extension resource
* `last_updated_by` - Last updated by user
* `metric_list` - List of metrics which are part of this metric extension
	* `compute_expression` - Compute Expression to calculate the value of this metric
	* `data_type` - Data type of value of this metric
	* `display_name` - Display name of the metric.
	* `is_dimension` - Current metric need to be included as dimension or not
	* `is_hidden` - Flag to marks whether a metric has to be uploaded or not. When isHidden = false -> Metric is uploaded, isHidden = true -> Metric is NOT uploaded
	* `metric_category` - Metric category
	* `name` - Name of the metric.
	* `unit` - Unit of metric value
* `name` - Metric Extension resource name
* `query_properties` - Collection method and query properties details of metric extension
	* `arguments` - Arguments required by either command or script
	* `auto_row_prefix` - Prefix for an auto generated metric, in case multiple rows with non unique key values are returned
	* `collection_method` - Type of possible collection methods.
	* `command` - OS command to execute without arguments
	* `delimiter` - Character used to delimit multiple metric values in single line of output
	* `identity_metric` - Semi-colon separated list of key properties from Managed Bean ObjectName to be used as key metrics
	* `in_param_details` - List of values and position of PL/SQL procedure IN parameters
		* `in_param_position` - Position of IN parameter
		* `in_param_value` - Value of IN parameter
	* `is_metric_service_enabled` - Indicates if Metric Service is enabled on server domain
	* `jmx_attributes` - List of JMX attributes or Metric Service Table columns separated by semi-colon
	* `managed_bean_query` - JMX Managed Bean Query or Metric Service Table name
	* `out_param_details` - Position and SQL Type of PL/SQL OUT parameter
		* `out_param_position` - Position of PL/SQL procedure OUT parameter
		* `out_param_type` - SQL Type of PL/SQL procedure OUT parameter
	* `script_details` - Script details applicable to any OS Command based Metric Extension which needs to run a script to collect data
		* `content` - Content of the script file as base64 encoded string
		* `name` - Name of the script file
	* `sql_details` - Details of Sql content which needs to execute to collect Metric Extension data
		* `content` - Sql statement or script file content as base64 encoded string
		* `script_file_name` - If a script needs to be executed, then provide file name of the script
	* `sql_type` - Type of SQL data collection method i.e. either a Statement or SQL Script File
	* `starts_with` - String prefix used to identify metric output of the OS Command
* `resource_type` - Resource type to which Metric Extension applies
* `resource_uri` - The URI path that the user can do a GET on to access the metric extension metadata
* `state` - The current lifecycle state of the metric extension
* `status` - The current status of the metric extension i.e. whether it is Draft or Published
* `tenant_id` - Tenant Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `time_created` - Metric Extension creation time. An RFC3339 formatted datetime string.
* `time_updated` - Metric Extension update time. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Metric Extension
	* `update` - (Defaults to 20 minutes), when updating the Metric Extension
	* `delete` - (Defaults to 20 minutes), when destroying the Metric Extension


## Import

MetricExtensions can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_metric_extension.test_metric_extension "id"
```

