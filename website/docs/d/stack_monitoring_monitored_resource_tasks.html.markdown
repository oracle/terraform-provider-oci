---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource_tasks"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitored_resource_tasks"
description: |-
  Provides the list of Monitored Resource Tasks in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitored_resource_tasks
This data source provides the list of Monitored Resource Tasks in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of stack monitoring resource tasks in the compartment.


## Example Usage

```hcl
data "oci_stack_monitoring_monitored_resource_tasks" "test_monitored_resource_tasks" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	status = var.monitored_resource_task_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for which  stack monitoring resource tasks should be listed. 
* `status` - (Optional) A filter to return only resources that matches with lifecycleState given.


## Attributes Reference

The following attributes are exported:

* `monitored_resource_tasks_collection` - The list of monitored_resource_tasks_collection.

### MonitoredResourceTask Reference

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

