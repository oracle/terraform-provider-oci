# oci\_identity\_compartment

## Compartment Resource

### Compartment Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the compartment.
* `description` - The description you assign to the compartment. Does not have to be unique, and it's changeable.
* `id` - The OCID of the compartment.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the compartment during creation. The name must be unique across all  compartments in the tenancy. 
* `state` - The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the compartment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new compartment in your tenancy.

If a compartment with the given `name` already exists, then that compartment will be used instead of creating a new compartment.

**Important:** Compartments cannot be deleted.

You must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy
is simply the root compartment. For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the compartment, which must be unique across all compartments in
your tenancy. You can use this name or the OCID when writing policies that apply
to the compartment. For more information about policies, see
[How Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policies.htm).

You must also specify a *description* for the compartment (although it can be an empty string). It does
not have to be unique, and you can change it anytime with 
[UpdateCompartment](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Compartment/UpdateCompartment).

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
object, first make sure its `lifecycleState` has changed to ACTIVE.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the compartment.
* `description` - (Required) The description you assign to the compartment during creation. Does not have to be unique, and it's changeable. 
* `name` - (Required) The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy. 


### Update Operation
Updates the specified compartment's description or name. You can't update the root compartment.

The following arguments support updates:
* `description` - The description you assign to the compartment during creation. Does not have to be unique, and it's changeable. 
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_identity_compartment" "test_compartment" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.compartment_description}"
	name = "${var.compartment_name}"
}
```

# oci\_identity\_compartments

## Compartment DataSource

Gets a list of compartments.

### List Operation
Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
for the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


The following attributes are exported:

* `compartments` - The list of compartments.

### Example Usage

```
data "oci_identity_compartments" "test_compartments" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```