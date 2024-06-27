---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegation_subscription"
sidebar_current: "docs-oci-datasource-delegate_access_control-delegation_subscription"
description: |-
  Provides details about a specific Delegation Subscription in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_delegation_subscription
This data source provides details about a specific Delegation Subscription resource in Oracle Cloud Infrastructure Delegate Access Control service.

Gets a DelegationSubscription by identifier

## Example Usage

```hcl
data "oci_delegate_access_control_delegation_subscription" "test_delegation_subscription" {
	#Required
	delegation_subscription_id = oci_delegate_access_control_delegation_subscription.test_delegation_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `delegation_subscription_id` - (Required) unique Delegation Subscription identifier


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

