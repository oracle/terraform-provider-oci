---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_statement_email_recipients_group"
sidebar_current: "docs-oci-resource-metering_computation-usage_statement_email_recipients_group"
description: |-
  Provides the Usage Statement Email Recipients Group resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_usage_statement_email_recipients_group
This resource provides the Usage Statement Email Recipients Group resource in Oracle Cloud Infrastructure Metering Computation service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/usage/latest/UsageStatementEmailRecipientsGroup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/metering_computation

Creates a list of email recipients that can receive usage statements for the subscription.


## Example Usage

```hcl
resource "oci_metering_computation_usage_statement_email_recipients_group" "test_usage_statement_email_recipients_group" {
	#Required
	compartment_id = var.compartment_id
	recipients_list {
		#Required
		email_id = oci_metering_computation_email.test_email.id
		state = var.usage_statement_email_recipients_group_recipients_list_state

		#Optional
		first_name = var.usage_statement_email_recipients_group_recipients_list_first_name
		last_name = var.usage_statement_email_recipients_group_recipients_list_last_name
	}
	subscription_id = oci_onesubscription_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The customer tenancy.
* `recipients_list` - (Required) (Updatable) The list of recipients that will receive usage statement emails.
	* `email_id` - (Required) (Updatable) The recipient email address.
	* `first_name` - (Optional) (Updatable) The recipient first name.
	* `last_name` - (Optional) (Updatable) The recipient last name.
	* `state` - (Required) (Updatable) The email recipient lifecycle state.
* `subscription_id` - (Required) The usage statement subscription unique OCID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Usage Statement Email Recipients Group
	* `update` - (Defaults to 20 minutes), when updating the Usage Statement Email Recipients Group
	* `delete` - (Defaults to 20 minutes), when destroying the Usage Statement Email Recipients Group


## Import

UsageStatementEmailRecipientsGroups can be imported using the `id`, e.g.

```
$ terraform import oci_metering_computation_usage_statement_email_recipients_group.test_usage_statement_email_recipients_group "usageStatements/{subscriptionId}/emailRecipientsGroups/{emailRecipientsGroupId}/compartmentId/{compartmentId}" 
```

