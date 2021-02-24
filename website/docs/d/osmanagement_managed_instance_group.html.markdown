---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_group"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instance_group"
description: |-
  Provides details about a specific Managed Instance Group in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_managed_instance_group
This data source provides details about a specific Managed Instance Group resource in Oracle Cloud Infrastructure OS Management service.

Returns a specific Managed Instance Group.


## Example Usage

```hcl
data "oci_osmanagement_managed_instance_group" "test_managed_instance_group" {
	#Required
	managed_instance_group_id = oci_osmanagement_managed_instance_group.test_managed_instance_group.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) OCID for the managed instance group


## Attributes Reference

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

