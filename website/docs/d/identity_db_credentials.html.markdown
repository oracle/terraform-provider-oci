---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_db_credentials"
sidebar_current: "docs-oci-datasource-identity-db_credentials"
description: |-
  Provides the list of Db Credentials in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_db_credentials
This data source provides the list of Db Credentials in Oracle Cloud Infrastructure Identity service.

Lists the DB credentials for the specified user. The returned object contains the credential's OCID


## Example Usage

```hcl
data "oci_identity_db_credentials" "test_db_credentials" {
	#Required
	user_id = oci_identity_user.test_user.id

	#Optional
	name = var.db_credential_name
	state = var.db_credential_state
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) A filter to only return resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `db_credentials` - The list of db_credentials.

### DbCredential Reference

The following attributes are exported:

* `description` - The description you assign to the DB credential. Does not have to be unique, and it's changeable.
* `id` - The OCID of the DB credential.
* `state` - The credential's current state. After creating a DB credential, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `DbCredential` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this credential will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the DB credential belongs to.

