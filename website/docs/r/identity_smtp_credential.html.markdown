---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_smtp_credential"
sidebar_current: "docs-oci-resource-identity-smtp_credential"
description: |-
  Provides the Smtp Credential resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_smtp_credential
This resource provides the Smtp Credential resource in Oracle Cloud Infrastructure Identity service.

Creates a new SMTP credential for the specified user. An SMTP credential has an SMTP user name and an SMTP password.
You must specify a *description* for the SMTP credential (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateSmtpCredential](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/SmtpCredentialSummary/UpdateSmtpCredential).


## Example Usage

```hcl
resource "oci_identity_smtp_credential" "test_smtp_credential" {
	#Required
	description = var.smtp_credential_description
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) (Updatable) The description you assign to the SMTP credentials during creation. Does not have to be unique, and it's changeable. 
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - The description you assign to the SMTP credential. Does not have to be unique, and it's changeable.
* `id` - The OCID of the SMTP credential.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `password` - The SMTP password. 
* `state` - The credential's current state.
* `time_created` - Date and time the `SmtpCredential` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this credential will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the SMTP credential belongs to.
* `username` - The SMTP user name. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Smtp Credential
	* `update` - (Defaults to 20 minutes), when updating the Smtp Credential
	* `delete` - (Defaults to 20 minutes), when destroying the Smtp Credential


## Import

SmtpCredentials can be imported using the `id`, e.g.

```
$ terraform import oci_identity_smtp_credential.test_smtp_credential "users/{userId}/smtpCredentials/{smtpCredentialId}" 
```

