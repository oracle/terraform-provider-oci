---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_smtp_credentials"
sidebar_current: "docs-oci-datasource-identity-smtp_credentials"
description: |-
  Provides the list of Smtp Credentials in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_smtp_credentials
This data source provides the list of Smtp Credentials in Oracle Cloud Infrastructure Identity service.

Lists the SMTP credentials for the specified user. The returned object contains the credential's OCID, 
the SMTP user name but not the SMTP password. The SMTP password is returned only upon creation.


## Example Usage

```hcl
data "oci_identity_smtp_credentials" "test_smtp_credentials" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `smtp_credentials` - The list of smtp_credentials.

### SmtpCredential Reference

The following attributes are exported:

* `description` - The description you assign to the SMTP credential. Does not have to be unique, and it's changeable.
* `id` - The OCID of the SMTP credential.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The credential's current state.
* `time_created` - Date and time the `SmtpCredential` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this credential will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the SMTP credential belongs to.
* `username` - The SMTP user name. 

