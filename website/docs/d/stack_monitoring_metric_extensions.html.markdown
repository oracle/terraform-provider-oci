---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_metric_extensions"
sidebar_current: "docs-oci-datasource-stack_monitoring-metric_extensions"
description: |-
  Provides the list of Metric Extensions in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_metric_extensions
This data source provides the list of Metric Extensions in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of metric extensions

## Example Usage

```hcl
data "oci_stack_monitoring_metric_extensions" "test_metric_extensions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	enabled_on_resource_id = oci_usage_proxy_resource.test_resource.id
	name = var.metric_extension_name
	resource_type = var.metric_extension_resource_type
	state = var.metric_extension_state
	status = var.metric_extension_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which data is listed.
* `enabled_on_resource_id` - (Optional) A filter to return metric extensions based on input resource Id on which metric extension is enabled
* `name` - (Optional) A filter to return resources based on name.
* `resource_type` - (Optional) A filter to return resources based on resource type.
* `state` - (Optional) A filter to return metric extensions based on Lifecycle State
* `status` - (Optional) A filter to return resources based on status e.g. Draft or Published


## Attributes Reference

The following attributes are exported:

* `metric_extension_collection` - The list of metric_extension_collection.

### MetricExtension Reference

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

