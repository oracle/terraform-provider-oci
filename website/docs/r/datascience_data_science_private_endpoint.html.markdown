---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_private_endpoint"
sidebar_current: "docs-oci-resource-datascience-data_science_private_endpoint"
description: |-
  Provides the Data Science Private Endpoint resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_private_endpoint
This resource provides the Data Science Private Endpoint resource in Oracle Cloud Infrastructure Data Science service.

Creates a Data Science private endpoint to be used by a Data Science resource.


## Example Usage

```hcl
resource "oci_datascience_private_endpoint" "test_data_science_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	data_science_resource_type = var.data_science_private_endpoint_data_science_resource_type
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.data_science_private_endpoint_description
	display_name = var.data_science_private_endpoint_display_name
	freeform_tags = {"Department"= "Finance"}
	nsg_ids = var.data_science_private_endpoint_nsg_ids
	sub_domain = var.data_science_private_endpoint_sub_domain
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the private endpoint.
* `data_science_resource_type` - (Required) Data Science resource type.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A user friendly description. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `nsg_ids` - (Optional) (Updatable) An array of network security group OCIDs. 
* `sub_domain` - (Optional) Subdomain for a private endpoint FQDN.
* `subnet_id` - (Required) The OCID of the subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create private endpoint.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user that created the private endpoint.
* `data_science_resource_type` - Data Science resource type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user friendly description. Avoid entering confidential information. 
* `display_name` - A user friendly name. It doesn't have to be unique. Avoid entering confidential information. 
* `fqdn` - Accesing the Data Science resource using FQDN. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of a private endpoint. 
* `lifecycle_details` - Details of the state of Data Science private endpoint.
* `nsg_ids` - An array of network security group OCIDs. 
* `state` - State of the Data Science private endpoint.
* `subnet_id` - The OCID of a subnet. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the Data Science private endpoint was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time that the Data Science private endpoint was updated expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Science Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Data Science Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Data Science Private Endpoint


## Import

DataSciencePrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_private_endpoint.test_data_science_private_endpoint "id"
```

