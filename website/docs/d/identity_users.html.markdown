---
layout: "oci"
page_title: "OCI: oci_identity_users"
sidebar_current: "docs-oci-datasource-identity-users"
description: |-
  Provides a list of Users
---

# Data Source: oci_identity_users
The `oci_identity_users` data source allows access to the list of OCI users

Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_users" "test_users" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `users` - The list of users.

### User Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the user.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the user.
* `inactive_state` - Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive:  - bit 0: SUSPENDED (reserved for future use) - bit 1: DISABLED (reserved for future use) - bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console) 
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. 
* `state` - The user's current state. After creating a user, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the user was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

