---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_version_sets"
sidebar_current: "docs-oci-datasource-datascience-model_version_sets"
description: |-
  Provides the list of Model Version Sets in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_version_sets
This data source provides the list of Model Version Sets in Oracle Cloud Infrastructure Data Science service.

Lists model version sets in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_model_version_sets" "test_model_version_sets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.model_version_set_created_by
	id = var.model_version_set_id
	name = var.model_version_set_name
	project_id = oci_datascience_project.test_project.id
	state = var.model_version_set_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `model_version_sets` - The list of model_version_sets.

### ModelVersionSet Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model version set compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model version set.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the model version set.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model version set.
* `name` - A user-friendly name for the resource. It must be unique and can't be modified.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model version set.
* `state` - The state of the model version set.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 
* `time_updated` - The date and time that the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

