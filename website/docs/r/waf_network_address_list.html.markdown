---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_network_address_list"
sidebar_current: "docs-oci-resource-waf-network_address_list"
description: |-
  Provides the Network Address List resource in Oracle Cloud Infrastructure Waf service
---

# oci_waf_network_address_list
This resource provides the Network Address List resource in Oracle Cloud Infrastructure Waf service.

Creates a new NetworkAddressList.


## Example Usage

```hcl
resource "oci_waf_network_address_list" "test_network_address_list" {
	#Required
	compartment_id = var.compartment_id
	type = var.network_address_list_type

	#Optional
	addresses = var.network_address_list_addresses
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.network_address_list_display_name
	freeform_tags = {"bar-key"= "value"}
	system_tags = var.network_address_list_system_tags
	vcn_addresses {

		#Optional
		addresses = var.network_address_list_vcn_addresses_addresses
		vcn_id = oci_core_vcn.test_vcn.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `addresses` - (Required when type=ADDRESSES) (Updatable) A list of IP address prefixes in CIDR notation. To specify all addresses, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) NetworkAddressList display name, can be renamed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `system_tags` - (Optional) (Updatable) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `type` - (Required) (Updatable) Type of NetworkAddressList.
* `vcn_addresses` - (Required when type=VCN_ADDRESSES) (Updatable) A list of private address prefixes, each associated with a particular VCN. To specify all addresses in a VCN, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6. 
	* `addresses` - (Required when type=VCN_ADDRESSES) (Updatable) A private IP address or CIDR IP address range.
	* `vcn_id` - (Required when type=VCN_ADDRESSES) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `addresses` - A list of IP address prefixes in CIDR notation. To specify all addresses, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - NetworkAddressList display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAddressList.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `state` - The current state of the NetworkAddressList.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the NetworkAddressList was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the NetworkAddressList was updated. An RFC3339 formatted datetime string.
* `type` - Type of NetworkAddressList.
* `vcn_addresses` - A list of private address prefixes, each associated with a particular VCN. To specify all addresses in a VCN, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6. 
	* `addresses` - A private IP address or CIDR IP address range.
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Address List
	* `update` - (Defaults to 20 minutes), when updating the Network Address List
	* `delete` - (Defaults to 20 minutes), when destroying the Network Address List


## Import

NetworkAddressLists can be imported using the `id`, e.g.

```
$ terraform import oci_waf_network_address_list.test_network_address_list "id"
```

