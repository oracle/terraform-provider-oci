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

Creates a new resolver endpoint. Requires a `PRIVATE` scope query parameter.


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
	endpoint_type = var.resolver_endpoint_endpoint_type
	forwarding_address = var.resolver_endpoint_forwarding_address
	listening_address = var.resolver_endpoint_listening_address
	nsg_ids = var.resolver_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `endpoint_type` - (Optional) (Updatable) The type of resolver endpoint. VNIC is currently the only supported type. 
* `forwarding_address` - (Optional) An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true. 
* `is_forwarding` - (Required) A Boolean flag indicating whether or not the resolver endpoint is for forwarding. 
* `is_listening` - (Required) A Boolean flag indicating whether or not the resolver endpoint is for listening. 
* `listening_address` - (Optional) An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified when isListening is true. 
* `name` - (Required) The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver. 
* `nsg_ids` - (Optional) An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the resolver endpoint is a part of. 
* `resolver_id` - (Required) The OCID of the target resolver.
* `scope` - (Required) Value must be `PRIVATE` when creating private name resolver endpoints. 
* `subnet_id` - (Required) The OCID of a subnet. Must be part of the VCN that the resolver is attached to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under and will be updated if the resolver's compartment is changed. 
* `endpoint_type` - The type of resolver endpoint. VNIC is currently the only supported type. 
* `forwarding_address` - An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true. 
* `is_forwarding` - A Boolean flag indicating whether or not the resolver endpoint is for forwarding. 
* `is_listening` - A Boolean flag indicating whether or not the resolver endpoint is for listening. 
* `listening_address` - An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified when isListening is true. 
* `name` - The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver. 
* `nsg_ids` - An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the resolver endpoint is a part of. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `subnet_id` - The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Resolver Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Resolver Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Resolver Endpoint


## Import

For legacy ResolverEndpoints created without `scope`, these ResolverEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_dns_resolver_endpoint.test_resolver_endpoint "resolverId/{resolverId}/name/{resolverEndpointName}" 
```

For ResolverEndpoints created using `scope`, these ResolverEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_dns_resolver_endpoint.test_resolver_endpoint "resolverId/{resolverId}/name/{name}/scope/{scope}"
```

