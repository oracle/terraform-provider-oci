---
subcategory: "Service Connector Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_sch_service_connector"
sidebar_current: "docs-oci-datasource-sch-service_connector"
description: |-
  Provides details about a specific Service Connector in Oracle Cloud Infrastructure Service Connector Hub service
---

# Data Source: oci_sch_service_connector
This data source provides details about a specific Service Connector resource in Oracle Cloud Infrastructure Service Connector Hub service.

Gets the specified connector's configuration information.
For more information, see
[Getting a Connector](https://docs.cloud.oracle.com/iaas/Content/connector-hub/get-service-connector.htm).


## Example Usage

```hcl
data "oci_sch_service_connector" "test_service_connector" {
	#Required
	service_connector_id = oci_sch_service_connector.test_service_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `service_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connector. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the connector. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connector. 
* `lifecycle_details` - A message describing the current state in more detail. For example, the message might provide actionable information for a resource in a `FAILED` state. 
* `lifecyle_details` - *Please note this property is deprecated and will be removed on January 27, 2026. Use `lifecycleDetails` instead.* A message describing the current state in more detail. For example, the message might provide actionable information for a resource in a `FAILED` state. 
* `source` - 
	* `config_map` - The configuration map for the connector plugin. This map includes parameters specific to the connector plugin type.  For example, for `QueueSource`, the map lists the OCID of the selected queue. To find the parameters for a connector plugin, get the plugin using [GetConnectorPlugin](https://docs.cloud.oracle.com/iaas/api/#/en/serviceconnectors/latest/ConnectorPlugin/GetConnectorPlugin) and review its schema value. 
	* `cursor` - The [read setting](https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector-streaming-source.htm), which determines where in the stream to start moving data. For configuration instructions, see [Creating a Connector with a Streaming Source](https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector-streaming-source.htm). 
		* `kind` - The type discriminator. 
	* `log_sources` - The logs for this Logging source. 
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the log source. 
		* `log_group_id` - Identifier of the log group. Either `_Audit` or the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group. Note: For the Notifications target, only `_Audit` is allowed. 
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log. 
	* `monitoring_sources` - One or more compartment-specific lists of metric namespaces to retrieve data from. 
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a compartment containing metric namespaces you want to use for the Monitoring source. 
		* `namespace_details` - Discriminator for namespaces in the compartment-specific list. 
			* `kind` - The type discriminator. 
			* `namespaces` - The namespaces for the compartment-specific list. 
				* `metrics` - The metrics to query for the specified metric namespace. 
					* `kind` - The type discriminator. 
				* `namespace` - The source service or application to use when querying for metric data points. Must begin with `oci_`.  Example: `oci_computeagent` 
	* `plugin_name` - The name of the connector plugin. This name indicates the service to be called by the connector plugin. For example, `QueueSource` indicates the Queue service. To find names of connector plugins, list the plugin using [ListConnectorPlugin](https://docs.cloud.oracle.com/iaas/api/#/en/serviceconnectors/latest/ConnectorPluginSummary/ListConnectorPlugins). 
	* `private_endpoint_metadata` - The private endpoint metadata for the connector's source or target. 
		* `rce_dns_proxy_ip_address` - The reverse connection endpoint (RCE) IP address for DNS lookups. 
		* `rce_traffic_ip_address` - The reverse connection endpoint (RCE) IP address for primary flow of traffic in the subnet. 
	* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream. 
* `state` - The current state of the connector. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `target` - 
	* `batch_rollover_size_in_mbs` - The batch rollover size in megabytes. 
	* `batch_rollover_time_in_ms` - The batch rollover time in milliseconds. 
	* `batch_size_in_kbs` - The batch rollover size in kilobytes. Only one size option can be specified: `batchSizeInKbs` or `batchSizeInNum`. 
	* `batch_size_in_num` - The batch rollover size in number of messages. Only one size option can be specified: `batchSizeInKbs` or `batchSizeInNum`. 
	* `batch_time_in_sec` - The batch rollover time in seconds. 
	* `bucket` - The name of the bucket. Valid characters are letters (upper or lower case), numbers, hyphens (-), underscores(_), and periods (.). Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. Example: my-new-bucket1 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric. 
	* `dimensions` - List of dimension names and values. 
		* `dimension_value` - Instructions for extracting the value corresponding to the specified dimension key: Either extract the value as-is (static) or derive the value from a path (evaluated). 
			* `kind` - The type of dimension value: static or evaluated. 
			* `path` - The location to use for deriving the dimension value (evaluated). The path must start with `logContent` in an acceptable notation style with supported [JMESPath selectors](https://jmespath.org/specification.html): expression with dot and index operator (`.` and `[]`). Example with dot notation: `logContent.data` Example with index notation: `logContent.data[0].content` For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). The returned value depends on the results of evaluation. If the evaluated value is valid, then the evaluated value is returned without double quotes. (Any front or trailing double quotes are trimmed before returning the value. For example, the evaluated value `"compartmentId"` is returned as `compartmentId`.) If the evaluated value is invalid, then the returned value is `SCH_EVAL_INVALID_VALUE`. If the evaluated value is empty, then the returned value is `SCH_EVAL_VALUE_EMPTY`. 
			* `value` - The data extracted from the specified dimension value (passed as-is). Unicode characters only. For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). 
		* `name` - Dimension key. A valid dimension key includes only printable ASCII, excluding periods (.) and spaces. Custom dimension keys are acceptable. Avoid entering confidential information. Due to use by Connector Hub, the following dimension names are reserved: `connectorId`, `connectorName`, `connectorSourceType`. For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). Example: `type` 
	* `enable_formatted_messaging` - Whether to apply a simplified, user-friendly format to the message. Applies only when friendly formatting is supported by the connector source and the subscription protocol. Example: `true` 
	* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function. 
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Logging Analytics log group. 
	* `log_source_identifier` - Identifier of the log source that you want to use for processing data received from the connector source. Applies to `StreamingSource` only. Equivalent to `name` at [LogAnalyticsSource](https://docs.cloud.oracle.com/iaas/api/#/en/logan-api-spec/latest/LogAnalyticsSource/). 
	* `metric` - The name of the metric. Example: `CpuUtilization` 
	* `metric_namespace` - The namespace of the metric. Example: `oci_computeagent` 
	* `namespace` - The namespace. 
	* `object_name_prefix` - The prefix of the objects. Avoid entering confidential information. 
	* `private_endpoint_metadata` - The private endpoint metadata for the connector's source or target. 
		* `rce_dns_proxy_ip_address` - The reverse connection endpoint (RCE) IP address for DNS lookups. 
		* `rce_traffic_ip_address` - The reverse connection endpoint (RCE) IP address for primary flow of traffic in the subnet. 
	* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream. 
	* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic. 
* `tasks` - The list of tasks. 
	* `batch_size_in_kbs` - Size limit (kilobytes) for batch sent to invoke the function. 
	* `batch_time_in_sec` - Time limit (seconds) for batch sent to invoke the function. 
	* `condition` - A filter or mask to limit the source used in the flow defined by the connector. 
	* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function to be used as a task. 
	* `private_endpoint_metadata` - The private endpoint metadata for the connector's source or target. 
		* `rce_dns_proxy_ip_address` - The reverse connection endpoint (RCE) IP address for DNS lookups. 
		* `rce_traffic_ip_address` - The reverse connection endpoint (RCE) IP address for primary flow of traffic in the subnet. 
* `time_created` - The date and time when the connector was created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 
* `time_updated` - The date and time when the connector was updated. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 

