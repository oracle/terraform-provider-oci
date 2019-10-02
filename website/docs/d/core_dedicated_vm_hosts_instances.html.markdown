---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dedicated_vm_hosts_instances"
sidebar_current: "docs-oci-datasource-core-dedicated_vm_hosts_instances"
description: |-
  Provides the list of Dedicated Vm Hosts Instances in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_dedicated_vm_hosts_instances
This data source provides the list of Dedicated Vm Hosts Instances in Oracle Cloud Infrastructure Core service.

Returns the list of instances on the dedicated virtual machine hosts that match the specified criteria.


## Example Usage

```hcl
data "oci_core_dedicated_vm_hosts_instances" "test_dedicated_vm_hosts_instances" {
	#Required
	compartment_id = "${var.compartment_id}"
	dedicated_vm_host_id = "${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}"

	#Optional
	availability_domain = "${var.dedicated_vm_hosts_instance_availability_domain}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `dedicated_vm_host_id` - (Required) The OCID of the dedicated VM host.


## Attributes Reference

The following attributes are exported:

* `dedicated_vm_host_instances` - The list of dedicated_vm_host_instances.

### DedicatedVmHostsInstance Reference

The following attributes are exported:

* `availability_domain` - The availability domain the virtual machine instance is running in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the virtual machine instance. 
* `instance_id` - The OCID of the virtual machine instance. 
* `shape` - The shape of the VM instance. 
* `time_created` - The date and time the virtual machine instance was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

