---
subcategory: "Datascience"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_models"
sidebar_current: "docs-oci-datasource-datascience-models"
description: |-
  Provides the list of Models in Oracle Cloud Infrastructure Datascience service
---

# Data Source: oci_datascience_models
This data source provides the list of Models in Oracle Cloud Infrastructure Datascience service.

Lists models in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_models" "test_models" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	created_by = "${var.model_created_by}"
	display_name = "${var.model_display_name}"
	id = "${var.model_id}"
	project_id = "${oci_datascience_project.test_project.id}"
	state = "${var.model_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `models` - The list of models.

### Model Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the user who created the model.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short blurb describing the model.
* `display_name` - A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project associated with the model.
* `state` - The state of the model.
* `time_created` - The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

