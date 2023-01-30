---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_project"
sidebar_current: "docs-oci-resource-dataintegration-workspace_project"
description: |-
  Provides the Workspace Project resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_project
This resource provides the Workspace Project resource in Oracle Cloud Infrastructure Data Integration service.

Creates a project. Projects are organizational constructs within a workspace that you use to organize your design-time resources, such as tasks or data flows. Projects can be organized into folders.


## Example Usage

```hcl
resource "oci_dataintegration_workspace_project" "test_workspace_project" {
	#Required
	identifier = var.workspace_project_identifier
	name = var.workspace_project_name
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	description = var.workspace_project_description
	key = var.workspace_project_key
	model_version = var.workspace_project_model_version
	object_status = var.workspace_project_object_status
	registry_metadata {

		#Optional
		aggregator_key = var.workspace_project_registry_metadata_aggregator_key
		is_favorite = var.workspace_project_registry_metadata_is_favorite
		key = var.workspace_project_registry_metadata_key
		labels = var.workspace_project_registry_metadata_labels
		registry_version = var.workspace_project_registry_metadata_registry_version
	}
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) (Updatable) A user defined description for the project.
* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `key` - (Optional) (Updatable) Generated key that can be used in API calls to identify project.
* `model_version` - (Optional) (Updatable) The model version of an object.
* `name` - (Required) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `registry_metadata` - (Optional) (Updatable) Information about the object and its parent.
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

* `description` - A user defined description for the project.
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `key` - Generated key that can be used in API calls to identify project.
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
	* `create` - (Defaults to 20 minutes), when creating the Workspace Project
	* `update` - (Defaults to 20 minutes), when updating the Workspace Project
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Project


## Import

WorkspaceProjects can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_project.test_workspace_project "workspaces/{workspaceId}/projects/{projectKey}" 
```

