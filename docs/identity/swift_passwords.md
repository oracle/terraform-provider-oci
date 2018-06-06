# oci_identity_swift_password

## SwiftPassword Resource
Deprecated. Use AuthToken instead.

### SwiftPassword Reference

The following attributes are exported:

* `description` - The description you assign to the Swift password. Does not have to be unique, and it's changeable.
* `expires_on` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `id` - The OCID of the Swift password.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `password` - The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not for `ListSwiftPasswords` or `UpdateSwiftPassword`. 
* `state` - The password's current state. After creating a password, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `SwiftPassword` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.



### Create Operation
**Deprecated. Use [CreateAuthToken](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/AuthToken/CreateAuthToken) instead.**

Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
[Managing User Credentials](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingcredentials.htm).

You must specify a *description* for the Swift password (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateSwiftPassword](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SwiftPassword/UpdateSwiftPassword).

Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
does not need to write a policy to give users this ability. To compare, administrators who have permission to the
tenancy can use this operation to create a Swift password for any user, including themselves.


The following arguments are supported:

* `description` - (Required) The description you assign to the Swift password during creation. Does not have to be unique, and it's changeable. 
* `user_id` - (Required) The OCID of the user.


### Update Operation
**Deprecated. Use [UpdateAuthToken](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/AuthToken/UpdateAuthToken) instead.**

Updates the specified Swift password's description.


The following arguments support updates:
* `description` - The description you assign to the Swift password during creation. Does not have to be unique, and it's changeable. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_swift_password" "test_swift_password" {
	#Required
	description = "${var.swift_password_description}"
	user_id = "${oci_identity_user.test_user.id}"
}
```

# oci_identity_swift_passwords

## SwiftPassword DataSource

Gets a list of swift_passwords.

### List Operation
**Deprecated. Use [ListAuthTokens](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/AuthToken/ListAuthTokens) instead.**

Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
the password itself. The actual password is returned only upon creation.

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


The following attributes are exported:

* `passwords` - The list of passwords.

### Example Usage

```hcl
data "oci_identity_swift_passwords" "test_swift_passwords" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```