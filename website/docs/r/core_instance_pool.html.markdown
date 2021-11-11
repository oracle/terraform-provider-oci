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

Create an instance pool.

## Example Usage

```hcl
resource "oci_core_instance_pool" "test_instance_pool" {
	#Required
	compartment_id = var.compartment_id
	instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
	placement_configurations {
		#Required
		availability_domain = var.instance_pool_placement_configurations_availability_domain
		primary_subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		fault_domains = var.instance_pool_placement_configurations_fault_domains
		secondary_vnic_subnets {
			#Required
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			display_name = var.instance_pool_placement_configurations_secondary_vnic_subnets_display_name
		}
	}
	size = var.instance_pool_size

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_pool_display_name
	freeform_tags = {"Department"= "Finance"}
	load_balancers {
		#Required
		backend_set_name = oci_load_balancer_backend_set.test_backend_set.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
		port = var.instance_pool_load_balancers_port
		vnic_selection = var.instance_pool_load_balancers_vnic_selection
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
* `load_balancers` - (Optional) The load balancers to attach to the instance pool. 
	* `backend_set_name` - (Required) The name of the backend set on the load balancer to add instances to.
	* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to attach to the instance pool. 
	* `port` - (Required) The port value to use when creating the backend set.
	* `vnic_selection` - (Required) Indicates which VNIC on each instance in the pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool. 
* `placement_configurations` - (Required) (Updatable) The placement configurations for the instance pool. Provide one placement configuration for each availability domain.

	To use the instance pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. 
	* `availability_domain` - (Required) (Updatable) The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
	* `fault_domains` - (Optional) (Updatable) The fault domains to place instances.

		If you don't provide any values, the system makes a best effort to distribute instances across all fault domains based on capacity.

		To distribute the instances evenly across selected fault domains, provide a set of fault domains. For example, you might want instances to be evenly distributed if your applications require high availability.

		To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

		Example: `[FAULT-DOMAIN-1, FAULT-DOMAIN-2, FAULT-DOMAIN-3]` 
	* `primary_subnet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. 
	* `secondary_vnic_subnets` - (Optional) (Updatable) The set of secondary VNIC data for instances in the pool.
		* `display_name` - (Optional) (Updatable) The display name of the VNIC. This is also use to match against the instance configuration defined secondary VNIC. 
		* `subnet_id` - (Required) (Updatable) The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
* `size` - (Required) (Updatable) The number of instances that should be in the instance pool. Modifying this value will override the size of the instance pool. If the instance pool is linked with autoscaling configuration, autoscaling configuration could resize the instance pool at a later point. The instance pool's actual size may differ from the configured size if it is associated with an autoscaling configuration. For the actual size of the instance pool, refer to the `actual_size` attribute.
* `state` - (Optional) (Updatable) The target state for the instance pool update operation (ignored at create time and should not be set). Could be set to RUNNING or STOPPED.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

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
	* `primary_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. 
	* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
		* `display_name` - The display name of the VNIC. This is also use to match against the instance configuration defined secondary VNIC. 
		* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
* `size` - The number of instances that should be in the instance pool.
* `actual_size` - The number of actual instances in the instance pool on the cloud. This attribute will be different when instance pool is used along with autoScaling Configuration.
* `state` - The current state of the instance pool.
* `time_created` - The date and time the instance pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Instance Pool
	* `update` - (Defaults to 1 hours), when updating the Instance Pool
	* `delete` - (Defaults to 1 hours), when destroying the Instance Pool


## Import

InstancePools can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_pool.test_instance_pool "id"
```

