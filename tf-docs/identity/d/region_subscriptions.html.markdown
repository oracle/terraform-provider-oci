---
layout: "oci"
page_title: "OCI: oci_identity_region_subscriptions"
sidebar_current: "docs-oci-datasource-region_subscriptions"
description: |-
Provides a list of RegionSubscriptions
---
# Data Source: oci_identity_region_subscriptions
The RegionSubscriptions data source allows access to the list of OCI region_subscriptions

Lists the region subscriptions for the specified tenancy.

## Example Usage

```hcl
data "oci_identity_region_subscriptions" "test_region_subscriptions" {
	#Required
	tenancy_id = "${oci_identity_tenancy.test_tenancy.id}"
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
* `region_key` - The region's key.  Allowed values are: - `PHX` - `IAD` - `FRA` - `LHR` 
* `region_name` - The region's name.  Allowed values are: - `us-phoenix-1` - `us-ashburn-1` - `eu-frankurt-1` - `uk-london-1` 
* `state` - The region subscription state.

