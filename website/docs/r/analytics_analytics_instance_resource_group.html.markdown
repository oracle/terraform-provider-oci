---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance_resource_group"
sidebar_current: "docs-oci-resource-analytics-analytics_instance_resource_group"
description: |-
  Provides the Analytics Instance Resource Group resource in Oracle Cloud Infrastructure Analytics service
---

# oci_analytics_analytics_instance_resource_group
This resource provides the Analytics Instance Resource Group resource in Oracle Cloud Infrastructure Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/analytics/latest/ResourceGroup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/analytics

Create a new resource group for the instance


## Example Usage

```hcl
resource "oci_analytics_analytics_instance_resource_group" "test_analytics_instance_resource_group" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id
	capacity = var.analytics_instance_resource_group_capacity
	resource_name = oci_cloud_guard_resource.test_resource.name

	#Optional
	description = var.analytics_instance_resource_group_description
	display_name = var.analytics_instance_resource_group_display_name
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the Analytics instance. 
* `capacity` - (Required) (Updatable) The capacity (in OCPU's) to be allocated for this resource.
* `description` - (Optional) (Updatable) Optional description of the resource group
* `display_name` - (Optional) (Updatable) Meaningful name of resource group for end user
* `resource_name` - (Required) (Updatable) Meaningful name of resource group for end user


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `capacity` - The capacity (in OCPU's) to be allocated for this resource.
* `description` - Optional description of the resource group
* `display_name` - Meaningful name of resource group for end user
* `id` - Unique identifier and name of resource group.  Must be unique within the instance
* `resource_name` - Meaningful name of resource group for end user

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Analytics Instance Resource Group
	* `update` - (Defaults to 20 minutes), when updating the Analytics Instance Resource Group
	* `delete` - (Defaults to 20 minutes), when destroying the Analytics Instance Resource Group


## Import

AnalyticsInstanceResourceGroups can be imported using the `id`, e.g.

```
$ terraform import oci_analytics_analytics_instance_resource_group.test_analytics_instance_resource_group "analyticsInstances/{analyticsInstanceId}/resourceGroups/{analyticsInstanceResourceGroupId}" 
```

