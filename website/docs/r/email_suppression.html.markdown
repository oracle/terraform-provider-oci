---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_suppression"
sidebar_current: "docs-oci-resource-email-suppression"
description: |-
  Provides the Suppression resource in Oracle Cloud Infrastructure Email service
---

# oci_email_suppression
This resource provides the Suppression resource in Oracle Cloud Infrastructure Email service.

Adds recipient email addresses to the suppression list for a tenancy.
Addresses added to the suppression list via the API are denoted as
"MANUAL" in the `reason` field. *Note:* All email addresses added to the
suppression list are normalized to include only lowercase letters.


## Example Usage

```hcl
resource "oci_email_suppression" "test_suppression" {
	#Required
	compartment_id = var.tenancy_ocid
	email_address = var.suppression_email_address
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID. 
* `email_address` - (Required) The recipient email address of the suppression.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID. 
* `email_address` - The email address of the suppression.
* `error_detail` - The specific error message returned by a system that resulted in the suppression. This message is usually an SMTP error code with additional descriptive text. Not provided for all types of suppressions. 
* `error_source` - DNS name of the source of the error that caused the suppression. Will be set to either the remote-mta or reporting-mta field from a delivery status notification (RFC 3464) when available. Not provided for all types of suppressions, and not always known.

	Note: Most SMTP errors that cause suppressions come from software run by email receiving systems rather than from Oracle Cloud Infrastructure email delivery itself. 
* `id` - The unique OCID of the suppression.
* `message_id` - The value of the Message-ID header from the email that triggered a suppression. This value is as defined in RFC 5322 section 3.6.4, excluding angle-brackets. Not provided for all types of suppressions. 
* `reason` - The reason that the email address was suppressed. For more information on the types of bounces, see [Suppression List](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components).
* `time_created` - The date and time a recipient's email address was added to the suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_last_suppressed` - The last date and time the suppression prevented submission in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Suppression
	* `update` - (Defaults to 20 minutes), when updating the Suppression
	* `delete` - (Defaults to 20 minutes), when destroying the Suppression


## Import

Suppressions can be imported using the `id`, e.g.

```
$ terraform import oci_email_suppression.test_suppression "id"
```

