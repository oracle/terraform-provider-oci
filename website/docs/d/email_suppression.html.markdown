---
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
	suppression_id = "${oci_email_suppression.test_suppression.id}"
}
```

## Argument Reference

The following arguments are supported:

* `suppression_id` - (Required) The unique OCID of the suppression.


## Attributes Reference

The following attributes are exported:

* `email_address` - The email address of the suppression.
* `id` - The unique OCID of the suppression.
* `reason` - The reason that the email address was suppressed. For more information on the types of bounces, see [Suppresion List](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/emaildeliveryoverview.htm#suppressionlist).
* `time_created` - The date and time a recipient's email address was added to the suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 

