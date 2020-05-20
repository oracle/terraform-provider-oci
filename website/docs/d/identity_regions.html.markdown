---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_regions"
sidebar_current: "docs-oci-datasource-identity-regions"
description: |-
  Provides the list of Regions in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_regions
This data source provides the list of Regions in Oracle Cloud Infrastructure Identity service.

Lists all the regions offered by Oracle Cloud Infrastructure.

## Example Usage

```hcl
data "oci_identity_regions" "test_regions" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `regions` - The list of regions.

### Region Reference

The following attributes are exported:

* `key` - The key of the region. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported 3-letter region codes.  Example: `PHX` 
* `name` - The name of the region. See [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm) for the full list of supported region names.  Example: `us-phoenix-1` 

