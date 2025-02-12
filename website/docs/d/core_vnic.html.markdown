---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vnic"
sidebar_current: "docs-oci-datasource-core-vnic"
description: |-
  Provides details about a specific Vnic in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_vnic
This data source provides details about a specific Vnic resource in Oracle Cloud Infrastructure Core service.

Gets the information for the specified virtual network interface card (VNIC).
You can get the VNIC [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) from the
[ListVnicAttachments](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/VnicAttachment/ListVnicAttachments)
operation.


## Example Usage

```hcl
data "oci_core_vnic" "test_vnic" {
	#Required
	vnic_id = oci_core_vnic.test_vnic.id
}
```

## Argument Reference

The following arguments are supported:

* `vnic_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The VNIC's availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VNIC.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `bminstance1` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.
* `ipv6addresses` - List of IPv6 addresses assigned to the VNIC.  Example: `2001:DB8::` 
* `is_primary` - Whether the VNIC is the primary VNIC (the VNIC that is automatically created and attached during instance launch). 
* `mac_address` - The MAC address of the VNIC.

	If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution, the MAC address is learned. If the VNIC belongs to a subnet, the MAC address is a static, Oracle-provided value.

	Example: `00:00:00:00:00:01` 
* `nsg_ids` - A list of the OCIDs of the network security groups that the VNIC belongs to.

	If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution (instead of belonging to a subnet), the value of the `nsgIds` attribute is ignored. Instead, the VNIC belongs to the NSGs that are associated with the VLAN itself. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

	For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
* `private_ip_address` - The private IP address of the primary `privateIp` object on the VNIC. The address is within the CIDR of the VNIC's subnet.  Example: `10.0.3.3`
* `public_ip_address` - The public IP address of the VNIC, if one is assigned. 
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see [Source Based Routing](https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing). 
* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

	 If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution (instead of belonging to a subnet), the `skipSourceDestCheck` attribute is `true`. This is because the source/destination check is always disabled for VNICs in a VLAN.

	Example: `true` 
* `state` - The current state of the VNIC.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
* `time_created` - The date and time the VNIC was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vlan_id` - If the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution (instead of belonging to a subnet), the `vlanId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN the VNIC is in. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). If the VNIC is instead in a subnet, `subnetId` has a value. 

