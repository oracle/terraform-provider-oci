---
subcategory: "Email"
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
	compartment_id = var.compartment_id

	#Optional
	domain = var.sender_domain
	email_address = var.sender_email_address
	state = var.sender_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID for the compartment.
* `domain` - (Optional) A filter to only return resources that match the given domain exactly.
* `email_address` - (Optional) The email address of the approved sender.
* `state` - (Optional) The current state of a sender.


## Attributes Reference

The following attributes are exported:

* `senders` - The list of senders.

### Sender Reference

The following attributes are exported:

* `compartment_id` - The OCID for the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `email_address` - The email address of the sender.
* `email_domain_id` - The email domain used to assert responsibility for emails sent from this sender. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The unique OCID of the sender.
* `is_spf` - Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components). 
* `state` - The current status of the approved sender.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). 

