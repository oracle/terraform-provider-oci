---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_group"
sidebar_current: "docs-oci-resource-osmanagement-managed_instance_group"
description: |-
  Provides the Managed Instance Group resource in Oracle Cloud Infrastructure OS Management service
---

# oci_osmanagement_managed_instance_group
This resource provides the Managed Instance Group resource in Oracle Cloud Infrastructure OS Management service.

Creates a new Managed Instance Group on the management system.
This will not contain any managed instances after it is first created,
and they must be added later.


## Example Usage

```hcl
resource "oci_osmanagement_managed_instance_group" "test_managed_instance_group" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.managed_instance_group_description
	freeform_tags = {"bar-key"= "value"}
	os_family = var.managed_instance_group_os_family
	managed_instance_ids = var.managed_instance_group_managed_instance_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) OCID for the Compartment
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Information specified by the user about the managed instance group
* `display_name` - (Required) (Updatable) Managed Instance Group identifier
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `os_family` - (Optional) The Operating System type of the managed instance(s) on which this scheduled job will operate. If not specified, this defaults to Linux. 
* `managed_instance_ids` - (Optional) The list of managed instance OCIDs to be added to the managed instance group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID for the Compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Information specified by the user about the managed instance group
* `display_name` - Managed Instance Group identifier
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID for the managed instance group
* `managed_instance_ids` - The list of managed instances OCIDs attached to the managed instance group.
* `managed_instances` - list of Managed Instances in the group
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `os_family` - The Operating System type of the managed instance.
* `state` - The current state of the Software Source.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group


## Import

ManagedInstanceGroups can be imported using the `id`, e.g.

```
$ terraform import oci_osmanagement_managed_instance_group.test_managed_instance_group "id"
```

