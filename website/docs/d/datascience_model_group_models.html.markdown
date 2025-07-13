---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_group_models"
sidebar_current: "docs-oci-datasource-datascience-model_group_models"
description: |-
  Provides the list of Model Group Models in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_group_models
This data source provides the list of Model Group Models in Oracle Cloud Infrastructure Data Science service.

Lists all models associated with the modelGroup in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_model_group_models" "test_model_group_models" {
	#Required
	compartment_id = var.compartment_id
	model_group_id = oci_datascience_model_group.test_model_group.id

	#Optional
	created_by = var.model_group_model_created_by
	display_name = var.model_group_model_display_name
	id = var.model_group_model_id
	state = var.model_group_model_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `model_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroup.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `model_group_models` - The list of model_group_models.

### ModelGroupModel Reference

The following attributes are exported:

* `category` - The category of the model.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name of the model.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group.
* `is_model_by_reference` - Identifier to indicate whether a model artifact resides in the Service Tenancy or Customer Tenancy.
* `lifecycle_details` - Details about the lifecycle state of the model.
* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model.
* `state` - The state of the model.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

