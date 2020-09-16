---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dedicated_vm_host_shapes"
sidebar_current: "docs-oci-datasource-core-dedicated_vm_host_shapes"
description: |-
  Provides the list of Dedicated Vm Host Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_dedicated_vm_host_shapes
This data source provides the list of Dedicated Vm Host Shapes in Oracle Cloud Infrastructure Core service.

Lists the shapes that can be used to launch a dedicated virtual machine host within the specified compartment.


## Example Usage

```hcl
data "oci_core_dedicated_vm_host_shapes" "test_dedicated_vm_host_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.dedicated_vm_host_shape_availability_domain
	instance_shape_name = var.dedicated_vm_host_shape_instance_shape_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `instance_shape_name` - (Optional) The name for the instance's shape. 


## Attributes Reference

The following attributes are exported:

* `dedicated_vm_host_shapes` - The list of dedicated_vm_host_shapes.

### DedicatedVmHostShape Reference

The following attributes are exported:

* `availability_domain` - The shape's availability domain. 
* `dedicated_vm_host_shape` - The name of the dedicated vm host shape. You can enumerate all available shapes by calling [ListDedicatedVmHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/dedicatedVmHostShapes). 

