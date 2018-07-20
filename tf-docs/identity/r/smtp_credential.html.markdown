---
layout: "oci"
page_title: "OCI: oci_identity_smtp_credential"
sidebar_current: "docs-oci-resource-smtp_credential"
description: |-
Creates and manages an OCI SmtpCredential
---

# oci_identity_smtp_credential
The `oci_identity_smtp_credential` resource creates and manages an OCI SmtpCredential

Creates a new SMTP credential for the specified user. An SMTP credential has an SMTP user name and an SMTP password.
You must specify a *description* for the SMTP credential (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateSmtpCredential](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SmtpCredentialSummary/UpdateSmtpCredential).


## Example Usage

```hcl
resource "oci_identity_smtp_credential" "test_smtp_credential" {
	#Required
	description = "${var.smtp_credential_description}"
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) (Updatable) The description you assign to the SMTP credentials during creation. Does not have to be unique, and it's changeable. 


## Attributes Reference

The following attributes are exported:

* `description` - The description you assign to the SMTP credential. Does not have to be unique, and it's changeable.
* `id` - The OCID of the SMTP credential.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The credential's current state. After creating a SMTP credential, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `SmtpCredential` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this credential will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the SMTP credential belongs to.
* `username` - The SMTP user name. 
