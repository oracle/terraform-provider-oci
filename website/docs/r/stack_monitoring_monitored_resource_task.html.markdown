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
		namespace = var.monitored_resource_task_task_details_namespace
		source = var.monitored_resource_task_task_details_source
		type = var.monitored_resource_task_task_details_type

		#Optional
		availability_proxy_metric_collection_interval = var.monitored_resource_task_task_details_availability_proxy_metric_collection_interval
		availability_proxy_metrics = var.monitored_resource_task_task_details_availability_proxy_metrics
		console_path_prefix = var.monitored_resource_task_task_details_console_path_prefix
		external_id_mapping = var.monitored_resource_task_task_details_external_id_mapping
		lifecycle_status_mappings_for_up_status = var.monitored_resource_task_task_details_lifecycle_status_mappings_for_up_status
		resource_group = var.monitored_resource_task_task_details_resource_group
		resource_name_filter = var.monitored_resource_task_task_details_resource_name_filter
		resource_name_mapping = var.monitored_resource_task_task_details_resource_name_mapping
		resource_type_filter = var.monitored_resource_task_task_details_resource_type_filter
		resource_type_mapping = var.monitored_resource_task_task_details_resource_type_mapping
		service_base_url = var.monitored_resource_task_task_details_service_base_url
		should_use_metrics_flow_for_status = var.monitored_resource_task_task_details_should_use_metrics_flow_for_status
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
	* `availability_proxy_metric_collection_interval` - (Optional) Metrics collection interval in seconds used when calculating the availability of the  resource based on metrics specified using the property 'availabilityProxyMetrics'. 
	* `availability_proxy_metrics` - (Optional) List of metrics to be used to calculate the availability of the resource. Resource is considered to be up if at least one of the specified metrics is available for the resource during the specified interval using the property 'availabilityProxyMetricCollectionIntervalInSeconds'. If no metrics are specified, availability will not be calculated for the resource. 
	* `console_path_prefix` - (Optional) The console path prefix to use for providing service home url page navigation.  For example if the prefix provided is 'security/bastion/bastions', the URL used for navigation will be https://<cloudhostname>/security/bastion/bastions/<resourceOcid>. If not provided, service home page link  will not be shown in the stack monitoring home page. 
	* `external_id_mapping` - (Optional) The external resource identifier property in the metric dimensions.  Resources imported will be using this property value for external id. 
	* `lifecycle_status_mappings_for_up_status` - (Optional) Lifecycle states of the external resource which reflects the status of the resource being up. 
	* `namespace` - (Required) Name space to be used for Oracle Cloud Infrastructure Native service resources discovery.
	* `resource_group` - (Optional) The resource group to use while fetching metrics from telemetry. If not specified, resource group will be skipped in the list metrics request. 
	* `resource_name_filter` - (Optional) The resource name filter. Resources matching with the resource name filter will be imported. Regular expressions will be accepted. 
	* `resource_name_mapping` - (Optional) The resource name property in the metric dimensions.  Resources imported will be using this property value for resource name. 
	* `resource_type_filter` - (Optional) The resource type filter. Resources matching with the resource type filter will be imported. Regular expressions will be accepted. 
	* `resource_type_mapping` - (Optional) The resource type property in the metric dimensions.  Resources imported will be using this property value for resource type. If not specified, namespace will be used for resource type. 
	* `service_base_url` - (Optional) The base URL of the Oracle Cloud Infrastructure service to which the resource belongs to. Also this property is applicable only when source is OCI_TELEMETRY_NATIVE. 
	* `should_use_metrics_flow_for_status` - (Optional) Flag to indicate whether status is calculated using metrics or  LifeCycleState attribute of the resource in Oracle Cloud Infrastructure service. 
	* `source` - (Required) Source from where the metrics pushed to telemetry. Possible values:
		* OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from Oracle Cloud Infrastructure Native Services.
		* OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus. 
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
	* `availability_proxy_metric_collection_interval` - Metrics collection interval in seconds used when calculating the availability of the  resource based on metrics specified using the property 'availabilityProxyMetrics'. 
	* `availability_proxy_metrics` - List of metrics to be used to calculate the availability of the resource. Resource is considered to be up if at least one of the specified metrics is available for the resource during the specified interval using the property 'availabilityProxyMetricCollectionIntervalInSeconds'. If no metrics are specified, availability will not be calculated for the resource. 
	* `console_path_prefix` - The console path prefix to use for providing service home url page navigation.  For example if the prefix provided is 'security/bastion/bastions', the URL used for navigation will be https://<cloudhostname>/security/bastion/bastions/<resourceOcid>. If not provided, service home page link  will not be shown in the stack monitoring home page. 
	* `external_id_mapping` - The external resource identifier property in the metric dimensions.  Resources imported will be using this property value for external id. 
	* `lifecycle_status_mappings_for_up_status` - Lifecycle states of the external resource which reflects the status of the resource being up. 
	* `namespace` - Name space to be used for Oracle Cloud Infrastructure Native service resources discovery.
	* `resource_group` - The resource group to use while fetching metrics from telemetry. If not specified, resource group will be skipped in the list metrics request. 
	* `resource_name_filter` - The resource name filter. Resources matching with the resource name filter will be imported. Regular expressions will be accepted. 
	* `resource_name_mapping` - The resource name property in the metric dimensions.  Resources imported will be using this property value for resource name. 
	* `resource_type_filter` - The resource type filter. Resources matching with the resource type filter will be imported. Regular expressions will be accepted. 
	* `resource_type_mapping` - The resource type property in the metric dimensions.  Resources imported will be using this property value for resource type. If not specified, namespace will be used for resource type. 
	* `service_base_url` - The base URL of the Oracle Cloud Infrastructure service to which the resource belongs to. Also this property is applicable only when source is OCI_TELEMETRY_NATIVE. 
	* `should_use_metrics_flow_for_status` - Flag to indicate whether status is calculated using metrics or  LifeCycleState attribute of the resource in Oracle Cloud Infrastructure service. 
	* `source` - Source from where the metrics pushed to telemetry. Possible values:
		* OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from Oracle Cloud Infrastructure Native Services.
		* OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus. 
	* `type` - Task type.
* `tenant_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy. 
* `time_created` - The date and time when the stack monitoring resource task was created, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time when the stack monitoring resource task was last updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
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

