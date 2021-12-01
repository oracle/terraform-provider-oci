---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dedicated_vm_host"
sidebar_current: "docs-oci-resource-core-dedicated_vm_host"
description: |-
  Provides the Dedicated Vm Host resource in Oracle Cloud Infrastructure Core service
---

# oci_core_dedicated_vm_host
This resource provides the Dedicated Vm Host resource in Oracle Cloud Infrastructure Core service.

Creates a new dedicated virtual machine host in the specified compartment and the specified availability domain.
Dedicated virtual machine hosts enable you to run your Compute virtual machine (VM) instances on dedicated servers
that are a single tenant and not shared with other customers.
For more information, see [Dedicated Virtual Machine Hosts](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/dedicatedvmhosts.htm).


## Example Usage

```hcl
resource "oci_core_dedicated_vm_host" "test_dedicated_vm_host" {
	#Required
	availability_domain = var.dedicated_vm_host_availability_domain
	compartment_id = var.compartment_id
	dedicated_vm_host_shape = var.dedicated_vm_host_dedicated_vm_host_shape

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.dedicated_vm_host_display_name
	fault_domain = var.dedicated_vm_host_fault_domain
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the dedicated virtual machine host.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `dedicated_vm_host_shape` - (Required) The dedicated virtual machine host shape. The shape determines the number of CPUs and other resources available for VM instances launched on the dedicated virtual machine host. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - (Optional) The fault domain for the dedicated virtual machine host's assigned instances. For more information, see [Fault Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm#fault). If you do not specify the fault domain, the system selects one for you. To change the fault domain for a dedicated virtual machine host, delete it and create a new dedicated virtual machine host in the preferred fault domain.

	To get a list of fault domains, use the `ListFaultDomains` operation in the [Identity and Access Management Service API](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/).

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dedicated Vm Host
	* `update` - (Defaults to 20 minutes), when updating the Dedicated Vm Host
	* `delete` - (Defaults to 20 minutes), when destroying the Dedicated Vm Host


## Import

DedicatedVmHosts can be imported using the `id`, e.g.

```
$ terraform import oci_core_dedicated_vm_host.test_dedicated_vm_host "id"
```

