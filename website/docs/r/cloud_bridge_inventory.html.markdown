---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_inventory"
sidebar_current: "docs-oci-resource-cloud_bridge-inventory"
description: |-
  Provides the Inventory resource in Oracle Cloud Infrastructure Cloud Bridge service
---

# oci_cloud_bridge_inventory
This resource provides the Inventory resource in Oracle Cloud Infrastructure Cloud Bridge service.

Creates an inventory.

## Example Usage

```hcl
resource "oci_cloud_bridge_inventory" "test_inventory" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.inventory_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenantId.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Inventory displayName.
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenantId.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Inventory display name.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Inventory OCID.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the inventory.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the inventory was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the inventory was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Inventory
	* `update` - (Defaults to 20 minutes), when updating the Inventory
	* `delete` - (Defaults to 20 minutes), when destroying the Inventory


## Import

Inventories can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_bridge_inventory.test_inventory "id"
```

