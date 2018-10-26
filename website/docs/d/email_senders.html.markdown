---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_senders"
sidebar_current: "docs-oci-datasource-email-senders"
description: |-
  Provides the list of Senders in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_senders
This data source provides the list of Senders in Oracle Cloud Infrastructure Email service.

Gets a collection of approved sender email addresses and sender IDs.


## Example Usage

```hcl
data "oci_email_senders" "test_senders" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	email_address = "${var.sender_email_address}"
	state = "${var.sender_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment.
* `email_address` - (Optional) The email address of the approved sender.
* `state` - (Optional) The current state of a sender.


## Attributes Reference

The following attributes are exported:

* `senders` - The list of senders.

### Sender Reference

The following attributes are exported:

* `email_address` - The email address of the sender.
* `id` - The unique OCID of the sender.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/emaildeliveryoverview.htm#spf). 
* `state` - The current status of the approved sender.
* `time_created` - The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 

