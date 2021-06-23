---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_private_endpoint"
sidebar_current: "docs-oci-resource-dataflow-private_endpoint"
description: |-
  Provides the Private Endpoint resource in Oracle Cloud Infrastructure Data Flow service
---

# oci_dataflow_private_endpoint
This resource provides the Private Endpoint resource in Oracle Cloud Infrastructure Data Flow service.

Creates a private endpoint to be used by applications.


## Example Usage

```hcl
resource "oci_dataflow_private_endpoint" "test_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	dns_zones = var.private_endpoint_dns_zones
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.private_endpoint_description
	display_name = var.private_endpoint_display_name
	freeform_tags = {"Department"= "Finance"}
	max_host_count = var.private_endpoint_max_host_count
	nsg_ids = var.private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of a compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A user-friendly description. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique. Avoid entering confidential information. 
* `dns_zones` - (Required) (Updatable) An array of DNS zone names. Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `max_host_count` - (Optional) (Updatable) The maximum number of hosts to be accessed through the private endpoint. This value is used to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up to 512. 
* `nsg_ids` - (Optional) (Updatable) An array of network security group OCIDs. 
* `subnet_id` - (Required) The OCID of a subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of a compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user-friendly description. Avoid entering confidential information. 
* `display_name` - A user-friendly name. It does not have to be unique. Avoid entering confidential information. 
* `dns_zones` - An array of DNS zone names. Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of a private endpoint. 
* `lifecycle_details` - The detailed messages about the lifecycle state. 
* `max_host_count` - The maximum number of hosts to be accessed through the private endpoint. This value is used to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up to 512. 
* `nsg_ids` - An array of network security group OCIDs. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `state` - The current state of this private endpoint. 
* `subnet_id` - The OCID of a subnet. 
* `time_created` - The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Private Endpoint


## Import

PrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_dataflow_private_endpoint.test_private_endpoint "id"
```

## Note

When a Private Endpoint resource is created it will be in `INACTIVE` state. When user runs an application using a Private Endpoint resource only then it moves to `ACTIVE` state. Also if there is already a Private Endpoint resource that is in `ACTIVE` state then on running the new application, the new Private Endpoint will be moved to `ACTIVE` state while the old one will be moved to `INACTIVE` state by the service. To update these states in your terraform state file user needs to do a `terraform refresh`.
