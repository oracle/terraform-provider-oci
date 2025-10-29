---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_software_images"
sidebar_current: "docs-oci-datasource-database-autonomous_database_software_images"
description: |-
  Provides the list of Autonomous Database Software Images in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_software_images
This data source provides the list of Autonomous Database Software Images in Oracle Cloud Infrastructure Database service.

Gets a list of the Autonomous AI Database Software Images in the specified compartment.


## Example Usage

```hcl
data "oci_database_autonomous_database_software_images" "test_autonomous_database_software_images" {
	#Required
	compartment_id = var.compartment_id
	image_shape_family = var.autonomous_database_software_image_image_shape_family

	#Optional
	display_name = var.autonomous_database_software_image_display_name
	state = var.autonomous_database_software_image_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `image_shape_family` - (Required) A filter to return only resources that match the given image shape family exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_database_software_image_collection` - The list of autonomous_database_software_image_collection.

### AutonomousDatabaseSoftwareImage Reference

The following attributes are exported:

* `autonomous_dsi_one_off_patches` - One-off patches included in the Autonomous AI Database Software Image
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_version` - The database version with which the Autonomous AI Database Software Image is to be built.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `display_name` - The user-friendly name for the Autonomous AI Database Software Image. The name does not have to be unique.
* `dst_file_version` - DST Time-Zone File version of the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database Software Image.
* `image_shape_family` - To what shape the image is meant for.
* `lifecycle_details` - Detailed message for the lifecycle state.
* `release_update` - The Release Updates.
* `state` - The current state of the Autonomous AI Database Software Image.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Autonomous AI Database Software Image was created.

