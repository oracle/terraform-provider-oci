---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vlan"
sidebar_current: "docs-oci-resource-core-vlan"
description: |-
  Provides the Vlan resource in Oracle Cloud Infrastructure Core service
---

# oci_core_vlan
This resource provides the Vlan resource in Oracle Cloud Infrastructure Core service.

Creates a VLAN in the specified VCN and the specified compartment.


## Example Usage

```hcl
resource "oci_core_vlan" "test_vlan" {
	#Required
	availability_domain = "${var.vlan_availability_domain}"
	cidr_block = "${var.vlan_cidr_block}"
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.vlan_display_name}"
	freeform_tags = {"Department"= "Finance"}
	nsg_ids = "${var.vlan_nsg_ids}"
	route_table_id = "${oci_core_route_table.test_route_table.id}"
	vlan_tag = "${var.vlan_vlan_tag}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the VLAN.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - (Required) The range of IPv4 addresses that will be used for layer 3 communication with hosts outside the VLAN.  Example: `192.0.2.0/24` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment to contain the VLAN.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A descriptive name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `nsg_ids` - (Optional) (Updatable) A list of the OCIDs of the network security groups (NSGs) to add all VNICs in the VLAN to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
* `route_table_id` - (Optional) (Updatable) The OCID of the route table the VLAN will use. If you don't provide a value, the VLAN uses the VCN's default route table. 
* `vcn_id` - (Required) The OCID of the VCN to contain the VLAN.
* `vlan_tag` - (Optional) The IEEE 802.1Q VLAN tag for this VLAN. The value must be unique across all VLANs in the VCN. If you don't provide a value, Oracle assigns one. You cannot change the value later. VLAN tag 0 is reserved for use by Oracle. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the VLAN.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - The range of IPv4 addresses that will be used for layer 3 communication with hosts outside the VLAN.  Example: `192.168.1.0/24` 
* `compartment_id` - The OCID of the compartment containing the VLAN.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The VLAN's Oracle ID (OCID).
* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to use with this VLAN. All VNICs in the VLAN belong to these NSGs. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
* `route_table_id` - The OCID of the route table that the VLAN uses.
* `state` - The VLAN's current state.
* `time_created` - The date and time the VLAN was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the VLAN is in.
* `vlan_tag` - The IEEE 802.1Q VLAN tag of this VLAN.  Example: `100` 

## Import

Vlans can be imported using the `id`, e.g.

```
$ terraform import oci_core_vlan.test_vlan "id"
```

