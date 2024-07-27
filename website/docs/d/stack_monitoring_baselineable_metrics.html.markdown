---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_baselineable_metrics"
sidebar_current: "docs-oci-datasource-stack_monitoring-baselineable_metrics"
description: |-
  Provides the list of Baselineable Metrics in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_baselineable_metrics
This data source provides the list of Baselineable Metrics in Oracle Cloud Infrastructure Stack Monitoring service.

List of summary of baseline-able metrics for a given resource group if specified.

## Example Usage

```hcl
data "oci_stack_monitoring_baselineable_metrics" "test_baselineable_metrics" {

	#Optional
	baselineable_metric_id = oci_stack_monitoring_baselineable_metric.test_baselineable_metric.id
	compartment_id = var.compartment_id
	is_out_of_box = var.baselineable_metric_is_out_of_box
	metric_namespace = var.baselineable_metric_metric_namespace
	name = var.baselineable_metric_name
	resource_group = var.baselineable_metric_resource_group
	resource_type = var.baselineable_metric_resource_type
}
```

## Argument Reference

The following arguments are supported:

* `baselineable_metric_id` - (Optional) Identifier for the metric
* `compartment_id` - (Optional) The ID of the compartment in which data is listed.
* `is_out_of_box` - (Optional) Is the baseline enabled metric defined out of box by Oracle or by end-user 
* `metric_namespace` - (Optional) A filter to return monitored resource types that has the matching namespace. 
* `name` - (Optional) Metric Name
* `resource_group` - (Optional) Resource Group
* `resource_type` - (Optional) Resource Type


## Attributes Reference

The following attributes are exported:

* `baselineable_metric_summary_collection` - The list of baselineable_metric_summary_collection.

### BaselineableMetric Reference

The following attributes are exported:

* `column` - metric column name
* `compartment_id` - OCID of the compartment
* `created_by` - Created user id
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID of the metric
* `is_out_of_box` - Is the metric created out of box, default false
* `last_updated_by` - last Updated user id
* `name` - name of the metric
* `namespace` - namespace of the metric
* `resource_group` - Resource group of the metric
* `resource_type` - Resource type of the metric
* `state` - The current lifecycle state of the metric extension
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenancy_id` - OCID of the tenancy
* `time_created` - creation date
* `time_last_updated` - last updated time

