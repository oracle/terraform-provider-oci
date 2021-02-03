---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_nat_gateway"
sidebar_current: "docs-oci-datasource-core-nat_gateway"
description: |-
  Provides details about a specific Nat Gateway in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_nat_gateway
This data source provides details about a specific Nat Gateway resource in Oracle Cloud Infrastructure Core service.

Gets the specified NAT gateway's information.

## Example Usage

```hcl
data "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	nat_gateway_id = oci_core_nat_gateway.test_nat_gateway.id
}
```

## Argument Reference

The following arguments are supported:

* `nat_gateway_id` - (Required) The NAT gateway's [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `block_traffic` - Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `true` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the NAT gateway. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NAT gateway. 
* `nat_ip` - The IP address associated with the NAT gateway. 
* `public_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the NAT gateway. 
* `state` - The NAT gateway's current state.
* `time_created` - The date and time the NAT gateway was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the NAT gateway belongs to. 

