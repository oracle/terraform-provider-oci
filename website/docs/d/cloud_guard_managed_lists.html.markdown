---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_managed_lists"
sidebar_current: "docs-oci-datasource-cloud_guard-managed_lists"
description: |-
  Provides the list of Managed Lists in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_managed_lists
This data source provides the list of Managed Lists in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all ManagedList resources in a compartment, identified by compartmentId.
The ListManagedLists operation returns only the managed lists in `compartmentId` passed.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return ManagedLists in only
those compartments for which the requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListManagedLists on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_cloud_guard_managed_lists" "test_managed_lists" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.managed_list_access_level
	compartment_id_in_subtree = var.managed_list_compartment_id_in_subtree
	display_name = var.managed_list_display_name
	list_type = var.managed_list_list_type
	resource_metadata_only = var.managed_list_resource_metadata_only
	state = var.managed_list_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the setting of `accessLevel`. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `list_type` - (Optional) The type of managed list.
* `resource_metadata_only` - (Optional) Default is false. When set to true, the list of all Oracle-managed resources metadata supported by Cloud Guard is returned. 
* `state` - (Optional) The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `managed_list_collection` - The list of managed_list_collection.

### ManagedList Reference

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

