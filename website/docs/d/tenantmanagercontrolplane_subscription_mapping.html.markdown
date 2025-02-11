---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_subscription_mapping"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-subscription_mapping"
description: |-
  Provides details about a specific Subscription Mapping in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_subscription_mapping
This data source provides details about a specific Subscription Mapping resource in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Get the subscription mapping details by subscription mapping ID.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_subscription_mapping" "test_subscription_mapping" {
	#Required
	subscription_mapping_id = oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription_mapping_id` - (Required) OCID of the subscriptionMappingId.


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

