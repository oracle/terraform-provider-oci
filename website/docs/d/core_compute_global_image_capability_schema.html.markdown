---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_global_image_capability_schema"
sidebar_current: "docs-oci-datasource-core-compute_global_image_capability_schema"
description: |-
  Provides details about a specific Compute Global Image Capability Schema in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_global_image_capability_schema
This data source provides details about a specific Compute Global Image Capability Schema resource in Oracle Cloud Infrastructure Core service.

Gets the specified Compute Global Image Capability Schema

## Example Usage

```hcl
data "oci_core_compute_global_image_capability_schema" "test_compute_global_image_capability_schema" {
	#Required
	compute_global_image_capability_schema_id = oci_core_compute_global_image_capability_schema.test_compute_global_image_capability_schema.id
}
```

## Argument Reference

The following arguments are supported:

* `compute_global_image_capability_schema_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute global image capability schema


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the compute global image capability schema 
* `current_version_name` - The name of the global capabilities version resource that is considered the current version.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute global image capability schema 
* `time_created` - The date and time the compute global image capability schema was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

