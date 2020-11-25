---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_public_ip_pools"
sidebar_current: "docs-oci-datasource-core-public_ip_pools"
description: |-
  Provides the list of Public Ip Pools in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_public_ip_pools
This data source provides the list of Public Ip Pools in Oracle Cloud Infrastructure Core service.

Lists the public IP pools in the specified compartment.
You can filter the list using query parameters.


## Example Usage

```hcl
data "oci_core_public_ip_pools" "test_public_ip_pools" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	byoip_range_id = oci_core_byoip_range.test_byoip_range.id
	display_name = var.public_ip_pool_display_name
}
```

## Argument Reference

The following arguments are supported:

* `byoip_range_id` - (Optional) A filter to return only resources that match the given BYOIP CIDR block. 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `public_ip_pool_collection` - The list of public_ip_pool_collection.

### PublicIpPool Reference

The following attributes are exported:

* `cidr_blocks` - The CIDR blocks added to this pool. This could be all or a portion of a BYOIP CIDR block.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this pool. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
* `state` - The public IP pool's current state.
* `time_created` - The date and time the public IP pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

