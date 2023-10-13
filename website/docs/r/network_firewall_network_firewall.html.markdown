---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall"
description: |-
  Provides the Network Firewall resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall
This resource provides the Network Firewall resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new NetworkFirewall.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall" "test_network_firewall" {
	#Required
	compartment_id = var.compartment_id
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	availability_domain = var.network_firewall_availability_domain
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.network_firewall_display_name
	freeform_tags = {"Department"= "Finance"}
	ipv4address = var.network_firewall_ipv4address
	ipv6address = var.network_firewall_ipv6address
	network_security_group_ids = var.network_firewall_network_security_group_ids
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) Availability Domain where Network Firewall instance is created. To get a list of availability domains for a tenancy, use [ListAvailabilityDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/AvailabilityDomain/ListAvailabilityDomains) operation. Example: `kIdk:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Network Firewall.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the Network Firewall. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `ipv4address` - (Optional) IPv4 address for the Network Firewall.
* `ipv6address` - (Optional) IPv6 address for the Network Firewall.
* `network_firewall_policy_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall Policy.
* `network_security_group_ids` - (Optional) (Updatable) An array of network security groups [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the Network Firewall.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Network Firewall.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - Availability Domain where Network Firewall instance is created. To get a list of availability domains for a tenancy, use the [ListAvailabilityDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/AvailabilityDomain/ListAvailabilityDomains) operation. Example: `kIdk:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Network Firewall.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the Network Firewall. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall resource.
* `ipv4address` - IPv4 address for the Network Firewall.
* `ipv6address` - IPv6 address for the Network Firewall.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in 'FAILED' state.
* `network_firewall_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall Policy.
* `network_security_group_ids` - An array of network security groups [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the Network Firewall.
* `state` - The current state of the Network Firewall.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Network Firewall.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the Network Firewall was created in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The time at which the Network Firewall was updated in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall


## Import

NetworkFirewalls can be imported using the `id`, e.g.

```
$ terraform import oci_network_firewall_network_firewall.test_network_firewall "id"
```

