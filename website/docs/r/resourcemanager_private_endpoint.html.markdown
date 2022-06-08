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

Creates a a private endpoint in the specified compartment.


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
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint details.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the private endpoint. Avoid entering confidential information.
* `display_name` - (Required) (Updatable) The private endpoint display name. Avoid entering confidential information.
* `dns_zones` - (Optional) (Updatable) DNS Proxy forwards any DNS FQDN queries over into the consumer DNS resolver if the DNS FQDN is included in the dns zones list otherwise it goes to service provider VCN resolver. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_used_with_configuration_source_provider` - (Optional) (Updatable) When `true`, allows the private endpoint to be used with a configuration source provider.
* `nsg_id_list` - (Optional) (Updatable) An array of network security group (NSG) [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the private endpoint. Order does not matter.
* `subnet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
* `vcn_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint details.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the private endpoint. Avoid entering confidential information.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_zones` - DNS Proxy forwards any DNS FQDN queries over into the consumer DNS resolver if the DNS FQDN is included in the dns zones list otherwise it goes to service provider VCN resolver. 
* `freeform_tags` - Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the private endpoint details.
* `is_used_with_configuration_source_provider` - When `true`, allows the private endpoint to be used with a configuration source provider.
* `nsg_id_list` - An array of network security groups (NSG) that the customer can optionally provide.
* `source_ips` - The source IPs which resource manager service will use to connect to customer's network. Automatically assigned by Resource Manager Service.
* `state` - The current lifecycle state of the private endpoint. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
* `time_created` - The date and time at which the private endpoint was created. Format is defined by RFC3339. Example: `2020-11-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Private Endpoint


## Import

PrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_resourcemanager_private_endpoint.test_private_endpoint "id"
```

