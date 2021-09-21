---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_network_address_lists"
sidebar_current: "docs-oci-datasource-waf-network_address_lists"
description: |-
  Provides the list of Network Address Lists in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_network_address_lists
This data source provides the list of Network Address Lists in Oracle Cloud Infrastructure Waf service.

Gets a list of all NetworkAddressLists in a compartment.


## Example Usage

```hcl
data "oci_waf_network_address_lists" "test_network_address_lists" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.network_address_list_display_name
	id = var.network_address_list_id
	state = var.network_address_list_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only the NetworkAddressList with the given [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `network_address_list_collection` - The list of network_address_list_collection.

### NetworkAddressList Reference

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

