---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegation_subscription"
sidebar_current: "docs-oci-resource-delegate_access_control-delegation_subscription"
description: |-
  Provides the Delegation Subscription resource in Oracle Cloud Infrastructure Delegate Access Control service
---

# oci_delegate_access_control_delegation_subscription
This resource provides the Delegation Subscription resource in Oracle Cloud Infrastructure Delegate Access Control service.

Creates Delegation Subscription in Delegation Control.


## Example Usage

```hcl
resource "oci_delegate_access_control_delegation_subscription" "test_delegation_subscription" {
	#Required
	compartment_id = var.compartment_id
	service_provider_id = oci_delegate_access_control_service_provider.test_service_provider.id
	subscribed_service_type = var.delegation_subscription_subscribed_service_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.delegation_subscription_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the Delegation Control.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the Delegation Subscription. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `service_provider_id` - (Required) Unique identifier of the Service Provider.
* `subscribed_service_type` - (Required) Subscribed Service Provider Service Type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the Delegation Subscription.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the Delegation Subscription. 
* `display_name` - Display name
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - Unique identifier for the Delegation Subscription.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `service_provider_id` - Unique identifier of the Service Provider.
* `state` - The current lifecycle state of the Service Provider.
* `subscribed_service_type` - Subscribed Service Provider Service Type.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Service Provider was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Service Provider was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Delegation Subscription
	* `update` - (Defaults to 20 minutes), when updating the Delegation Subscription
	* `delete` - (Defaults to 20 minutes), when destroying the Delegation Subscription


## Import

DelegationSubscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_delegate_access_control_delegation_subscription.test_delegation_subscription "id"
```

