---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_region_subscriptions"
sidebar_current: "docs-oci-datasource-identity-region_subscriptions"
description: |-
  Provides the list of Region Subscriptions in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_region_subscriptions
This data source provides the list of Region Subscriptions in Oracle Cloud Infrastructure Identity service.

Lists the region subscriptions for the specified tenancy.

## Example Usage

```hcl
data "oci_identity_region_subscriptions" "test_region_subscriptions" {
	#Required
	tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `tenancy_id` - (Required) The OCID of the tenancy.


## Attributes Reference

The following attributes are exported:

* `region_subscriptions` - The list of region_subscriptions.

### RegionSubscription Reference

The following attributes are exported:

* `is_home_region` - Indicates if the region is the home region or not.
* `region_key` - The region's key. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported 3-letter region codes.  Example: `PHX` 
* `region_name` - The region's name. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 
* `status` - The region subscription status.

