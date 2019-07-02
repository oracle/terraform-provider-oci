---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ons_subscription"
sidebar_current: "docs-oci-datasource-ons-subscription"
description: |-
  Provides details about a specific Subscription in Oracle Cloud Infrastructure Ons service
---

# Data Source: oci_ons_subscription
This data source provides details about a specific Subscription resource in Oracle Cloud Infrastructure Ons service.

Gets the specified subscription's configuration information.

Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.


## Example Usage

```hcl
data "oci_ons_subscription" "test_subscription" {
	#Required
	subscription_id = "${oci_ons_subscription.test_subscription.id}"
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription to retrieve. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription. 
* `created_time` - The time when this suscription was created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `delivery_policy` - The delivery policy of the subscription. Stored as a JSON string.
* `endpoint` - The endpoint of the subscription. Valid values depend on the protocol. For EMAIL, only an email address is valid. For HTTPS, only a PagerDuty URL is valid. A URL cannot exceed 512 characters. 
* `etag` - For optimistic concurrency control. See `if-match`. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription. 
* `protocol` - The protocol used for the subscription. Valid values: EMAIL, HTTPS. 
* `state` - The lifecycle state of the subscription. The status of a new subscription is PENDING; when confirmed, the subscription status changes to ACTIVE. 
* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic. 

