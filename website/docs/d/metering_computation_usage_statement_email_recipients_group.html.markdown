---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_statement_email_recipients_group"
sidebar_current: "docs-oci-datasource-metering_computation-usage_statement_email_recipients_group"
description: |-
  Provides details about a specific Usage Statement Email Recipients Group in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_usage_statement_email_recipients_group
This data source provides details about a specific Usage Statement Email Recipients Group resource in Oracle Cloud Infrastructure Metering Computation service.

Return the saved usage statement email recipient group.


## Example Usage

```hcl
data "oci_metering_computation_usage_statement_email_recipients_group" "test_usage_statement_email_recipients_group" {
	#Required
	compartment_id = var.compartment_id
	email_recipients_group_id = oci_identity_group.test_group.id
	subscription_id = oci_onesubscription_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment ID in which to list resources.
* `email_recipients_group_id` - (Required) The email recipient group OCID.
* `subscription_id` - (Required) The UsageStatement Subscription unique OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The customer tenancy OCID.
* `id` - The usage statement email recipients group OCID.
* `recipients_list` - The list of recipient will receive the usage statement email.
	* `email_id` - the email of the recipient.
	* `first_name` - the first name of the recipient.
	* `last_name` - the last name of the recipient.
	* `state` - The email recipient lifecycle state.
* `state` - The email recipient group lifecycle state.

