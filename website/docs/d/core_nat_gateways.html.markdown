---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_nat_gateways"
sidebar_current: "docs-oci-datasource-core-nat_gateways"
description: |-
  Provides the list of Nat Gateways in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_nat_gateways
This data source provides the list of Nat Gateways in Oracle Cloud Infrastructure Core service.

Lists the NAT gateways in the specified compartment. You may optionally specify a VCN OCID
to filter the results by VCN.


## Example Usage

```hcl
data "oci_core_nat_gateways" "test_nat_gateways" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.nat_gateway_display_name
	state = var.nat_gateway_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `nat_gateways` - The list of nat_gateways.

### NatGateway Reference

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

