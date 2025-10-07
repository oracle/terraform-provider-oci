---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_private_endpoint"
sidebar_current: "docs-oci-resource-resourcemanager-private_endpoint"
description: |-
  Provides the Private Endpoint resource in Oracle Cloud Infrastructure Resource Manager service
---

# oci_resourcemanager_private_endpoint
This resource provides the Private Endpoint resource in Oracle Cloud Infrastructure Resource Manager service.

Creates a private endpoint in the specified compartment.
For more information, see
[Creating a Private Endpoint](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/create-private-endpoints.htm).


## Example Usage

```hcl
resource "oci_resourcemanager_private_endpoint" "test_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.private_endpoint_display_name
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.private_endpoint_description
	dns_zones = var.private_endpoint_dns_zones
	freeform_tags = {"Department"= "Finance"}
	is_used_with_configuration_source_provider = var.private_endpoint_is_used_with_configuration_source_provider
	nsg_id_list = var.private_endpoint_nsg_id_list
	security_attributes = var.private_endpoint_security_attributes
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the private endpoint. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_zones` - (Optional) (Updatable) DNS Proxy forwards any DNS FQDN queries over into the consumer DNS resolver if the DNS FQDN is included in the dns zones list otherwise it goes to service provider VCN resolver. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_used_with_configuration_source_provider` - (Optional) (Updatable) When `true`, allows the private endpoint to be used with a configuration source provider.
* `nsg_id_list` - (Optional) (Updatable) The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of [network security groups (NSGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm) for the private endpoint. Order does not matter. 
* `security_attributes` - (Optional) (Updatable) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `subnet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
* `vcn_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the private endpoint. Avoid entering confidential information.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_zones` - DNS zones to use for accessing private Git servers. For private Git server instructions, see [Private Git Server](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/private-endpoints.htm#private-git). Specify DNS fully qualified domain names (FQDNs); DNS Proxy forwards related DNS FQDN queries to the consumer DNS resolver. For DNS FQDNs not specified, queries go to service provider VCN resolver. Example: `abc.oraclevcn.com` 
* `freeform_tags` - Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
* `is_used_with_configuration_source_provider` - When `true`, allows the private endpoint to be used with a configuration source provider.
* `nsg_id_list` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of [network security groups (NSGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm) for the private endpoint. Order does not matter. 
* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `source_ips` - The source IP addresses that Resource Manager uses to connect to your network. Automatically assigned by Resource Manager.
* `state` - The current lifecycle state of the private endpoint. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time at which the private endpoint was created. Format is defined by RFC3339. Example: `2020-11-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Private Endpoint


## Import

PrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_resourcemanager_private_endpoint.test_private_endpoint "id"
```

