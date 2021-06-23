---
subcategory: "Notifications"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_subscription"
sidebar_current: "docs-oci-resource-ons-subscription"
description: |-
  Provides the Subscription resource in Oracle Cloud Infrastructure Notifications service
---

# oci_ons_subscription
This resource provides the Subscription resource in Oracle Cloud Infrastructure Notifications service.

Creates a subscription for the specified topic and sends a subscription confirmation URL to the endpoint. The subscription remains in "Pending" status until it has been confirmed.
For information about confirming subscriptions, see
[To confirm a subscription](https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#confirmSub).

Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.


## Example Usage

```hcl
resource "oci_ons_subscription" "test_subscription" {
	#Required
	compartment_id = var.compartment_id
	endpoint = var.subscription_endpoint
	protocol = var.subscription_protocol
	topic_id = oci_ons_notification_topic.test_notification_topic.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `endpoint` - (Required) A locator that corresponds to the subscription protocol. For example, an email address for a subscription that uses the `EMAIL` protocol, or a URL for a subscription that uses an HTTP-based protocol. HTTP-based protocols use URL endpoints that begin with "http:" or "https:". A URL cannot exceed 512 characters. Avoid entering confidential information.

	For protocol-specific endpoint formats and steps to get or create endpoints, see [To create a subscription](https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#createSub). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `protocol` - (Required) The protocol used for the subscription.

	Allowed values:
	* `CUSTOM_HTTPS`
	* `EMAIL`
	* `HTTPS` (deprecated; for PagerDuty endpoints, use `PAGERDUTY`)
	* `ORACLE_FUNCTIONS`
	* `PAGERDUTY`
	* `SLACK`
	* `SMS`

	For information about subscription protocols, see [To create a subscription](https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#createSub). 
* `topic_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic for the subscription. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription. 
* `created_time` - The time when this suscription was created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `delivery_policy` - The delivery policy of the subscription. Stored as a JSON string.
* `endpoint` - A locator that corresponds to the subscription protocol. For example, an email address for a subscription that uses the `EMAIL` protocol, or a URL for a subscription that uses an HTTP-based protocol. 
* `etag` - For optimistic concurrency control. See `if-match`. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription. 
* `protocol` - The protocol used for the subscription.

	Allowed values:
	* `CUSTOM_HTTPS`
	* `EMAIL`
	* `HTTPS` (deprecated; for PagerDuty endpoints, use `PAGERDUTY`)
	* `ORACLE_FUNCTIONS`
	* `PAGERDUTY`
	* `SLACK`
	* `SMS`

	For information about subscription protocols, see [To create a subscription](https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#createSub). 
* `state` - The lifecycle state of the subscription. The status of a new subscription is PENDING; when confirmed, the subscription status changes to ACTIVE. 
* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscription
	* `update` - (Defaults to 20 minutes), when updating the Subscription
	* `delete` - (Defaults to 20 minutes), when destroying the Subscription


## Import

Subscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_ons_subscription.test_subscription "id"
```

