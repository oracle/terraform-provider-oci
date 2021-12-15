---
subcategory: "Service Connector Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_sch_service_connector"
sidebar_current: "docs-oci-resource-sch-service_connector"
description: |-
  Provides the Service Connector resource in Oracle Cloud Infrastructure Service Connector Hub service
---

# oci_sch_service_connector
This resource provides the Service Connector resource in Oracle Cloud Infrastructure Service Connector Hub service.

Creates a new service connector in the specified compartment.
A service connector is a logically defined flow for moving data from
a source service to a destination service in Oracle Cloud Infrastructure.
For instructions, see
[To create a service connector](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create).
For general information about service connectors, see
[Service Connector Hub Overview](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).

For purposes of access control, you must provide the
[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where
you want the service connector to reside. Notice that the service connector
doesn't have to be in the same compartment as the source or target services.
For information about access control and compartments, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

After you send your request, the new service connector's state is temporarily
CREATING. When the state changes to ACTIVE, data begins transferring from the
source service to the target service. For instructions on deactivating and
activating service connectors, see
[To activate or deactivate a service connector](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).


## Example Usage

```hcl
resource "oci_sch_service_connector" "test_service_connector" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.service_connector_display_name
	source {
		#Required
		kind = var.service_connector_source_kind

		#Optional
		cursor {

			#Optional
			kind = var.service_connector_source_cursor_kind
		}
		log_sources {

			#Optional
			compartment_id = var.compartment_id
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
		stream_id = oci_streaming_stream.test_stream.id
	}
	target {
		#Required
		kind = var.service_connector_target_kind

		#Optional
		batch_rollover_size_in_mbs = var.service_connector_target_batch_rollover_size_in_mbs
		batch_rollover_time_in_ms = var.service_connector_target_batch_rollover_time_in_ms
		bucket = var.service_connector_target_bucket
		compartment_id = var.compartment_id
		dimensions {

			#Optional
			dimension_value {
				#Required
				kind = var.service_connector_target_dimensions_dimension_value_kind

				#Optional
				path = var.service_connector_target_dimensions_dimension_value_path
				value = var.service_connector_target_dimensions_dimension_value_value
			}
			name = var.service_connector_target_dimensions_name
		}
		enable_formatted_messaging = var.service_connector_target_enable_formatted_messaging
		function_id = oci_functions_function.test_function.id
		log_group_id = oci_logging_log_group.test_log_group.id
		metric = var.service_connector_target_metric
		metric_namespace = var.service_connector_target_metric_namespace
		namespace = var.service_connector_target_namespace
		object_name_prefix = var.service_connector_target_object_name_prefix
		stream_id = oci_streaming_stream.test_stream.id
		topic_id = oci_ons_notification_topic.test_notification_topic.id
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.service_connector_description
	freeform_tags = {"bar-key"= "value"}
	tasks {
		#Required
		kind = var.service_connector_tasks_kind

		#Optional
		batch_size_in_kbs = var.service_connector_tasks_batch_size_in_kbs
		batch_time_in_sec = var.service_connector_tasks_batch_time_in_sec
		condition = var.service_connector_tasks_condition
		function_id = oci_functions_function.test_function.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the comparment to create the service connector in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The description of the resource. Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `source` - (Required) (Updatable) An object that represents the source of the flow defined by the service connector. An example source is the VCNFlow logs within the NetworkLogs group. For more information about flows defined by service connectors, see [Service Connector Hub Overview](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm). For configuration instructions, see [To create a service connector](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create). 
	* `cursor` - (Applicable when kind=streaming) (Updatable) The type of [cursor](https://docs.cloud.oracle.com/iaas/Content/Streaming/Tasks/using_a_single_consumer.htm#usingcursors), which determines the starting point from which the stream will be consumed. 
		* `kind` - (Required when kind=streaming) (Updatable) The type descriminator. 
	* `kind` - (Required) (Updatable) The type descriminator. 
	* `log_sources` - (Required when kind=logging) (Updatable) The logs for this Logging source. 
		* `compartment_id` - (Required when kind=logging) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the log source. 
		* `log_group_id` - (Applicable when kind=logging) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group. 
		* `log_id` - (Applicable when kind=logging) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log. 
	* `stream_id` - (Required when kind=streaming) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream. 
* `target` - (Required) (Updatable) An object that represents the target of the flow defined by the service connector. An example target is a stream (Streaming service). For more information about flows defined by service connectors, see [Service Connector Hub Overview](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm). For configuration instructions, see [To create a service connector](https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create). 
	* `batch_rollover_size_in_mbs` - (Applicable when kind=objectStorage) (Updatable) The batch rollover size in megabytes. 
	* `batch_rollover_time_in_ms` - (Applicable when kind=objectStorage) (Updatable) The batch rollover time in milliseconds. 
	* `bucket` - (Required when kind=objectStorage) (Updatable) The name of the bucket. Avoid entering confidential information. 
	* `compartment_id` - (Required when kind=monitoring) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric. 
	* `dimensions` - (Applicable when kind=monitoring) (Updatable) List of dimension names and values. 
		* `dimension_value` - (Required when kind=monitoring) (Updatable) Instructions for extracting the value corresponding to the specified dimension key: Either extract the value as-is (static) or derive the value from a path (evaluated). 
			* `kind` - (Required) (Updatable) The type of dimension value: static or evaluated. 
			* `path` - (Required when kind=jmesPath) (Updatable) The location to use for deriving the dimension value (evaluated). The path must start with `logContent` in an acceptable notation style with supported [JMESPath selectors](https://jmespath.org/specification.html): expression with dot and index operator (`.`, and `[]`). Example with dot notation: `logContent.data` Example with index notation: `logContent.data[0].content` For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). The returned value depends on the results of evaluation. If the evaluated value is valid, then the evaluated value is returned without double quotes. (Any front or trailing double quotes are trimmed before returning the value. For example, the evaluated value `"compartmentId"` is returned as `compartmentId`.) If the evaluated value is invalid, then the returned value is `SCH_EVAL_INVALID_VALUE`. If the evaluated value is empty, then the returned value is `SCH_EVAL_VALUE_EMPTY`. 
			* `value` - (Required when kind=static) (Updatable) The data extracted from the specified dimension value (passed as-is). Unicode characters only. For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). 
		* `name` - (Required when kind=monitoring) (Updatable) Dimension key. A valid dimension key includes only printable ASCII, excluding periods (.) and spaces. Custom dimension keys are acceptable. Avoid entering confidential information. Due to use by Service Connector Hub, the following dimension names are reserved: `connectorId`, `connectorName`, `connectorSourceType`. For information on valid dimension keys and values, see [MetricDataDetails Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/MetricDataDetails). Example: `type` 
	* `enable_formatted_messaging` - (Applicable when kind=notifications) (Updatable) Whether to apply a simplified, user-friendly format to the message. Applies only when friendly formatting is supported by the service connector source and the subscription protocol.  Example: `true` 
	* `function_id` - (Required when kind=functions) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function. 
	* `kind` - (Required) (Updatable) The type descriminator. 
	* `log_group_id` - (Required when kind=loggingAnalytics) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Logging Analytics log group. 
	* `metric` - (Required when kind=monitoring) (Updatable) The name of the metric.  Example: `CpuUtilization` 
	* `metric_namespace` - (Required when kind=monitoring) (Updatable) The namespace of the metric.  Example: `oci_computeagent` 
	* `namespace` - (Applicable when kind=objectStorage) (Updatable) The namespace. 
	* `object_name_prefix` - (Applicable when kind=objectStorage) (Updatable) The prefix of the objects. Avoid entering confidential information. 
	* `stream_id` - (Required when kind=streaming) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream. 
	* `topic_id` - (Required when kind=notifications) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic. 
* `tasks` - (Optional) (Updatable) The list of tasks. 
	* `batch_size_in_kbs` - (Applicable when kind=function) (Updatable) Size limit (kilobytes) for batch sent to invoke the function. 
	* `batch_time_in_sec` - (Applicable when kind=function) (Updatable) Time limit (seconds) for batch sent to invoke the function. 
	* `condition` - (Required when kind=logRule) (Updatable) A filter or mask to limit the source used in the flow defined by the service connector. 
	* `function_id` - (Required when kind=function) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function to be used as a task. 
	* `kind` - (Required) (Updatable) The type descriminator. 
* `state` - (Optional) (Updatable) The target state for the service connector. Could be set to `ACTIVE` or `INACTIVE`.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Service Connector
	* `update` - (Defaults to 20 minutes), when updating the Service Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Service Connector


## Import

ServiceConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_sch_service_connector.test_service_connector "id"
```

