---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_folder"
sidebar_current: "docs-oci-resource-dataintegration-workspace_folder"
description: |-
  Provides the Workspace Folder resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_folder
This resource provides the Workspace Folder resource in Oracle Cloud Infrastructure Data Integration service.

Creates a folder in a project or in another folder, limited to two levels of folders. |
Folders are used to organize your design-time resources, such as tasks or data flows.


## Example Usage

```hcl
resource "oci_dataintegration_workspace_folder" "test_workspace_folder" {
	#Required
	identifier = var.workspace_folder_identifier
	name = var.workspace_folder_name
	registry_metadata {

		#Optional
		aggregator_key = var.workspace_folder_registry_metadata_aggregator_key
		is_favorite = var.workspace_folder_registry_metadata_is_favorite
		key = var.workspace_folder_registry_metadata_key
		labels = var.workspace_folder_registry_metadata_labels
		registry_version = var.workspace_folder_registry_metadata_registry_version
	}
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	category_name = oci_marketplace_category.test_category.name
	description = var.workspace_folder_description
	key = var.workspace_folder_key
	model_version = var.workspace_folder_model_version
	object_status = var.workspace_folder_object_status
}
```

## Argument Reference

The following arguments are supported:

* `category_name` - (Optional) (Updatable) The category name.
* `description` - (Optional) (Updatable) A user defined description for the folder.
* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `key` - (Optional) (Updatable) Currently not used on folder creation. Reserved for future.
* `model_version` - (Optional) (Updatable) The model version of an object.
* `name` - (Required) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `registry_metadata` - (Required) (Updatable) Information about the object and its parent.
	* `aggregator_key` - (Optional) (Updatable) The owning object's key for this object.
	* `is_favorite` - (Optional) (Updatable) Specifies whether this object is a favorite or not.
	* `key` - (Optional) (Updatable) The identifying key for the object.
	* `labels` - (Optional) (Updatable) Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	* `registry_version` - (Optional) (Updatable) The registry version.
* `workspace_id` - (Required) The workspace ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `category_name` - The category name.
* `description` - A user defined description for the folder.
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `key` - Generated key that can be used in API calls to identify folder.
* `key_map` - A key map. If provided, the key is replaced with generated key. This structure provides mapping between user provided key and generated key.
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
* `model_type` - The type of the object.
* `model_version` - The model version of an object.
* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - The version of the object that is used to track changes in the object instance.
* `parent_ref` - A reference to the object's parent.
	* `parent` - Key of the parent object.
	* `root_doc_id` - Key of the root document object.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Workspace Folder
	* `update` - (Defaults to 20 minutes), when updating the Workspace Folder
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Folder


## Import

WorkspaceFolders can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_folder.test_workspace_folder "workspaces/{workspaceId}/folders/{folderKey}" 
```

