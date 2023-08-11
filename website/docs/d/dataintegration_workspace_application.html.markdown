---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_application"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_application"
description: |-
  Provides details about a specific Workspace Application in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_application
This data source provides details about a specific Workspace Application resource in Oracle Cloud Infrastructure Data Integration service.

Retrieves an application using the specified identifier.

## Example Usage

```hcl
data "oci_dataintegration_workspace_application" "test_workspace_application" {
	#Required
	application_key = var.workspace_application_application_key
	workspace_id = oci_dataintegration_workspace.test_workspace.id
}
```

## Argument Reference

The following arguments are supported:

* `application_key` - (Required) The application key.
* `workspace_id` - (Required) The workspace ID.


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

