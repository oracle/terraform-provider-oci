---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_group"
sidebar_current: "docs-oci-resource-datascience-model_group"
description: |-
  Provides the Model Group resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model_group
This resource provides the Model Group resource in Oracle Cloud Infrastructure Data Science service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-science/latest/ModelGroup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datascience

Create a new Model Group resource.

## Example Usage

```hcl
resource "oci_datascience_model_group" "test_model_group" {
	#Required
	compartment_id = var.compartment_id
	create_type = var.model_group_create_type
	project_id = oci_datascience_project.test_project.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.model_group_description
	display_name = var.model_group_display_name
	freeform_tags = {"Department"= "Finance"}
	member_model_entries {

		#Optional
		member_model_details {

			#Optional
			inference_key = var.model_group_member_model_entries_member_model_details_inference_key
			model_id = oci_datascience_model.test_model.id
		}
	}
	model_group_clone_source_details {
		#Required
		model_group_clone_source_type = var.model_group_model_group_clone_source_details_model_group_clone_source_type
		source_id = oci_datascience_source.test_source.id

		#Optional
		modify_model_group_details {

			#Optional
			defined_tags = var.model_group_model_group_clone_source_details_modify_model_group_details_defined_tags
			description = var.model_group_model_group_clone_source_details_modify_model_group_details_description
			display_name = var.model_group_model_group_clone_source_details_modify_model_group_details_display_name
			freeform_tags = var.model_group_model_group_clone_source_details_modify_model_group_details_freeform_tags
			model_group_details {
				#Required
				type = var.model_group_model_group_clone_source_details_modify_model_group_details_model_group_details_type

				#Optional
				base_model_id = oci_datascience_model.test_model.id
				custom_metadata_list {

					#Optional
					category = var.model_group_model_group_clone_source_details_modify_model_group_details_model_group_details_custom_metadata_list_category
					description = var.model_group_model_group_clone_source_details_modify_model_group_details_model_group_details_custom_metadata_list_description
					key = var.model_group_model_group_clone_source_details_modify_model_group_details_model_group_details_custom_metadata_list_key
					value = var.model_group_model_group_clone_source_details_modify_model_group_details_model_group_details_custom_metadata_list_value
				}
			}
			model_group_version_history_id = oci_datascience_model_group_version_history.test_model_group_version_history.id
			version_label = var.model_group_model_group_clone_source_details_modify_model_group_details_version_label
		}
		patch_model_group_member_model_details {

			#Optional
			items {
				#Required
				operation = var.model_group_model_group_clone_source_details_patch_model_group_member_model_details_items_operation
				values {
					#Required
					model_id = oci_datascience_model.test_model.id

					#Optional
					inference_key = var.model_group_model_group_clone_source_details_patch_model_group_member_model_details_items_values_inference_key
				}
			}
		}
	}
	model_group_details {
		#Required
		type = var.model_group_model_group_details_type

		#Optional
		base_model_id = oci_datascience_model.test_model.id
		custom_metadata_list {

			#Optional
			category = var.model_group_model_group_details_custom_metadata_list_category
			description = var.model_group_model_group_details_custom_metadata_list_description
			key = var.model_group_model_group_details_custom_metadata_list_key
			value = var.model_group_model_group_details_custom_metadata_list_value
		}
	}
	model_group_version_history_id = oci_datascience_model_group_version_history.test_model_group_version_history.id
	version_label = var.model_group_version_label
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the modelGroup in.
* `create_type` - (Required) The type of the model group create operation.
* `defined_tags` - (Applicable when create_type=CREATE) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Applicable when create_type=CREATE) (Updatable) A short description of the modelGroup.
* `display_name` - (Applicable when create_type=CREATE) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My ModelGroup` 
* `freeform_tags` - (Applicable when create_type=CREATE) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `member_model_entries` - (Required when create_type=CREATE) List of member models (inferenceKey & modelId) to be associated with the model group.
	* `member_model_details` - (Applicable when create_type=CREATE) Each List item contains inference key and model ocid.
		* `inference_key` - (Applicable when create_type=CREATE) SaaS friendly name of the model.
		* `model_id` - (Required when create_type=CREATE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `model_group_clone_source_details` - (Required when create_type=CLONE) Model Group clone source details.
	* `model_group_clone_source_type` - (Required) Source resource for model group clone operation.
	* `modify_model_group_details` - (Optional) Overwrites the properties of the source modelGroup.
		* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `description` - (Optional) A short description of the modelGroup.
		* `display_name` - (Optional) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My ModelGroup` 
		* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
		* `model_group_details` - (Optional) The model group details.
			* `base_model_id` - (Required when type=STACKED) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model in the group that represents the base model for stacked deployment.
			* `custom_metadata_list` - (Applicable when model_group_clone_source_type=MODEL_GROUP | MODEL_GROUP_VERSION_HISTORY) An array of custom metadata details for the model group.
				* `category` - (Applicable when model_group_clone_source_type=MODEL_GROUP | MODEL_GROUP_VERSION_HISTORY) Category of the metadata.
				* `description` - (Applicable when model_group_clone_source_type=MODEL_GROUP | MODEL_GROUP_VERSION_HISTORY) Description of model metadata.
				* `key` - (Applicable when model_group_clone_source_type=MODEL_GROUP | MODEL_GROUP_VERSION_HISTORY) Key of the metadata.
				* `value` - (Applicable when model_group_clone_source_type=MODEL_GROUP | MODEL_GROUP_VERSION_HISTORY) Value of the metadata.
			* `type` - (Required) The type of the model group.
		* `model_group_version_history_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group version history to which the modelGroup is associated.
		* `version_label` - (Optional) An additional description of the lifecycle state of the model group.
	* `patch_model_group_member_model_details` - (Optional) Specifies the list of new models to be added and list of models from source model group to be removed for cloning.
		* `items` - (Optional) Array of patch instructions.
			* `operation` - (Required) A single instruction to be included as part of Patch request content. Enum type (INSERT and REMOVE).
			* `values` - (Required) Array of inference key and model OCID.
				* `inference_key` - (Optional) SaaS friendly name of the model.
				* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
	* `source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group version history.
* `model_group_details` - (Required when create_type=CREATE) The model group details.
	* `base_model_id` - (Required when type=STACKED) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model in the group that represents the base model for stacked deployment.
	* `custom_metadata_list` - (Applicable when create_type=CREATE) An array of custom metadata details for the model group.
		* `category` - (Applicable when create_type=CREATE) Category of the metadata.
		* `description` - (Applicable when create_type=CREATE) Description of model metadata.
		* `key` - (Applicable when create_type=CREATE) Key of the metadata.
		* `value` - (Applicable when create_type=CREATE) Value of the metadata.
	* `type` - (Required) The type of the model group.
* `model_group_version_history_id` - (Applicable when create_type=CREATE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model group version history to which the modelGroup is associated.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the modelGroup.
* `version_label` - (Applicable when create_type=CREATE) (Updatable) An additional description of the lifecycle state of the model group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model Group
	* `update` - (Defaults to 20 minutes), when updating the Model Group
	* `delete` - (Defaults to 20 minutes), when destroying the Model Group


## Import

ModelGroups can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_group.test_model_group "id"
```

