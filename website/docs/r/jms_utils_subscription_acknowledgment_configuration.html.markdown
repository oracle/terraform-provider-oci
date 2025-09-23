---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_subscription_acknowledgment_configuration"
sidebar_current: "docs-oci-resource-jms_utils-subscription_acknowledgment_configuration"
description: |-
  Provides the Subscription Acknowledgment Configuration resource in Oracle Cloud Infrastructure Jms Utils service
---

# oci_jms_utils_subscription_acknowledgment_configuration
This resource provides the Subscription Acknowledgment Configuration resource in Oracle Cloud Infrastructure Jms Utils service.

Updates the configuration for subscription acknowledgment.

## Example Usage

```hcl
resource "oci_jms_utils_subscription_acknowledgment_configuration" "test_subscription_acknowledgment_configuration" {

	#Optional
	compartment_id = var.compartment_id
	is_acknowledged = var.subscription_acknowledgment_configuration_is_acknowledged
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `is_acknowledged` - (Optional) (Updatable) Flag to determine whether the subscription was acknowledged or not.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `acknowledged_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal that ackwnoledged the subscription.
* `is_acknowledged` - Flag to determine whether the subscription was acknowledged or not.
* `time_acknowledged` - The date and time the subscription was acknowledged (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscription Acknowledgment Configuration
	* `update` - (Defaults to 20 minutes), when updating the Subscription Acknowledgment Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Subscription Acknowledgment Configuration


## Import

SubscriptionAcknowledgmentConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_jms_utils_subscription_acknowledgment_configuration.test_subscription_acknowledgment_configuration "id"
```

