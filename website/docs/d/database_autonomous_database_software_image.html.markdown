---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_software_image"
sidebar_current: "docs-oci-datasource-database-autonomous_database_software_image"
description: |-
  Provides details about a specific Autonomous Database Software Image in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_software_image
This data source provides details about a specific Autonomous Database Software Image resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Autonomous Database Software Image.

## Example Usage

```hcl
data "oci_database_autonomous_database_software_image" "test_autonomous_database_software_image" {
	#Required
	autonomous_database_software_image_id = oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_software_image_id` - (Required) The Autonomous Database Software Image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_dsi_one_off_patches` - One-off patches included in the Autonomous Database Software Image
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_version` - The database version with which the Autonomous Database Software Image is to be built.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Database Software Image. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database Software Image.
* `image_shape_family` - To what shape the image is meant for.
* `lifecycle_details` - Detailed message for the lifecycle state.
* `release_update` - The Release Updates.
* `state` - The current state of the Autonomous Database Software Image.
* `time_created` - The date and time the Autonomous Database Software Image was created.

