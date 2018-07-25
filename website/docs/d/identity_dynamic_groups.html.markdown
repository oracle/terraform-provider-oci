---
layout: "oci"
page_title: "OCI: oci_identity_dynamic_groups"
sidebar_current: "docs-oci-datasource-identity-dynamic_groups"
description: |-
  Provides a list of DynamicGroups
---

# Data Source: oci_identity_dynamic_groups
The `oci_identity_dynamic_groups` data source allows access to the list of OCI dynamic_groups

Lists the dynamic groups in your tenancy. You must specify your tenancy's OCID as the value for
the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_dynamic_groups" "test_dynamic_groups" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `dynamic_groups` - The list of dynamic_groups.

### DynamicGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the group.
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `id` - The OCID of the group.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `matching_rule` - A rule string that defines which instance certificates will be matched. For syntax, see [Managing Dynamic Groups](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingdynamicgroups.htm). 
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 
* `state` - The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

