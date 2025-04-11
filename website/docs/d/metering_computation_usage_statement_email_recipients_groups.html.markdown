---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_statement_email_recipients_groups"
sidebar_current: "docs-oci-datasource-metering_computation-usage_statement_email_recipients_groups"
description: |-
  Provides the list of Usage Statement Email Recipients Groups in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_usage_statement_email_recipients_groups
This data source provides the list of Usage Statement Email Recipients Groups in Oracle Cloud Infrastructure Metering Computation service.

Returns the saved usage statement email recipients group.


## Example Usage

```hcl
data "oci_metering_computation_usage_statement_email_recipients_groups" "test_usage_statement_email_recipients_groups" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_onesubscription_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment ID in which to list resources.
* `subscription_id` - (Required) The usage statement subscription unique OCID.


## Attributes Reference

The following attributes are exported:

* `email_recipients_group_collection` - The list of email_recipients_group_collection.

### UsageStatementEmailRecipientsGroup Reference

The following attributes are exported:

* `compartment_id` - The customer tenancy OCID.
* `id` - The usage statement email recipients group OCID.
* `recipients_list` - The list of recipients that will receive usage statement emails.
	* `email_id` - The recipient email address.
	* `first_name` - The recipient first name.
	* `last_name` - The recipient last name.
	* `state` - The email recipient lifecycle state.
* `state` - The email recipients group lifecycle state.

