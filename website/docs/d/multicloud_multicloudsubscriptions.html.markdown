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

Lists activated Multicloud subscriptions in the specified compartment. For more information, see
[Listing Multicloud Subscriptions](https://docs.cloud.oracle.com/iaas/Content/multicloud-hub/list-subscriptions.htm).


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
	* `classic_subscription_id` - Subscription ID for Oracle Cloud Infrastructure and partner cloud in classic format.
	* `csp_additional_properties` - Properties specific to the cloud service provider. For example, AzureSubnetId for Azure. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `partner_cloud_account_identifier` - The partner cloud account ID.
	* `partner_cloud_tenant_identifier` - The partner cloud tenant ID.
	* `payment_plan` - Payment plan for the subscription.
	* `service_name` - The cloud service provider.
	* `lifecycle_state` - The current state of the subscription.
	* `subscription_id` - URL to the subscription details page. Example: `https://{console-url}/org-mgmt/subscription/ocid1.organizationssubscription.oc1.iad.exampleuniqueid?region=us-ashburn-1`. 
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time that the subscription was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_end_date` - The end date and time for the subscription, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_linked_date` - The date and time that the Multicloud base compartment was created, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The date and time that the subscription was updated, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).

