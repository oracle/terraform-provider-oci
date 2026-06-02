---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_resolver_endpoint"
sidebar_current: "docs-oci-datasource-dns-resolver_endpoint"
description: |-
  Provides details about a specific Resolver Endpoint in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_resolver_endpoint
This data source provides details about a specific Resolver Endpoint resource in Oracle Cloud Infrastructure DNS service.

Gets information about a specific resolver endpoint.

Note that attempting to get a resolver endpoint in the DELETED lifecycle state will result
in a `404` response to be consistent with other operations of the API.


## Example Usage

```hcl
data "oci_dns_resolver_endpoint" "test_resolver_endpoint" {
	#Required
	resolver_endpoint_name = oci_dns_resolver_endpoint.test_resolver_endpoint.name
	resolver_id = oci_dns_resolver.test_resolver.id
	scope = "PRIVATE"

	#Optional
}
```

## Argument Reference

The following arguments are supported:

* `resolver_endpoint_name` - (Required) The name of the target resolver endpoint.
* `resolver_id` - (Required) The OCID of the target resolver.
* `scope` - (Required) Value must be `PRIVATE` when listing private name resolver endpoints.


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

