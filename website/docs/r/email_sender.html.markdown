---
layout: "oci"
page_title: "OCI: oci_email_sender"
sidebar_current: "docs-oci-resource-email-sender"
description: |-
  Creates and manages an OCI Sender
---

# oci_email_sender
The `oci_email_sender` resource creates and manages an OCI Sender

Creates a sender for a tenancy in a given compartment.

## Example Usage

```hcl
resource "oci_email_sender" "test_sender" {
	#Required
	compartment_id = "${var.compartment_id}"
	email_address = "${var.sender_email_address}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the sender.
* `email_address` - (Required) The email address of the sender.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `email_address` - The email address of the sender.
* `id` - The unique OCID of the sender.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.us-phoenix-1.oraclecloud.com/Content/Email/Concepts/emaildeliveryoverview.htm#spf). 
* `state` - The current status of the approved sender.
* `time_created` - The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 

## Import

Senders can be imported using the `id`, e.g.

```
$ terraform import oci_email_sender.test_sender "id"
```
