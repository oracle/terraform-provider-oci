---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_resolver_endpoints"
sidebar_current: "docs-oci-datasource-dns-resolver_endpoints"
description: |-
  Provides the list of Resolver Endpoints in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_resolver_endpoints
This data source provides the list of Resolver Endpoints in Oracle Cloud Infrastructure DNS service.

Gets a list of all endpoints within a resolver. The collection can be filtered by name or lifecycle state.
It can be sorted on creation time or name both in ASC or DESC order. Note that when no lifecycleState
query parameter is provided, the collection does not include resolver endpoints in the DELETED
lifecycle state to be consistent with other operations of the API.


## Example Usage

```hcl
data "oci_dns_resolver_endpoints" "test_resolver_endpoints" {
	#Required
	resolver_id = oci_dns_resolver.test_resolver.id
	scope = "PRIVATE"

	#Optional
	name = var.resolver_endpoint_name
	state = var.resolver_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of a resource.
* `resolver_id` - (Required) The OCID of the target resolver.
* `scope` - (Required) Value must be `PRIVATE` when listing private name resolver endpoints.
* `state` - (Optional) The state of a resource.


## Attributes Reference

The following attributes are exported:

* `resolver_endpoints` - The list of resolver_endpoints.

### ResolverEndpoint Reference

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

