---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dedicated_vm_hosts"
sidebar_current: "docs-oci-datasource-core-dedicated_vm_hosts"
description: |-
  Provides the list of Dedicated Vm Hosts in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_dedicated_vm_hosts
This data source provides the list of Dedicated Vm Hosts in Oracle Cloud Infrastructure Core service.

Returns the list of dedicated virtual machine hosts that match the specified criteria in the specified compartment.

You can limit the list by specifying a dedicated virtual machine host display name. The list will include all the identically-named
dedicated virtual machine hosts in the compartment.


## Example Usage

```hcl
data "oci_core_dedicated_vm_hosts" "test_dedicated_vm_hosts" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.dedicated_vm_host_availability_domain
	display_name = var.dedicated_vm_host_display_name
	instance_shape_name = var.dedicated_vm_host_instance_shape_name
	remaining_memory_in_gbs_greater_than_or_equal_to = var.dedicated_vm_host_remaining_memory_in_gbs_greater_than_or_equal_to
	remaining_ocpus_greater_than_or_equal_to = var.dedicated_vm_host_remaining_ocpus_greater_than_or_equal_to
	state = var.dedicated_vm_host_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `instance_shape_name` - (Optional) The name for the instance's shape. 
* `remaining_memory_in_gbs_greater_than_or_equal_to` - (Optional) The remaining memory of the dedicated VM host, in GBs.
* `remaining_ocpus_greater_than_or_equal_to` - (Optional) The available OCPUs of the dedicated VM host.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `dedicated_vm_hosts` - The list of dedicated_vm_hosts.

### DedicatedVmHost Reference

The following attributes are exported:

* `availability_domain` - The availability domain the dedicated virtual machine host is running in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the dedicated virtual machine host.
* `dedicated_vm_host_shape` - The dedicated virtual machine host shape. The shape determines the number of CPUs and other resources available for VMs. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - The fault domain for the dedicated virtual machine host's assigned instances. For more information, see [Fault Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm#fault).

	If you do not specify the fault domain, the system selects one for you. To change the fault domain for a dedicated virtual machine host, delete it, and then create a new dedicated virtual machine host in the preferred fault domain.

	To get a list of fault domains, use the `ListFaultDomains` operation in the [Identity and Access Management Service API](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/).

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated VM host. 
* `remaining_memory_in_gbs` - The current available memory of the dedicated VM host, in GBs. 
* `remaining_ocpus` - The current available OCPUs of the dedicated VM host. 
* `state` - The current state of the dedicated VM host. 
* `time_created` - The date and time the dedicated VM host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `total_memory_in_gbs` - The current total memory of the dedicated VM host, in GBs. 
* `total_ocpus` - The current total OCPUs of the dedicated VM host. 

