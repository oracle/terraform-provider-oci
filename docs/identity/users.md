# oci_identity_user

## User Resource

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



### Create Operation
Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
IAM Service components, see [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).

You must specify your tenancy's OCID as the compartment ID in the request object (remember that the
tenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and
some policies) reside within the tenancy itself, unlike cloud resources such as compute instances,
which typically reside within compartments inside the tenancy. For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the user, which must be unique across all users in your tenancy
and cannot be changed. Allowed characters: No spaces. Only letters, numerals, hyphens, periods,
underscores, +, and @. If you specify a name that's already in use, you'll get a 409 error.
This name will be the user's login to the Console. You might want to pick a
name that your company's own identity system (e.g., Active Directory, LDAP, etc.) already uses.
If you delete a user and then create a new user with the same name, they'll be considered different
users because they have different OCIDs.

You must also specify a *description* for the user (although it can be an empty string).
It does not have to be unique, and you can change it anytime with
[UpdateUser](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/UpdateUser). You can use the field to provide the user's
full name, a description, a nickname, or other information to generally identify the user.

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before
using the object, first make sure its `lifecycleState` has changed to ACTIVE.

A new user has no permissions until you place the user in one or more groups (see
[AddUserToGroup](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/AddUserToGroup)). If the user needs to
access the Console, you need to provide the user a password (see
[CreateOrResetUIPassword](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UIPassword/CreateOrResetUIPassword)).
If the user needs to access the Oracle Cloud Infrastructure REST API, you need to upload a
public API signing key for that user (see
[Required Keys and OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm) and also
[UploadApiKey](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/ApiKey/UploadApiKey)).

**Important:** Make sure to inform the new user which compartment(s) they have access to.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the user.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) The description you assign to the user during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. 


### Update Operation
Updates the description of the specified user.

The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the user during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_user" "test_user" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.user_description}"
	name = "${var.user_name}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_identity_users

## User DataSource

Gets a list of users.

### List Operation
Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


The following attributes are exported:

* `users` - The list of users.

### Example Usage

```hcl
data "oci_identity_users" "test_users" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```