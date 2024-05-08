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

Creates a new ManagedList resource.


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

* `compartment_id` - (Required) (Updatable) Compartment OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Managed list description

	Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) Managed list display name.

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `list_items` - (Optional) (Updatable) List of items in the managed list
* `list_type` - (Optional) Type of information stored in the list
* `source_managed_list_id` - (Optional) OCID of the source managed list


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID where the resource is created
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Managed list description
* `display_name` - Managed list display name
* `feed_provider` - Provider of the managed list feed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that can't be changed after creation
* `is_editable` - Is this list editable?
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. [DEPRECATE]
* `list_items` - List of items in the managed list
* `list_type` - Type of information contained in the managed list
* `source_managed_list_id` - OCID of the source managed list
* `state` - The current lifecycle state of the resource
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the managed list was created. Format defined by RFC3339.
* `time_updated` - The date and time the managed list was last updated. Format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed List
	* `update` - (Defaults to 20 minutes), when updating the Managed List
	* `delete` - (Defaults to 20 minutes), when destroying the Managed List


## Import

ManagedLists can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_managed_list.test_managed_list "id"
```

