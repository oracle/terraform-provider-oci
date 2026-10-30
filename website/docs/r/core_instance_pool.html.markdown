---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool"
sidebar_current: "docs-oci-resource-core-instance_pool"
description: |-
  Provides the Instance Pool resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_pool
This resource provides the Instance Pool resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/InstancePool

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates an instance pool.

To determine whether capacity is available for a specific shape before you create an instance pool,
use the [CreateComputeCapacityReport](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ComputeCapacityReport/CreateComputeCapacityReport)
operation.
You can optionally include `overridePolicy` to define availability domain-scoped
launch overrides for future launches. Existing pool members are not modified.


## Example Usage

```hcl
resource "oci_core_instance_pool" "test_instance_pool" {
	#Required
	compartment_id = var.compartment_id
	instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
	placement_configurations {
		#Required
		availability_domain = var.instance_pool_placement_configurations_availability_domain

		#Optional
		compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
		fault_domains = var.instance_pool_placement_configurations_fault_domains
		primary_subnet_id = oci_core_subnet.test_subnet.id
		primary_vnic_subnets {
			#Required
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			ipv6address_ipv6subnet_cidr_pair_details {

				#Optional
				ipv6subnet_cidr = var.instance_pool_placement_configurations_primary_vnic_subnets_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
			}
			is_assign_ipv6ip = var.instance_pool_placement_configurations_primary_vnic_subnets_is_assign_ipv6ip
		}
		secondary_vnic_subnets {
			#Required
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			display_name = var.instance_pool_placement_configurations_secondary_vnic_subnets_display_name
			ipv6address_ipv6subnet_cidr_pair_details {

				#Optional
				ipv6subnet_cidr = var.instance_pool_placement_configurations_secondary_vnic_subnets_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
			}
			is_assign_ipv6ip = var.instance_pool_placement_configurations_secondary_vnic_subnets_is_assign_ipv6ip
		}
	}
	size = var.instance_pool_size

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_pool_display_name
	freeform_tags = {"Department"= "Finance"}
	instance_display_name_formatter = var.instance_pool_instance_display_name_formatter
	instance_hostname_formatter = var.instance_pool_instance_hostname_formatter
	lifecycle_management {
		#Required
		lifecycle_actions {

			#Optional
			pre_termination {
				#Required
				is_enabled = var.instance_pool_lifecycle_management_lifecycle_actions_pre_termination_is_enabled
				on_timeout {
					#Required
					preserve_block_volume_mode = var.instance_pool_lifecycle_management_lifecycle_actions_pre_termination_on_timeout_preserve_block_volume_mode
					preserve_boot_volume_mode = var.instance_pool_lifecycle_management_lifecycle_actions_pre_termination_on_timeout_preserve_boot_volume_mode
				}
				timeout = var.instance_pool_lifecycle_management_lifecycle_actions_pre_termination_timeout
			}
		}
	}
	load_balancers {
		#Required
		backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		port = var.instance_pool_load_balancers_port
		vnic_selection = var.instance_pool_load_balancers_vnic_selection
	}
	override_policy {
		#Required
		override_statements {
			#Required
			name = var.instance_pool_override_policy_override_statements_name
			override_options {
				#Required
				name = var.instance_pool_override_policy_override_statements_override_options_name
				override {

					#Optional
					capacity_reservation_id = oci_core_capacity_reservation.test_capacity_reservation.id
					shape = var.instance_pool_override_policy_override_statements_override_options_override_shape
					shape_config_details {

						#Optional
						baseline_ocpu_utilization = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_baseline_ocpu_utilization
						local_volume_size_in_gbs = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_local_volume_size_in_gbs
						memory_in_gbs = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_memory_in_gbs
						nvmes = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_nvmes
						ocpus = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_ocpus
						resource_management = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_resource_management
						vcpus = var.instance_pool_override_policy_override_statements_override_options_override_shape_config_details_vcpus
					}
				}
			}
			selector {
				#Required
				availability_domain = var.instance_pool_override_policy_override_statements_selector_availability_domain
			}
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_configuration_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration associated with the instance pool. 
* `instance_display_name_formatter` - (Optional) (Updatable) A user-friendly formatter for the instance pool's instances. Instance displaynames follow the format. The formatter does not retroactively change instance's displaynames, only instance displaynames in the future follow the format 
* `instance_hostname_formatter` - (Optional) (Updatable) A user-friendly formatter for the instance pool's instances. Instance hostnames follow the format. The formatter does not retroactively change instance's hostnames, only instance hostnames in the future follow the format 
* `lifecycle_management` - (Optional) (Updatable) The lifecycle management options for the instance pool. 
	* `lifecycle_actions` - (Required) (Updatable) The lifecycle actions for the instance pool. 
		* `pre_termination` - (Optional) (Updatable) The data for pre-termination action for an instance pool 
			* `is_enabled` - (Required) (Updatable) Whether pre-termination action is enabled or not. 
			* `on_timeout` - (Required) (Updatable) Options to handle timeout for pre-termination action. 
				* `preserve_block_volume_mode` - (Required) (Updatable) Whether the block volume should be preserved after termination. 
				* `preserve_boot_volume_mode` - (Required) (Updatable) Whether the boot volume should be preserved after termination. 
			* `timeout` - (Required) (Updatable) The timeout in seconds for pre-termination action for an instance pool(min = 0 sec, max = 7200 secs).
* `load_balancers` - (Optional) The load balancers to attach to the instance pool. (Note: From 6.16.0 load_balancers field in oci_core_instance_pool is changed from TypeList to TypeSet - to support load balancer insert operation. Also, LB cant by accessed by index)
	* `backend_set_name` - (Required) The name of the backend set on the load balancer to add instances to.
	* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to attach to the instance pool. 
	* `port` - (Required) The port value to use when creating the backend set.
	* `vnic_selection` - (Required) Indicates which VNIC on each instance in the pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool. 
* `override_policy` - (Optional) (Updatable) Optional pool-level launch override policy for future instance launches.

	The policy is evaluated only after placement is resolved for a launch. In V1, provide between 1 and 3 availability domain-scoped statements and no more than one statement for the same availability domain.

	For a matched availability domain, request order defines the canonical override candidate sequence. Preferred-candidate reuse may try one current option once as a fast path. If that fast path misses, canonical evaluation restarts from option 1. During canonical evaluation, the worker advances to the next option only for `out-of-host-capacity`, `LimitExceeded`, and `QuotaExceeded`, and only after same-option retries are exhausted. Other failures are terminal for that launch attempt.

	When used on create, omitting `overridePolicy` preserves the existing base-plus-placement behavior. When used on update, omitting `overridePolicy` leaves the stored policy unchanged, a non-null value replaces the stored policy, and `null` clears it. Read surfaces omit `overridePolicy` when no policy is stored on the pool. 
	* `override_statements` - (Required) (Updatable) The ordered set of availability domain-scoped override statements for this policy. 
		* `name` - (Required) (Updatable) A user-friendly name for the override statement. 
		* `override_options` - (Required) (Updatable) The ordered list of candidate override options for the matched availability domain.

			Request order defines the canonical candidate sequence. Override options within the same statement must normalize to distinct override payloads. 
			* `name` - (Required) (Updatable) A user-friendly name for the override option. 
			* `override` - (Required) (Updatable) The concrete launch mutation payload for an override option.

				In V1, the override payload is limited to `shape`, `shapeConfigDetails`, and `capacityReservationId`.

				Each override option must include at least one of `shape` or `capacityReservationId`, so a `capacityReservationId`-only override is valid. `shapeConfigDetails` may be supplied only when `shape` is supplied. If `shape` is a Flex or Fungible shape, `shapeConfigDetails` is required.

				Placement-owned and system-owned fields are not supported in this payload. Any field outside this V1 payload shape is rejected. 
				* `capacity_reservation_id` - (Optional) (Updatable) The OCID of the compute capacity reservation to use for the override option. 
				* `shape` - (Optional) (Updatable) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

					You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
				* `shape_config_details` - (Optional) (Updatable) The shape configuration requested for the override option. 
					* `baseline_ocpu_utilization` - (Optional) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

						The following values are supported:
						* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
						* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
						* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
					* `local_volume_size_in_gbs` - (Optional) (Updatable) The NVMe-backed local storage capacity, in GB, for flexible dense (DenseLV) VM shapes. If the selected shape  is DenseLV, the value must be greater than 0. For all other shapes, the value must be null (if specified);  any non-null value for a non-DenseLV shape results in an error. 
					* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to the instance, in gigabytes. 
					* `nvmes` - (Optional) (Updatable) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available. 
					* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the instance. 
					* `resource_management` - (Optional) (Updatable) This field is reserved for internal use. 
					* `vcpus` - (Optional) (Updatable) The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2. 
		* `selector` - (Required) (Updatable) The selector for an override statement.

			In V1, the selector contains only the scalar `availabilityDomain`. Additional selector fields are rejected. 
			* `availability_domain` - (Required) (Updatable) The availability domain that this statement applies to.

				The value must match an availability domain present in the pool's `placementConfigurations`. 
* `placement_configurations` - (Required) (Updatable) The placement configurations for the instance pool. Provide one placement configuration for each availability domain.

	To use the instance pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. To use compute cluster with instance pool, provide a single placement configuration. 
	* `availability_domain` - (Required) (Updatable) The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
	* `compute_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) that the instance will be created in.

		Make sure the compute cluster belongs to the same availability domain as specified in placement configuration otherwise the request will be rejected with 400. Once this field is set, it cannot be updated. Also any update to the availability domain in placement configuration will be blocked. 
	* `fault_domains` - (Optional) (Updatable) The fault domains to place instances.

		If you don't provide any values, the system makes a best effort to distribute instances across all fault domains based on capacity.

		To distribute the instances evenly across selected fault domains, provide a set of fault domains. For example, you might want instances to be evenly distributed if your applications require high availability.

		To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

		Example: `[FAULT-DOMAIN-1, FAULT-DOMAIN-2, FAULT-DOMAIN-3]`
	* `primary_subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. This field is deprecated. Use `primaryVnicSubnets` instead to set VNIC data for instances in the pool. 
	* `primary_vnic_subnets` - (Optional) (Updatable) Details about the IPv6 primary subnet.
		* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) (Updatable) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6subnet_cidr` - (Optional) (Updatable) Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
		* `is_assign_ipv6ip` - (Optional) (Updatable) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
		* `subnet_id` - (Required) (Updatable) The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	* `secondary_vnic_subnets` - (Optional) (Updatable) The set of secondary VNIC data for instances in the pool.
		* `display_name` - (Optional) (Updatable) The display name of the VNIC. This is also used to match against the instance configuration defined secondary VNIC.
		* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) (Updatable) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6subnet_cidr` - (Optional) (Updatable) Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation. 
		* `is_assign_ipv6ip` - (Optional) (Updatable) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
		* `subnet_id` - (Required) (Updatable) The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
* `size` - (Required) (Updatable) The number of instances that should be in the instance pool. Modifying this value will override the size of the instance pool. If the instance pool is linked with autoscaling configuration, autoscaling configuration could resize the instance pool at a later point. The instance pool's actual size may differ from the configured size if it is associated with an autoscaling configuration, instance pool's actual size will be reflected in this size attribute. For a GMC-enabled resource pool, this is the number of GMC resources that should be in the pool.
* `state` - (Optional) (Updatable) The target state for the instance pool update operation (ignored at create time and should not be set). Could be set to RUNNING or STOPPED.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `load_balancers` - The load balancers attached to the instance pool. (Note: From 6.16.0 load_balancers field in oci_core_instance_pool is changed from TypeList to TypeSet - to support load balancer insert operation. Also, LB cant by accessed by index)
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
* `size` - The number of instances that should be in the instance pool.
* `actual_size` - The number of actual instances in the instance pool on the cloud. This attribute will be different when instance pool is used along with autoScaling Configuration.
* `state` - The current state of the instance pool.
* `time_created` - The date and time the instance pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Instance Pool
	* `update` - (Defaults to 1 hours), when updating the Instance Pool
	* `delete` - (Defaults to 1 hours), when destroying the Instance Pool


## Import

InstancePools can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_pool.test_instance_pool "id"
```

