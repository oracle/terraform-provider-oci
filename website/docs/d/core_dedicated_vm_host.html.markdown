---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dedicated_vm_host"
sidebar_current: "docs-oci-datasource-core-dedicated_vm_host"
description: |-
  Provides details about a specific Dedicated Vm Host in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_dedicated_vm_host
This data source provides details about a specific Dedicated Vm Host resource in Oracle Cloud Infrastructure Core service.

Gets information about the specified dedicated virtual machine host.

## Example Usage

```hcl
data "oci_core_dedicated_vm_host" "test_dedicated_vm_host" {
	#Required
	dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
}
```

## Argument Reference

The following arguments are supported:

* `dedicated_vm_host_id` - (Required) The OCID of the dedicated VM host.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the dedicated virtual machine host is running in.  Example: `Uocm:PHX-AD-1` 
* `capacity_bins` - A list of total and remaining CPU and memory per capacity bucket. 
	* `capacity_index` - Zero-based index for the corresponding capacity bucket. 
	* `remaining_memory_in_gbs` - The remaining memory of the capacity bucket, in GBs. 
	* `remaining_ocpus` - The available OCPUs of the capacity bucket. 
	* `supported_shapes` - List of VMI shapes supported on each capacity bucket. 
	* `total_memory_in_gbs` - The total memory of the capacity bucket, in GBs. 
	* `total_ocpus` - The total OCPUs of the capacity bucket. 
* `capacity_config` - The capacity configuration selected to be configured for the Dedicated Virtual Machine host.  Run [ListDedicatedVmHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DedicatedVmHostShapeSummary/ListDedicatedVmHostShapes) API to see details of this capacity configuration. 
* `compartment_id` - The OCID of the compartment that contains the dedicated virtual machine host.
* `compute_bare_metal_host_id` - The compute bare metal host OCID of the dedicated virtual machine host. 
* `dedicated_vm_host_shape` - The dedicated virtual machine host shape. The shape determines the number of CPUs and other resources available for VMs. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - The fault domain for the dedicated virtual machine host's assigned instances. For more information, see [Fault Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm#fault).

	If you do not specify the fault domain, the system selects one for you. To change the fault domain for a dedicated virtual machine host, delete it, and then create a new dedicated virtual machine host in the preferred fault domain.

	To get a list of fault domains, use the `ListFaultDomains` operation in the [Identity and Access Management Service API](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/).

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated VM host. 
* `is_memory_encryption_enabled` - Specifies if the Dedicated Virtual Machine Host (DVMH) is restricted to running only Confidential VMs. If `true`, only Confidential VMs can be launched. If `false`, Confidential VMs cannot be launched. 
* `placement_constraint_details` - The details for providing placement constraints. 
	* `compute_bare_metal_host_id` - The OCID of the compute bare metal host. This is only available for dedicated capacity customers.
	* `compute_host_group_id` - The OCID of the compute host group. This is only available for dedicated capacity customers.
	* `type` - The type for the placement constraints. Use `COMPUTE_BARE_METAL_HOST` when specifying the compute bare metal host OCID. Use `HOST_GROUP` when specifying the compute host group OCID. 
* `remaining_memory_in_gbs` - The current available memory of the dedicated VM host, in GBs. 
* `remaining_ocpus` - The current available OCPUs of the dedicated VM host. 
* `state` - The current state of the dedicated VM host. 
* `time_created` - The date and time the dedicated VM host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `total_memory_in_gbs` - The current total memory of the dedicated VM host, in GBs. 
* `total_ocpus` - The current total OCPUs of the dedicated VM host. 

