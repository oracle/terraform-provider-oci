---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegation_subscriptions"
sidebar_current: "docs-oci-datasource-delegate_access_control-delegation_subscriptions"
description: |-
  Provides the list of Delegation Subscriptions in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_delegation_subscriptions
This data source provides the list of Delegation Subscriptions in Oracle Cloud Infrastructure Delegate Access Control service.

Lists the Delegation Subscriptions in Delegation Control.


## Example Usage

```hcl
data "oci_delegate_access_control_delegation_subscriptions" "test_delegation_subscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.delegation_subscription_display_name
	state = var.delegation_subscription_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return Delegation Subscription resources that match the given display name.
* `state` - (Optional) A filter to return only Delegation Subscription resources whose lifecycleState matches the given Delegation Subscription lifecycle state.


## Attributes Reference

The following attributes are exported:

* `delegation_subscription_summary_collection` - The list of delegation_subscription_summary_collection.

### DelegationSubscription Reference

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

