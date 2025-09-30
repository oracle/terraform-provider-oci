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
* `capacity_configs` - A list of capacity configs that are supported by this dedicated VM host shape. 
	* `capacity_bins` - A list of total CPU and memory per capacity bucket. 
		* `capacity_index` - Zero-based index for the corresponding capacity bucket. 
		* `supported_shapes` - List of VMI shapes supported on each capacity bucket. 
		* `total_memory_in_gbs` - The total memory of the capacity bucket, in GBs. 
		* `total_ocpus` - The total OCPUs of the capacity bucket. 
	* `capacity_config_name` - The name of each capacity config. 
	* `is_default` - Whether this capacity config is the default config. 
	* `supported_capabilities` - Specifies the capabilities that the Dedicated Virtual Machine Host (DVMH) Shape or Virtual Machine Instance Shape could support. 
		* `is_memory_encryption_supported` - Whether the DVMH shape could support confidential VMs or the VM instance shape could be confidential. 
* `dedicated_vm_host_shape` - The name of the dedicated VM host shape. You can enumerate all available shapes by calling [ListDedicatedVmHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DedicatedVmHostShapeSummary/ListDedicatedVmHostShapes). 

