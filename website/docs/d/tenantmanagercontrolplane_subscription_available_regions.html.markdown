---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_subscription_available_regions"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-subscription_available_regions"
description: |-
  Provides the list of Subscription Available Regions in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_subscription_available_regions
This data source provides the list of Subscription Available Regions in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

List the available regions based on subscription ID.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_subscription_available_regions" "test_subscription_available_regions" {
	#Required
	subscription_id = oci_tenantmanagercontrolplane_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) OCID of the subscription.


## Attributes Reference

The following attributes are exported:

* `available_region_collection` - The list of available_region_collection.

### SubscriptionAvailableRegion Reference

The following attributes are exported:

* `items` - Array containing available region items.
	* `region_name` - Region availability for the subscription.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

