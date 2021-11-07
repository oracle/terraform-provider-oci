---
subcategory: "Service Connector Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_sch_service_connectors"
sidebar_current: "docs-oci-datasource-sch-service_connectors"
description: |-
  Provides the list of Service Connectors in Oracle Cloud Infrastructure Service Connector Hub service
---

# Data Source: oci_sch_service_connectors
This data source provides the list of Service Connectors in Oracle Cloud Infrastructure Service Connector Hub service.

Lists service connectors in the specified compartment.


## Example Usage

```hcl
data "oci_sch_service_connectors" "test_service_connectors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.service_connector_display_name
	state = var.service_connector_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for this request. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.  Example: `example_service_connector` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `ACTIVE` 


## Attributes Reference

The following attributes are exported:

* `service_connector_collection` - The list of service_connector_collection.

### ServiceConnector Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the service connector. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service connector. 
* `lifecyle_details` - A message describing the current state in more detail. For example, the message might provide actionable information for a resource in a `FAILED` state. 
* `source` - An object that represents the source of the flow defined by the service connector. An example source is the VCNFlow logs within the NetworkLogs group. For more information about flows defined by service connectors, see [Service Connector Hub Overview](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm). For configuration instructions, see [To create a service connector](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create). 
	* `cursor` - The type of [cursor](https://docs.cloud.oracle.com/iaas/Content/Streaming/Tasks/using_a_single_consumer.htm#usingcursors), which determines the starting point from which the stream will be consumed. 
		* `kind` - The type descriminator. 
	* `kind` - The type descriminator. 
	* `log_sources` - The logs for this Logging source. 
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the log source. 
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group. 
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log. 
	* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream. 
* `state` - The current state of the service connector. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `target` - An object that represents the target of the flow defined by the service connector. An example target is a stream (Streaming service). For more information about flows defined by service connectors, see [Service Connector Hub Overview](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm). For configuration instructions, see [To create a service connector](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create). 
	* `batch_rollover_size_in_mbs` - The batch rollover size in megabytes. 
	* `batch_rollover_time_in_ms` - The batch rollover time in milliseconds. 
	* `bucket` - The name of the bucket. Avoid entering confidential information. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric. 
	* `dimensions` - List of dimension names and values. 
		* `dimension_value` - Instructions for extracting the value corresponding to the specified dimension key: Either extract the value as-is (static) or derive the value from a path (evaluated). 
			* `kind` - The type of dimension value: static or evaluated. 
			* `path` - The location to use for deriving the dimension value (evaluated). The path must start with `logContent` in an acceptable notation style with supported [JMESPath selectors](https://jmespath.org/specification.html): expression with dot and index operator (`.`, and `[]`). Example with dot notation: `logContent.data` Example with index notation: `logContent.data[0].content` For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). The returned value depends on the results of evaluation. If the evaluated value is valid, then the evaluated value is returned without double quotes. (Any front or trailing double quotes are trimmed before returning the value. For example, the evaluated value `"compartmentId"` is returned as `compartmentId`.) If the evaluated value is invalid, then the returned value is `SCH_EVAL_INVALID_VALUE`. If the evaluated value is empty, then the returned value is `SCH_EVAL_VALUE_EMPTY`. 
			* `value` - The data extracted from the specified dimension value (passed as-is). Unicode characters only. For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). 
		* `name` - Dimension key. A valid dimension key includes only printable ASCII, excluding periods (.) and spaces. Custom dimension keys are acceptable. Avoid entering confidential information. Due to use by Service Connector Hub, the following dimension names are reserved: `connectorId`, `connectorName`, `connectorSourceType`. For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). Example: `type` 
	* `enable_formatted_messaging` - Whether to apply a simplified, user-friendly format to the message. Applies only when friendly formatting is supported by the service connector source and the subscription protocol.  Example: `true` 
	* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function. 
	* `kind` - The type descriminator. 
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Logging Analytics log group. 
	* `metric` - The name of the metric.  Example: `CpuUtilization` 
	* `metric_namespace` - The namespace of the metric.  Example: `oci_computeagent` 
	* `namespace` - The namespace. 
	* `object_name_prefix` - The prefix of the objects. Avoid entering confidential information. 
	* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream. 
	* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic. 
* `tasks` - The list of tasks. 
	* `batch_size_in_kbs` - Size limit (kilobytes) for batch sent to invoke the function. 
	* `batch_time_in_sec` - Time limit (seconds) for batch sent to invoke the function. 
	* `condition` - A filter or mask to limit the source used in the flow defined by the service connector. 
	* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function to be used as a task. 
	* `kind` - The type descriminator. 
* `time_created` - The date and time when the service connector was created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 
* `time_updated` - The date and time when the service connector was updated. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 

