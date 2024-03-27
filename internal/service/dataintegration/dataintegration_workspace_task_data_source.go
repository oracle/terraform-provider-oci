// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceTaskDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["expand_references"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["workspace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataintegrationWorkspaceTaskResource(), fieldMap, readSingularDataintegrationWorkspaceTask)
}

func readSingularDataintegrationWorkspaceTask(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceTaskDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceTaskDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.GetTaskResponse
}

func (s *DataintegrationWorkspaceTaskDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceTaskDataSourceCrud) Get() error {
	request := oci_dataintegration.GetTaskRequest{}

	if expandReferences, ok := s.D.GetOkExists("expand_references"); ok {
		tmp := expandReferences.(string)
		request.ExpandReferences = &tmp
	}

	if taskKey, ok := s.D.GetOkExists("key"); ok {
		tmp := taskKey.(string)
		request.TaskKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.GetTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataintegrationWorkspaceTaskDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceTaskDataSource-", DataintegrationWorkspaceTaskDataSource(), s.D))
	switch v := (s.Res.Task).(type) {
	case oci_dataintegration.TaskFromDataLoaderTaskDetails:
		s.D.Set("model_type", "DATA_LOADER_TASK")

		if v.ConditionalCompositeFieldMap != nil {
			s.D.Set("conditional_composite_field_map", []interface{}{ConditionalCompositeFieldMapToMap(v.ConditionalCompositeFieldMap)})
		} else {
			s.D.Set("conditional_composite_field_map", nil)
		}

		if v.DataFlow != nil {
			s.D.Set("data_flow", []interface{}{DataFlowToMap(v.DataFlow)})
		} else {
			s.D.Set("data_flow", nil)
		}

		if v.IsSingleLoad != nil {
			s.D.Set("is_single_load", *v.IsSingleLoad)
		}

		if v.ParallelLoadLimit != nil {
			s.D.Set("parallel_load_limit", *v.ParallelLoadLimit)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{ObjectMetadataToMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

		if v.RegistryMetadata != nil {
			s.D.Set("registry_metadata", []interface{}{DataIntegration_Task_RegistryMetadataToMap(v.RegistryMetadata)})
		} else {
			s.D.Set("registry_metadata", nil)
		}
	case oci_dataintegration.TaskFromIntegrationTaskDetails:
		s.D.Set("model_type", "INTEGRATION_TASK")

		if v.DataFlow != nil {
			s.D.Set("data_flow", []interface{}{DataFlowToMap(v.DataFlow)})
		} else {
			s.D.Set("data_flow", nil)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{ObjectMetadataToMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

		if v.RegistryMetadata != nil {
			s.D.Set("registry_metadata", []interface{}{DataIntegration_Task_RegistryMetadataToMap(v.RegistryMetadata)})
		} else {
			s.D.Set("registry_metadata", nil)
		}
	case oci_dataintegration.TaskFromOciDataflowTaskDetails:
		s.D.Set("model_type", "OCI_DATAFLOW_TASK")

		if v.DataflowApplication != nil {
			s.D.Set("dataflow_application", []interface{}{DataflowApplicationToMap(v.DataflowApplication)})
		} else {
			s.D.Set("dataflow_application", nil)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{ObjectMetadataToMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

		if v.RegistryMetadata != nil {
			s.D.Set("registry_metadata", []interface{}{DataIntegration_Task_RegistryMetadataToMap(v.RegistryMetadata)})
		} else {
			s.D.Set("registry_metadata", nil)
		}
	case oci_dataintegration.TaskFromPipelineTaskDetails:
		s.D.Set("model_type", "PIPELINE_TASK")

		if v.Pipeline != nil {
			s.D.Set("pipeline", []interface{}{PipelineToMap(v.Pipeline)})
		} else {
			s.D.Set("pipeline", nil)
		}

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{ObjectMetadataToMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

		if v.RegistryMetadata != nil {
			s.D.Set("registry_metadata", []interface{}{DataIntegration_Task_RegistryMetadataToMap(v.RegistryMetadata)})
		} else {
			s.D.Set("registry_metadata", nil)
		}
	case oci_dataintegration.TaskFromRestTaskDetails:
		s.D.Set("model_type", "REST_TASK")

		s.D.Set("api_call_mode", v.ApiCallMode)

		if v.AuthConfig != nil {
			authConfigArray := []interface{}{}
			if authConfigMap := AuthConfigToMap(&v.AuthConfig); authConfigMap != nil {
				authConfigArray = append(authConfigArray, authConfigMap)
			}
			s.D.Set("auth_config", authConfigArray)
		} else {
			s.D.Set("auth_config", nil)
		}

		if v.CancelRestCallConfig != nil {
			s.D.Set("cancel_rest_call_config", []interface{}{CancelRestCallConfigToMap(v.CancelRestCallConfig)})
		} else {
			s.D.Set("cancel_rest_call_config", nil)
		}

		if v.ExecuteRestCallConfig != nil {
			s.D.Set("execute_rest_call_config", []interface{}{ExecuteRestCallConfigToMap(v.ExecuteRestCallConfig)})
		} else {
			s.D.Set("execute_rest_call_config", nil)
		}

		if v.PollRestCallConfig != nil {
			s.D.Set("poll_rest_call_config", []interface{}{PollRestCallConfigToMap(v.PollRestCallConfig)})
		} else {
			s.D.Set("poll_rest_call_config", nil)
		}

		typedExpressions := []interface{}{}
		for _, item := range v.TypedExpressions {
			typedExpressions = append(typedExpressions, TypedExpressionToMap(item))
		}
		s.D.Set("typed_expressions", typedExpressions)

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{DataIntegration_Task_ObjectMetadataToMap(v.Metadata)})
			s.D.Set("registry_metadata", []interface{}{DataintegrationTaskObjectMetadataToRegistryMetadataMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
			s.D.Set("registry_metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

	case oci_dataintegration.TaskFromSqlTaskDetails:
		s.D.Set("model_type", "SQL_TASK")

		if v.Operation != nil {
			s.D.Set("operation", v.Operation)
		} else {
			s.D.Set("operation", nil)
		}

		if v.Script != nil {
			s.D.Set("script", []interface{}{ScriptToMap(v.Script)})
		} else {
			s.D.Set("script", nil)
		}

		s.D.Set("sql_script_type", v.SqlScriptType)

		if v.ConfigProviderDelegate != nil {
			s.D.Set("config_provider_delegate", []interface{}{ConfigProviderToMap(v.ConfigProviderDelegate)})
		} else {
			s.D.Set("config_provider_delegate", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Identifier != nil {
			s.D.Set("identifier", *v.Identifier)
		}

		inputPorts := []interface{}{}
		for _, item := range v.InputPorts {
			inputPorts = append(inputPorts, InputPortToMap(item))
		}
		s.D.Set("input_ports", inputPorts)

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		s.D.Set("key_map", v.KeyMap)
		//s.D.Set("key_map", v.KeyMap)

		if v.Metadata != nil {
			s.D.Set("metadata", []interface{}{ObjectMetadataToMap(v.Metadata)})
		} else {
			s.D.Set("metadata", nil)
		}

		if v.ModelVersion != nil {
			s.D.Set("model_version", *v.ModelVersion)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ObjectStatus != nil {
			s.D.Set("object_status", *v.ObjectStatus)
		}

		if v.ObjectVersion != nil {
			s.D.Set("object_version", *v.ObjectVersion)
		}

		if v.OpConfigValues != nil {
			s.D.Set("op_config_values", []interface{}{ConfigValuesToMap(v.OpConfigValues)})
		} else {
			s.D.Set("op_config_values", nil)
		}

		outputPorts := []interface{}{}
		for _, item := range v.OutputPorts {
			outputPorts = append(outputPorts, OutputPortToMap(item))
		}
		s.D.Set("output_ports", outputPorts)

		parameters := []interface{}{}
		for _, item := range v.Parameters {
			parameters = append(parameters, ParameterToMap(item))
		}
		s.D.Set("parameters", parameters)

		if v.ParentRef != nil {
			s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(v.ParentRef)})
		} else {
			s.D.Set("parent_ref", nil)
		}

		if v.RegistryMetadata != nil {
			s.D.Set("registry_metadata", []interface{}{DataIntegration_Task_RegistryMetadataToMap(v.RegistryMetadata)})
		} else {
			s.D.Set("registry_metadata", nil)
		}
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", s.Res.Task)
		return nil
	}

	return nil
}
