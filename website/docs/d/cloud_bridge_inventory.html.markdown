---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_inventory"
sidebar_current: "docs-oci-datasource-cloud_bridge-inventory"
description: |-
  Provides details about a specific Inventory in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_inventory
This data source provides details about a specific Inventory resource in Oracle Cloud Infrastructure Cloud Bridge service.

Gets an inventory by identifier.

## Example Usage

```hcl
data "oci_cloud_bridge_inventory" "test_inventory" {
	#Required
	inventory_id = oci_cloud_bridge_inventory.test_inventory.id
}
```

## Argument Reference

The following arguments are supported:

* `inventory_id` - (Required) Inventory OCID.


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

