---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_database_private_endpoint"
sidebar_current: "docs-oci-datasource-distributed_database-distributed_database_private_endpoint"
description: |-
  Provides details about a specific Distributed Database Private Endpoint in Oracle Cloud Infrastructure Distributed Database service
---

# Data Source: oci_distributed_database_distributed_database_private_endpoint
This data source provides details about a specific Distributed Database Private Endpoint resource in Oracle Cloud Infrastructure Distributed Database service.

Get the DistributedDatabasePrivateEndpoint resource.


## Example Usage

```hcl
data "oci_distributed_database_distributed_database_private_endpoint" "test_distributed_database_private_endpoint" {
	#Required
	distributed_database_private_endpoint_id = oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `distributed_database_private_endpoint_id` - (Required) Distributed Database PrivateEndpoint identifier


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

