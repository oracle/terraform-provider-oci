---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_resolver_endpoint"
sidebar_current: "docs-oci-resource-dns-resolver_endpoint"
description: |-
  Provides the Resolver Endpoint resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_resolver_endpoint
This resource provides the Resolver Endpoint resource in Oracle Cloud Infrastructure DNS service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/dns/latest/ResolverEndpoint

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dns

Creates a new resolver endpoint in the same compartment as the resolver.


## Example Usage

```hcl
resource "oci_dns_resolver_endpoint" "test_resolver_endpoint" {
	#Required
	is_forwarding = var.resolver_endpoint_is_forwarding
	is_listening = var.resolver_endpoint_is_listening
	name = var.resolver_endpoint_name
	resolver_id = oci_dns_resolver.test_resolver.id
	subnet_id = oci_core_subnet.test_subnet.id
	scope = "PRIVATE"

	#Optional
	defined_tags = var.resolver_endpoint_defined_tags
	endpoint_type = var.resolver_endpoint_endpoint_type
	forwarding_address = var.resolver_endpoint_forwarding_address
	freeform_tags = var.resolver_endpoint_freeform_tags
	listening_address = var.resolver_endpoint_listening_address
	nsg_ids = var.resolver_endpoint_nsg_ids
	security_attributes = var.resolver_endpoint_security_attributes
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

  **Example:** `{"Operations": {"CostCenter": "42"}}`
* `endpoint_type` - (Optional) The type of resolver endpoint. VNIC is currently the only supported type. 
* `forwarding_address` - (Optional) An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `is_forwarding` - (Required) A Boolean flag indicating whether or not the resolver endpoint is for forwarding. 
* `is_listening` - (Required) A Boolean flag indicating whether or not the resolver endpoint is for listening. 
* `listening_address` - (Optional) An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified when isListening is true. 
* `name` - (Required) The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver. 
* `nsg_ids` - (Optional) An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the resolver endpoint is a part of. 
* `resolver_id` - (Required) The OCID of the target resolver.
* `security_attributes` - (Optional) (Updatable) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
* `scope` - (Required) Value must be `PRIVATE` when creating private name resolver endpoints.
* `subnet_id` - (Required) The OCID of a subnet. Must be part of the VCN that the resolver is attached to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under and will be updated if the resolver's compartment is changed. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

  **Example:** `{"Operations": {"CostCenter": "42"}}`
* `endpoint_type` - The type of resolver endpoint. VNIC is currently the only supported type. 
* `forwarding_address` - An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `id` - The Terraform ID of the resolver endpoint.
* `is_forwarding` - A Boolean flag indicating whether or not the resolver endpoint is for forwarding. 
* `is_listening` - A Boolean flag indicating whether or not the resolver endpoint is for listening. 
* `listening_address` - An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified when isListening is true. 
* `name` - The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver. 
* `nsg_ids` - An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the resolver endpoint is a part of. 
* `pe_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint resource that this resolver endpoint corresponds to.
* `resolver_id` - The OCID of the resolver.
* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `subnet_id` - The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC resource that this resolver endpoint corresponds to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Resolver Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Resolver Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Resolver Endpoint


## Import

For ResolverEndpoints created using `scope`, these ResolverEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_dns_resolver_endpoint.test_resolver_endpoint "resolverId/{resolverId}/name/{name}/scope/{scope}"
```

