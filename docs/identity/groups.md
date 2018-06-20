# oci_identity_group

## Group Resource

### Group Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the group.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 
* `state` - The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new group in your tenancy.

You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
reside within compartments inside the tenancy. For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the group, which must be unique across all groups in your tenancy and
cannot be changed. You can use this name or the OCID when writing policies that apply to the group. For more
information about policies, see [How Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policies.htm).

You must also specify a *description* for the group (although it can be an empty string). It does not
have to be unique, and you can change it anytime with [UpdateGroup](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Group/UpdateGroup).

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
object, first make sure its `lifecycleState` has changed to ACTIVE.

After creating the group, you need to put users in it and write policies for it.
See [AddUserToGroup](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/AddUserToGroup) and
[CreatePolicy](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/CreatePolicy).


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the group.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) The description you assign to the group during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 


### Update Operation
Updates the specified group.

The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the group during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_group" "test_group" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.group_description}"
	name = "${var.group_name}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_identity_groups

## Group DataSource

Gets a list of groups.

### List Operation
Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


The following attributes are exported:

* `groups` - The list of groups.

### Example Usage

```hcl
data "oci_identity_groups" "test_groups" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```