---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_supported_host_shapes"
sidebar_current: "docs-oci-datasource-ocvp-supported_host_shapes"
description: |-
  Provides the list of Supported Host Shapes in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_supported_host_shapes
This data source provides the list of Supported Host Shapes in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists supported compute shapes for ESXi hosts.


## Example Usage

```hcl
data "oci_ocvp_supported_host_shapes" "test_supported_host_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	initial_host_shape_name = oci_core_shape.test_shape.name
	is_single_host_sddc_supported = var.supported_host_shape_is_single_host_sddc_supported
	name = var.supported_host_shape_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `initial_host_shape_name` - (Optional) A filter to return only the shapes compatible with the initial host shape of the Cluster. 
* `is_single_host_sddc_supported` - (Optional) A filter to return only resources that support single host SDDC.
* `name` - (Optional) A filter to return only resources that match the given name exactly.
* `sddc_type` - (**Deprecated**) (Optional) A filter to return only resources that match the given SDDC type exactly.


## Attributes Reference

The following attributes are exported:

* `items` - A list of the supported compute shapes for ESXi hosts.
	* `default_ocpu_count` - The default OCPU count of the shape. 
	* `description` - Description of the shape. 
    * `is_support_monthly_sku` - (**Deprecated**) Whether the shape supports "MONTH" SKU.
	* `is_single_host_sddc_supported` - Indicates whether the shape supports single host SDDCs. 
	* `is_support_monthly_commitment` - Whether the shape supports "MONTH" Commitment.
	* `is_support_shielded_instances` - Indicates whether the shape supports shielded instances.
	* `name` - The name of the supported compute shape. 
	* `shape_family` - The family of the shape. ESXi hosts of one SDDC must have the same shape family. 
	* `supported_ocpu_count` - Support OCPU count of the shape. 
	* `supported_operations` - The operations where you can use the shape. The operations can be CREATE_SDDC or CREATE_ESXI_HOST. 
    * `supported_sddc_types` - (**Deprecated**) The supported SDDC types for the shape.
	* `supported_vmware_software_versions` - The VMware software versions supported by the shape. 

