---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_multicloudpolicies"
sidebar_current: "docs-oci-datasource-multicloud-multicloudpolicies"
description: |-
  Provides the list of Multicloudpolicies in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_multicloudpolicies
This data source provides the list of Multicloudpolicies in Oracle Cloud Infrastructure Multicloud service.

Gets a list of Multicloud IAM Policies.


## Example Usage

```hcl
data "oci_multicloud_multicloudpolicies" "test_multicloudpolicies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.multicloudpolicy_display_name
	is_force_refresh = var.multicloudpolicy_is_force_refresh
	subscription_id = oci_onesubscription_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `is_force_refresh` - (Optional) Refresh the policies.
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.


## Attributes Reference

The following attributes are exported:

* `multicloud_policy_collection` - The list of multicloud_policy_collection.

### Multicloudpolicy Reference

The following attributes are exported:

* `compartment_id` - Tenancy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the query.
* `items` - List of MulticloudPolicySummary.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `groups` - groups required for the particular subscriptionType IAM policy statements required.
	* `policies` - Missing policy definitions.
		* `compartment_id` - Compartment The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the policy is configured.
		* `compartment_name` - Description of the compartment e.g. Base Compartment, Root Compartment
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `description` - Description of the policy purpose.
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `name` - Name of the missing policy.
		* `lifecycle_state` - The current state of the Multicloud Policy.
		* `statements` - IAM policy statements required.
		* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `lifecycle_state` - The current state of the Multicloud Network Alert.
	* `subscription_id` - Compartment The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Subscription
	* `subscription_type` - Oracle Cloud Infrastructure Subscription Type.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 

