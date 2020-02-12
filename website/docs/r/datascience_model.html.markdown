---
subcategory: "Datascience"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model"
sidebar_current: "docs-oci-resource-datascience-model"
description: |-
  Provides the Model resource in Oracle Cloud Infrastructure Datascience service
---

# oci_datascience_model
This resource provides the Model resource in Oracle Cloud Infrastructure Datascience service.

Creates a new model.

## Example Usage

```hcl
resource "oci_datascience_model" "test_model" {
	#Required
	artifact_content_length = "${var.artifact_content_length}"
    model_artifact = "${var.model_artifact}"
	compartment_id = "${var.compartment_id}"
	project_id = "${oci_datascience_project.test_project.id}"

	#Optional
	artifact_content_disposition = "${var.content_disposition}" 
	defined_tags = {"Operations.CostCenter"= "42"}
	description = "${var.model_description}"
	display_name = "${var.model_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `artifact_content_length` - (Required) The length of the artifact to upload.
* `artifact_content_disposition` - (Optional) The content disposition of the artifact to upload.
* `model_artifact` - (Required) The model artifact to upload.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the compartment to create the model in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short blurb describing the model.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My Model` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project to associate with the model.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `artifact_content_length` - The length of the artifact.
* `artifact_content_disposition` - The content disposition of the artifact.
* `artifact_content_md5` - The base-64 encoded MD5 hash of the artifact.
* `artifact_last_modified` - The artifact modification time.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model’s compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the user who created the model.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short blurb describing the model.
* `display_name` - A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project associated with the model.
* `state` - The state of the model.
* `time_created` - The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

## Import

Import is not supported for this resource.

