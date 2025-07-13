---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_groups"
sidebar_current: "docs-oci-datasource-datascience-model_groups"
description: |-
  Provides the list of Model Groups in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_groups
This data source provides the list of Model Groups in Oracle Cloud Infrastructure Data Science service.

Lists all the modelGroups in the specified compartment. The query must include compartmentId.

## Example Usage

```hcl
data "oci_datascience_model_groups" "test_model_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.model_group_created_by
	display_name = var.model_group_display_name
	id = var.model_group_id
	model_group_version_history_id = oci_datascience_model_group_version_history.test_model_group_version_history.id
	project_id = oci_datascience_project.test_project.id
	state = var.model_group_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Applicable when create_type=CREATE) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `model_group_version_history_id` - (Applicable when create_type=CREATE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroupVersionHistory.
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
* `state` - (Optional) A filter to return resources matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `model_groups` - The list of model_groups.

### ModelGroup Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroup's compartment.
* `create_type` - The type of the model group create operation.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the modelGroup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the modelGroup.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroup.
* `lifecycle_details` - Details about the lifecycle state of the model group.
* `member_model_entries` - List of member models (inferenceKey & modelId) to be associated with the model group.
	* `member_model_details` - Each List item contains inference key and model ocid.
		* `inference_key` - SaaS friendly name of the model.
		* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `model_group_details` - The model group details.
	* `base_model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model in the group that represents the base model for stacked deployment.
	* `custom_metadata_list` - An array of custom metadata details for the model group.
		* `category` - Category of the metadata.
		* `description` - Description of model metadata.
		* `key` - Key of the metadata.
		* `value` - Value of the metadata.
	* `type` - The type of the model group.
* `model_group_version_history_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group version history to which the modelGroup is associated.
* `model_group_version_history_name` - The name of the model group version history to which the model group is associated.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the modelGroup.
* `source_model_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group used for the clone operation.
* `state` - The state of the modelGroup.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 
* `time_updated` - The date and time the resource was last updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 
* `version_id` - Unique identifier assigned to each version of the model group. It would be auto-incremented number generated by service.
* `version_label` - An additional description of the lifecycle state of the model group.

