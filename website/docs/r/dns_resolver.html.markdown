---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_resolver"
sidebar_current: "docs-oci-resource-dns-resolver"
description: |-
  Provides the Resolver resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_resolver
This resource provides the Resolver resource in Oracle Cloud Infrastructure DNS service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/dns/latest/Resolver

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/dns

Updates the specified resolver with your new information.

Note: Resolvers are associated with VCNs and created when a VCN is created. Wait until the created VCN's state shows as Available in the OCI Console before updating DNS resolver properties.
A VCN cannot be deleted while its resolver has resolver endpoints. Additionally, a resolver endpoint cannot be deleted if it is referenced in the resolver's rules. To remove rules from a resolver, update the resolver resource.

Destroy behavior: This resource does not delete the underlying DNS Resolver. The resolver itself is deleted only when the attached VCN is deleted. When this Terraform resource is destroyed, managed properties on the resolver (for example, attached views and rules) are cleared so the VCN can be deleted.

Default view behavior on VCN delete: If the resolver's default view contains customer-created zones, deleting the VCN (which deletes the resolver) can convert that default view into a non-protected regular view. That view may persist even if it was never imported into Terraform state. To avoid orphaned resources, either delete the zones from the default view before deleting the VCN, or plan to clean up the resulting view afterward.

## Example Usage

```hcl
resource "oci_dns_resolver" "test_resolver" {
	#Required
	resolver_id = oci_dns_resolver.test_resolver.id

	#Optional
	scope = "PRIVATE"
	attached_views {
		#Required
		view_id = oci_dns_view.test_view.id
	}
	defined_tags = var.resolver_defined_tags
	display_name = var.resolver_display_name
	freeform_tags = var.resolver_freeform_tags
	rules {
		#Required
		action = var.resolver_rules_action
		destination_addresses = var.resolver_rules_destination_addresses
		source_endpoint_name = oci_data_connectivity_endpoint.test_endpoint.name

		#Optional
		client_address_conditions = var.resolver_rules_client_address_conditions
		qname_cover_conditions = var.resolver_rules_qname_cover_conditions
	}
}
```

## Argument Reference

The following arguments are supported:

* `attached_views` - (Optional) (Updatable) The attached views. Views are evaluated in order.
	* `view_id` - (Required) (Updatable) The OCID of the view.
* `compartment_id` - (Optional) (Updatable) The OCID of the owning compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `display_name` - (Optional) (Updatable) The display name of the resolver. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `resolver_id` - (Required) The OCID of the target resolver.
* `rules` - (Optional) (Updatable) Rules for the resolver. Rules are evaluated in order, and only the first matching rule will have its action applied. 
	* `action` - (Required) (Updatable) The action determines the behavior of the rule. If a query matches a supplied condition, the action will apply. If there are no conditions on the rule, all queries are subject to the specified action.
		* `FORWARD` - Matching requests will be forwarded from the source interface to the destination address. 
	* `client_address_conditions` - (Optional) (Updatable) A list of CIDR blocks. In order for the rule action to apply, the query must come from a client within one of the CIDR blocks. 
	* `destination_addresses` - (Required) (Updatable) IP addresses to which queries should be forwarded. Currently limited to a single address. 
	* `qname_cover_conditions` - (Optional) (Updatable) A list of domain names. In order for the rule action to apply, the query must either match or be a subdomain of one of the listed domains. 
	* `source_endpoint_name` - (Required) (Updatable) Case-insensitive name of an endpoint, that is a sub-resource of the resolver, to use as the forwarding interface. The endpoint must have isForwarding set to true. 
* `scope` - (Optional) Specifies to operate only on resources that have a matching DNS scope. 

## Attributes Reference

The following attributes are exported:

* `attached_vcn_id` - The OCID of the attached VCN. 
* `attached_views` - The attached views. Views are evaluated in order.
	* `view_id` - The OCID of the view.
* `compartment_id` - The OCID of the owning compartment.
* `default_view_id` - The OCID of the default view. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `display_name` - The display name of the resolver. 
* `endpoints` - Read-only array of endpoints for the resolver. 
	* `compartment_id` - The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under and will be updated if the resolver's compartment is changed. 
	* `endpoint_type` - The type of resolver endpoint. VNIC is currently the only supported type. 
	* `forwarding_address` - An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true. 
	* `is_forwarding` - A Boolean flag indicating whether or not the resolver endpoint is for forwarding. 
	* `is_listening` - A Boolean flag indicating whether or not the resolver endpoint is for listening. 
	* `listening_address` - An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified when isListening is true. 
	* `name` - The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver. 
	* `self` - The canonical absolute URL of the resource.
	* `state` - The current state of the resource.
	* `subnet_id` - The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
	* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

		**Example:** `2016-07-22T17:23:59:60Z` 
	* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

		**Example:** `2016-07-22T17:23:59:60Z` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the resolver.
* `is_protected` - A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. 
* `rules` - Rules for the resolver. Rules are evaluated in order, and only the first matching rule will have its action applied. 
	* `action` - The action determines the behavior of the rule. If a query matches a supplied condition, the action will apply. If there are no conditions on the rule, all queries are subject to the specified action.
		* `FORWARD` - Matching requests will be forwarded from the source interface to the destination address. 
	* `client_address_conditions` - A list of CIDR blocks. In order for the rule action to apply, the query must come from a client within one of the CIDR blocks. 
	* `destination_addresses` - IP addresses to which queries should be forwarded. Currently limited to a single address. 
	* `qname_cover_conditions` - A list of domain names. In order for the rule action to apply, the query must either match or be a subdomain of one of the listed domains. 
	* `source_endpoint_name` - Case-insensitive name of an endpoint, that is a sub-resource of the resolver, to use as the forwarding interface. The endpoint must have isForwarding set to true. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Resolver
	* `update` - (Defaults to 20 minutes), when updating the Resolver
	* `delete` - (Defaults to 20 minutes), when destroying the Resolver


## Import

Resolvers can be imported using their OCID, e.g.

```
$ terraform import oci_dns_resolver.test_resolver "id"
```
