---
layout: "oci"
page_title: "OCI: oci_email_suppression"
sidebar_current: "docs-oci-datasource-email-suppression"
description: |-
Provides details about a specific Suppression
---

# Data Source: oci_email_suppression
The Suppression data source provides details about a specific Suppression

Gets the details of a suppressed recipient email address for a given
`suppressionId`. Each suppression is given a unique OCID.


## Example Usage

```hcl
data "oci_email_suppression" "test_suppression" {
	#Required
	suppression_id = "${var.suppression_suppression_id}"
}
```

## Argument Reference

The following arguments are supported:

* `suppression_id` - (Required) The unique OCID of the suppression.


## Attributes Reference

The following attributes are exported:

* `email_address` - The email address of the suppression.
* `id` - The unique OCID of the suppression.
* `reason` - The reason that the email address was suppressed. For more information on the types of bounces, see [Suppresion List](https://docs.us-phoenix-1.oraclecloud.com/Content/Email/Concepts/emaildeliveryoverview.htm#suppressionlist).
* `time_created` - The date and time a recipient's email address was added to the suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 

