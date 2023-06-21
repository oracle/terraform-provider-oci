---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_groups"
sidebar_current: "docs-oci-datasource-identity-groups"
description: |-
  Provides the list of Groups in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_groups
This data source provides the list of Groups in Oracle Cloud Infrastructure Identity service.

Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_groups" "test_groups" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	name = var.group_name
	state = var.group_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `groups` - The list of groups.

### Group Reference

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

