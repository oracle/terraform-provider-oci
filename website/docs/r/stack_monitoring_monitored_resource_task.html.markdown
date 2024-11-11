---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource_task"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resource_task"
description: |-
  Provides the Monitored Resource Task resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resource_task
This resource provides the Monitored Resource Task resource in Oracle Cloud Infrastructure Stack Monitoring service.

Create a new stack monitoring resource task.

## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resource_task" "test_monitored_resource_task" {
	#Required
	compartment_id = var.compartment_id
	task_details {
		#Required
		type = var.monitored_resource_task_task_details_type

		#Optional
		agent_id = oci_cloud_bridge_agent.test_agent.id
		availability_proxy_metric_collection_interval = var.monitored_resource_task_task_details_availability_proxy_metric_collection_interval
		availability_proxy_metrics = var.monitored_resource_task_task_details_availability_proxy_metrics
		console_path_prefix = var.monitored_resource_task_task_details_console_path_prefix
		external_id_mapping = var.monitored_resource_task_task_details_external_id_mapping
		handler_type = var.monitored_resource_task_task_details_handler_type
		is_enable = var.monitored_resource_task_task_details_is_enable
		lifecycle_status_mappings_for_up_status = var.monitored_resource_task_task_details_lifecycle_status_mappings_for_up_status
		namespace = var.monitored_resource_task_task_details_namespace
		receiver_properties {

			#Optional
			listener_port = var.monitored_resource_task_task_details_receiver_properties_listener_port
		}
		resource_group = var.monitored_resource_task_task_details_resource_group
		resource_name_filter = var.monitored_resource_task_task_details_resource_name_filter
		resource_name_mapping = var.monitored_resource_task_task_details_resource_name_mapping
		resource_type_filter = var.monitored_resource_task_task_details_resource_type_filter
		resource_type_mapping = var.monitored_resource_task_task_details_resource_type_mapping
		resource_types_configuration {

			#Optional
			availability_metrics_config {

				#Optional
				collection_interval_in_seconds = var.monitored_resource_task_task_details_resource_types_configuration_availability_metrics_config_collection_interval_in_seconds
				metrics = var.monitored_resource_task_task_details_resource_types_configuration_availability_metrics_config_metrics
			}
			handler_config {

				#Optional
				collectd_resource_name_config {

					#Optional
					exclude_properties = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_collectd_resource_name_config_exclude_properties
					include_properties = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_collectd_resource_name_config_include_properties
					suffix = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_collectd_resource_name_config_suffix
				}
				collector_types = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_collector_types
				handler_properties {

					#Optional
					name = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_handler_properties_name
					value = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_handler_properties_value
				}
				metric_mappings {

					#Optional
					collector_metric_name = oci_monitoring_metric.test_metric.name
					is_skip_upload = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_metric_mappings_is_skip_upload
					metric_upload_interval_in_seconds = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_metric_mappings_metric_upload_interval_in_seconds
					telemetry_metric_name = oci_monitoring_metric.test_metric.name
				}
				metric_name_config {

					#Optional
					exclude_pattern_on_prefix = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_metric_name_config_exclude_pattern_on_prefix
					is_prefix_with_collector_type = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_metric_name_config_is_prefix_with_collector_type
				}
				metric_upload_interval_in_seconds = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_metric_upload_interval_in_seconds
				telegraf_resource_name_config {

					#Optional
					exclude_tags = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_telegraf_resource_name_config_exclude_tags
					include_tags = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_telegraf_resource_name_config_include_tags
					is_use_tags_only = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_telegraf_resource_name_config_is_use_tags_only
				}
				telemetry_resource_group = var.monitored_resource_task_task_details_resource_types_configuration_handler_config_telemetry_resource_group
			}
			resource_type = var.monitored_resource_task_task_details_resource_types_configuration_resource_type
		}
		service_base_url = var.monitored_resource_task_task_details_service_base_url
		should_use_metrics_flow_for_status = var.monitored_resource_task_task_details_should_use_metrics_flow_for_status
		source = var.monitored_resource_task_task_details_source
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	name = var.monitored_resource_task_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment identifier. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Optional) Name of the task. If not provided by default the following names will be taken Oracle Cloud Infrastructure tasks - namespace plus timestamp. 
* `task_details` - (Required) The request details for the performing the task.
	* `agent_id` - (Required when type=UPDATE_AGENT_RECEIVER) Management Agent Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `availability_proxy_metric_collection_interval` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) Metrics collection interval in seconds used when calculating the availability of the  resource based on metrics specified using the property 'availabilityProxyMetrics'. 
	* `availability_proxy_metrics` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) List of metrics to be used to calculate the availability of the resource. Resource is considered to be up if at least one of the specified metrics is available for  the resource during the specified interval using the property  'availabilityProxyMetricCollectionInterval'. If no metrics are specified, availability will not be calculated for the resource. 
	* `console_path_prefix` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The console path prefix to use for providing service home url page navigation.  For example if the prefix provided is 'security/bastion/bastions', the URL used for navigation will be https://<cloudhostname>/security/bastion/bastions/<resourceOcid>. If not provided, service home page link  will not be shown in the stack monitoring home page. 
	* `external_id_mapping` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The external resource identifier property in the metric dimensions.  Resources imported will be using this property value for external id. 
	* `handler_type` - (Required when type=UPDATE_AGENT_RECEIVER | UPDATE_RESOURCE_TYPE_CONFIGS) Type of the handler.
	* `is_enable` - (Required when type=UPDATE_AGENT_RECEIVER) True to enable the receiver and false to disable the receiver on the agent. 
	* `lifecycle_status_mappings_for_up_status` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) Lifecycle states of the external resource which reflects the status of the resource being up. 
	* `namespace` - (Required when type=IMPORT_OCI_TELEMETRY_RESOURCES) Name space to be used for Oracle Cloud Infrastructure Native service resources discovery.
	* `receiver_properties` - (Applicable when type=UPDATE_AGENT_RECEIVER) Properties for agent receiver.
		* `listener_port` - (Required when type=UPDATE_AGENT_RECEIVER) Receiver listener port.
	* `resource_group` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The resource group to use while fetching metrics from telemetry. If not specified, resource group will be skipped in the list metrics request. 
	* `resource_name_filter` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The resource name filter. Resources matching with the resource name filter will be imported. Regular expressions will be accepted. 
	* `resource_name_mapping` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The resource name property in the metric dimensions.  Resources imported will be using this property value for resource name. 
	* `resource_type_filter` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The resource type filter. Resources matching with the resource type filter will be imported. Regular expressions will be accepted. 
	* `resource_type_mapping` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The resource type property in the metric dimensions.  Resources imported will be using this property value for resource type. If not specified, namespace will be used for resource type. 
	* `resource_types_configuration` - (Required when type=UPDATE_RESOURCE_TYPE_CONFIGS) A collection of resource type configuration details. User can provide  availability proxy metrics list for resource types along with the  telegraf/collectd handler configuration for the resource types. 
		* `availability_metrics_config` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Availability metrics details.
			* `collection_interval_in_seconds` - (Required when type=UPDATE_RESOURCE_TYPE_CONFIGS) Availability metric collection internal in seconds.
			* `metrics` - (Required when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of metrics used for availability calculation for the resource.
		* `handler_config` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Specific resource mapping configurations for Agent Extension Handlers.
			* `collectd_resource_name_config` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Resource name generation overriding configurations for collectd resource types. 
				* `exclude_properties` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of property names to be excluded.
				* `include_properties` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of property names to be included.
				* `suffix` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) String to be suffixed to the resource name.
			* `collector_types` - (Required when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of collector/plugin names.
			* `handler_properties` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of handler configuration properties
				* `name` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Property name.
				* `value` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Property value.
			* `metric_mappings` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of AgentExtensionHandlerMetricMappingDetails.
				* `collector_metric_name` - (Required when type=UPDATE_RESOURCE_TYPE_CONFIGS) Metric name as defined by the collector.
				* `is_skip_upload` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Is ignoring this metric.
				* `metric_upload_interval_in_seconds` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Metric upload interval in seconds. Any metric sent by telegraf/collectd before the  configured interval expires will be dropped. 
				* `telemetry_metric_name` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Metric name to be upload to telemetry.
			* `metric_name_config` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Metric name generation overriding configurations.
				* `exclude_pattern_on_prefix` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) String pattern to be removed from the prefix of the metric name.
				* `is_prefix_with_collector_type` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) is prefixing the metric with collector type.
			* `metric_upload_interval_in_seconds` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Metric upload interval in seconds. Any metric sent by telegraf/collectd before the  configured interval expires will be dropped. 
			* `telegraf_resource_name_config` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Resource name generation overriding configurations for telegraf resource types. 
				* `exclude_tags` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of tag names to be excluded.
				* `include_tags` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) List of tag names to be included.
				* `is_use_tags_only` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Flag to indicate if only tags will be used for resource name generation.
			* `telemetry_resource_group` - (Applicable when type=UPDATE_RESOURCE_TYPE_CONFIGS) Resource group string; if not specified, the resource group string will be generated by the handler.
		* `resource_type` - (Required when type=UPDATE_RESOURCE_TYPE_CONFIGS) Resource type.
	* `service_base_url` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) The base URL of the Oracle Cloud Infrastructure service to which the resource belongs to. Also this property is applicable only when source is OCI_TELEMETRY_NATIVE. 
	* `should_use_metrics_flow_for_status` - (Applicable when type=IMPORT_OCI_TELEMETRY_RESOURCES) Flag to indicate whether status is calculated using metrics or  LifeCycleState attribute of the resource in Oracle Cloud Infrastructure service. 
	* `source` - (Required when type=IMPORT_OCI_TELEMETRY_RESOURCES) Source from where the metrics pushed to telemetry. Possible values:
		* OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from Oracle Cloud Infrastructure Native Services.
		* OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus.
		* OCI_TELEMETRY_TELEGRAF    - The metrics are pushed to telemetry from Telegraf receiver.
		* OCI_TELEMETRY_COLLECTD    - The metrics are pushed to telemetry from CollectD receiver. 
	* `type` - (Required) Task type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
	* `namespace` - Name space to be used for Oracle Cloud Infrastructure Native service resources discovery.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resource Task
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resource Task
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resource Task


## Import

MonitoredResourceTasks can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task "id"
```

