---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource_task"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitored_resource_task"
description: |-
  Provides details about a specific Monitored Resource Task in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitored_resource_task
This data source provides details about a specific Monitored Resource Task resource in Oracle Cloud Infrastructure Stack Monitoring service.

Gets stack monitoring resource task details by identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

## Example Usage

```hcl
data "oci_stack_monitoring_monitored_resource_task" "test_monitored_resource_task" {
	#Required
	monitored_resource_task_id = oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task.id
}
```

## Argument Reference

The following arguments are supported:

* `monitored_resource_task_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of stack monitoring resource task.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment identifier. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Task identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - Name of the task.
* `state` - The current state of the stack monitoring resource task.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `task_details` - The request details for the performing the task.
	* `agent_id` - Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `availability_proxy_metric_collection_interval` - Metrics collection interval in seconds used when calculating the availability of the  resource based on metrics specified using the property 'availabilityProxyMetrics'. 
	* `availability_proxy_metrics` - List of metrics to be used to calculate the availability of the resource. Resource is considered to be up if at least one of the specified metrics is available for  the resource during the specified interval using the property  'availabilityProxyMetricCollectionInterval'. If no metrics are specified, availability will not be calculated for the resource. 
	* `console_path_prefix` - The console path prefix to use for providing service home url page navigation.  For example if the prefix provided is 'security/bastion/bastions', the URL used for navigation will be https://<cloudhostname>/security/bastion/bastions/<resourceOcid>. If not provided, service home page link  will not be shown in the stack monitoring home page. 
	* `external_id_mapping` - The external resource identifier property in the metric dimensions.  Resources imported will be using this property value for external id. 
	* `handler_type` - Type of the handler.
	* `is_enable` - True to enable the receiver and false to disable the receiver on the agent. 
	* `lifecycle_status_mappings_for_up_status` - Lifecycle states of the external resource which reflects the status of the resource being up. 
	* `namespace` - Name space to be used for Oracle Cloud Infrastructure Native service resources' import.
	* `receiver_properties` - Properties for agent receiver.
		* `listener_port` - Receiver listener port.
	* `resource_group` - The resource group to use while fetching metrics from telemetry. If not specified, resource group will be skipped in the list metrics request. 
	* `resource_name_filter` - The resource name filter. Resources matching with the resource name filter will be imported. Regular expressions will be accepted. 
	* `resource_name_mapping` - The resource name property in the metric dimensions.  Resources imported will be using this property value for resource name. 
	* `resource_type_filter` - The resource type filter. Resources matching with the resource type filter will be imported. Regular expressions will be accepted. 
	* `resource_type_mapping` - The resource type property in the metric dimensions.  Resources imported will be using this property value for resource type. If not specified, namespace will be used for resource type. 
	* `resource_types_configuration` - A collection of resource type configuration details. User can provide  availability proxy metrics list for resource types along with the  telegraf/collectd handler configuration for the resource types. 
		* `availability_metrics_config` - Availability metrics details.
			* `collection_interval_in_seconds` - Availability metric collection internal in seconds.
			* `metrics` - List of metrics used for availability calculation for the resource.
		* `handler_config` - Specific resource mapping configurations for Agent Extension Handlers.
			* `collectd_resource_name_config` - Resource name generation overriding configurations for collectd resource types. 
				* `exclude_properties` - List of property names to be excluded.
				* `include_properties` - List of property names to be included.
				* `suffix` - String to be suffixed to the resource name.
			* `collector_types` - List of collector/plugin names.
			* `handler_properties` - List of handler configuration properties
				* `name` - Property name.
				* `value` - Property value.
			* `metric_mappings` - List of AgentExtensionHandlerMetricMappingDetails.
				* `collector_metric_name` - Metric name as defined by the collector.
				* `is_skip_upload` - Is ignoring this metric.
				* `metric_upload_interval_in_seconds` - Metric upload interval in seconds. Any metric sent by telegraf/collectd before the  configured interval expires will be dropped. 
				* `telemetry_metric_name` - Metric name to be upload to telemetry.
			* `metric_name_config` - Metric name generation overriding configurations.
				* `exclude_pattern_on_prefix` - String pattern to be removed from the prefix of the metric name.
				* `is_prefix_with_collector_type` - is prefixing the metric with collector type.
			* `metric_upload_interval_in_seconds` - Metric upload interval in seconds. Any metric sent by telegraf/collectd before the  configured interval expires will be dropped. 
			* `telegraf_resource_name_config` - Resource name generation overriding configurations for telegraf resource types. 
				* `exclude_tags` - List of tag names to be excluded.
				* `include_tags` - List of tag names to be included.
				* `is_use_tags_only` - Flag to indicate if only tags will be used for resource name generation.
			* `telemetry_resource_group` - Resource group string; if not specified, the resource group string will be generated by the handler.
		* `resource_type` - Resource type.
	* `service_base_url` - The base URL of the Oracle Cloud Infrastructure service to which the resource belongs to. Also this property is applicable only when source is OCI_TELEMETRY_NATIVE. 
	* `should_use_metrics_flow_for_status` - Flag to indicate whether status is calculated using metrics or  LifeCycleState attribute of the resource in Oracle Cloud Infrastructure service. 
	* `source` - Source from where the metrics pushed to telemetry. Possible values:
		* OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from Oracle Cloud Infrastructure Native Services.
		* OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus.
		* OCI_TELEMETRY_TELEGRAF    - The metrics are pushed to telemetry from Telegraf receiver.
		* OCI_TELEMETRY_COLLECTD    - The metrics are pushed to telemetry from CollectD receiver. 
	* `type` - Task type.
* `tenant_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy. 
* `time_created` - The date and time when the stack monitoring resource task was created, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time when the stack monitoring resource task was last updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `type` - Type of the task.
* `work_request_ids` - Identifiers [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for work requests submitted for this task. 

