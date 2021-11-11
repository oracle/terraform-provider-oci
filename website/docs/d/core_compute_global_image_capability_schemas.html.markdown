---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_global_image_capability_schemas"
sidebar_current: "docs-oci-datasource-core-compute_global_image_capability_schemas"
description: |-
  Provides the list of Compute Global Image Capability Schemas in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_global_image_capability_schemas
This data source provides the list of Compute Global Image Capability Schemas in Oracle Cloud Infrastructure Core service.

Lists Compute Global Image Capability Schema in the specified compartment.


## Example Usage

```hcl
data "oci_core_compute_global_image_capability_schemas" "test_compute_global_image_capability_schemas" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.compute_global_image_capability_schema_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) A filter to return only resources that match the given compartment OCID exactly. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_global_image_capability_schemas` - The list of compute_global_image_capability_schemas.

### ComputeGlobalImageCapabilitySchema Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the compute global image capability schema 
* `current_version_name` - The name of the global capabilities version resource that is considered the current version.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute global image capability schema 
* `time_created` - The date and time the compute global image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

