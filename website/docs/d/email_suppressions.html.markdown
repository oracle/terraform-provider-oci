---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_suppressions"
sidebar_current: "docs-oci-datasource-email-suppressions"
description: |-
  Provides the list of Suppressions in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_suppressions
This data source provides the list of Suppressions in Oracle Cloud Infrastructure Email service.

Gets a list of suppressed recipient email addresses for a user. The
`compartmentId` for suppressions must be a tenancy OCID. The returned list
is sorted by creation time in descending order.


## Example Usage

```hcl
data "oci_email_suppressions" "test_suppressions" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	email_address = var.suppression_email_address
	time_created_greater_than_or_equal_to = var.suppression_time_created_greater_than_or_equal_to
	time_created_less_than = var.suppression_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment.
* `email_address` - (Optional) The email address of the suppression.
* `time_created_greater_than_or_equal_to` - (Optional) Search for suppressions that were created within a specific date range, using this parameter to specify the earliest creation date for the returned list (inclusive). Specifying this parameter without the corresponding `timeCreatedLessThan` parameter will retrieve suppressions created from the given `timeCreatedGreaterThanOrEqualTo` to the current time, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for suppressions that were created within a specific date range, using this parameter to specify the latest creation date for the returned list (exclusive). Specifying this parameter without the corresponding `timeCreatedGreaterThanOrEqualTo` parameter will retrieve all suppressions created before the specified end date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `suppressions` - The list of suppressions.

### Suppression Reference

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

