---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_inventories"
sidebar_current: "docs-oci-datasource-cloud_bridge-inventories"
description: |-
  Provides the list of Inventories in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_inventories
This data source provides the list of Inventories in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of inventories.


## Example Usage

```hcl
data "oci_cloud_bridge_inventories" "test_inventories" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	state = var.inventory_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `state` - (Optional) A filter to return inventory if the lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `inventory_collection` - The list of inventory_collection.

### Inventory Reference

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

