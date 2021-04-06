---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_sender"
sidebar_current: "docs-oci-datasource-email-sender"
description: |-
  Provides details about a specific Sender in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_sender
This data source provides details about a specific Sender resource in Oracle Cloud Infrastructure Email service.

Gets an approved sender for a given `senderId`.

## Example Usage

```hcl
data "oci_email_sender" "test_sender" {
	#Required
	sender_id = oci_email_sender.test_sender.id
}
```

## Argument Reference

The following arguments are supported:

* `sender_id` - (Required) The unique OCID of the sender.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID for the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `email_address` - The email address of the sender.
* `email_domain_id` - The email domain used to assert responsibility for emails sent from this sender. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The unique OCID of the sender.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components). 
* `state` - The current status of the approved sender.
* `time_created` - The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 

