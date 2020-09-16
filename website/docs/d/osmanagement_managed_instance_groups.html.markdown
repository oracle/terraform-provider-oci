---
subcategory: "Osmanagement"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_groups"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instance_groups"
description: |-
  Provides the list of Managed Instance Groups in Oracle Cloud Infrastructure Osmanagement service
---

# Data Source: oci_osmanagement_managed_instance_groups
This data source provides the list of Managed Instance Groups in Oracle Cloud Infrastructure Osmanagement service.

Returns a list of all Managed Instance Groups.


## Example Usage

```hcl
data "oci_osmanagement_managed_instance_groups" "test_managed_instance_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.managed_instance_group_display_name
	os_family = var.managed_instance_group_os_family
	state = var.managed_instance_group_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `os_family` - (Optional) The OS family for which to list resources.
* `state` - (Optional) The current lifecycle state for the object.


## Attributes Reference

The following attributes are exported:

* `managed_instance_groups` - The list of managed_instance_groups.

### ManagedInstanceGroup Reference

The following attributes are exported:

* `compartment_id` - OCID for the Compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Information specified by the user about the managed instance group
* `display_name` - Managed Instance Group identifier
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID for the managed instance group
* `managed_instances` - list of Managed Instances in the group
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `os_family` - The Operating System type of the managed instance.
* `state` - The current state of the Software Source.

