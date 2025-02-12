---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_subscription_mappings"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-subscription_mappings"
description: |-
  Provides the list of Subscription Mappings in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_subscription_mappings
This data source provides the list of Subscription Mappings in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Lists the subscription mappings for all the subscriptions owned by a given compartmentId. Only the root compartment is allowed.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_subscription_mappings" "test_subscription_mappings" {
	#Required
	subscription_id = oci_tenantmanagercontrolplane_subscription.test_subscription.id

	#Optional
	compartment_id = var.compartment_id
	state = var.subscription_mapping_state
	subscription_mapping_id = oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `state` - (Optional) The lifecycle state of the resource.
* `subscription_id` - (Required) OCID of the subscription.
* `subscription_mapping_id` - (Optional) A unique ID for subscription and tenancy mapping.


## Attributes Reference

The following attributes are exported:

* `subscription_mapping_collection` - The list of subscription_mapping_collection.

### SubscriptionMapping Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment. Always a tenancy OCID.
* `id` - OCID of the mapping between subscription and compartment identified by the tenancy.
* `is_explicitly_assigned` - Denotes if the subscription is explicity assigned to the root compartment or tenancy.
* `state` - Lifecycle state of the subscriptionMapping.
* `subscription_id` - OCID of the subscription.
* `time_created` - Date-time when subscription mapping was created.
* `time_terminated` - Date-time when subscription mapping was terminated.
* `time_updated` - Date-time when subscription mapping was updated.

