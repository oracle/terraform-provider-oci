// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "workspace_task_identifier" {
  default = "TERSI_TEST_REST_TASK"
}

variable "workspace_task_model_type" {
  default = "REST_TASK"
}

variable "workspace_task_name" {
  default = "TERSI_TEST_REST_TASK"
}

variable "workspace_task_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_task_registry_metadata_labels" {
  default = ["labels"]
}

variable "workspace_task_api_call_mode" {
  default = "ASYNC_GENERIC"
}

variable "workspace_task_auth_config_model_type" {
  default = "OCI_RESOURCE_AUTH_CONFIG"
}

variable "workspace_task_auth_config_resource_principal_source" {
  default = "WORKSPACE"
}

variable "workspace_task_cancel_rest_call_config_method_type" {
  default = "DELETE"
}

variable "workspace_task_cancel_rest_call_config_request_headers" {
  default = {Content-Type: "application/json"}
}

variable "workspace_task_cancel_rest_call_config_model_type" {
  default = "CANCEL_REST_CALL_CONFIG"
}

variable "workspace_task_cancel_rest_call_config_config_values_config_param_values_request_url_string_value" {
default = "http://den03cyq.us.oracle.com:8086/20200430/workspaces/"
}

variable "workspace_task_cancel_rest_call_config_config_values_config_param_values_request_payload_parameter_value" {
 default = ""
}

variable "workspace_task_cancel_rest_call_config_config_values_config_param_values_request_payload_dataparam_string_value" {
 default = "{\n    \"modelType\": \"USER_PROJECT\",\n    \"name\":\"PROJECT_NAME\",\n    \"identifier\":\"PROJECT_NAME\",\n    \"description\":\"Project created using REST task.\"\n}"
}

variable "workspace_task_config_provider_delegate_bindings_key" {
 default = "PARAMETER_20230920_094229"
}

variable "workspace_task_config_provider_delegate_bindings_simple_value" {
 default = "12"
}

variable "workspace_task_description" {
  default = "description"
}

variable "workspace_task_execute_rest_call_config_model_type" {
  default = "REST_CALL_CONFIG"
}

variable "workspace_task_execute_rest_call_config_method_type" {
  default = "GET"
}

variable "workspace_task_execute_rest_call_config_request_headers" {
  default = {Content-Type: "application/json"}
}

variable "workspace_task_execute_rest_call_config_config_values_config_param_values_request_url_string_value" {
  default = "http://den03cyq.us.oracle.com:8086/20200430/workspaces/"
}

variable "workspace_task_execute_rest_call_config_config_values_config_param_values_request_url_parameter_value" {
  default = ""
}

variable "workspace_task_execute_rest_call_config_config_values_config_param_values_request_payload_dataparam_string_value" {
  default = "{\n    \"modelType\": \"USER_PROJECT\",\n    \"name\":\"PROJECT_NAME\",\n    \"identifier\":\"PROJECT_NAME\",\n    \"description\":\"Project created using REST task.\"\n}"
}

variable "workspace_task_op_config_values_config_param_value_key" {
  default = "successCondition"
}

variable "workspace_task_op_config_values_config_param_value_string_value" {
  default = "true"
}

variable "workspace_task_parameters_default_value" {
  default = "1234"
}

variable "workspace_task_parameters_description" {
  default = "description"
}

variable "workspace_task_parameters_is_input" {
  default = true
}

variable "workspace_task_parameters_is_output" {
  default = false
}


variable "workspace_task_parameters_model_type" {
  default = "PARAMETER"
}

variable "workspace_task_parameters_name" {
  default = "WORKSPACE_ID"
}

variable "workspace_task_parameters_type" {
  default = "Seeded:/typeSystems/PLATFORM/dataTypes/STRING"
}

variable "workspace_task_parameters_type_name" {
  default = "STRING"
}

variable "workspace_task_parameters_config_values_config_key" {
  default = "length"
}

variable "workspace_task_parameters_config_values_config_param_values_int_value" {
  default = "100"
}

variable "workspace_task_poll_rest_call_config_method_type" {
  default = "GET"
}

variable "workspace_task_poll_rest_call_config_request_headers" {
  default = {Content-Type: "application/json"}
}

variable "workspace_task_poll_rest_call_config_config_values_config_param_values_poll_max_duration" {
  default = "140"
}

variable "workspace_task_poll_rest_call_config_config_values_config_param_values_poll_max_duration_unit" {
  default = "MINUTES"
}

variable "workspace_task_poll_rest_call_config_config_values_config_param_values_poll_interval" {
  default = "2"
}

variable "workspace_task_poll_rest_call_config_config_values_config_param_values_poll_interval_unit" {
  default = "MINUTES"
}

variable "workspace_task_poll_rest_call_config_config_values_config_param_values_request_url" {
  default = "http://den03cyq.us.oracle.com:8086/20200430/workspaces/WORKSPACE_ID/projects/PROJECT_KEY"
}

variable "workspace_task_typed_expressions_expression" {
  default = "CAST(json_path(SYS.RESPONSE_PAYLOAD, 'key') AS String)"
}

variable "workspace_task_typed_expressions_model_type" {
  default = "TYPED_EXPRESSION"
}


variable "workspace_task_typed_expressions_name" {
  default = "PROJECT_KEY"
}

variable "workspace_task_typed_expressions_type" {
  default = "Seeded:/typeSystems/PLATFORM/dataTypes/STRING"
}

variable "workspace_task_typed_expressions_config_values_config_param_values_int_value" {
  default = "2000"
}

variable "workspace_task_folder_id" {
  default = ""
}

variable "workspace_task_type" {
  default = ["REST_TASK"]
}

variable "workspace_task_key" {
  default = []
}

variable "workspace_task_fields" {
  default = ["metadata"]
}

variable "workspace_task_identifiers" {
  default = ["TERSI_TEST_REST_TASK"]
}

variable "compartment_ocid" {
default = ""
}

variable "workspace_task_cancel_rest_call_config_config_values_config_param_values_request_payload_ref_value_model_type" {
    default = "JSON_TEXT"
}


variable "workspace_task_execute_rest_call_config_config_values_config_param_values_request_payload_ref_value_model_type" {
    default = "JSON_TEXT"
}

variable "workspace_project_identifier" {
  default = ["TESTWORKSPACEPROJECT"]
}

variable "workspace_project_name" {
  default = "TestWorkspaceProject"
}
variable "workspace_task_poll_rest_call_config_config_values_config_param_values_poll_condition_model_type" {
    default = "EXPRESSION"
}

variable "workspace_task_poll_rest_call_config_config_values_config_param_values_poll_condition_expr_string" {
    default = "CAST(json_path(SYS.RESPONSE_PAYLOAD, 'name') AS String) != 'PROJECT_TEST"
}

resource "oci_dataintegration_workspace" "test_workspace" {
  #Required
  display_name = "TfTestWorkspace"
  compartment_id = var.compartment_ocid
  is_private_network_enabled = false
}

resource "oci_dataintegration_workspace_project" "test_workspace_project" {
  #Required
  identifier   = element(var.workspace_project_identifier, 0)
  name         = var.workspace_project_name
  workspace_id = oci_dataintegration_workspace.test_workspace.id
}


//Example for Rest Task
resource "oci_dataintegration_workspace_task" "test_workspace_task" {
  #Required
  identifier = var.workspace_task_identifier
  model_type = var.workspace_task_model_type
  name       = var.workspace_task_name
  registry_metadata {

    #Optional
    aggregator_key   = oci_dataintegration_workspace_project.test_workspace_project.key
    is_favorite      = var.workspace_task_registry_metadata_is_favorite
    labels           = var.workspace_task_registry_metadata_labels
  }
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  api_call_mode = var.workspace_task_api_call_mode
  auth_config {

    #Optional
    model_type    = var.workspace_task_auth_config_model_type
    resource_principal_source = var.workspace_task_auth_config_resource_principal_source
  }
  cancel_rest_call_config {

    #Optional
    config_values {

      #Optional
      config_param_values {

        #Optional
        request_url {
            string_value      = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_request_url_string_value
        }
        request_payload {
            parameter_value   = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_request_payload_parameter_value
            ref_value  {
                model_type = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_request_payload_ref_value_model_type
                config_values {
                    config_param_values {
                        data_param {
                            string_value = var.workspace_task_cancel_rest_call_config_config_values_config_param_values_request_payload_dataparam_string_value
                        }
                    }
                }
            }

        }
      }
    }
    method_type     = var.workspace_task_cancel_rest_call_config_method_type
    request_headers = var.workspace_task_cancel_rest_call_config_request_headers
    model_type      = var.workspace_task_cancel_rest_call_config_model_type
  }
  config_provider_delegate {

    #Optional
    bindings {

      #Optional
      key = var.workspace_task_config_provider_delegate_bindings_key
      parameter_values {
           simple_value      = var.workspace_task_config_provider_delegate_bindings_simple_value
        }
    }
  }
  description = var.workspace_task_description
  execute_rest_call_config {

    #Optional
    config_values {

      #Optional
      config_param_values {
        request_url {
            string_value      = var.workspace_task_execute_rest_call_config_config_values_config_param_values_request_url_string_value
        }
        request_payload {
            parameter_value   = var.workspace_task_execute_rest_call_config_config_values_config_param_values_request_url_parameter_value
            ref_value  {
                model_type = var.workspace_task_execute_rest_call_config_config_values_config_param_values_request_payload_ref_value_model_type
                config_values {
                    config_param_values {
                        data_param {
                            string_value = var.workspace_task_execute_rest_call_config_config_values_config_param_values_request_payload_dataparam_string_value
                        }
                    }
                }
            }
        }
      }
    }
    model_type      = var.workspace_task_execute_rest_call_config_model_type
    method_type     = var.workspace_task_execute_rest_call_config_method_type
    request_headers = var.workspace_task_execute_rest_call_config_request_headers
  }
  op_config_values {

    #Optional
    config_param_values {
       #Optional
       key = var.workspace_task_op_config_values_config_param_value_key
       config_param_value {
          string_value      = var.workspace_task_op_config_values_config_param_value_string_value
       }
    }

  }
  parameters {
    #Required
    model_type = var.workspace_task_parameters_model_type

    #Optional
    config_values {

      #Optional
      config_param_values {
        key = var.workspace_task_parameters_config_values_config_key
        config_param_value {
          int_value      = var.workspace_task_parameters_config_values_config_param_values_int_value
        }
      }

    }
    default_value           = var.workspace_task_parameters_default_value
    description             = var.workspace_task_parameters_description
    is_input                = var.workspace_task_parameters_is_input
    is_output               = var.workspace_task_parameters_is_output
    name                    = var.workspace_task_parameters_name
    type                    = var.workspace_task_parameters_type
    type_name               = var.workspace_task_parameters_type_name
  }

  poll_rest_call_config {

    #Optional
    config_values {

      #Optional
      config_param_values {
        poll_max_duration {
            object_value      = var.workspace_task_poll_rest_call_config_config_values_config_param_values_poll_max_duration
        }
        poll_max_duration_unit {
            string_value      = var.workspace_task_poll_rest_call_config_config_values_config_param_values_poll_max_duration_unit
        }
        poll_interval {
            object_value      = var.workspace_task_poll_rest_call_config_config_values_config_param_values_poll_interval
        }
        poll_interval_unit {
            string_value      = var.workspace_task_poll_rest_call_config_config_values_config_param_values_poll_interval_unit
        }
        request_url {
            string_value      = var.workspace_task_poll_rest_call_config_config_values_config_param_values_request_url
        }
        poll_condition {
            ref_value  {
                model_type  = var.workspace_task_poll_rest_call_config_config_values_config_param_values_poll_condition_model_type
                expr_string = var.workspace_task_poll_rest_call_config_config_values_config_param_values_poll_condition_expr_string
            }
        }
      }
    }
    method_type     = var.workspace_task_poll_rest_call_config_method_type
    request_headers = var.workspace_task_poll_rest_call_config_request_headers
  }
  typed_expressions {

    #Optional
    config_values {

      #Optional
      config_param_values {
        length {
            int_value      = var.workspace_task_typed_expressions_config_values_config_param_values_int_value
        }
      }
    }
    expression    = var.workspace_task_typed_expressions_expression
    model_type    = var.workspace_task_typed_expressions_model_type
    name          = var.workspace_task_typed_expressions_name
    type          = var.workspace_task_typed_expressions_type
  }
}

data "oci_dataintegration_workspace_tasks" "test_workspace_tasks" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  fields     = var.workspace_task_fields
  folder_id  = oci_dataintegration_workspace_project.test_workspace_project.key
  identifier = var.workspace_task_identifiers
  name       = var.workspace_task_name
  type       = var.workspace_task_type
}
