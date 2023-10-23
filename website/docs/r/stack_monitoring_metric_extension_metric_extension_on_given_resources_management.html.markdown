---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management"
sidebar_current: "docs-oci-resource-stack_monitoring-metric_extension_metric_extension_on_given_resources_management"
description: |-
  Provides the Metric Extension Metric Extension On Given Resources Management resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management
This resource provides the Metric Extension Metric Extension On Given Resources Management resource in Oracle Cloud Infrastructure Stack Monitoring service.

Submits a request to enable matching metric extension Id for the given Resource IDs


## Example Usage

```hcl
resource "oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management" "test_metric_extension_metric_extension_on_given_resources_management" {
	#Required
	metric_extension_id = oci_stack_monitoring_metric_extension.test_metric_extension.id
	resource_ids = var.metric_extension_metric_extension_on_given_resources_management_resource_ids
	enable_metric_extension_on_given_resources = var.enable_metric_extension_on_given_resources
}
```

## Argument Reference

The following arguments are supported:

* `metric_extension_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the metric extension resource.
* `resource_ids` - (Required) List of Resource IDs [OCIDs]. Currently, supports only one resource id per request.
* `enable_metric_extension_on_given_resources` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Metric Extension Metric Extension On Given Resources Management
	* `update` - (Defaults to 20 minutes), when updating the Metric Extension Metric Extension On Given Resources Management
	* `delete` - (Defaults to 20 minutes), when destroying the Metric Extension Metric Extension On Given Resources Management
