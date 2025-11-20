---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_multicloudsubscriptions"
sidebar_current: "docs-oci-datasource-multicloud-multicloudsubscriptions"
description: |-
  Provides the list of Multicloud subscriptions in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_multicloudsubscriptions
This data source provides the list of Multicloud subscriptions in Oracle Cloud Infrastructure Multicloud service.

Gets a list of Multicloud subscriptions.


## Example Usage

```hcl
data "oci_multicloud_multicloudsubscriptions" "test_multicloudsubscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.


## Attributes Reference

The following attributes are exported:

* `multicloud_subscription_collection` - The list of Multicloudsubscription.

### Multicloudsubscription Reference

The following attributes are exported:

* `items` - List of MulticloudSubscriptionSummary.
	* `active_commitment` - Total value for the subscription.
	* `classic_subscription_id` - Subscription ID for Oracle Cloud Infrastructure and Partner cloud in classic format.
	* `csp_additional_properties` - CSP Specific Additional Properties, AzureSubnetId for Azure
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `partner_cloud_account_identifier` - The partner cloud account ID.
	* `payment_plan` - Payment plan for the subscription.
	* `service_name` - The serviceName that externalLocation map object belongs to.
	* `lifecycle_state` - The current state of the subscription.
	* `subscription_id` - URL to the subscription page https://{console-url}/org-mgmt/subscription/ocid1.organizationssubscription.oc1.iad.amaaaaaapf266qyaqohz27zvh45jzaielgwojo53bh24s7cy5q5g7fiknpxa?region=us-ashburn-1.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the subscription was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_end_date` - The date and time when the subscription is finishing, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_linked_date` - The date and time when the multicloud link was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The date and time the subscription was updated, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

