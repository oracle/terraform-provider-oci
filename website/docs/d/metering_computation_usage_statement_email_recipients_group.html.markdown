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

Returns the saved usage statement email recipients group.


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
* `email_recipients_group_id` - (Required) The email recipients group OCID.
* `subscription_id` - (Required) The usage statement subscription unique OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The customer tenancy OCID.
* `id` - The usage statement email recipients group OCID.
* `recipients_list` - The list of recipients that will receive usage statement emails.
	* `email_id` - The recipient email address.
	* `first_name` - The recipient first name.
	* `last_name` - The recipient last name.
	* `state` - The email recipient lifecycle state.
* `state` - The email recipients group lifecycle state.

