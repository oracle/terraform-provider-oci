---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_group"
sidebar_current: "docs-oci-datasource-identity-group"
description: |-
  Provides details about a specific Group in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_group
This data source provides details about a specific Group resource in Oracle Cloud Infrastructure Identity service.

Gets the specified group's information.

This operation does not return a list of all the users in the group. To do that, use
[ListUserGroupMemberships](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/UserGroupMembership/ListUserGroupMemberships) and
provide the group's OCID as a query parameter in the request.


## Example Usage

```hcl
data "oci_identity_group" "test_group" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) The OCID of the group.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the group.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 
* `state` - The group's current state.
* `time_created` - Date and time the group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

