---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_software_image"
sidebar_current: "docs-oci-resource-database-autonomous_database_software_image"
description: |-
  Provides the Autonomous Database Software Image resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_software_image
This resource provides the Autonomous Database Software Image resource in Oracle Cloud Infrastructure Database service.

create Autonomous Database Software Image in the specified compartment.


## Example Usage

```hcl
resource "oci_database_autonomous_database_software_image" "test_autonomous_database_software_image" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.autonomous_database_software_image_display_name
	image_shape_family = var.autonomous_database_software_image_image_shape_family
	source_cdb_id = oci_database_source_cdb.test_source_cdb.id

	#Optional
	defined_tags = var.autonomous_database_software_image_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the Autonomous Database Software Image. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `image_shape_family` - (Required) To what shape the image is meant for.
* `source_cdb_id` - (Required) The source Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) from which to create Autonomous Database Software Image.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Database Software Image
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Database Software Image
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Database Software Image


## Import

AutonomousDatabaseSoftwareImages can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database_software_image.test_autonomous_database_software_image "id"
```

