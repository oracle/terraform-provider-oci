---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_baselineable_metric"
sidebar_current: "docs-oci-datasource-stack_monitoring-baselineable_metric"
description: |-
  Provides details about a specific Baselineable Metric in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_baselineable_metric
This data source provides details about a specific Baselineable Metric resource in Oracle Cloud Infrastructure Stack Monitoring service.

Get the Baseline-able metric for the given id

## Example Usage

```hcl
data "oci_stack_monitoring_baselineable_metric" "test_baselineable_metric" {
	#Required
	baselineable_metric_id = oci_stack_monitoring_baselineable_metric.test_baselineable_metric.id
}
```

## Argument Reference

The following arguments are supported:

* `baselineable_metric_id` - (Required) Identifier for the metric


## Attributes Reference

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
* `state` - The current lifecycle state of the metric extension
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenancy_id` - OCID of the tenancy
* `time_created` - creation date
* `time_last_updated` - last updated time

