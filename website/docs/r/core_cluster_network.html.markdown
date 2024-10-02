---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cluster_network"
sidebar_current: "docs-oci-resource-core-cluster_network"
description: |-
  Provides the Cluster Network resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cluster_network
This resource provides the Cluster Network resource in Oracle Cloud Infrastructure Core service.

Creates a [cluster network with instance pools](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm).
A cluster network is a group of high performance computing (HPC), GPU, or optimized bare metal
instances that are connected with an ultra low-latency remote direct memory access (RDMA) network.
Cluster networks with instance pools use instance pools to manage groups of identical instances.

Use cluster networks with instance pools when you want predictable capacity for a specific number of identical
instances that are managed as a group.

If you want to manage instances in the RDMA network independently of each other or use different types of instances
in the network group, create a compute cluster by using the [CreateComputeCluster](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ComputeCluster/CreateComputeCluster)
operation.

To determine whether capacity is available for a specific shape before you create a cluster network,
use the [CreateComputeCapacityReport](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ComputeCapacityReport/CreateComputeCapacityReport)
operation.


## Example Usage

```hcl
resource "oci_core_cluster_network" "test_cluster_network" {
	#Required
	compartment_id = var.compartment_id
	instance_pools {
		#Required
		instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
		size = var.cluster_network_instance_pools_size

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		display_name = var.cluster_network_instance_pools_display_name
		freeform_tags = {"Department"= "Finance"}
	}
	placement_configuration {
		#Required
		availability_domain = var.cluster_network_placement_configuration_availability_domain
		primary_vnic_subnets {
			#Required
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			ipv6address_ipv6subnet_cidr_pair_details {

				#Optional
				ipv6subnet_cidr = var.cluster_network_placement_configuration_primary_vnic_subnets_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
			}
			is_assign_ipv6ip = var.cluster_network_placement_configuration_primary_vnic_subnets_is_assign_ipv6ip
		}
		secondary_vnic_subnets {
			#Required
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			display_name = var.cluster_network_placement_configuration_secondary_vnic_subnets_display_name
			ipv6address_ipv6subnet_cidr_pair_details {

				#Optional
				ipv6subnet_cidr = var.cluster_network_placement_configuration_secondary_vnic_subnets_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
			}
			is_assign_ipv6ip = var.cluster_network_placement_configuration_secondary_vnic_subnets_is_assign_ipv6ip
		}
	}

	#Optional
	cluster_configuration {
		#Required
		hpc_island_id = oci_core_hpc_island.test_hpc_island.id

		#Optional
		network_block_ids = var.cluster_network_cluster_configuration_network_block_ids
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cluster_network_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_configuration` - (Optional) The HPC cluster configuration requested when launching instances of a cluster network.

	If the parameter is provided, instances will only be placed within the HPC island and list of network blocks that you specify. If a list of network blocks are missing or not provided, the instances will be placed in any HPC blocks in the HPC island that you specify. If the values of HPC island or network block that you provide are not valid, an error is returned. 
	* `hpc_island_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HPC island. 
	* `network_block_ids` - (Optional) The list of network block OCIDs.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the cluster network. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_pools` - (Required) (Updatable) The data to create the instance pools in the cluster network.

	Each cluster network can have one instance pool. 
	* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool. 
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `instance_configuration_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration associated with the instance pool. 
	* `size` - (Required) (Updatable) The number of instances that should be in the instance pool. 
* `placement_configuration` - (Required) The location for where the instance pools in a cluster network will place instances.
	* `availability_domain` - (Required) The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. This field is deprecated. Use `primaryVnicSubnets` instead to set VNIC data for instances in the pool. 
	* `primary_vnic_subnets` - (Optional) Details about the IPv6 primary subnet.
		* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6subnet_cidr` - (Optional) Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
		* `is_assign_ipv6ip` - (Optional) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
		* `subnet_id` - (Required) The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	* `secondary_vnic_subnets` - (Optional) The set of secondary VNIC data for instances in the pool.
		* `display_name` - (Optional) The display name of the VNIC. This is also used to match against the instance configuration defined secondary VNIC.
		* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6subnet_cidr` - (Optional) Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
		* `is_assign_ipv6ip` - (Optional) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
		* `subnet_id` - (Required) The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the cluster netowrk. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HPC island used by the cluster network.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster network.
* `instance_pools` - The instance pools in the cluster network.

	Each cluster network can have one instance pool. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
	* `instance_configuration_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration associated with the instance pool. 
 	* `load_balancers` - The load balancers attached to the instance pool. 
		* `backend_set_name` - The name of the backend set on the load balancer.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attachment.
		* `instance_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool of the load balancer attachment. 
		* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attached to the instance pool. 
		* `port` - The port value used for the backends.
		* `state` - The status of the interaction between the instance pool and the load balancer.
		* `vnic_selection` - Indicates which VNIC on each instance in the instance pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool. 
	* `placement_configurations` - The placement configurations for the instance pool.
		* `availability_domain` - The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
		* `fault_domains` - The fault domains to place instances.

			If you don't provide any values, the system makes a best effort to distribute instances across all fault domains based on capacity.

			To distribute the instances evenly across selected fault domains, provide a set of fault domains. For example, you might want instances to be evenly distributed if your applications require high availability.

			To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `[FAULT-DOMAIN-1, FAULT-DOMAIN-2, FAULT-DOMAIN-3]`
		* `primary_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. This field is deprecated. Use `primaryVnicSubnets` instead to set VNIC data for instances in the pool. 
		* `primary_vnic_subnets` - Details about the IPv6 primary subnet.
			* `ipv6address_ipv6subnet_cidr_pair_details` - A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
				* `ipv6subnet_cidr` - Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
			* `is_assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
			* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
		* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
			* `display_name` - The display name of the VNIC. This is also used to match against the instance configuration defined secondary VNIC.
			* `ipv6address_ipv6subnet_cidr_pair_details` - A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
				* `ipv6subnet_cidr` - Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
			* `is_assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
			* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	* `size` - The number of instances that should be in the instance pool.
	* `state` - The current state of the instance pool.
	* `time_created` - The date and time the instance pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `network_block_ids` - The list of network block OCIDs of the HPC island.
* `placement_configuration` - The location for where the instance pools in a cluster network will place instances.
	* `availability_domain` - The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. This field is deprecated. Use `primaryVnicSubnets` instead to set VNIC data for instances in the pool. 
	* `primary_vnic_subnets` - Details about the IPv6 primary subnet.
		* `ipv6address_ipv6subnet_cidr_pair_details` - A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6subnet_cidr` - Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
		* `is_assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
		* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
		* `display_name` - The display name of the VNIC. This is also used to match against the instance configuration defined secondary VNIC.
		* `ipv6address_ipv6subnet_cidr_pair_details` - A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6subnet_cidr` - Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
		* `is_assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
		* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
* `state` - The current state of the cluster network.
* `time_created` - The date and time the resource was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the resource was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Network
	* `update` - (Defaults to 20 minutes), when updating the Cluster Network
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Network


## Import

ClusterNetworks can be imported using the `id`, e.g.

```
$ terraform import oci_core_cluster_network.test_cluster_network "id"
```

