---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_managed_list"
sidebar_current: "docs-oci-resource-cloud_guard-managed_list"
description: |-
  Provides the Managed List resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_managed_list
This resource provides the Managed List resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a new ManagedList.


## Example Usage

```hcl
resource "oci_cloud_guard_managed_list" "test_managed_list" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.managed_list_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.managed_list_description
	freeform_tags = {"bar-key"= "value"}
	list_items = var.managed_list_list_items
	list_type = var.managed_list_list_type
	source_managed_list_id = oci_cloud_guard_managed_list.test_managed_list.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) ManagedList description
* `display_name` - (Required) (Updatable) ManagedList display name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `list_items` - (Optional) (Updatable) List of ManagedListItem
* `list_type` - (Optional) type of the list
* `source_managed_list_id` - (Optional) OCID of the Source ManagedList


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier where the resource is created
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - ManagedList description
* `display_name` - ManagedList display name
* `feed_provider` - provider of the feed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `is_editable` - If this list is editable or not
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `list_items` - List of ManagedListItem
* `list_type` - type of the list
* `source_managed_list_id` - OCID of the Source ManagedList
* `state` - The current state of the resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the managed list was created. Format defined by RFC3339.
* `time_updated` - The date and time the managed list was updated. Format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed List
	* `update` - (Defaults to 20 minutes), when updating the Managed List
	* `delete` - (Defaults to 20 minutes), when destroying the Managed List


## Import

ManagedLists can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_managed_list.test_managed_list "id"
```

