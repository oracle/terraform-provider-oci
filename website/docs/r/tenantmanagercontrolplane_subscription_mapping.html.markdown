---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_subscription_mapping"
sidebar_current: "docs-oci-resource-tenantmanagercontrolplane-subscription_mapping"
description: |-
  Provides the Subscription Mapping resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# oci_tenantmanagercontrolplane_subscription_mapping
This resource provides the Subscription Mapping resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/organizations/latest/SubscriptionMapping

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/tenantmanagercontrolplane

Assign the tenancy record identified by the compartment ID to the given subscription ID.

## Example Usage

```hcl
resource "oci_tenantmanagercontrolplane_subscription_mapping" "test_subscription_mapping" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_tenantmanagercontrolplane_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) OCID of the compartment. Always a tenancy OCID.
* `subscription_id` - (Required) OCID of Subscription.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment. Always a tenancy OCID.
* `id` - OCID of the mapping between subscription and compartment identified by the tenancy.
* `is_explicitly_assigned` - Denotes if the subscription is explicity assigned to the root compartment or tenancy.
* `state` - Lifecycle state of the subscriptionMapping.
* `subscription_id` - OCID of the subscription.
* `time_created` - Date-time when subscription mapping was created.
* `time_terminated` - Date-time when subscription mapping was terminated.
* `time_updated` - Date-time when subscription mapping was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscription Mapping
	* `update` - (Defaults to 20 minutes), when updating the Subscription Mapping
	* `delete` - (Defaults to 20 minutes), when destroying the Subscription Mapping


## Import

SubscriptionMappings can be imported using the `id`, e.g.

```
$ terraform import oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping "id"
```

