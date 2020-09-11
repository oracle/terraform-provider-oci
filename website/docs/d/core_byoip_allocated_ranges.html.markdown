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

Lists the ByoipAllocatedRange objects for the ByoipRange.
Each ByoipAllocatedRange object has a CIDR block part of the ByoipRange and the PublicIpPool it is assigned to.


## Example Usage

```hcl
data "oci_core_byoip_allocated_ranges" "test_byoip_allocated_ranges" {
	#Required
	byoip_range_id = oci_core_byoip_range.test_byoip_range.id
}
```

## Argument Reference

The following arguments are supported:

* `byoip_range_id` - (Required) The OCID of the Byoip Range object.


## Attributes Reference

The following attributes are exported:

* `byoip_allocated_range_collection` - The list of byoip_allocated_range_collection.

### ByoipAllocatedRange Reference

The following attributes are exported:

* `items` - list of Byoip allocated ranges as part of public IP pool
	* `cidr_block` - The address range part of the ByoipRange which is used for a publicIpPool.
	* `public_ip_pool_id` - The OCID of the PublicIpPool containing the part of the Byoip range. 

