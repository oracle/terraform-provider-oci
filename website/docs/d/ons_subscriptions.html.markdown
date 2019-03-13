---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_subscriptions"
sidebar_current: "docs-oci-datasource-ons-subscriptions"
description: |-
  Provides the list of Subscriptions in Oracle Cloud Infrastructure Ons service
---

# Data Source: oci_ons_subscriptions
This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Ons service.

Lists the subscriptions in the specified compartment or topic. 


## Example Usage

```hcl
data "oci_ons_subscriptions" "test_subscriptions" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	topic_id = "${oci_ons_notification_topic.test_notification_topic.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `topic_id` - (Optional) Return all subscriptions that are subscribed to the given topic OCID. Either this query parameter or the compartmentId query parameter must be set. 


## Attributes Reference

The following attributes are exported:

* `subscriptions` - The list of subscriptions.

### Subscription Reference

The following attributes are exported:

* `created_time` - The time when this suscription was created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `delivery_policy` - 
	* `backoff_retry_policy` - 
		* `max_retry_duration` - The maximum retry duration in milliseconds.
		* `policy_type` - The type of delivery policy. Default value: EXPONENTIAL. 
* `endpoint` - The endpoint of the subscription. Valid values depend on the protocol.  For EMAIL, only an email address is valid. For HTTPS, only a PagerDuty URL is valid. A URL cannot exceed 512 characters. Avoid entering confidential information. 
* `etag` - For optimistic concurrency control. See `if-match`. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription. 
* `protocol` - The protocol used for the subscription. Valid values: EMAIL, HTTPS. 
* `state` - The lifecycle state of the subscription. Default value for a newly created subscription: PENDING. 
* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic. 

