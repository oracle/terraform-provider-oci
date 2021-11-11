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

