---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_user"
sidebar_current: "docs-oci-datasource-identity-user"
description: |-
  Provides details about a specific User in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_user
This data source provides details about a specific User resource in Oracle Cloud Infrastructure Identity service.

Gets the specified user's information.

## Example Usage

```hcl
data "oci_identity_user" "test_user" {
	#Required
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `capabilities` - Properties indicating how the user is allowed to authenticate.
	* `can_use_api_keys` - Indicates if the user can use API keys.
	* `can_use_auth_tokens` - Indicates if the user can use SWIFT passwords / auth tokens.
	* `can_use_console_password` - Indicates if the user can log in to the console.
	* `can_use_customer_secret_keys` - Indicates if the user can use SigV4 symmetric keys.
	* `can_use_db_credentials` - Indicates if the user can use DB passwords. 
	* `can_use_oauth2client_credentials` - Indicates if the user can use OAuth2 credentials and tokens. 
	* `can_use_smtp_credentials` - Indicates if the user can use SMTP passwords.
* `compartment_id` - The OCID of the tenancy containing the user.
* `db_user_name` - DB username of the DB credential. Has to be unique across the tenancy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable.
* `email` - The email address you assign to the user. The email address must be unique across all users in the tenancy. 
* `email_verified` - Whether the email address has been validated.
* `external_identifier` - Identifier of the user in the identity provider
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the user.
* `identity_provider_id` - The OCID of the `IdentityProvider` this user belongs to.
* `inactive_state` - Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive:
	* bit 0: SUSPENDED (reserved for future use)
	* bit 1: DISABLED (reserved for future use)
	* bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console) 
* `last_successful_login_time` - The date and time of when the user most recently logged in the format defined by RFC3339 (ex. `2016-08-25T21:10:29.600Z`). If there is no login history, this field is null.

	For illustrative purposes, suppose we have a user who has logged in at July 1st, 2020 at 1200 PST and logged out 30 minutes later. They then login again on July 2nd, 2020 at 1500 PST.

	Their previousSuccessfulLoginTime would be `2020-07-01:19:00.000Z`.

	Their lastSuccessfulLoginTime would be `2020-07-02:22:00.000Z`. 
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. 
* `previous_successful_login_time` - The date and time of when the user most recently logged in the format defined by RFC3339 (ex. `2016-08-25T21:10:29.600Z`). If there is no login history, this field is null.

	For illustrative purposes, suppose we have a user who has logged in at July 1st, 2020 at 1200 PST and logged out 30 minutes later. They then login again on July 2nd, 2020 at 1500 PST.

	Their previousSuccessfulLoginTime would be `2020-07-01:19:00.000Z`.

	Their lastSuccessfulLoginTime would be `2020-07-02:22:00.000Z`. 
* `state` - The user's current state.
* `time_created` - Date and time the user was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

