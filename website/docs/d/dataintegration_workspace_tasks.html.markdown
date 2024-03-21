---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_tasks"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_tasks"
description: |-
  Provides the list of Workspace Tasks in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_tasks
This data source provides the list of Workspace Tasks in Oracle Cloud Infrastructure Data Integration service.

Retrieves a list of all tasks in a specified project or folder.


## Example Usage

```hcl
data "oci_dataintegration_workspace_tasks" "test_workspace_tasks" {
	#Required
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	fields = var.workspace_task_fields
	folder_id = oci_dataintegration_folder.test_folder.id
	identifier = var.workspace_task_identifier
	key = var.workspace_task_key
	name = var.workspace_task_name
	type = var.workspace_task_type
}
```

## Argument Reference

The following arguments are supported:

* `fields` - (Optional) Specifies the fields to get for an object.
* `folder_id` - (Optional) Unique key of the folder.
* `identifier` - (Optional) Used to filter by the identifier of the object.
* `key` - (Optional) Used to filter by the key of the object.
* `name` - (Optional) Used to filter by the name of the object.
* `type` - (Optional) Used to filter by the object type of the object. It can be suffixed with an optional filter operator InSubtree. If this operator is not specified, then exact match is considered. <br><br><B>Examples:</B><br> <ul> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=false</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=true</B> returns all objects of type data loader task</li> </ul>
* `workspace_id` - (Required) The workspace ID.


## Attributes Reference

The following attributes are exported:

* `task_summary_collection` - The list of task_summary_collection.

### WorkspaceTask Reference

The following attributes are exported:

* `api_call_mode` - The REST invocation pattern to use. ASYNC_OCI_WORKREQUEST is being deprecated as well as cancelEndpoint/MethodType.
* `auth_config` - Authentication configuration for Generic REST invocation.
	* `key` - Generated key that can be used in API calls to identify this object.
	* `model_type` - The specific authentication configuration to be used for Generic REST invocation.
	* `model_version` - The model version of an object.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `resource_principal_source` - The Oracle Cloud Infrastructure resource type that will supply the authentication token
* `auth_details` - Authentication type to be used for Generic REST invocation. This is deprecated.
	* `key` - Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
	* `model_type` - The authentication mode to be used for Generic REST invocation.
	* `model_version` - The model version of an object.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
* `cancel_endpoint` - An expression node.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `expr_string` - The expression string for the object.
	* `key` - The object key.
	* `model_type` - The object type.
	* `model_version` - The object's model version.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
* `cancel_method_type` - The REST method to use for canceling the original request.
* `cancel_rest_call_config` - The REST API configuration for cancelling the task.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `method_type` - The REST method to use.
	* `request_headers` - The headers for the REST call.
* `conditional_composite_field_map` - A conditional composite field map.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `description` - Detailed description for the object.
	* `field_map_scope` - An array of projection rules.
		* `config_values` - Configuration values can be string, objects, or parameters.
			* `config_param_values` - The configuration parameter values.
				* `int_value` - An integer value of the parameter.
				* `object_value` - An object value of the parameter.
				* `parameter_value` - Reference to the parameter by its key.
				* `ref_value` - The root object reference value.
				* `root_object_value` - The root object value, used in custom parameters.
				* `string_value` - A string value of the parameter.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
		* `description` - A user defined description for the object.
		* `from_name` - The attribute name that needs to be renamed.
		* `is_cascade` - Specifies whether to cascade or not.
		* `is_case_sensitive` - Specifies if the rule is case sensitive.
		* `is_java_regex_syntax` - Specifies whether the rule uses a java regex syntax.
		* `is_skip_remaining_rules_on_match` - Specifies whether to skip remaining rules when a match is found.
		* `key` - The key of the object.
		* `matching_strategy` - The pattern matching strategy.
		* `model_type` - The type of the project rule.
		* `model_version` - The model version of an object.
		* `name` - Name of the group.
		* `names` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
		* `pattern` - The rule pattern.
		* `rule_type` - The rule type.
		* `scope` - Reference to a typed object. This can be either a key value to an object within the document, a shall referenced to a `TypedObject`, or a full `TypedObject` definition.
		* `to_name` - The new attribute name.
		* `types` - An array of types.
	* `field_maps` - An array of field maps.
	* `key` - The object key.
	* `model_type` - The model type for the field map.
	* `model_version` - The object's model version.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
* `config_provider_delegate` - The information about the configuration provider.
* `data_flow` - The data flow type contains the audit summary information and the definition of the data flow.
	* `description` - Detailed description for the object.
	* `flow_config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `key` - Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
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
	* `model_type` - The type of the object.
	* `model_version` - The model version of an object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `nodes` - An array of nodes.
		* `config_provider_delegate` - The information about the configuration provider.
		* `description` - Detailed description for the object.
		* `input_links` - An array of input links.
			* `description` - Detailed description for the object.
			* `field_map` - A field map is a way to map a source row shape to a target row shape that may be different.
			* `from_link` - The from link reference.
			* `key` - The key of the object.
			* `model_type` - The model type of the object.
			* `model_version` - The model version of an object.
			* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
			* `port` - Key of FlowPort reference
		* `key` - The key of the object.
		* `model_type` - The type of the object.
		* `model_version` - The model version of an object.
		* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `operator` - An operator defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
		* `output_links` - An array of output links.
			* `description` - Detailed description for the object.
			* `key` - The key of the object.
			* `model_type` - The model type of the object.
			* `model_version` - The model version of an object.
			* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
			* `port` - Key of FlowPort reference
			* `to_links` - The links from this output link to connect to other links in flow.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
		* `ui_properties` - The UI properties of the object.
			* `coordinate_x` - The X coordinate of the object.
			* `coordinate_y` - The Y coordinate of the object.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - The version of the object that is used to track changes in the object instance.
	* `parameters` - An array of parameters.
		* `config_values` - Configuration values can be string, objects, or parameters.
			* `config_param_values` - The configuration parameter values.
				* `int_value` - An integer value of the parameter.
				* `object_value` - An object value of the parameter.
				* `parameter_value` - Reference to the parameter by its key.
				* `ref_value` - The root object reference value.
				* `root_object_value` - The root object value, used in custom parameters.
				* `string_value` - A string value of the parameter.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
		* `default_value` - The default value of the parameter.
		* `description` - Detailed description for the object.
		* `is_input` - Specifies whether the parameter is input value.
		* `is_output` - Specifies whether the parameter is output value.
		* `key` - The key of the object.
		* `model_type` - The type of the types object.
		* `model_version` - The model version of an object.
		* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `output_aggregation_type` - The output aggregation type.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
		* `root_object_default_value` - The default value of the parameter which can be an object in DIS, such as a data entity.
		* `type` - This can either be a string value referencing the type or a BaseType object.
		* `type_name` - The type of value the parameter was created for.
		* `used_for` - The param name for which parameter is created for for eg. driver Shape, Operation etc.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `target_field_map_summary` - A hash map that maps TypedObject keys to a field map that maps to the typed object as a target, for java sdk.
		* `field_map` - A field map is a way to map a source row shape to a target row shape that may be different.
	* `typed_object_map` - A hash map that maps TypedObject keys to the object itself, for java sdk.
		* `typed_object` - The `TypedObject` class is a base class for any model object that has a type.
* `dataflow_application` - Minimum information required to recognize a Dataflow Application object.
	* `application_id` - The application id for which Oracle Cloud Infrastructure data flow task is to be created.
	* `compartment_id` - The compartmentId id under which Oracle Cloud Infrastructure dataflow application lies.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
* `description` - Detailed description for the object.
* `endpoint` - An expression node.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `expr_string` - The expression string for the object.
	* `key` - The object key.
	* `model_type` - The object type.
	* `model_version` - The object's model version.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
* `execute_rest_call_config` - The REST API configuration for execution.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `method_type` - The REST method to use.
	* `request_headers` - The headers for the REST call.
* `headers` - The headers for the REST call. This property is deprecated, use ExecuteRestCallConfig's headers property instead.
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `input_ports` - An array of input ports.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `description` - Detailed description for the object.
	* `fields` - An array of fields.
	* `key` - The key of the object.
	* `model_type` - The type of the types object.
	* `model_version` - The model version of an object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `port_type` - The port details for the data asset.Type.
* `is_single_load` - Defines whether Data Loader task is used for single load or multiple
* `json_data` - JSON data for payload body. This property is deprecated, use ExecuteRestCallConfig's payload config param instead.
* `key` - Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
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
* `method_type` - The REST method to use. This property is deprecated, use ExecuteRestCallConfig's methodType property instead.
* `model_type` - The type of the task.
* `model_version` - The object's model version.
* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - The version of the object that is used to track changes in the object instance.
* `op_config_values` - Configuration values can be string, objects, or parameters.
	* `config_param_values` - The configuration parameter values.
		* `int_value` - An integer value of the parameter.
		* `object_value` - An object value of the parameter.
		* `parameter_value` - Reference to the parameter by its key.
		* `ref_value` - The root object reference value.
		* `root_object_value` - The root object value, used in custom parameters.
		* `string_value` - A string value of the parameter.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
* `operation` - Describes the shape of the execution result
* `output_ports` - An array of output ports.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `description` - Detailed description for the object.
	* `fields` - An array of fields.
	* `key` - The key of the object.
	* `model_type` - The type of the types object.
	* `model_version` - The model version of an object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `port_type` - The port details for the data asset.Type.
* `parallel_load_limit` - Defines the number of entities being loaded in parallel at a time for a Data Loader task
* `parameters` - An array of parameters.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `default_value` - The default value of the parameter.
	* `description` - Detailed description for the object.
	* `is_input` - Specifies whether the parameter is input value.
	* `is_output` - Specifies whether the parameter is output value.
	* `key` - The key of the object.
	* `model_type` - The type of the types object.
	* `model_version` - The model version of an object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `output_aggregation_type` - The output aggregation type.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `root_object_default_value` - The default value of the parameter which can be an object in DIS, such as a data entity.
	* `type` - This can either be a string value referencing the type or a BaseType object.
	* `type_name` - The type of value the parameter was created for.
	* `used_for` - The param name for which parameter is created for for eg. driver Shape, Operation etc.
* `parent_ref` - A reference to the object's parent.
	* `parent` - Key of the parent object.
	* `root_doc_id` - Key of the root document object.
* `pipeline` - A pipeline is a logical grouping of tasks that together perform a higher level operation. For example, a pipeline could contain a set of tasks that load and clean data, then execute a dataflow to analyze the data. The pipeline allows you to manage the activities as a unit instead of individually. Users can also schedule the pipeline instead of the tasks independently.
	* `description` - Detailed description for the object.
	* `flow_config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `key` - Generated key that can be used in API calls to identify pipeline. On scenarios where reference to the pipeline is needed, a value can be passed in create.
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
	* `model_version` - This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `nodes` - A list of nodes attached to the pipeline.
		* `config_provider_delegate` - The information about the configuration provider.
		* `description` - Detailed description for the object.
		* `input_links` - An array of input links.
			* `description` - Detailed description for the object.
			* `field_map` - A field map is a way to map a source row shape to a target row shape that may be different.
			* `from_link` - The from link reference.
			* `key` - The key of the object.
			* `model_type` - The model type of the object.
			* `model_version` - The model version of an object.
			* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
			* `port` - Key of FlowPort reference
		* `key` - The key of the object.
		* `model_type` - The type of the object.
		* `model_version` - The model version of an object.
		* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `operator` - An operator defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
		* `output_links` - An array of output links.
			* `description` - Detailed description for the object.
			* `key` - The key of the object.
			* `model_type` - The model type of the object.
			* `model_version` - The model version of an object.
			* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
			* `port` - Key of FlowPort reference
			* `to_links` - The links from this output link to connect to other links in flow.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
		* `ui_properties` - The UI properties of the object.
			* `coordinate_x` - The X coordinate of the object.
			* `coordinate_y` - The Y coordinate of the object.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	* `parameters` - A list of parameters for the pipeline, this allows certain aspects of the pipeline to be configured when the pipeline is executed.
		* `config_values` - Configuration values can be string, objects, or parameters.
			* `config_param_values` - The configuration parameter values.
				* `int_value` - An integer value of the parameter.
				* `object_value` - An object value of the parameter.
				* `parameter_value` - Reference to the parameter by its key.
				* `ref_value` - The root object reference value.
				* `root_object_value` - The root object value, used in custom parameters.
				* `string_value` - A string value of the parameter.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
		* `default_value` - The default value of the parameter.
		* `description` - Detailed description for the object.
		* `is_input` - Specifies whether the parameter is input value.
		* `is_output` - Specifies whether the parameter is output value.
		* `key` - The key of the object.
		* `model_type` - The type of the types object.
		* `model_version` - The model version of an object.
		* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `output_aggregation_type` - The output aggregation type.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
		* `root_object_default_value` - The default value of the parameter which can be an object in DIS, such as a data entity.
		* `type` - This can either be a string value referencing the type or a BaseType object.
		* `type_name` - The type of value the parameter was created for.
		* `used_for` - The param name for which parameter is created for for eg. driver Shape, Operation etc.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `variables` - The list of variables required in pipeline, variables can be used to store values that can be used as inputs to tasks in the pipeline.
		* `config_values` - Configuration values can be string, objects, or parameters.
			* `config_param_values` - The configuration parameter values.
				* `int_value` - An integer value of the parameter.
				* `object_value` - An object value of the parameter.
				* `parameter_value` - Reference to the parameter by its key.
				* `ref_value` - The root object reference value.
				* `root_object_value` - The root object value, used in custom parameters.
				* `string_value` - A string value of the parameter.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
		* `default_value` - A default value for the vairable.
		* `description` - Detailed description for the object.
		* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
		* `key` - Generated key that can be used in API calls to identify variable. On scenarios where reference to the variable is needed, a value can be passed in create.
		* `model_type` - The type of the object.
		* `model_version` - This is a version number that is used by the service to upgrade objects if needed through releases of the service.
		* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `object_version` - This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
		* `root_object_default_value` - A base class for all model types, including First Class and its contained objects.
			* `key` - The key of the object.
			* `model_type` - The type of the object.
			* `model_version` - The model version of an object.
			* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - A reference to the object's parent.
				* `parent` - Key of the parent object.
				* `root_doc_id` - Key of the root document object.
		* `type` - Base type for the type system.
* `poll_rest_call_config` - The REST API configuration for polling.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `method_type` - The REST method to use.
	* `request_headers` - The headers for the REST call.
* `registry_metadata` - Information about the object and its parent.
	* `aggregator_key` - The owning object's key for this object.
	* `is_favorite` - Specifies whether this object is a favorite or not.
	* `key` - The identifying key for the object.
	* `labels` - Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	* `registry_version` - The registry version.
* `script` - The script object.
	* `key` - The key of the object.
	* `model_type` - The type of the object.
	* `model_version` - The model version of an object.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
* `sql_script_type` - Indicates whether the task is invoking a custom SQL script or stored procedure.
* `typed_expressions` - List of typed expressions.
	* `config_values` - Configuration values can be string, objects, or parameters.
		* `config_param_values` - The configuration parameter values.
			* `int_value` - An integer value of the parameter.
			* `object_value` - An object value of the parameter.
			* `parameter_value` - Reference to the parameter by its key.
			* `ref_value` - The root object reference value.
			* `root_object_value` - The root object value, used in custom parameters.
			* `string_value` - A string value of the parameter.
		* `parent_ref` - A reference to the object's parent.
			* `parent` - Key of the parent object.
			* `root_doc_id` - Key of the root document object.
	* `description` - Detailed description for the object.
	* `expression` - The expression string for the object.
	* `key` - The key of the object.
	* `model_type` - The type of the types object.
	* `model_version` - The model version of an object.
	* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - A reference to the object's parent.
		* `parent` - Key of the parent object.
		* `root_doc_id` - Key of the root document object.
	* `type` - The object type.

