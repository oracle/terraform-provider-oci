---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_generative_ai_private_endpoint"
sidebar_current: "docs-oci-resource-generative_ai-generative_ai_private_endpoint"
description: |-
  Provides the Generative Ai Private Endpoint resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_generative_ai_private_endpoint
This resource provides the Generative Ai Private Endpoint resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/GenerativeAiPrivateEndpoint

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a Generative AI private endpoint.


## Example Usage

```hcl
resource "oci_generative_ai_generative_ai_private_endpoint" "test_generative_ai_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	dns_prefix = var.generative_ai_private_endpoint_dns_prefix
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.generative_ai_private_endpoint_description
	display_name = var.generative_ai_private_endpoint_display_name
	freeform_tags = {"Department"= "Finance"}
	nsg_ids = var.generative_ai_private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the private endpoint is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A description of this private endpoint. 
* `display_name` - (Optional) (Updatable) A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `dns_prefix` - (Required) (Updatable) dnsPrefix of the private endpoint FQDN. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `nsg_ids` - (Optional) (Updatable) A list of the OCIDs of the network security groups (NSGs) to add the private endpoint's VNIC to. 
* `subnet_id` - (Required) The OCID of the customer's subnet where the private endpoint VNIC will reside. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the private endpoint. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of this private endpoint. 
* `display_name` - A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `fqdn` - Fully qualified domain name the customer will use for access (for eg: xyz.oraclecloud.com) 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of a private endpoint. 
* `lifecycle_details` - The detailed messages about the lifecycle state 
* `nsg_ids` - A list of the OCIDs of the network security groups that the private endpoint's VNIC belongs to. 
* `previous_state` - Generative AI private endpoint. 
* `private_endpoint_ip` - The private IP address (in the customer's VCN) that represents the access point for the associated endpoint service. 
* `state` - The current state of the Generative AI Private Endpoint. 
* `subnet_id` - The OCID of the subnet that the private endpoint belongs to. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the Generative AI private endpoint was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time that the Generative AI private endpoint was updated expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Generative Ai Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Generative Ai Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Generative Ai Private Endpoint


## Import

GenerativeAiPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint "id"
```

