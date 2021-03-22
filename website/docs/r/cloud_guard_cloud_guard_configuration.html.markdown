---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_cloud_guard_configuration"
sidebar_current: "docs-oci-resource-cloud_guard-cloud_guard_configuration"
description: |-
  Provides the Cloud Guard Configuration resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_cloud_guard_configuration
This resource provides the Cloud Guard Configuration resource in Oracle Cloud Infrastructure Cloud Guard service.

Enable/Disable Cloud Guard. The reporting region cannot be updated once created.


## Example Usage

```hcl
resource "oci_cloud_guard_cloud_guard_configuration" "test_cloud_guard_configuration" {
	#Required
	compartment_id = var.compartment_id
	reporting_region = var.cloud_guard_configuration_reporting_region
	status = var.cloud_guard_configuration_status

	#Optional
	self_manage_resources = var.cloud_guard_configuration_self_manage_resources
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The ID of the compartment in which to list resources.
* `reporting_region` - (Required) (Updatable) The reporting region value
* `self_manage_resources` - (Optional) (Updatable) Identifies if Oracle managed resources will be created by customers. If no value is specified false is the default. 
* `status` - (Required) (Updatable) Status of Cloud Guard Tenant


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `reporting_region` - The reporting region value
* `self_manage_resources` - Identifies if Oracle managed resources were created by customers 
* `status` - Status of Cloud Guard Tenant

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Guard Configuration
	* `update` - (Defaults to 20 minutes), when updating the Cloud Guard Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Guard Configuration


## Import

Import is not supported for this resource.

