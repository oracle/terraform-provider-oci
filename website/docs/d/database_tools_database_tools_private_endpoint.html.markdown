---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_private_endpoint"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_private_endpoint"
description: |-
  Provides details about a specific Database Tools Private Endpoint in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_private_endpoint
This data source provides details about a specific Database Tools Private Endpoint resource in Oracle Cloud Infrastructure Database Tools service.

Gets details of a specified Database Tools private endpoint.

## Example Usage

```hcl
data "oci_database_tools_database_tools_private_endpoint" "test_database_tools_private_endpoint" {
	#Required
	database_tools_private_endpoint_id = oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools private endpoint.


## Attributes Reference

The following attributes are exported:

* `additional_fqdns` - A list of additional FQDNs that can be also be used for the private endpoint.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools private endpoint.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the Database Tools private endpoint.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `endpoint_fqdn` - Then FQDN to use for the private endpoint.
* `endpoint_service_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Endpoint Service.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools private endpoint.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `nsg_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups that the private endpoint's VNIC belongs to.  For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
* `private_endpoint_ip` - The private IP address that represents the access point for the associated endpoint service.
* `private_endpoint_vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint's VNIC.
* `reverse_connection_configuration` - Reverse connection configuration details of the private endpoint.
	* `reverse_connections_source_ips` - A list of IP addresses in the customer VCN to be used as the source IPs for reverse connection packets traveling from the service's VCN to the customer's VCN. 
		* `source_ip` - The IP address in the customer's VCN to be used as the source IP for reverse connection packets traveling from the customer's VCN to the service's VCN. 
* `state` - The current state of the Database Tools private endpoint.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Database Tools private endpoint was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Database Tools private endpoint was updated. An RFC3339 formatted datetime string
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN that the private endpoint belongs to.

