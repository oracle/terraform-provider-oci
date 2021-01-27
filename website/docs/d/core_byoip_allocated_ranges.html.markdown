---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_byoip_allocated_ranges"
sidebar_current: "docs-oci-datasource-core-byoip_allocated_ranges"
description: |-
  Provides the list of Byoip Allocated Ranges in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_byoip_allocated_ranges
This data source provides the list of Byoip Allocated Ranges in Oracle Cloud Infrastructure Core service.

Lists the subranges of a BYOIP CIDR block currently allocated to an IP pool.
Each `ByoipAllocatedRange` object also lists the IP pool where it is allocated.


## Example Usage

```hcl
data "oci_core_byoip_allocated_ranges" "test_byoip_allocated_ranges" {
	#Required
	byoip_range_id = oci_core_byoip_range.test_byoip_range.id
}
```

## Argument Reference

The following arguments are supported:

* `byoip_range_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource containing the BYOIP CIDR block.


## Attributes Reference

The following attributes are exported:

* `byoip_allocated_range_collection` - The list of byoip_allocated_range_collection.

### ByoipAllocatedRange Reference

The following attributes are exported:

* `items` - A list of subranges of a BYOIP CIDR block allocated to an IP pool.
	* `cidr_block` - The BYOIP CIDR block range or subrange allocated to an IP pool. This could be all or part of a BYOIP CIDR block.
	* `public_ip_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IP pool containing the CIDR block. 

