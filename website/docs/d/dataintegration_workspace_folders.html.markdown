---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_folders"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_folders"
description: |-
  Provides the list of Workspace Folders in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_folders
This data source provides the list of Workspace Folders in Oracle Cloud Infrastructure Data Integration service.

Retrieves a list of folders in a project and provides options to filter the list.


## Example Usage

```hcl
data "oci_dataintegration_workspace_folders" "test_workspace_folders" {
	#Required
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	aggregator_key = var.workspace_folder_aggregator_key
	fields = var.workspace_folder_fields
	identifier = var.workspace_folder_identifier
	name = var.workspace_folder_name
	name_contains = var.workspace_folder_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `aggregator_key` - (Optional) Used to filter by the project or the folder object. 
* `fields` - (Optional) Specifies the fields to get for an object.
* `identifier` - (Optional) Used to filter by the identifier of the object.
* `name` - (Optional) Used to filter by the name of the object.
* `name_contains` - (Optional) This parameter can be used to filter objects by the names that match partially or fully with the given value.
* `workspace_id` - (Required) The workspace ID.


## Attributes Reference

The following attributes are exported:

* `folder_summary_collection` - The list of folder_summary_collection.

### WorkspaceFolder Reference

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

