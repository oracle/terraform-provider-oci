---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_task"
sidebar_current: "docs-oci-resource-dataintegration-workspace_task"
description: |-
  Provides the Workspace Task resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_task
This resource provides the Workspace Task resource in Oracle Cloud Infrastructure Data Integration service.

Creates a new task ready for performing data integrations. There are specialized types of tasks that include data loader and integration tasks.


## Example Usage

```hcl
resource "oci_dataintegration_workspace_task" "test_workspace_task" {
	#Required
	identifier = var.workspace_task_identifier
	model_type = var.workspace_task_model_type
	name = var.workspace_task_name
	registry_metadata {

		#Optional
		aggregator_key = var.workspace_task_registry_metadata_aggregator_key
		is_favorite = var.workspace_task_registry_metadata_is_favorite
		key = var.workspace_task_registry_metadata_key
		labels = var.workspace_task_registry_metadata_labels
		registry_version = var.workspace_task_registry_metadata_registry_version
	}
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	api_call_mode = var.workspace_task_api_call_mode
	auth_config {

		#Optional
		key = var.workspace_task_auth_config_key
		model_type = var.workspace_task_auth_config_model_type
		model_version = var.workspace_task_auth_config_model_version
		parent_ref {

			#Optional
			parent = var.workspace_task_auth_config_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		resource_principal_source = var.workspace_task_auth_config_resource_principal_source
	}
	auth_details {

		#Optional
		key = var.workspace_task_auth_details_key
		model_type = var.workspace_task_auth_details_model_type
		model_version = var.workspace_task_auth_details_model_version
		parent_ref {

			#Optional
			parent = var.workspace_task_auth_details_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
	}
	cancel_endpoint {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_cancel_endpoint_config_values_config_param_values_int_value
				object_value = var.workspace_task_cancel_endpoint_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_cancel_endpoint_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_cancel_endpoint_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_cancel_endpoint_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_cancel_endpoint_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_cancel_endpoint_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		expr_string = var.workspace_task_cancel_endpoint_expr_string
		key = var.workspace_task_cancel_endpoint_key
		model_type = var.workspace_task_cancel_endpoint_model_type
		model_version = var.workspace_task_cancel_endpoint_model_version
		object_status = var.workspace_task_cancel_endpoint_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_cancel_endpoint_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
	}
	cancel_method_type = var.workspace_task_cancel_method_type
	cancel_rest_call_config {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_int_value
				object_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_cancel_rest_call_config_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		method_type = var.workspace_task_cancel_rest_call_config_method_type
		request_headers = var.workspace_task_cancel_rest_call_config_request_headers
	}
	conditional_composite_field_map {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_conditional_composite_field_map_config_values_config_param_values_int_value
				object_value = var.workspace_task_conditional_composite_field_map_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_conditional_composite_field_map_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_conditional_composite_field_map_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_conditional_composite_field_map_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_conditional_composite_field_map_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_conditional_composite_field_map_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		description = var.workspace_task_conditional_composite_field_map_description
		field_map_scope {
			#Required
			model_type = var.workspace_task_conditional_composite_field_map_field_map_scope_model_type

			#Optional
			config_values {

				#Optional
				config_param_values {

					#Optional
					int_value = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_config_param_values_int_value
					object_value = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_config_param_values_object_value
					parameter_value = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_config_param_values_parameter_value
					ref_value = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_config_param_values_ref_value
					root_object_value = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_config_param_values_root_object_value
					string_value = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_config_param_values_string_value
				}
				parent_ref {

					#Optional
					parent = var.workspace_task_conditional_composite_field_map_field_map_scope_config_values_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
			}
			description = var.workspace_task_conditional_composite_field_map_field_map_scope_description
			from_name = var.workspace_task_conditional_composite_field_map_field_map_scope_from_name
			is_cascade = var.workspace_task_conditional_composite_field_map_field_map_scope_is_cascade
			is_case_sensitive = var.workspace_task_conditional_composite_field_map_field_map_scope_is_case_sensitive
			is_java_regex_syntax = var.workspace_task_conditional_composite_field_map_field_map_scope_is_java_regex_syntax
			is_skip_remaining_rules_on_match = var.workspace_task_conditional_composite_field_map_field_map_scope_is_skip_remaining_rules_on_match
			key = var.workspace_task_conditional_composite_field_map_field_map_scope_key
			matching_strategy = var.workspace_task_conditional_composite_field_map_field_map_scope_matching_strategy
			model_version = var.workspace_task_conditional_composite_field_map_field_map_scope_model_version
			name = var.workspace_task_conditional_composite_field_map_field_map_scope_name
			names = var.workspace_task_conditional_composite_field_map_field_map_scope_names
			object_status = var.workspace_task_conditional_composite_field_map_field_map_scope_object_status
			parent_ref {

				#Optional
				parent = var.workspace_task_conditional_composite_field_map_field_map_scope_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
			pattern = var.workspace_task_conditional_composite_field_map_field_map_scope_pattern
			rule_type = var.workspace_task_conditional_composite_field_map_field_map_scope_rule_type
			scope = var.workspace_task_conditional_composite_field_map_field_map_scope_scope
			to_name = var.workspace_task_conditional_composite_field_map_field_map_scope_to_name
			types = var.workspace_task_conditional_composite_field_map_field_map_scope_types
		}
		field_maps = var.workspace_task_conditional_composite_field_map_field_maps
		key = var.workspace_task_conditional_composite_field_map_key
		model_type = var.workspace_task_conditional_composite_field_map_model_type
		model_version = var.workspace_task_conditional_composite_field_map_model_version
		object_status = var.workspace_task_conditional_composite_field_map_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_conditional_composite_field_map_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
	}
	config_provider_delegate {

		#Optional
		bindings {

			#Optional
			root_object_value = var.workspace_task_config_provider_delegate_bindings_root_object_value
			simple_value = var.workspace_task_config_provider_delegate_bindings_simple_value
		}
	}
	data_flow {

		#Optional
		description = var.workspace_task_data_flow_description
		flow_config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_data_flow_flow_config_values_config_param_values_int_value
				object_value = var.workspace_task_data_flow_flow_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_data_flow_flow_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_data_flow_flow_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_data_flow_flow_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_data_flow_flow_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_data_flow_flow_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		identifier = var.workspace_task_data_flow_identifier
		key = var.workspace_task_data_flow_key
		key_map = var.workspace_task_data_flow_key_map
		metadata {

			#Optional
			aggregator {

				#Optional
				description = var.workspace_task_data_flow_metadata_aggregator_description
				identifier = var.workspace_task_data_flow_metadata_aggregator_identifier
				key = var.workspace_task_data_flow_metadata_aggregator_key
				name = var.workspace_task_data_flow_metadata_aggregator_name
				type = var.workspace_task_data_flow_metadata_aggregator_type
			}
			aggregator_key = var.workspace_task_data_flow_metadata_aggregator_key
			count_statistics {

				#Optional
				object_type_count_list {

					#Optional
					object_count = var.workspace_task_data_flow_metadata_count_statistics_object_type_count_list_object_count
					object_type = var.workspace_task_data_flow_metadata_count_statistics_object_type_count_list_object_type
				}
			}
			created_by = var.workspace_task_data_flow_metadata_created_by
			created_by_name = var.workspace_task_data_flow_metadata_created_by_name
			identifier_path = var.workspace_task_data_flow_metadata_identifier_path
			info_fields = var.workspace_task_data_flow_metadata_info_fields
			is_favorite = var.workspace_task_data_flow_metadata_is_favorite
			labels = var.workspace_task_data_flow_metadata_labels
			registry_version = var.workspace_task_data_flow_metadata_registry_version
			time_created = var.workspace_task_data_flow_metadata_time_created
			time_updated = var.workspace_task_data_flow_metadata_time_updated
			updated_by = var.workspace_task_data_flow_metadata_updated_by
			updated_by_name = var.workspace_task_data_flow_metadata_updated_by_name
		}
		model_type = var.workspace_task_data_flow_model_type
		model_version = var.workspace_task_data_flow_model_version
		name = var.workspace_task_data_flow_name
		nodes {

			#Optional
			config_provider_delegate = var.workspace_task_data_flow_nodes_config_provider_delegate
			description = var.workspace_task_data_flow_nodes_description
			input_links {

				#Optional
				description = var.workspace_task_data_flow_nodes_input_links_description
				field_map = var.workspace_task_data_flow_nodes_input_links_field_map
				from_link = var.workspace_task_data_flow_nodes_input_links_from_link
				key = var.workspace_task_data_flow_nodes_input_links_key
				model_type = var.workspace_task_data_flow_nodes_input_links_model_type
				model_version = var.workspace_task_data_flow_nodes_input_links_model_version
				object_status = var.workspace_task_data_flow_nodes_input_links_object_status
				parent_ref {

					#Optional
					parent = var.workspace_task_data_flow_nodes_input_links_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
				port = var.workspace_task_data_flow_nodes_input_links_port
			}
			key = var.workspace_task_data_flow_nodes_key
			model_type = var.workspace_task_data_flow_nodes_model_type
			model_version = var.workspace_task_data_flow_nodes_model_version
			name = var.workspace_task_data_flow_nodes_name
			object_status = var.workspace_task_data_flow_nodes_object_status
			operator = var.workspace_task_data_flow_nodes_operator
			output_links {

				#Optional
				description = var.workspace_task_data_flow_nodes_output_links_description
				key = var.workspace_task_data_flow_nodes_output_links_key
				model_type = var.workspace_task_data_flow_nodes_output_links_model_type
				model_version = var.workspace_task_data_flow_nodes_output_links_model_version
				object_status = var.workspace_task_data_flow_nodes_output_links_object_status
				parent_ref {

					#Optional
					parent = var.workspace_task_data_flow_nodes_output_links_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
				port = var.workspace_task_data_flow_nodes_output_links_port
				to_links = var.workspace_task_data_flow_nodes_output_links_to_links
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_data_flow_nodes_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
			ui_properties {

				#Optional
				coordinate_x = var.workspace_task_data_flow_nodes_ui_properties_coordinate_x
				coordinate_y = var.workspace_task_data_flow_nodes_ui_properties_coordinate_y
			}
		}
		object_status = var.workspace_task_data_flow_object_status
		object_version = var.workspace_task_data_flow_object_version
		parameters {

			#Optional
			config_values {

				#Optional
				config_param_values {

					#Optional
					int_value = var.workspace_task_data_flow_parameters_config_values_config_param_values_int_value
					object_value = var.workspace_task_data_flow_parameters_config_values_config_param_values_object_value
					parameter_value = var.workspace_task_data_flow_parameters_config_values_config_param_values_parameter_value
					ref_value = var.workspace_task_data_flow_parameters_config_values_config_param_values_ref_value
					root_object_value = var.workspace_task_data_flow_parameters_config_values_config_param_values_root_object_value
					string_value = var.workspace_task_data_flow_parameters_config_values_config_param_values_string_value
				}
				parent_ref {

					#Optional
					parent = var.workspace_task_data_flow_parameters_config_values_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
			}
			default_value = var.workspace_task_data_flow_parameters_default_value
			description = var.workspace_task_data_flow_parameters_description
			is_input = var.workspace_task_data_flow_parameters_is_input
			is_output = var.workspace_task_data_flow_parameters_is_output
			key = var.workspace_task_data_flow_parameters_key
			model_type = var.workspace_task_data_flow_parameters_model_type
			model_version = var.workspace_task_data_flow_parameters_model_version
			name = var.workspace_task_data_flow_parameters_name
			object_status = var.workspace_task_data_flow_parameters_object_status
			output_aggregation_type = var.workspace_task_data_flow_parameters_output_aggregation_type
			parent_ref {

				#Optional
				parent = var.workspace_task_data_flow_parameters_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
			root_object_default_value = var.workspace_task_data_flow_parameters_root_object_default_value
			type = var.workspace_task_data_flow_parameters_type
			type_name = var.workspace_task_data_flow_parameters_type_name
			used_for = var.workspace_task_data_flow_parameters_used_for
		}
		parent_ref {

			#Optional
			parent = var.workspace_task_data_flow_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		target_field_map_summary {

			#Optional
			field_map = var.workspace_task_data_flow_target_field_map_summary_field_map
		}
		typed_object_map {

			#Optional
			typed_object = var.workspace_task_data_flow_typed_object_map_typed_object
		}
	}
	dataflow_application {

		#Optional
		application_id = oci_dataflow_application.test_application.id
		compartment_id = var.compartment_id
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_dataflow_application_config_values_config_param_values_int_value
				object_value = var.workspace_task_dataflow_application_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_dataflow_application_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_dataflow_application_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_dataflow_application_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_dataflow_application_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_dataflow_application_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
	}
	description = var.workspace_task_description
	endpoint {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_endpoint_config_values_config_param_values_int_value
				object_value = var.workspace_task_endpoint_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_endpoint_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_endpoint_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_endpoint_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_endpoint_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_endpoint_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		expr_string = var.workspace_task_endpoint_expr_string
		key = var.workspace_task_endpoint_key
		model_type = var.workspace_task_endpoint_model_type
		model_version = var.workspace_task_endpoint_model_version
		object_status = var.workspace_task_endpoint_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_endpoint_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
	}
	execute_rest_call_config {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_int_value
				object_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_execute_rest_call_config_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		method_type = var.workspace_task_execute_rest_call_config_method_type
		request_headers = var.workspace_task_execute_rest_call_config_request_headers
	}
	headers = var.workspace_task_headers
	input_ports {
		#Required
		model_type = var.workspace_task_input_ports_model_type

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_input_ports_config_values_config_param_values_int_value
				object_value = var.workspace_task_input_ports_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_input_ports_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_input_ports_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_input_ports_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_input_ports_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_input_ports_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		description = var.workspace_task_input_ports_description
		fields = var.workspace_task_input_ports_fields
		key = var.workspace_task_input_ports_key
		model_version = var.workspace_task_input_ports_model_version
		name = var.workspace_task_input_ports_name
		object_status = var.workspace_task_input_ports_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_input_ports_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		port_type = var.workspace_task_input_ports_port_type
	}
	is_single_load = var.workspace_task_is_single_load
	json_data = var.workspace_task_json_data
	key = var.workspace_task_key
	method_type = var.workspace_task_method_type
	model_version = var.workspace_task_model_version
	object_status = var.workspace_task_object_status
	op_config_values {

		#Optional
		config_param_values {

			#Optional
			int_value = var.workspace_task_op_config_values_config_param_values_int_value
			object_value = var.workspace_task_op_config_values_config_param_values_object_value
			parameter_value = var.workspace_task_op_config_values_config_param_values_parameter_value
			ref_value = var.workspace_task_op_config_values_config_param_values_ref_value
			root_object_value = var.workspace_task_op_config_values_config_param_values_root_object_value
			string_value = var.workspace_task_op_config_values_config_param_values_string_value
		}
		parent_ref {

			#Optional
			parent = var.workspace_task_op_config_values_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
	}
	operation = var.workspace_task_operation
	output_ports {
		#Required
		model_type = var.workspace_task_output_ports_model_type

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_output_ports_config_values_config_param_values_int_value
				object_value = var.workspace_task_output_ports_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_output_ports_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_output_ports_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_output_ports_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_output_ports_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_output_ports_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		description = var.workspace_task_output_ports_description
		fields = var.workspace_task_output_ports_fields
		key = var.workspace_task_output_ports_key
		model_version = var.workspace_task_output_ports_model_version
		name = var.workspace_task_output_ports_name
		object_status = var.workspace_task_output_ports_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_output_ports_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		port_type = var.workspace_task_output_ports_port_type
	}
	parallel_load_limit = var.workspace_task_parallel_load_limit
	parameters {
		#Required
		model_type = var.workspace_task_parameters_model_type

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_parameters_config_values_config_param_values_int_value
				object_value = var.workspace_task_parameters_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_parameters_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_parameters_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_parameters_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_parameters_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_parameters_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		default_value = var.workspace_task_parameters_default_value
		description = var.workspace_task_parameters_description
		is_input = var.workspace_task_parameters_is_input
		is_output = var.workspace_task_parameters_is_output
		key = var.workspace_task_parameters_key
		model_version = var.workspace_task_parameters_model_version
		name = var.workspace_task_parameters_name
		object_status = var.workspace_task_parameters_object_status
		output_aggregation_type = var.workspace_task_parameters_output_aggregation_type
		parent_ref {

			#Optional
			parent = var.workspace_task_parameters_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		root_object_default_value = var.workspace_task_parameters_root_object_default_value
		type = var.workspace_task_parameters_type
		type_name = var.workspace_task_parameters_type_name
		used_for = var.workspace_task_parameters_used_for
	}
	parent_ref {

		#Optional
		parent = var.workspace_task_parent_ref_parent
		root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
	}
	pipeline {

		#Optional
		description = var.workspace_task_pipeline_description
		flow_config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_pipeline_flow_config_values_config_param_values_int_value
				object_value = var.workspace_task_pipeline_flow_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_pipeline_flow_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_pipeline_flow_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_pipeline_flow_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_pipeline_flow_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_pipeline_flow_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		identifier = var.workspace_task_pipeline_identifier
		key = var.workspace_task_pipeline_key
		metadata {

			#Optional
			aggregator {

				#Optional
				description = var.workspace_task_pipeline_metadata_aggregator_description
				identifier = var.workspace_task_pipeline_metadata_aggregator_identifier
				key = var.workspace_task_pipeline_metadata_aggregator_key
				name = var.workspace_task_pipeline_metadata_aggregator_name
				type = var.workspace_task_pipeline_metadata_aggregator_type
			}
			aggregator_key = var.workspace_task_pipeline_metadata_aggregator_key
			count_statistics {

				#Optional
				object_type_count_list {

					#Optional
					object_count = var.workspace_task_pipeline_metadata_count_statistics_object_type_count_list_object_count
					object_type = var.workspace_task_pipeline_metadata_count_statistics_object_type_count_list_object_type
				}
			}
			created_by = var.workspace_task_pipeline_metadata_created_by
			created_by_name = var.workspace_task_pipeline_metadata_created_by_name
			identifier_path = var.workspace_task_pipeline_metadata_identifier_path
			info_fields = var.workspace_task_pipeline_metadata_info_fields
			is_favorite = var.workspace_task_pipeline_metadata_is_favorite
			labels = var.workspace_task_pipeline_metadata_labels
			registry_version = var.workspace_task_pipeline_metadata_registry_version
			time_created = var.workspace_task_pipeline_metadata_time_created
			time_updated = var.workspace_task_pipeline_metadata_time_updated
			updated_by = var.workspace_task_pipeline_metadata_updated_by
			updated_by_name = var.workspace_task_pipeline_metadata_updated_by_name
		}
		model_type = var.workspace_task_pipeline_model_type
		model_version = var.workspace_task_pipeline_model_version
		name = var.workspace_task_pipeline_name
		nodes {

			#Optional
			config_provider_delegate = var.workspace_task_pipeline_nodes_config_provider_delegate
			description = var.workspace_task_pipeline_nodes_description
			input_links {

				#Optional
				description = var.workspace_task_pipeline_nodes_input_links_description
				field_map = var.workspace_task_pipeline_nodes_input_links_field_map
				from_link = var.workspace_task_pipeline_nodes_input_links_from_link
				key = var.workspace_task_pipeline_nodes_input_links_key
				model_type = var.workspace_task_pipeline_nodes_input_links_model_type
				model_version = var.workspace_task_pipeline_nodes_input_links_model_version
				object_status = var.workspace_task_pipeline_nodes_input_links_object_status
				parent_ref {

					#Optional
					parent = var.workspace_task_pipeline_nodes_input_links_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
				port = var.workspace_task_pipeline_nodes_input_links_port
			}
			key = var.workspace_task_pipeline_nodes_key
			model_type = var.workspace_task_pipeline_nodes_model_type
			model_version = var.workspace_task_pipeline_nodes_model_version
			name = var.workspace_task_pipeline_nodes_name
			object_status = var.workspace_task_pipeline_nodes_object_status
			operator = var.workspace_task_pipeline_nodes_operator
			output_links {

				#Optional
				description = var.workspace_task_pipeline_nodes_output_links_description
				key = var.workspace_task_pipeline_nodes_output_links_key
				model_type = var.workspace_task_pipeline_nodes_output_links_model_type
				model_version = var.workspace_task_pipeline_nodes_output_links_model_version
				object_status = var.workspace_task_pipeline_nodes_output_links_object_status
				parent_ref {

					#Optional
					parent = var.workspace_task_pipeline_nodes_output_links_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
				port = var.workspace_task_pipeline_nodes_output_links_port
				to_links = var.workspace_task_pipeline_nodes_output_links_to_links
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_pipeline_nodes_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
			ui_properties {

				#Optional
				coordinate_x = var.workspace_task_pipeline_nodes_ui_properties_coordinate_x
				coordinate_y = var.workspace_task_pipeline_nodes_ui_properties_coordinate_y
			}
		}
		object_status = var.workspace_task_pipeline_object_status
		object_version = var.workspace_task_pipeline_object_version
		parameters {

			#Optional
			config_values {

				#Optional
				config_param_values {

					#Optional
					int_value = var.workspace_task_pipeline_parameters_config_values_config_param_values_int_value
					object_value = var.workspace_task_pipeline_parameters_config_values_config_param_values_object_value
					parameter_value = var.workspace_task_pipeline_parameters_config_values_config_param_values_parameter_value
					ref_value = var.workspace_task_pipeline_parameters_config_values_config_param_values_ref_value
					root_object_value = var.workspace_task_pipeline_parameters_config_values_config_param_values_root_object_value
					string_value = var.workspace_task_pipeline_parameters_config_values_config_param_values_string_value
				}
				parent_ref {

					#Optional
					parent = var.workspace_task_pipeline_parameters_config_values_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
			}
			default_value = var.workspace_task_pipeline_parameters_default_value
			description = var.workspace_task_pipeline_parameters_description
			is_input = var.workspace_task_pipeline_parameters_is_input
			is_output = var.workspace_task_pipeline_parameters_is_output
			key = var.workspace_task_pipeline_parameters_key
			model_type = var.workspace_task_pipeline_parameters_model_type
			model_version = var.workspace_task_pipeline_parameters_model_version
			name = var.workspace_task_pipeline_parameters_name
			object_status = var.workspace_task_pipeline_parameters_object_status
			output_aggregation_type = var.workspace_task_pipeline_parameters_output_aggregation_type
			parent_ref {

				#Optional
				parent = var.workspace_task_pipeline_parameters_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
			root_object_default_value = var.workspace_task_pipeline_parameters_root_object_default_value
			type = var.workspace_task_pipeline_parameters_type
			type_name = var.workspace_task_pipeline_parameters_type_name
			used_for = var.workspace_task_pipeline_parameters_used_for
		}
		parent_ref {

			#Optional
			parent = var.workspace_task_pipeline_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		variables {

			#Optional
			config_values {

				#Optional
				config_param_values {

					#Optional
					int_value = var.workspace_task_pipeline_variables_config_values_config_param_values_int_value
					object_value = var.workspace_task_pipeline_variables_config_values_config_param_values_object_value
					parameter_value = var.workspace_task_pipeline_variables_config_values_config_param_values_parameter_value
					ref_value = var.workspace_task_pipeline_variables_config_values_config_param_values_ref_value
					root_object_value = var.workspace_task_pipeline_variables_config_values_config_param_values_root_object_value
					string_value = var.workspace_task_pipeline_variables_config_values_config_param_values_string_value
				}
				parent_ref {

					#Optional
					parent = var.workspace_task_pipeline_variables_config_values_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
			}
			default_value = var.workspace_task_pipeline_variables_default_value
			description = var.workspace_task_pipeline_variables_description
			identifier = var.workspace_task_pipeline_variables_identifier
			key = var.workspace_task_pipeline_variables_key
			model_type = var.workspace_task_pipeline_variables_model_type
			model_version = var.workspace_task_pipeline_variables_model_version
			name = var.workspace_task_pipeline_variables_name
			object_status = var.workspace_task_pipeline_variables_object_status
			object_version = var.workspace_task_pipeline_variables_object_version
			parent_ref {

				#Optional
				parent = var.workspace_task_pipeline_variables_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
			root_object_default_value {

				#Optional
				key = var.workspace_task_pipeline_variables_root_object_default_value_key
				model_type = var.workspace_task_pipeline_variables_root_object_default_value_model_type
				model_version = var.workspace_task_pipeline_variables_root_object_default_value_model_version
				object_status = var.workspace_task_pipeline_variables_root_object_default_value_object_status
				parent_ref {

					#Optional
					parent = var.workspace_task_pipeline_variables_root_object_default_value_parent_ref_parent
					root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
				}
			}
			type = var.workspace_task_pipeline_variables_type
		}
	}
	poll_rest_call_config {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_poll_rest_call_config_config_values_config_param_values_int_value
				object_value = var.workspace_task_poll_rest_call_config_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_poll_rest_call_config_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_poll_rest_call_config_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_poll_rest_call_config_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_poll_rest_call_config_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_poll_rest_call_config_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		method_type = var.workspace_task_poll_rest_call_config_method_type
		request_headers = var.workspace_task_poll_rest_call_config_request_headers
	}
	script {

		#Optional
		key = var.workspace_task_script_key
		model_type = var.workspace_task_script_model_type
		model_version = var.workspace_task_script_model_version
		object_status = var.workspace_task_script_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_script_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
	}
	sql_script_type = var.workspace_task_sql_script_type
	typed_expressions {

		#Optional
		config_values {

			#Optional
			config_param_values {

				#Optional
				int_value = var.workspace_task_typed_expressions_config_values_config_param_values_int_value
				object_value = var.workspace_task_typed_expressions_config_values_config_param_values_object_value
				parameter_value = var.workspace_task_typed_expressions_config_values_config_param_values_parameter_value
				ref_value = var.workspace_task_typed_expressions_config_values_config_param_values_ref_value
				root_object_value = var.workspace_task_typed_expressions_config_values_config_param_values_root_object_value
				string_value = var.workspace_task_typed_expressions_config_values_config_param_values_string_value
			}
			parent_ref {

				#Optional
				parent = var.workspace_task_typed_expressions_config_values_parent_ref_parent
				root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
			}
		}
		description = var.workspace_task_typed_expressions_description
		expression = var.workspace_task_typed_expressions_expression
		key = var.workspace_task_typed_expressions_key
		model_type = var.workspace_task_typed_expressions_model_type
		model_version = var.workspace_task_typed_expressions_model_version
		name = var.workspace_task_typed_expressions_name
		object_status = var.workspace_task_typed_expressions_object_status
		parent_ref {

			#Optional
			parent = var.workspace_task_typed_expressions_parent_ref_parent
			root_doc_id = oci_dataintegration_root_doc.test_root_doc.id
		}
		type = var.workspace_task_typed_expressions_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `api_call_mode` - (Applicable when model_type=REST_TASK) (Updatable) The REST invocation pattern to use. ASYNC_OCI_WORKREQUEST is being deprecated as well as cancelEndpoint/MethodType.
* `auth_config` - (Applicable when model_type=REST_TASK) (Updatable) Authentication configuration for Generic REST invocation.
	* `key` - (Optional) (Updatable) Generated key that can be used in API calls to identify this object.
	* `model_type` - (Optional) (Updatable) The specific authentication configuration to be used for Generic REST invocation.
	* `model_version` - (Optional) (Updatable) The model version of an object.
	* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
		* `parent` - (Optional) (Updatable) Key of the parent object.
		* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `resource_principal_source` - (Optional) (Updatable) The Oracle Cloud Infrastructure resource type that will supply the authentication token
* `auth_details` - (Applicable when model_type=REST_TASK) (Updatable) Authentication type to be used for Generic REST invocation. This is deprecated.
	* `key` - (Applicable when model_type=REST_TASK) (Updatable) Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
	* `model_type` - (Applicable when model_type=REST_TASK) (Updatable) The authentication mode to be used for Generic REST invocation.
	* `model_version` - (Applicable when model_type=REST_TASK) (Updatable) The model version of an object.
	* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
* `cancel_endpoint` - (Applicable when model_type=REST_TASK) (Updatable) An expression node.
	* `config_values` - (Applicable when model_type=REST_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=REST_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=REST_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=REST_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=REST_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=REST_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `expr_string` - (Applicable when model_type=REST_TASK) (Updatable) The expression string for the object.
	* `key` - (Applicable when model_type=REST_TASK) (Updatable) The object key.
	* `model_type` - (Applicable when model_type=REST_TASK) (Updatable) The object type.
	* `model_version` - (Applicable when model_type=REST_TASK) (Updatable) The object's model version.
	* `object_status` - (Applicable when model_type=REST_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
* `cancel_method_type` - (Applicable when model_type=REST_TASK) (Updatable) The REST method to use for canceling the original request.
* `cancel_rest_call_config` - (Applicable when model_type=REST_TASK) (Updatable) The REST API configuration for cancelling the task.
	* `config_values` - (Applicable when model_type=REST_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=REST_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=REST_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=REST_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=REST_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=REST_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `method_type` - (Applicable when model_type=REST_TASK) (Updatable) The REST method to use.
	* `request_headers` - (Applicable when model_type=REST_TASK) (Updatable) The headers for the REST call.
* `conditional_composite_field_map` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A conditional composite field map.
	* `config_values` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the root document object.
	* `description` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Detailed description for the object.
	* `field_map_scope` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) An array of projection rules.
		* `config_values` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Configuration values can be string, objects, or parameters.
			* `config_param_values` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The configuration parameter values.
				* `int_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) An integer value of the parameter.
				* `object_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) An object value of the parameter.
				* `parameter_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Reference to the parameter by its key.
				* `ref_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The root object reference value.
				* `root_object_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The root object value, used in custom parameters.
				* `string_value` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A string value of the parameter.
			* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the root document object.
		* `description` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A user defined description for the object.
		* `from_name` - (Applicable when model_type=RENAME_RULE) (Updatable) The attribute name that needs to be renamed.
		* `is_cascade` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE | NAME_LIST_RULE | NAME_PATTERN_RULE | TYPED_NAME_PATTERN_RULE | TYPE_LIST_RULE) (Updatable) Specifies whether to cascade or not.
		* `is_case_sensitive` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE | NAME_LIST_RULE | NAME_PATTERN_RULE | TYPED_NAME_PATTERN_RULE | TYPE_LIST_RULE) (Updatable) Specifies if the rule is case sensitive.
		* `is_java_regex_syntax` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Specifies whether the rule uses a java regex syntax.
		* `is_skip_remaining_rules_on_match` - (Optional) (Updatable) Specifies whether to skip remaining rules when a match is found.
		* `key` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The key of the object.
		* `matching_strategy` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE | NAME_LIST_RULE | NAME_PATTERN_RULE | TYPED_NAME_PATTERN_RULE | TYPE_LIST_RULE) (Updatable) The pattern matching strategy.
		* `model_type` - (Required) (Updatable) The type of the project rule.
		* `model_version` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The model version of an object.
		* `name` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE) (Updatable) Name of the group.
		* `names` - (Applicable when model_type=NAME_LIST_RULE | TYPED_NAME_PATTERN_RULE) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the root document object.
		* `pattern` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE | NAME_PATTERN_RULE | TYPED_NAME_PATTERN_RULE) (Updatable) The rule pattern.
		* `rule_type` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE | NAME_LIST_RULE | NAME_PATTERN_RULE | TYPED_NAME_PATTERN_RULE | TYPE_LIST_RULE) (Updatable) The rule type.
		* `scope` - (Applicable when model_type=GROUPED_NAME_PATTERN_RULE | NAME_LIST_RULE | NAME_PATTERN_RULE | TYPED_NAME_PATTERN_RULE | TYPE_LIST_RULE) (Updatable) Reference to a typed object. This can be either a key value to an object within the document, a shall referenced to a `TypedObject`, or a full `TypedObject` definition.
		* `to_name` - (Applicable when model_type=RENAME_RULE) (Updatable) The new attribute name.
		* `types` - (Applicable when model_type=TYPED_NAME_PATTERN_RULE | TYPE_LIST_RULE) (Updatable) An array of types.
	* `field_maps` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) An array of field maps.
	* `key` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The object key.
	* `model_type` - (Required when model_type=DATA_LOADER_TASK) (Updatable) The model type for the field map.
	* `model_version` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The object's model version.
	* `object_status` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Key of the root document object.
* `config_provider_delegate` - (Optional) (Updatable) The type to create a config provider.
	* `bindings` - (Optional) bindings
		* `root_object_value` - (Optional) This can be any object such as a file entity, a schema, or a table.
		* `simple_value` - (Optional) A simple value for the parameter.
* `data_flow` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The data flow type contains the audit summary information and the definition of the data flow.
	* `description` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Detailed description for the object.
	* `flow_config_values` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
	* `identifier` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
	* `key_map` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	* `metadata` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A summary type containing information about the object including its key, name and when/who created/updated it.
		* `aggregator` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A summary type containing information about the object's aggregator including its type, key, name and description.
			* `description` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The description of the aggregator.
			* `identifier` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The identifier of the aggregator.
			* `key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The key of the aggregator object.
			* `name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The name of the aggregator.
			* `type` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The type of the aggregator.
		* `aggregator_key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The owning object key for this object.
		* `count_statistics` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A count statistics.
			* `object_type_count_list` - (Required when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The array of statistics.
				* `object_count` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The value for the count statistic object.
				* `object_type` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The type of object for the count statistic object.
		* `created_by` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The user that created the object.
		* `created_by_name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The user that created the object.
		* `identifier_path` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The full path to identify this object.
		* `info_fields` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Information property fields.
		* `is_favorite` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Specifies whether this object is a favorite or not.
		* `labels` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Labels are keywords or tags that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
		* `registry_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The registry version of the object.
		* `time_created` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The date and time that the object was created.
		* `time_updated` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The date and time that the object was updated.
		* `updated_by` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The user that updated the object.
		* `updated_by_name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The user that updated the object.
	* `model_type` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The type of the object.
	* `model_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model version of an object.
	* `name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `nodes` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An array of nodes.
		* `config_provider_delegate` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The information about the configuration provider.
		* `description` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Detailed description for the object.
		* `input_links` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An array of input links.
			* `description` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Detailed description for the object.
			* `field_map` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A field map is a way to map a source row shape to a target row shape that may be different.
			* `from_link` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The from link reference.
			* `key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The key of the object.
			* `model_type` - (Required when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model type of the object.
			* `model_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model version of an object.
			* `object_status` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
			* `port` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of FlowPort reference
		* `key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The key of the object.
		* `model_type` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The type of the object.
		* `model_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model version of an object.
		* `name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `operator` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An operator defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
		* `output_links` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An array of output links.
			* `description` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Detailed description for the object.
			* `key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The key of the object.
			* `model_type` - (Required when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model type of the object.
			* `model_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model version of an object.
			* `object_status` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
			* `port` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of FlowPort reference
			* `to_links` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The links from this output link to connect to other links in flow.
		* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
		* `ui_properties` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The UI properties of the object.
			* `coordinate_x` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The X coordinate of the object.
			* `coordinate_y` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The Y coordinate of the object.
	* `object_status` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The version of the object that is used to track changes in the object instance.
	* `parameters` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An array of parameters.
		* `config_values` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Configuration values can be string, objects, or parameters.
			* `config_param_values` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The configuration parameter values.
				* `int_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An integer value of the parameter.
				* `object_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) An object value of the parameter.
				* `parameter_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Reference to the parameter by its key.
				* `ref_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The root object reference value.
				* `root_object_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The root object value, used in custom parameters.
				* `string_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A string value of the parameter.
			* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
		* `default_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The default value of the parameter.
		* `description` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Detailed description for the object.
		* `is_input` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Specifies whether the parameter is input value.
		* `is_output` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Specifies whether the parameter is output value.
		* `key` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The key of the object.
		* `model_type` - (Required when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The type of the types object.
		* `model_version` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The model version of an object.
		* `name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `output_aggregation_type` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The output aggregation type.
		* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
		* `root_object_default_value` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The default value of the parameter which can be an object in DIS, such as a data entity.
		* `type` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) This can either be a string value referencing the type or a BaseType object.
		* `type_name` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The type of value the parameter was created for.
		* `used_for` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The param name for which parameter is created for for eg. driver Shape, Operation etc.
	* `parent_ref` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) Key of the root document object.
	* `target_field_map_summary` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A hash map that maps TypedObject keys to a field map that maps to the typed object as a target, for java sdk.
		* `field_map` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A field map is a way to map a source row shape to a target row shape that may be different.
	* `typed_object_map` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) A hash map that maps TypedObject keys to the object itself, for java sdk.
		* `typed_object` - (Applicable when model_type=DATA_LOADER_TASK | INTEGRATION_TASK) (Updatable) The `TypedObject` class is a base class for any model object that has a type.
* `dataflow_application` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) Minimum information required to recognize a Dataflow Application object.
	* `application_id` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) The application id for which Oracle Cloud Infrastructure data flow task is to be created.
	* `compartment_id` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) The compartmentId id under which Oracle Cloud Infrastructure dataflow application lies.
	* `config_values` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=OCI_DATAFLOW_TASK) (Updatable) Key of the root document object.
* `description` - (Optional) (Updatable) Detailed description for the object.
* `endpoint` - (Applicable when model_type=REST_TASK) (Updatable) An expression node.
	* `config_values` - (Applicable when model_type=REST_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=REST_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=REST_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=REST_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=REST_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=REST_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `expr_string` - (Applicable when model_type=REST_TASK) (Updatable) The expression string for the object.
	* `key` - (Applicable when model_type=REST_TASK) (Updatable) The object key.
	* `model_type` - (Applicable when model_type=REST_TASK) (Updatable) The object type.
	* `model_version` - (Applicable when model_type=REST_TASK) (Updatable) The object's model version.
	* `object_status` - (Applicable when model_type=REST_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
* `execute_rest_call_config` - (Applicable when model_type=REST_TASK) (Updatable) The REST API configuration for execution.
	* `config_values` - (Applicable when model_type=REST_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=REST_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=REST_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=REST_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=REST_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=REST_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `method_type` - (Applicable when model_type=REST_TASK) (Updatable) The REST method to use.
	* `request_headers` - (Applicable when model_type=REST_TASK) (Updatable) The headers for the REST call.
* `headers` - (Applicable when model_type=REST_TASK) (Updatable) The headers for the REST call. This property is deprecated, use ExecuteRestCallConfig's headers property instead.
* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `input_ports` - (Optional) (Updatable) An array of input ports.
	* `config_values` - (Optional) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Optional) (Updatable) The configuration parameter values.
			* `int_value` - (Optional) (Updatable) An integer value of the parameter.
			* `object_value` - (Optional) (Updatable) An object value of the parameter.
			* `parameter_value` - (Optional) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Optional) (Updatable) The root object reference value.
			* `root_object_value` - (Optional) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Optional) (Updatable) A string value of the parameter.
		* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
			* `parent` - (Optional) (Updatable) Key of the parent object.
			* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `description` - (Optional) (Updatable) Detailed description for the object.
	* `fields` - (Optional) (Updatable) An array of fields.
	* `key` - (Optional) (Updatable) The key of the object.
	* `model_type` - (Required) (Updatable) The type of the types object.
	* `model_version` - (Optional) (Updatable) The model version of an object.
	* `name` - (Optional) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
		* `parent` - (Optional) (Updatable) Key of the parent object.
		* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `port_type` - (Optional) (Updatable) The port details for the data asset.Type.
* `is_single_load` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Defines whether Data Loader task is used for single load or multiple
* `json_data` - (Applicable when model_type=REST_TASK) (Updatable) JSON data for payload body. This property is deprecated, use ExecuteRestCallConfig's payload config param instead.
* `key` - (Optional) (Updatable) Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
* `method_type` - (Applicable when model_type=REST_TASK) (Updatable) The REST method to use. This property is deprecated, use ExecuteRestCallConfig's methodType property instead.
* `model_type` - (Required) (Updatable) The type of the task.
* `model_version` - (Optional) (Updatable) The object's model version.
* `name` - (Required) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `op_config_values` - (Optional) (Updatable) Configuration values can be string, objects, or parameters.
	* `config_param_values` - (Optional) (Updatable) The configuration parameter values.
		* `int_value` - (Optional) (Updatable) An integer value of the parameter.
		* `object_value` - (Optional) (Updatable) An object value of the parameter.
		* `parameter_value` - (Optional) (Updatable) Reference to the parameter by its key.
		* `ref_value` - (Optional) (Updatable) The root object reference value.
		* `root_object_value` - (Optional) (Updatable) The root object value, used in custom parameters.
		* `string_value` - (Optional) (Updatable) A string value of the parameter.
	* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
		* `parent` - (Optional) (Updatable) Key of the parent object.
		* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
* `operation` - (Applicable when model_type=SQL_TASK) (Updatable) Describes the shape of the execution result
* `output_ports` - (Optional) (Updatable) An array of output ports.
	* `config_values` - (Optional) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Optional) (Updatable) The configuration parameter values.
			* `int_value` - (Optional) (Updatable) An integer value of the parameter.
			* `object_value` - (Optional) (Updatable) An object value of the parameter.
			* `parameter_value` - (Optional) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Optional) (Updatable) The root object reference value.
			* `root_object_value` - (Optional) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Optional) (Updatable) A string value of the parameter.
		* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
			* `parent` - (Optional) (Updatable) Key of the parent object.
			* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `description` - (Optional) (Updatable) Detailed description for the object.
	* `fields` - (Optional) (Updatable) An array of fields.
	* `key` - (Optional) (Updatable) The key of the object.
	* `model_type` - (Required) (Updatable) The type of the types object.
	* `model_version` - (Optional) (Updatable) The model version of an object.
	* `name` - (Optional) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
		* `parent` - (Optional) (Updatable) Key of the parent object.
		* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `port_type` - (Optional) (Updatable) The port details for the data asset.Type.
* `parallel_load_limit` - (Applicable when model_type=DATA_LOADER_TASK) (Updatable) Defines the number of entities being loaded in parallel at a time for a Data Loader task
* `parameters` - (Optional) (Updatable) An array of parameters.
	* `config_values` - (Optional) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Optional) (Updatable) The configuration parameter values.
			* `int_value` - (Optional) (Updatable) An integer value of the parameter.
			* `object_value` - (Optional) (Updatable) An object value of the parameter.
			* `parameter_value` - (Optional) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Optional) (Updatable) The root object reference value.
			* `root_object_value` - (Optional) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Optional) (Updatable) A string value of the parameter.
		* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
			* `parent` - (Optional) (Updatable) Key of the parent object.
			* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `default_value` - (Optional) (Updatable) The default value of the parameter.
	* `description` - (Optional) (Updatable) Detailed description for the object.
	* `is_input` - (Optional) (Updatable) Specifies whether the parameter is input value.
	* `is_output` - (Optional) (Updatable) Specifies whether the parameter is output value.
	* `key` - (Optional) (Updatable) The key of the object.
	* `model_type` - (Required) (Updatable) The type of the types object.
	* `model_version` - (Optional) (Updatable) The model version of an object.
	* `name` - (Optional) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `output_aggregation_type` - (Optional) (Updatable) The output aggregation type.
	* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
		* `parent` - (Optional) (Updatable) Key of the parent object.
		* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
	* `root_object_default_value` - (Optional) (Updatable) The default value of the parameter which can be an object in DIS, such as a data entity.
	* `type` - (Optional) (Updatable) This can either be a string value referencing the type or a BaseType object.
	* `type_name` - (Optional) (Updatable) The type of value the parameter was created for.
	* `used_for` - (Optional) (Updatable) The param name for which parameter is created for for eg. driver Shape, Operation etc.
* `parent_ref` - (Optional) (Updatable) A reference to the object's parent.
	* `parent` - (Optional) (Updatable) Key of the parent object.
	* `root_doc_id` - (Optional) (Updatable) Key of the root document object.
* `pipeline` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A pipeline is a logical grouping of tasks that together perform a higher level operation. For example, a pipeline could contain a set of tasks that load and clean data, then execute a dataflow to analyze the data. The pipeline allows you to manage the activities as a unit instead of individually. Users can also schedule the pipeline instead of the tasks independently.
	* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Detailed description for the object.
	* `flow_config_values` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
	* `identifier` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Generated key that can be used in API calls to identify pipeline. On scenarios where reference to the pipeline is needed, a value can be passed in create.
	* `metadata` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A summary type containing information about the object including its key, name and when/who created/updated it.
		* `aggregator` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A summary type containing information about the object's aggregator including its type, key, name and description.
			* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The description of the aggregator.
			* `identifier` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The identifier of the aggregator.
			* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The key of the aggregator object.
			* `name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The name of the aggregator.
			* `type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of the aggregator.
		* `aggregator_key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The owning object key for this object.
		* `count_statistics` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A count statistics.
			* `object_type_count_list` - (Required when model_type=PIPELINE_TASK) (Updatable) The array of statistics.
				* `object_count` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The value for the count statistic object.
				* `object_type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of object for the count statistic object.
		* `created_by` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The user that created the object.
		* `created_by_name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The user that created the object.
		* `identifier_path` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The full path to identify this object.
		* `info_fields` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Information property fields.
		* `is_favorite` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Specifies whether this object is a favorite or not.
		* `labels` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Labels are keywords or tags that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
		* `registry_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The registry version of the object.
		* `time_created` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The date and time that the object was created.
		* `time_updated` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The date and time that the object was updated.
		* `updated_by` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The user that updated the object.
		* `updated_by_name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The user that updated the object.
	* `model_type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of the object.
	* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	* `name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `nodes` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A list of nodes attached to the pipeline.
		* `config_provider_delegate` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The information about the configuration provider.
		* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Detailed description for the object.
		* `input_links` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An array of input links.
			* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Detailed description for the object.
			* `field_map` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A field map is a way to map a source row shape to a target row shape that may be different.
			* `from_link` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The from link reference.
			* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The key of the object.
			* `model_type` - (Required when model_type=PIPELINE_TASK) (Updatable) The model type of the object.
			* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The model version of an object.
			* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
			* `port` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of FlowPort reference
		* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The key of the object.
		* `model_type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of the object.
		* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The model version of an object.
		* `name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `operator` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An operator defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
		* `output_links` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An array of output links.
			* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Detailed description for the object.
			* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The key of the object.
			* `model_type` - (Required when model_type=PIPELINE_TASK) (Updatable) The model type of the object.
			* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The model version of an object.
			* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
			* `port` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of FlowPort reference
			* `to_links` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The links from this output link to connect to other links in flow.
		* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
		* `ui_properties` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The UI properties of the object.
			* `coordinate_x` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The X coordinate of the object.
			* `coordinate_y` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The Y coordinate of the object.
	* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `object_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	* `parameters` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A list of parameters for the pipeline, this allows certain aspects of the pipeline to be configured when the pipeline is executed.
		* `config_values` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Configuration values can be string, objects, or parameters.
			* `config_param_values` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The configuration parameter values.
				* `int_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An integer value of the parameter.
				* `object_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An object value of the parameter.
				* `parameter_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Reference to the parameter by its key.
				* `ref_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The root object reference value.
				* `root_object_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The root object value, used in custom parameters.
				* `string_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A string value of the parameter.
			* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
		* `default_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The default value of the parameter.
		* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Detailed description for the object.
		* `is_input` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Specifies whether the parameter is input value.
		* `is_output` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Specifies whether the parameter is output value.
		* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The key of the object.
		* `model_type` - (Required when model_type=PIPELINE_TASK) (Updatable) The type of the types object.
		* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The model version of an object.
		* `name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `output_aggregation_type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The output aggregation type.
		* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
		* `root_object_default_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The default value of the parameter which can be an object in DIS, such as a data entity.
		* `type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) This can either be a string value referencing the type or a BaseType object.
		* `type_name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of value the parameter was created for.
		* `used_for` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The param name for which parameter is created for for eg. driver Shape, Operation etc.
	* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
	* `variables` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The list of variables required in pipeline, variables can be used to store values that can be used as inputs to tasks in the pipeline.
		* `config_values` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Configuration values can be string, objects, or parameters.
			* `config_param_values` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The configuration parameter values.
				* `int_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An integer value of the parameter.
				* `object_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) An object value of the parameter.
				* `parameter_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Reference to the parameter by its key.
				* `ref_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The root object reference value.
				* `root_object_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The root object value, used in custom parameters.
				* `string_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A string value of the parameter.
			* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
		* `default_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A default value for the vairable.
		* `description` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Detailed description for the object.
		* `identifier` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
		* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Generated key that can be used in API calls to identify variable. On scenarios where reference to the variable is needed, a value can be passed in create.
		* `model_type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of the object.
		* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) This is a version number that is used by the service to upgrade objects if needed through releases of the service.
		* `name` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
		* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
		* `object_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
		* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
		* `root_object_default_value` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A base class for all model types, including First Class and its contained objects.
			* `key` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The key of the object.
			* `model_type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The type of the object.
			* `model_version` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The model version of an object.
			* `object_status` - (Applicable when model_type=PIPELINE_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
			* `parent_ref` - (Applicable when model_type=PIPELINE_TASK) (Updatable) A reference to the object's parent.
				* `parent` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the parent object.
				* `root_doc_id` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Key of the root document object.
		* `type` - (Applicable when model_type=PIPELINE_TASK) (Updatable) Base type for the type system.
* `poll_rest_call_config` - (Applicable when model_type=REST_TASK) (Updatable) The REST API configuration for polling.
	* `config_values` - (Applicable when model_type=REST_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=REST_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=REST_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=REST_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=REST_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=REST_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `method_type` - (Applicable when model_type=REST_TASK) (Updatable) The REST method to use.
	* `request_headers` - (Applicable when model_type=REST_TASK) (Updatable) The headers for the REST call.
* `registry_metadata` - (Required) (Updatable) Information about the object and its parent.
	* `aggregator_key` - (Optional) (Updatable) The owning object's key for this object.
	* `is_favorite` - (Optional) (Updatable) Specifies whether this object is a favorite or not.
	* `key` - (Optional) (Updatable) The identifying key for the object.
	* `labels` - (Optional) (Updatable) Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	* `registry_version` - (Optional) (Updatable) The registry version.
* `script` - (Applicable when model_type=SQL_TASK) (Updatable) The script object.
	* `key` - (Applicable when model_type=SQL_TASK) (Updatable) The key of the object.
	* `model_type` - (Applicable when model_type=SQL_TASK) (Updatable) The type of the object.
	* `model_version` - (Applicable when model_type=SQL_TASK) (Updatable) The model version of an object.
	* `object_status` - (Applicable when model_type=SQL_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Applicable when model_type=SQL_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=SQL_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=SQL_TASK) (Updatable) Key of the root document object.
* `sql_script_type` - (Applicable when model_type=SQL_TASK) (Updatable) Indicates whether the task is invoking a custom SQL script or stored procedure.
* `typed_expressions` - (Applicable when model_type=REST_TASK) (Updatable) List of typed expressions.
	* `config_values` - (Applicable when model_type=REST_TASK) (Updatable) Configuration values can be string, objects, or parameters.
		* `config_param_values` - (Applicable when model_type=REST_TASK) (Updatable) The configuration parameter values.
			* `int_value` - (Applicable when model_type=REST_TASK) (Updatable) An integer value of the parameter.
			* `object_value` - (Applicable when model_type=REST_TASK) (Updatable) An object value of the parameter.
			* `parameter_value` - (Applicable when model_type=REST_TASK) (Updatable) Reference to the parameter by its key.
			* `ref_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object reference value.
			* `root_object_value` - (Applicable when model_type=REST_TASK) (Updatable) The root object value, used in custom parameters.
			* `string_value` - (Applicable when model_type=REST_TASK) (Updatable) A string value of the parameter.
		* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
			* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
			* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `description` - (Applicable when model_type=REST_TASK) (Updatable) Detailed description for the object.
	* `expression` - (Applicable when model_type=REST_TASK) (Updatable) The expression string for the object.
	* `key` - (Applicable when model_type=REST_TASK) (Updatable) The key of the object.
	* `model_type` - (Required when model_type=REST_TASK) (Updatable) The type of the types object.
	* `model_version` - (Applicable when model_type=REST_TASK) (Updatable) The model version of an object.
	* `name` - (Applicable when model_type=REST_TASK) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	* `object_status` - (Applicable when model_type=REST_TASK) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	* `parent_ref` - (Applicable when model_type=REST_TASK) (Updatable) A reference to the object's parent.
		* `parent` - (Applicable when model_type=REST_TASK) (Updatable) Key of the parent object.
		* `root_doc_id` - (Applicable when model_type=REST_TASK) (Updatable) Key of the root document object.
	* `type` - (Applicable when model_type=REST_TASK) (Updatable) The object type.
* `workspace_id` - (Required) The workspace ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Workspace Task
	* `update` - (Defaults to 20 minutes), when updating the Workspace Task
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Task


## Import

WorkspaceTasks can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_task.test_workspace_task "workspaces/{workspaceId}/tasks/{taskKey}" 
```

