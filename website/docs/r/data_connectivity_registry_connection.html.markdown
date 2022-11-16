---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registry_connection"
sidebar_current: "docs-oci-resource-data_connectivity-registry_connection"
description: |-
  Provides the Registry Connection resource in Oracle Cloud Infrastructure Data Connectivity service
---

# oci_data_connectivity_registry_connection
This resource provides the Registry Connection resource in Oracle Cloud Infrastructure Data Connectivity service.

Creates a connection under an existing data asset.

## Example Usage

```hcl
resource "oci_data_connectivity_registry_connection" "test_registry_connection" {
	#Required
	identifier = var.registry_connection_identifier
	name = var.registry_connection_name
	properties = var.registry_connection_properties
	registry_id = oci_data_connectivity_registry.test_registry.id
	type = var.registry_connection_type

	#Optional
	connection_properties {

		#Optional
		name = var.registry_connection_connection_properties_name
		value = var.registry_connection_connection_properties_value
	}
	description = var.registry_connection_description
	is_default = var.registry_connection_is_default
	key = var.registry_connection_key
	metadata {

		#Optional
		aggregator {

			#Optional
			description = var.registry_connection_metadata_aggregator_description
			identifier = var.registry_connection_metadata_aggregator_identifier
			key = var.registry_connection_metadata_aggregator_key
			name = var.registry_connection_metadata_aggregator_name
			type = var.registry_connection_metadata_aggregator_type
		}
		aggregator_key = var.registry_connection_metadata_aggregator_key
		created_by = var.registry_connection_metadata_created_by
		created_by_name = var.registry_connection_metadata_created_by_name
		identifier_path = var.registry_connection_metadata_identifier_path
		info_fields = var.registry_connection_metadata_info_fields
		is_favorite = var.registry_connection_metadata_is_favorite
		labels = var.registry_connection_metadata_labels
		registry_version = var.registry_connection_metadata_registry_version
		time_created = var.registry_connection_metadata_time_created
		time_updated = var.registry_connection_metadata_time_updated
		updated_by = var.registry_connection_metadata_updated_by
		updated_by_name = var.registry_connection_metadata_updated_by_name
	}
	model_type = var.registry_connection_model_type
	model_version = var.registry_connection_model_version
	object_status = var.registry_connection_object_status
	object_version = var.registry_connection_object_version
	primary_schema {
		#Required
		identifier = var.registry_connection_primary_schema_identifier
		key = var.registry_connection_primary_schema_key
		model_type = var.registry_connection_primary_schema_model_type
		name = var.registry_connection_primary_schema_name

		#Optional
		default_connection = var.registry_connection_primary_schema_default_connection
		description = var.registry_connection_primary_schema_description
		external_key = var.registry_connection_primary_schema_external_key
		is_has_containers = var.registry_connection_primary_schema_is_has_containers
		metadata {

			#Optional
			aggregator {

				#Optional
				description = var.registry_connection_primary_schema_metadata_aggregator_description
				identifier = var.registry_connection_primary_schema_metadata_aggregator_identifier
				key = var.registry_connection_primary_schema_metadata_aggregator_key
				name = var.registry_connection_primary_schema_metadata_aggregator_name
				type = var.registry_connection_primary_schema_metadata_aggregator_type
			}
			aggregator_key = var.registry_connection_primary_schema_metadata_aggregator_key
			created_by = var.registry_connection_primary_schema_metadata_created_by
			created_by_name = var.registry_connection_primary_schema_metadata_created_by_name
			identifier_path = var.registry_connection_primary_schema_metadata_identifier_path
			info_fields = var.registry_connection_primary_schema_metadata_info_fields
			is_favorite = var.registry_connection_primary_schema_metadata_is_favorite
			labels = var.registry_connection_primary_schema_metadata_labels
			registry_version = var.registry_connection_primary_schema_metadata_registry_version
			time_created = var.registry_connection_primary_schema_metadata_time_created
			time_updated = var.registry_connection_primary_schema_metadata_time_updated
			updated_by = var.registry_connection_primary_schema_metadata_updated_by
			updated_by_name = var.registry_connection_primary_schema_metadata_updated_by_name
		}
		model_version = var.registry_connection_primary_schema_model_version
		object_status = var.registry_connection_primary_schema_object_status
		object_version = var.registry_connection_primary_schema_object_version
		parent_ref {

			#Optional
			parent = var.registry_connection_primary_schema_parent_ref_parent
		}
		resource_name = var.registry_connection_primary_schema_resource_name
	}
	registry_metadata {

		#Optional
		aggregator_key = var.registry_connection_registry_metadata_aggregator_key
		created_by_user_id = oci_identity_user.test_user.id
		created_by_user_name = oci_identity_user.test_user.name
		is_favorite = var.registry_connection_registry_metadata_is_favorite
		key = var.registry_connection_registry_metadata_key
		labels = var.registry_connection_registry_metadata_labels
		registry_version = var.registry_connection_registry_metadata_registry_version
		time_created = var.registry_connection_registry_metadata_time_created
		time_updated = var.registry_connection_registry_metadata_time_updated
		updated_by_user_id = oci_identity_user.test_user.id
		updated_by_user_name = oci_identity_user.test_user.name
	}
}
```

## Argument Reference

The following arguments are supported:

* `connection_properties` - (Optional) (Updatable) The properties of the connection.
	* `name` - (Optional) (Updatable) Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `value` - (Optional) (Updatable) The value for the connection name property.
* `description` - (Optional) (Updatable) User-defined description of the connection.
* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
* `is_default` - (Optional) (Updatable) The default property of the connection.
* `key` - (Optional) Generated key that can be used in API calls to identify the connection. In scenarios where reference to the connection is required, a value can be passed in create.
* `metadata` - (Optional) (Updatable) A summary type containing information about the object including its key, name, the time that it was created or updated, and the user who created or updated it.
	* `aggregator` - (Optional) (Updatable) A summary type containing information about the object's aggregator including its type, key, name, and description.
		* `description` - (Optional) (Updatable) The description of the aggregator.
		* `identifier` - (Optional) (Updatable) The identifier of the aggregator.
		* `key` - (Optional) (Updatable) The key of the aggregator object.
		* `name` - (Optional) (Updatable) The name of the aggregator.
		* `type` - (Optional) (Updatable) The type of the aggregator.
	* `aggregator_key` - (Optional) (Updatable) The owning object key for this object.
	* `created_by` - (Optional) (Updatable) The user that created the object.
	* `created_by_name` - (Optional) (Updatable) The user that created the object.
	* `identifier_path` - (Optional) (Updatable) The full path to identify the object.
	* `info_fields` - (Optional) (Updatable) Information property fields.
	* `is_favorite` - (Optional) (Updatable) Specifies whether this object is a favorite.
	* `labels` - (Optional) (Updatable) Labels are keywords or tags that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
	* `registry_version` - (Optional) (Updatable) The registry version of the object.
	* `time_created` - (Optional) (Updatable) The date and time that the object was created.
	* `time_updated` - (Optional) (Updatable) The date and time that the object was updated.
	* `updated_by` - (Optional) (Updatable) The user that updated the object.
	* `updated_by_name` - (Optional) (Updatable) The user that updated the object.
* `model_type` - (Optional) (Updatable) The type of the object.
* `model_version` - (Optional) (Updatable) The model version of an object.
* `name` - (Required) (Updatable) Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - (Optional) (Updatable) The version of the object that is used to track changes in the object instance.
* `primary_schema` - (Optional) (Updatable) The schema object.
	* `default_connection` - (Optional) (Updatable) The default connection key.
	* `description` - (Optional) (Updatable) User-defined description for the schema.
	* `external_key` - (Optional) (Updatable) The external key of the object.
	* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	* `is_has_containers` - (Optional) (Updatable) Specifies whether the schema has containers.
	* `key` - (Required) (Updatable) The object key.
	* `metadata` - (Optional) (Updatable) A summary type containing information about the object including its key, name, the time that it was created or updated, and the user who created or updated it.
		* `aggregator` - (Optional) (Updatable) A summary type containing information about the object's aggregator including its type, key, name, and description.
			* `description` - (Optional) (Updatable) The description of the aggregator.
			* `identifier` - (Optional) (Updatable) The identifier of the aggregator.
			* `key` - (Optional) (Updatable) The key of the aggregator object.
			* `name` - (Optional) (Updatable) The name of the aggregator.
			* `type` - (Optional) (Updatable) The type of the aggregator.
		* `aggregator_key` - (Optional) (Updatable) The owning object key for this object.
		* `created_by` - (Optional) (Updatable) The user that created the object.
		* `created_by_name` - (Optional) (Updatable) The user that created the object.
		* `identifier_path` - (Optional) (Updatable) The full path to identify the object.
		* `info_fields` - (Optional) (Updatable) Information property fields.
		* `is_favorite` - (Optional) (Updatable) Specifies whether this object is a favorite.
		* `labels` - (Optional) (Updatable) Labels are keywords or tags that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
		* `registry_version` - (Optional) (Updatable) The registry version of the object.
		* `time_created` - (Optional) (Updatable) The date and time that the object was created.
		* `time_updated` - (Optional) (Updatable) The date and time that the object was updated.
		* `updated_by` - (Optional) (Updatable) The user that updated the object.
		* `updated_by_name` - (Optional) (Updatable) The user that updated the object.
	* `model_type` - (Required) (Updatable) The object type.
	* `model_version` - (Optional) (Updatable) The model version of the object.
	* `name` - (Required) (Updatable) Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - (Optional) (Updatable) The version of the object that is used to track changes in the object instance.
	* `parent_ref` - (Optional) (Updatable) A reference to the parent object.
		* `parent` - (Optional) (Updatable) Key of the parent object.
	* `resource_name` - (Optional) (Updatable) A resource name can have letters, numbers, and special characters. The value is editable and is restricted to 4000 characters.
* `properties` - (Required) (Updatable) All the properties of the connection in a key-value map format.
* `registry_id` - (Required) The registry OCID.
* `registry_metadata` - (Optional) (Updatable) Information about the object and its parent.
	* `aggregator_key` - (Optional) (Updatable) The owning object's key for this object.
	* `created_by_user_id` - (Optional) (Updatable) The ID of the user who created the object.
	* `created_by_user_name` - (Optional) (Updatable) The name of the user who created the object.
	* `is_favorite` - (Optional) (Updatable) Specifies whether the object is a favorite.
	* `key` - (Optional) (Updatable) The identifying key for the object.
	* `labels` - (Optional) (Updatable) Labels are keywords or labels that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
	* `registry_version` - (Optional) (Updatable) The registry version.
	* `time_created` - (Optional) (Updatable) The date and time that the object was created.
	* `time_updated` - (Optional) (Updatable) The date and time that the object was updated.
	* `updated_by_user_id` - (Optional) (Updatable) The ID of the user who updated the object.
	* `updated_by_user_name` - (Optional) (Updatable) The name of the user who updated the object.
* `type` - (Required) (Updatable) Specific Connection Type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `connection_properties` - The properties of the connection.
	* `name` - Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `value` - The value for the connection name property.
* `description` - User-defined description for the connection.
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
* `is_default` - The default property of the connection.
* `key` - Generated key that can be used in API calls to identify the connection. In scenarios where reference to the connection is required, a value can be passed in create.
* `metadata` - A summary type containing information about the object including its key, name, the time that it was created or updated, and the user who created or updated it.
	* `aggregator` - A summary type containing information about the object's aggregator including its type, key, name, and description.
		* `description` - The description of the aggregator.
		* `identifier` - The identifier of the aggregator.
		* `key` - The key of the aggregator object.
		* `name` - The name of the aggregator.
		* `type` - The type of the aggregator.
	* `aggregator_key` - The owning object key for this object.
	* `created_by` - The user that created the object.
	* `created_by_name` - The user that created the object.
	* `identifier_path` - The full path to identify the object.
	* `info_fields` - Information property fields.
	* `is_favorite` - Specifies whether this object is a favorite.
	* `labels` - Labels are keywords or tags that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
	* `registry_version` - The registry version of the object.
	* `time_created` - The date and time that the object was created.
	* `time_updated` - The date and time that the object was updated.
	* `updated_by` - The user that updated the object.
	* `updated_by_name` - The user that updated the object.
* `model_type` - The type of the object.
* `model_version` - The model version of an object.
* `name` - Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - The version of the object that is used to track changes in the object instance.
* `primary_schema` - The schema object.
	* `default_connection` - The default connection key.
	* `description` - User-defined description for the schema.
	* `external_key` - The external key of the object.
	* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	* `is_has_containers` - Specifies whether the schema has containers.
	* `key` - The object key.
	* `metadata` - A summary type containing information about the object including its key, name, the time that it was created or updated, and the user who created or updated it.
		* `aggregator` - A summary type containing information about the object's aggregator including its type, key, name, and description.
			* `description` - The description of the aggregator.
			* `identifier` - The identifier of the aggregator.
			* `key` - The key of the aggregator object.
			* `name` - The name of the aggregator.
			* `type` - The type of the aggregator.
		* `aggregator_key` - The owning object key for this object.
		* `created_by` - The user that created the object.
		* `created_by_name` - The user that created the object.
		* `identifier_path` - The full path to identify the object.
		* `info_fields` - Information property fields.
		* `is_favorite` - Specifies whether this object is a favorite.
		* `labels` - Labels are keywords or tags that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
		* `registry_version` - The registry version of the object.
		* `time_created` - The date and time that the object was created.
		* `time_updated` - The date and time that the object was updated.
		* `updated_by` - The user that updated the object.
		* `updated_by_name` - The user that updated the object.
	* `model_type` - The object type.
	* `model_version` - The model version of the object.
	* `name` - Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - The version of the object that is used to track changes in the object instance.
	* `parent_ref` - A reference to the parent object.
		* `parent` - Key of the parent object.
	* `resource_name` - A resource name can have letters, numbers, and special characters. The value is editable and is restricted to 4000 characters.
* `properties` - All the properties of the connection in a key-value map format.
* `registry_metadata` - Information about the object and its parent.
	* `aggregator_key` - The owning object's key for this object.
	* `created_by_user_id` - The ID of the user who created the object.
	* `created_by_user_name` - The name of the user who created the object.
	* `is_favorite` - Specifies whether the object is a favorite.
	* `key` - The identifying key for the object.
	* `labels` - Labels are keywords or labels that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
	* `registry_version` - The registry version.
	* `time_created` - The date and time that the object was created.
	* `time_updated` - The date and time that the object was updated.
	* `updated_by_user_id` - The ID of the user who updated the object.
	* `updated_by_user_name` - The name of the user who updated the object.
* `type` - Specific Connection Type

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Registry Connection
	* `update` - (Defaults to 20 minutes), when updating the Registry Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Registry Connection


## Import

RegistryConnections can be imported using the `id`, e.g.

```
$ terraform import oci_data_connectivity_registry_connection.test_registry_connection "registries/{registryId}/connections/{connectionKey}" 
```

