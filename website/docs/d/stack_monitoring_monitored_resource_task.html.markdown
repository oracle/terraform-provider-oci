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
	* `availability_proxy_metric_collection_interval` - Metrics collection interval in seconds used when calculating the availability of the  resource based on metrics specified using the property 'availabilityProxyMetrics'. 
	* `availability_proxy_metrics` - List of metrics to be used to calculate the availability of the resource. Resource is considered to be up if at least one of the specified metrics is available for  the resource during the specified interval using the property  'availabilityProxyMetricCollectionIntervalInSeconds'. If no metrics are specified, availability will not be calculated for the resource. 
	* `namespace` - Name space to be used for Oracle Cloud Infrastructure Native service resources' import.
	* `resource_group` - The resource group to use while fetching metrics from telemetry. If not specified, resource group will be skipped in the list metrics request. 
	* `source` - Source from where the metrics pushed to telemetry. Possible values:
		* OCI_TELEMETRY_NATIVE      - The metrics are pushed to telemetry from Oracle Cloud Infrastructure Native Services.
		* OCI_TELEMETRY_PROMETHEUS  - The metrics are pushed to telemetry from Prometheus. 
	* `type` - Task type.
* `tenant_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy. 
* `time_created` - The date and time when the stack monitoring resource task was created, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time when the stack monitoring resource task was last updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `work_request_ids` - Identifiers [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for work requests submitted for this task. 

