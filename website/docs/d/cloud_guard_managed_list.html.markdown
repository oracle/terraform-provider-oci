---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_managed_list"
sidebar_current: "docs-oci-datasource-cloud_guard-managed_list"
description: |-
  Provides details about a specific Managed List in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_managed_list
This data source provides details about a specific Managed List resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a managed list identified by managedListId.

## Example Usage

```hcl
data "oci_cloud_guard_managed_list" "test_managed_list" {
	#Required
	managed_list_id = oci_cloud_guard_managed_list.test_managed_list.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_list_id` - (Required) The cloudguard list OCID to be passed in the request.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier where the resource is created
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - ManagedList description.
* `display_name` - ManagedList display name.
* `feed_provider` - provider of the feed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
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

