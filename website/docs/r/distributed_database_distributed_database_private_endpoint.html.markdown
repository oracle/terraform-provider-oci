---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_database_private_endpoint"
sidebar_current: "docs-oci-resource-distributed_database-distributed_database_private_endpoint"
description: |-
  Provides the Distributed Database Private Endpoint resource in Oracle Cloud Infrastructure Distributed Database service
---

# oci_distributed_database_distributed_database_private_endpoint
This resource provides the Distributed Database Private Endpoint resource in Oracle Cloud Infrastructure Distributed Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/distributed_database

Creates a DistributedDatabasePrivateEndpoint.

  Patch operation to update VCN network security group (NSG) IDs associated with the DistributedDatabasePrivateEndpoint.


## Example Usage

```hcl
resource "oci_distributed_database_distributed_database_private_endpoint" "test_distributed_database_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.distributed_database_private_endpoint_display_name
	distributed_database_private_endpoint_id = var.distributed_database_private_endpoint_distributed_database_private_endpoint_id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.distributed_database_private_endpoint_description
	freeform_tags = {"bar-key"= "value"}
	patch_operations {
		#Required
		operation = var.distributed_database_private_endpoint_patch_operations_operation
		selection = var.distributed_database_private_endpoint_patch_operations_selection

		#Optional
		value = var.distributed_database_private_endpoint_patch_operations_value
	}
	vcn_nsg_ids {
		#Required
		nsg_ids = var.distributed_database_private_endpoint_vcn_nsg_ids_nsg_ids
		vcn_id = oci_core_vcn.test_vcn.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Identifier of the compartment where private endpoint is to be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) DistributedDatabasePrivateEndpoint description.
* `display_name` - (Required) (Updatable) Private endpoint display name.
* `distributed_database_private_endpoint_id` - (Required) 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `patch_operations` - (Optional) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `subnet_id` - (Required) Identifier of the customer subnet against which private endpoint is to be created.
* `vcn_nsg_ids` - (Optional) The list of network security group (NSG) details to be associated with the private endpoint.
	* `nsg_ids` - (Required) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
	* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.
* `reinstate_proxy_instance_trigger` - (Optional) (Updatable) An optional property when incremented triggers Reinstate Proxy Instance. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Identifier of the compartment in which private endpoint exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - DistributedDatabasePrivateEndpoint description.
* `display_name` - DistributedDatabasePrivateEndpoint display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The identifier of the Private Endpoint.
* `lifecycle_details` - Detailed message for the lifecycle state.
* `private_ip` - IP address of the Private Endpoint.
* `proxy_compute_instance_id` - The identifier of the proxy compute instance.
* `state` - Lifecycle states for private endpoint.
* `subnet_id` - Identifier of the subnet in which private endpoint exists.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the DistributedDatabasePrivateEndpoint was first created. An RFC3339 formatted datetime string
* `time_updated` - The time the Private Endpoint was last updated. An RFC3339 formatted datetime string
* `vcn_id` - Identifier of the VCN in which subnet exists.
* `vcn_nsg_ids` - The list of network security group (NSG) details associated with the private endpoint.
	* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Distributed Database Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Distributed Database Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Distributed Database Private Endpoint


## Import

DistributedDatabasePrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint "id"
```

