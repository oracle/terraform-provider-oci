---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dedicated_vm_host_instance_shapes"
sidebar_current: "docs-oci-datasource-core-dedicated_vm_host_instance_shapes"
description: |-
  Provides the list of Dedicated Vm Host Instance Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_dedicated_vm_host_instance_shapes
This data source provides the list of Dedicated Vm Host Instance Shapes in Oracle Cloud Infrastructure Core service.

Lists the shapes that can be used to launch a virtual machine instance on a dedicated virtual machine host within the specified compartment.
You can filter the list by compatibility with a specific dedicated virtual machine host shape.


## Example Usage

```hcl
data "oci_core_dedicated_vm_host_instance_shapes" "test_dedicated_vm_host_instance_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.dedicated_vm_host_instance_shape_availability_domain}"
	dedicated_vm_host_shape = "${var.dedicated_vm_host_instance_shape_dedicated_vm_host_shape}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `dedicated_vm_host_shape` - (Optional) Dedicated VM host shape name 


## Attributes Reference

The following attributes are exported:

* `dedicated_vm_host_instance_shapes` - The list of dedicated_vm_host_instance_shapes.

### DedicatedVmHostInstanceShape Reference

The following attributes are exported:

* `availability_domain` - The shape's availability domain. 
* `instance_shape_name` - The name of the virtual machine instance shapes that can be launched on a dedicated VM host. 

