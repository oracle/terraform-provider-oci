---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_application"
sidebar_current: "docs-oci-resource-dataintegration-workspace_application"
description: |-
  Provides the Workspace Application resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_application
This resource provides the Workspace Application resource in Oracle Cloud Infrastructure Data Integration service.

Creates an application.


## Example Usage

```hcl
resource "oci_dataintegration_workspace_application" "test_workspace_application" {
	#Required
	identifier = var.workspace_application_identifier
	name = var.workspace_application_name
	workspace_id = oci_dataintegration_workspace.test_workspace.id
	model_type = var.workspace_application_model_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.workspace_application_description
	display_name = var.workspace_application_display_name
	freeform_tags = {"bar-key"= "value"}
	key = var.workspace_application_key
	model_version = var.workspace_application_model_version
	object_status = var.workspace_application_object_status
	registry_metadata {

		#Optional
		aggregator_key = var.workspace_application_registry_metadata_aggregator_key
		is_favorite = var.workspace_application_registry_metadata_is_favorite
		key = var.workspace_application_registry_metadata_key
		labels = var.workspace_application_registry_metadata_labels
		registry_version = var.workspace_application_registry_metadata_registry_version
	}
	source_application_info {

		#Optional
		application_key = var.workspace_application_source_application_info_application_key
		copy_type = var.workspace_application_source_application_info_copy_type
		workspace_id = oci_dataintegration_workspace.test_workspace.id
	}
	state = var.workspace_application_state
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Detailed description for the object.
* `display_name` - (Optional) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `key` - (Optional) (Updatable) Currently not used on application creation. Reserved for future.
* `model_type` - (Required) (Updatable) The type of the application.
* `model_version` - (Optional) (Updatable) The object's model version.
* `name` - (Required) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `registry_metadata` - (Optional) Information about the object and its parent.
	* `aggregator_key` - (Optional) The owning object's key for this object.
	* `is_favorite` - (Optional) Specifies whether this object is a favorite or not.
	* `key` - (Optional) The identifying key for the object.
	* `labels` - (Optional) Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	* `registry_version` - (Optional) The registry version.
* `source_application_info` - (Optional) The information about the application.
	* `application_key` - (Optional) The source application key to use when creating the application.
	* `copy_type` - (Optional) Parameter to specify the link between SOURCE and TARGET application after copying. CONNECTED    - Indicate that TARGET application is conneced to SOURCE and can be synced after copy. DISCONNECTED - Indicate that TARGET application is not conneced to SOURCE and can evolve independently.
	* `workspace_id` - (Optional) The OCID of the workspace containing the application. This allows cross workspace deployment to publish an application from a different workspace into the current workspace specified in this operation.
* `state` - (Optional) (Updatable) The current state of the workspace.
* `workspace_id` - (Required) The workspace ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_version` - The application's version.
* `compartment_id` - OCID of the compartment that this resource belongs to. Defaults to compartment of the Workspace.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `dependent_object_metadata` - A list of dependent objects in this patch.
	* `action` - The patch action indicating if object was created, updated, or deleted.
	* `identifier` - Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `key` - The key of the object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `name_path` - The fully qualified path of the published object, which would include its project and folder.
	* `object_version` - The object version.
	* `type` - The type of the object in patch.
* `description` - Detailed description for the object.
* `display_name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID of the resource that is used to uniquely identify the application
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `key` - Generated key that can be used in API calls to identify application.
* `key_map` - A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
* `metadata` - A summary type containing information about the object including its key, name and when/who created/updated it.
	* `aggregator` - A summary type containing information about the object's aggregator including its type, key, name and description.
		* `description` - The description of the aggregator.
		* `identifier` - The identifier of the aggregator.
		* `key` - The key of the aggregator object.
		* `name` - The name of the aggregator.
		* `type` - The type of the aggregator.
	* `aggregator_key` - The owning object key for this object.
	* `count_statistics` - A count statistics.
		* `object_type_count_list` - The array of statistics.
			* `object_count` - The value for the count statistic object.
			* `object_type` - The type of object for the count statistic object.
	* `created_by` - The user that created the object.
	* `created_by_name` - The user that created the object.
	* `identifier_path` - The full path to identify this object.
	* `info_fields` - Information property fields.
	* `is_favorite` - Specifies whether this object is a favorite or not.
	* `labels` - Labels are keywords or tags that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
	* `registry_version` - The registry version of the object.
	* `time_created` - The date and time that the object was created.
	* `time_updated` - The date and time that the object was updated.
	* `updated_by` - The user that updated the object.
	* `updated_by_name` - The user that updated the object.
* `model_type` - The object type.
* `model_version` - The object's model version.
* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - The version of the object that is used to track changes in the object instance.
* `parent_ref` - A reference to the object's parent.
	* `parent` - Key of the parent object.
	* `root_doc_id` - Key of the root document object.
* `published_object_metadata` - A list of objects that are published or unpublished in this patch.
	* `action` - The patch action indicating if object was created, updated, or deleted.
	* `identifier` - Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `key` - The key of the object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `name_path` - The fully qualified path of the published object, which would include its project and folder.
	* `object_version` - The object version.
	* `type` - The type of the object in patch.
* `source_application_info` - The information about the application.
	* `application_key` - The source application key to use when creating the application.
	* `application_version` - The source application version of the application.
	* `last_patch_key` - The last patch key for the application.
	* `workspace_id` - The OCID of the workspace containing the application. This allows cross workspace deployment to publish an application from a different workspace into the current workspace specified in this operation.
* `state` - The current state of the workspace.
* `time_created` - The date and time the application was created, in the timestamp format defined by RFC3339. 
* `time_patched` - The date and time the application was patched, in the timestamp format defined by RFC3339. 
* `time_updated` - The date and time the application was updated, in the timestamp format defined by RFC3339. example: 2019-08-25T21:10:29.41Z 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Workspace Application
	* `update` - (Defaults to 20 minutes), when updating the Workspace Application
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Application


## Import

WorkspaceApplications can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_application.test_workspace_application "workspaces/{workspaceId}/applications/{applicationKey}" 
```

