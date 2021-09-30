---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub"
sidebar_current: "docs-oci-resource-opsi-awr_hub"
description: |-
  Provides the Awr Hub resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_awr_hub
This resource provides the Awr Hub resource in Oracle Cloud Infrastructure Opsi service.

Create a AWR hub resource for the tenant in Operations Insights.
This resource will be created in root compartment.


## Example Usage

```hcl
resource "oci_opsi_awr_hub" "test_awr_hub" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.awr_hub_display_name
	object_storage_bucket_name = oci_objectstorage_bucket.test_bucket.name
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) User-friedly name of AWR Hub that does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `object_storage_bucket_name` - (Required) Object Storage Bucket Name
* `operations_insights_warehouse_id` - (Required) OPSI Warehouse OCID


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `awr_mailbox_url` - Mailbox URL required for AWR hub and AWR source setup.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - User-friedly name of AWR Hub that does not have to be unique.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - AWR Hub OCID
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `object_storage_bucket_name` - Object Storage Bucket Name
* `operations_insights_warehouse_id` - OPSI Warehouse OCID
* `state` - Possible lifecycle states
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Awr Hub
	* `update` - (Defaults to 20 minutes), when updating the Awr Hub
	* `delete` - (Defaults to 20 minutes), when destroying the Awr Hub


## Import

AwrHubs can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_awr_hub.test_awr_hub "id"
```

