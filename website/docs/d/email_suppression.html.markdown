---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_suppression"
sidebar_current: "docs-oci-datasource-email-suppression"
description: |-
  Provides details about a specific Suppression in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_suppression
This data source provides details about a specific Suppression resource in Oracle Cloud Infrastructure Email service.

Gets the details of a suppressed recipient email address for a given
`suppressionId`. Each suppression is given a unique OCID.


## Example Usage

```hcl
data "oci_email_suppression" "test_suppression" {
	#Required
	suppression_id = oci_email_suppression.test_suppression.id
}
```

## Argument Reference

The following arguments are supported:

* `suppression_id` - (Required) The unique OCID of the suppression.


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

