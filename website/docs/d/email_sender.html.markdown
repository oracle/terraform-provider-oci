---
layout: "oci"
page_title: "OCI: oci_email_sender"
sidebar_current: "docs-oci-datasource-email-sender"
description: |-
  Provides details about a specific Sender
---

# Data Source: oci_email_sender
The `oci_email_sender` data source provides details about a specific Sender

Gets an approved sender for a given `senderId`.

## Example Usage

```hcl
data "oci_email_sender" "test_sender" {
	#Required
	sender_id = "${oci_email_sender.test_sender.id}"
}
```

## Argument Reference

The following arguments are supported:

* `sender_id` - (Required) The unique OCID of the sender.


## Attributes Reference

The following attributes are exported:

* `email_address` - The email address of the sender.
* `id` - The unique OCID of the sender.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.us-phoenix-1.oraclecloud.com/Content/Email/Concepts/emaildeliveryoverview.htm#spf). 
* `state` - The current status of the approved sender.
* `time_created` - The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 

