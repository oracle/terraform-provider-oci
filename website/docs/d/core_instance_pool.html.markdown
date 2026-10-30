---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool"
sidebar_current: "docs-oci-datasource-core-instance_pool"
description: |-
  Provides details about a specific Instance Pool in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_pool
This data source provides details about a specific Instance Pool resource in Oracle Cloud Infrastructure Core service.

Gets the specified instance pool

## Example Usage

```hcl
data "oci_core_instance_pool" "test_instance_pool" {
	#Required
	instance_pool_id = oci_core_instance_pool.test_instance_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool. 
* `current_size` - Count of instance in running state associated to the Instance Pool.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
* `instance_configuration_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration associated with the instance pool. 
* `instance_display_name_formatter` - A user-friendly formatter for the instance pool's instances. Instance displaynames follow the format. The formatter does not retroactively change instance's displaynames, only instance displaynames in the future follow the format 
* `instance_hostname_formatter` - A user-friendly formatter for the instance pool's instances. Instance hostnames follow the format. The formatter does not retroactively change instance's hostnames, only instance hostnames in the future follow the format 
* `lifecycle_management` - The lifecycle management options for the instance pool. 
	* `lifecycle_actions` - The lifecycle actions for the instance pool. 
		* `pre_termination` - The data for pre-termination action for an instance pool 
			* `is_enabled` - Whether pre-termination action is enabled or not. 
			* `on_timeout` - Options to handle timeout for pre-termination action. 
				* `preserve_block_volume_mode` - Whether the block volume should be preserved after termination. 
				* `preserve_boot_volume_mode` - Whether the boot volume should be preserved after termination. 
			* `timeout` - The timeout in seconds for pre-termination action for an instance pool(min = 0 sec, max = 7200 secs). 
* `load_balancers` - The load balancers attached to the instance pool. 
	* `backend_set_name` - The name of the backend set on the load balancer.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attachment.
	* `instance_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool of the load balancer attachment. 
	* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attached to the instance pool. 
	* `port` - The port value used for the backends.
	* `state` - The status of the interaction between the instance pool and the load balancer.
	* `vnic_selection` - Indicates which VNIC on each instance in the instance pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool. 
* `override_policy` - Optional pool-level launch override policy for future instance launches.

	The policy is evaluated only after placement is resolved for a launch. In V1, provide between 1 and 3 availability domain-scoped statements and no more than one statement for the same availability domain.

	For a matched availability domain, request order defines the canonical override candidate sequence. Preferred-candidate reuse may try one current option once as a fast path. If that fast path misses, canonical evaluation restarts from option 1. During canonical evaluation, the worker advances to the next option only for `out-of-host-capacity`, `LimitExceeded`, and `QuotaExceeded`, and only after same-option retries are exhausted. Other failures are terminal for that launch attempt.

	When used on create, omitting `overridePolicy` preserves the existing base-plus-placement behavior. When used on update, omitting `overridePolicy` leaves the stored policy unchanged, a non-null value replaces the stored policy, and `null` clears it. Read surfaces omit `overridePolicy` when no policy is stored on the pool. 
	* `override_statements` - The ordered set of availability domain-scoped override statements for this policy. 
		* `name` - A user-friendly name for the override statement. 
		* `override_options` - The ordered list of candidate override options for the matched availability domain.

			Request order defines the canonical candidate sequence. Override options within the same statement must normalize to distinct override payloads. 
			* `name` - A user-friendly name for the override option. 
			* `override` - The concrete launch mutation payload for an override option.

				In V1, the override payload is limited to `shape`, `shapeConfigDetails`, and `capacityReservationId`.

				Each override option must include at least one of `shape` or `capacityReservationId`, so a `capacityReservationId`-only override is valid. `shapeConfigDetails` may be supplied only when `shape` is supplied. If `shape` is a Flex or Fungible shape, `shapeConfigDetails` is required.

				Placement-owned and system-owned fields are not supported in this payload. Any field outside this V1 payload shape is rejected. 
				* `capacity_reservation_id` - The OCID of the compute capacity reservation to use for the override option. 
				* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

					You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
				* `shape_config_details` - The shape configuration requested for the override option. 
					* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

						The following values are supported:
						* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
						* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
						* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
					* `local_volume_size_in_gbs` - The NVMe-backed local storage capacity, in GB, for flexible dense (DenseLV) VM shapes. If the selected shape  is DenseLV, the value must be greater than 0. For all other shapes, the value must be null (if specified);  any non-null value for a non-DenseLV shape results in an error. 
					* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
					* `nvmes` - The number of NVMe drives to be used for storage. A single drive has 6.8 TB available. 
					* `ocpus` - The total number of OCPUs available to the instance. 
					* `resource_management` - This field is reserved for internal use. 
					* `vcpus` - The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2. 
		* `selector` - The selector for an override statement.

			In V1, the selector contains only the scalar `availabilityDomain`. Additional selector fields are rejected. 
			* `availability_domain` - The availability domain that this statement applies to.

				The value must match an availability domain present in the pool's `placementConfigurations`. 
* `placement_configurations` - The placement configurations for the instance pool.
	* `availability_domain` - The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
	* `compute_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) that the instance will be created in. 
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
* `pool_type` - The type of resources managed by the pool.
* `size` - The number of actual instances in the instance pool on the cloud. This attribute will be different when instance pool is used along with autoScaling Configuration.
* `state` - The current state of the instance pool.
* `time_created` - The date and time the instance pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

