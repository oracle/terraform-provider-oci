# oci_identity_dynamic_group

## DynamicGroup Resource

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



### Create Operation
Creates a new dynamic group in your tenancy.

You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
reside within compartments inside the tenancy. For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the dynamic group, which must be unique across all dynamic groups in your
tenancy, and cannot be changed. Note that this name has to be also unique accross all groups in your tenancy.
You can use this name or the OCID when writing policies that apply to the dynamic group. For more information
about policies, see [How Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policies.htm).

You must also specify a *description* for the dynamic group (although it can be an empty string). It does not
have to be unique, and you can change it anytime with [UpdateDynamicGroup](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/DynamicGroup/UpdateDynamicGroup).

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
object, first make sure its `lifecycleState` has changed to ACTIVE.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the group.
* `description` - (Required) The description you assign to the group during creation. Does not have to be unique, and it's changeable.
* `matching_rule` - (Required) The matching rule to dynamically match an instance certificate to this dynamic group. For rule syntax, see [Managing Dynamic Groups](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingdynamicgroups.htm). 
* `name` - (Required) The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 


### Update Operation
Updates the specified dynamic group.

The following arguments support updates:
* `description` - The description you assign to the group during creation. Does not have to be unique, and it's changeable.
* `matching_rule` - The matching rule to dynamically match an instance certificate to this dynamic group. For rule syntax, see [Managing Dynamic Groups](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingdynamicgroups.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_dynamic_group" "test_dynamic_group" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.dynamic_group_description}"
	matching_rule = "${var.dynamic_group_matching_rule}"
	name = "${var.dynamic_group_name}"
}
```

# oci_identity_dynamic_groups

## DynamicGroup DataSource

Gets a list of dynamic_groups.

### List Operation
Lists the dynamic groups in your tenancy. You must specify your tenancy's OCID as the value for
the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


The following attributes are exported:

* `dynamic_groups` - The list of dynamic_groups.

### Example Usage

```hcl
data "oci_identity_dynamic_groups" "test_dynamic_groups" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```