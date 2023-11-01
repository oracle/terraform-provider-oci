---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_metric_extensions_test_management"
sidebar_current: "docs-oci-resource-stack_monitoring-metric_extensions_test_management"
description: |-
  Provides the Metric Extensions Test Management resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_metric_extensions_test_management
This resource provides the Metric Extensions Test Management resource in Oracle Cloud Infrastructure Stack Monitoring service.

Performs test of Metric Extension on a specific resource Id

## Example Usage

```hcl
resource "oci_stack_monitoring_metric_extensions_test_management" "test_metric_extensions_test_management" {
	#Required
	metric_extension_id = oci_stack_monitoring_metric_extension.test_metric_extension.id
	resource_ids = var.metric_extensions_test_management_resource_ids
}
```

## Argument Reference

The following arguments are supported:

* `metric_extension_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the metric extension resource.
* `resource_ids` - (Required) List of Resource IDs [OCID]. Currently, supports only one resource id per request.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `test_run_id` - Test Run Id
* `test_run_metric_suffix` - Test Run Metric Suffix
* `test_run_namespace_name` - Test Run Namespace name
* `test_run_resource_group_name` - Test Run Resource Group name

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Metric Extensions Test Management
	* `update` - (Defaults to 20 minutes), when updating the Metric Extensions Test Management
	* `delete` - (Defaults to 20 minutes), when destroying the Metric Extensions Test Management


## Import

Import is not supported for this resource.

