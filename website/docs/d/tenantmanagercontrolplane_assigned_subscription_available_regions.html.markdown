---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_assigned_subscription_available_regions"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-assigned_subscription_available_regions"
description: |-
  Provides the list of Assigned Subscription Available Regions in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_assigned_subscription_available_regions
This data source provides the list of Assigned Subscription Available Regions in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

List of available regions for a given assigned subscription.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_assigned_subscription_available_regions" "test_assigned_subscription_available_regions" {
	#Required
	assigned_subscription_id = var.assigned_subscription_id
}
```

## Argument Reference

The following arguments are supported:

* `assigned_subscription_id` - (Required) OCID of the assigned Oracle Cloud Subscription.


## Attributes Reference

The following attributes are exported:

* `available_region_collection` - The list of available_region_collection.

### AssignedSubscriptionAvailableRegion Reference

The following attributes are exported:

* `items` - Array containing available region items.
	* `region_name` - Region availability for the subscription.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
