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
	name = var.supported_host_shape_name
	sddc_type = var.supported_host_shape_sddc_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) A filter to return only resources that match the given name exactly.
* `sddc_type` - (**Deprecated**) (Optional) A filter to return only resources that match the given SDDC type exactly.


## Attributes Reference

The following attributes are exported:
* `items` - The list of the supported compute shapes for ESXi hosts.

### Items Reference
  * `default_ocpu_count` - The default OCPU count of the shape. 
  * `description` - Description of the shape. 
  * `is_support_shielded_instances` - Indicates whether the shape supports shielded instances.
  * `name` - The name of the supported compute shape. 
  * `shape_family` - The family of the shape. ESXi hosts of one SDDC must have the same shape family. 
  * `supported_ocpu_count` - Support OCPU count of the shape. 
  * `supported_operations` - The operations where you can use the shape. The operations can be CREATE_SDDC or CREATE_ESXI_HOST. 
  * `supported_sddc_types` - The supported SDDC types for the shape. 
  * `supported_vmware_software_versions` - The VMware software versions supported by the shape.
* `supported_host_shape_collection` - The list of supported_host_shape_collection.

### SupportedHostShape Reference

The following attributes are exported:

* `items` - A list of the supported compute shapes for ESXi hosts.
	* `default_ocpu_count` - The default OCPU count of the shape. 
	* `description` - Description of the shape. 
	* `is_support_monthly_sku` - Whether the shape supports "MONTH" SKU.
	* `is_support_shielded_instances` - Indicates whether the shape supports shielded instances.
	* `name` - The name of the supported compute shape. 
	* `shape_family` - The family of the shape. ESXi hosts of one SDDC must have the same shape family. 
	* `supported_ocpu_count` - Support OCPU count of the shape. 
	* `supported_operations` - The operations where you can use the shape. The operations can be CREATE_SDDC or CREATE_ESXI_HOST. 
	* `supported_sddc_types` - The supported SDDC types for the shape. 

