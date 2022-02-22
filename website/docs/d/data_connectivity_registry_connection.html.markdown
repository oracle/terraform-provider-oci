---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registry_connection"
sidebar_current: "docs-oci-datasource-data_connectivity-registry_connection"
description: |-
  Provides details about a specific Registry Connection in Oracle Cloud Infrastructure Data Connectivity service
---

# Data Source: oci_data_connectivity_registry_connection
This data source provides details about a specific Registry Connection resource in Oracle Cloud Infrastructure Data Connectivity service.

Retrieves the connection details using the specified identifier.

## Example Usage

```hcl
data "oci_data_connectivity_registry_connection" "test_registry_connection" {
	#Required
	connection_key = var.registry_connection_connection_key
	registry_id = oci_data_connectivity_registry.test_registry.id
}
```

## Argument Reference

The following arguments are supported:

* `connection_key` - (Required) The connection key.
* `registry_id` - (Required) The registry Ocid.


## Attributes Reference

The following attributes are exported:

* `connection_properties` - The properties for the connection.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `value` - The value for the connection name property.
* `description` - User-defined description for the connection.
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `is_default` - The default property for the connection.
* `key` - Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
* `metadata` - A summary type containing information about the object including its key, name and when/who created/updated it.
	* `aggregator` - A summary type containing information about the object's aggregator including its type, key, name and description.
		* `description` - The description of the aggregator.
		* `identifier` - The identifier of the aggregator.
		* `key` - The key of the aggregator object.
		* `name` - The name of the aggregator.
		* `type` - The type of the aggregator.
	* `aggregator_key` - The owning object key for this object.
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
* `primary_schema` - The schema object.
	* `default_connection` - The default connection key.
	* `description` - User-defined description for the schema.
	* `external_key` - The external key for the object.
	* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `is_has_containers` - Specifies whether the schema has containers.
	* `key` - The object key.
	* `metadata` - A summary type containing information about the object including its key, name and when/who created/updated it.
		* `aggregator` - A summary type containing information about the object's aggregator including its type, key, name and description.
			* `description` - The description of the aggregator.
			* `identifier` - The identifier of the aggregator.
			* `key` - The key of the aggregator object.
			* `name` - The name of the aggregator.
			* `type` - The type of the aggregator.
		* `aggregator_key` - The owning object key for this object.
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
	* `model_type` - The object's type.
	* `model_version` - The object's model version.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - The version of the object that is used to track changes in the object instance.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
	* `resource_name` - A resource name can have letters, numbers, and special characters. The value is editable and is restricted to 4000 characters.
* `properties` - All the properties for the connection in a key-value map format.
* `registry_metadata` - Information about the object and its parent.
	* `aggregator_key` - The owning object's key for this object.
	* `created_by_user_id` - The id of the user who created the object.
	* `created_by_user_name` - The name of the user who created the object.
	* `is_favorite` - Specifies whether this object is a favorite or not.
	* `key` - The identifying key for the object.
	* `labels` - Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	* `registry_version` - The registry version.
	* `time_created` - The date and time that the object was created.
	* `time_updated` - The date and time that the object was updated.
	* `updated_by_user_id` - The id of the user who updated the object.
	* `updated_by_user_name` - The name of the user who updated the object.
* `type` - Specific Connection Type

