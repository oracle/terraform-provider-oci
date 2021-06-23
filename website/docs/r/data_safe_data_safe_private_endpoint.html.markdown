---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_data_safe_private_endpoint"
sidebar_current: "docs-oci-resource-data_safe-data_safe_private_endpoint"
description: |-
  Provides the Data Safe Private Endpoint resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_data_safe_private_endpoint
This resource provides the Data Safe Private Endpoint resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new Data Safe private endpoint.


## Example Usage

```hcl
resource "oci_data_safe_data_safe_private_endpoint" "test_data_safe_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.data_safe_private_endpoint_display_name
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.data_safe_private_endpoint_description
	freeform_tags = {"Department"= "Finance"}
	nsg_ids = var.data_safe_private_endpoint_nsg_ids
	private_endpoint_ip = var.data_safe_private_endpoint_private_endpoint_ip
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the private endpoint.
* `display_name` - (Required) (Updatable) The display name for the private endpoint. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `nsg_ids` - (Optional) (Updatable) The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_endpoint_ip` - (Optional) The private IP address of the private endpoint.
* `subnet_id` - (Required) The OCID of the subnet.
* `vcn_id` - (Required) The OCID of the VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the private endpoint.
* `display_name` - The display name of the private endpoint.
* `endpoint_fqdn` - The three-label fully qualified domain name (FQDN) of the private endpoint. The customer VCN's DNS records are updated with this FQDN.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Data Safe private endpoint.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_endpoint_id` - The OCID of the underlying private endpoint.
* `private_endpoint_ip` - The private IP address of the private endpoint. 
* `state` - The current state of the private endpoint.
* `subnet_id` - The OCID of the subnet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Safe Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Data Safe Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Data Safe Private Endpoint


## Import

DataSafePrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint "id"
```

