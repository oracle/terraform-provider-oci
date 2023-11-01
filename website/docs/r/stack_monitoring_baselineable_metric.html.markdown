---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_baselineable_metric"
sidebar_current: "docs-oci-resource-stack_monitoring-baselineable_metric"
description: |-
  Provides the Baselineable Metric resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_baselineable_metric
This resource provides the Baselineable Metric resource in Oracle Cloud Infrastructure Stack Monitoring service.

Creates the specified Baseline-able metric

## Example Usage

```hcl
resource "oci_stack_monitoring_baselineable_metric" "test_baselineable_metric" {
	#Required
	column = var.baselineable_metric_column
	compartment_id = var.compartment_id
	name = var.baselineable_metric_name
	namespace = var.baselineable_metric_namespace
	resource_group = var.baselineable_metric_resource_group
}
```

## Argument Reference

The following arguments are supported:

* `column` - (Required) (Updatable) metric column name
* `compartment_id` - (Required) (Updatable) OCID of the compartment
* `name` - (Required) (Updatable) name of the metric
* `namespace` - (Required) (Updatable) namespace of the metric
* `resource_group` - (Required) (Updatable) Resource group of the metric


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Baselineable Metric
	* `update` - (Defaults to 20 minutes), when updating the Baselineable Metric
	* `delete` - (Defaults to 20 minutes), when destroying the Baselineable Metric


## Import

BaselineableMetrics can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_baselineable_metric.test_baselineable_metric "id"
```

