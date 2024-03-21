// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataintegrationWorkspaceTaskRequiredOnlyResource = DataintegrationWorkspaceTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Required, acctest.Create, DataintegrationWorkspaceTaskRepresentation)

	DataintegrationWorkspaceTaskResourceConfig = DataintegrationWorkspaceTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Optional, acctest.Update, DataintegrationWorkspaceTaskRepresentation)

	DataintegrationWorkspaceTaskSingularDataSourceRepresentation = map[string]interface{}{
		"key":               acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_task.test_workspace_task.key}`},
		"workspace_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"expand_references": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	DataintegrationWorkspaceTaskDataSourceRepresentation = map[string]interface{}{
		"workspace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"fields":       acctest.Representation{RepType: acctest.Optional, Create: []string{`metadata`}},
		"folder_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataintegration_workspace_project.test_workspace_project.key}`},
		"identifier":   acctest.Representation{RepType: acctest.Optional, Create: []string{`TERSI_TEST_REST_TASK`}},
		"name":         acctest.Representation{RepType: acctest.Optional, Create: `TERSI_TEST_REST_TASK`},
		"type":         acctest.Representation{RepType: acctest.Optional, Create: []string{`REST_TASK`}},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceTaskDataSourceFilterRepresentation}}
	DataintegrationWorkspaceTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_task.test_workspace_task.name}`}},
	}
	DataintegrationWorkspaceTaskRepresentation = map[string]interface{}{
		"identifier":               acctest.Representation{RepType: acctest.Required, Create: `TERSI_TEST_REST_TASK`},
		"model_type":               acctest.Representation{RepType: acctest.Required, Create: `REST_TASK`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `TERSI_TEST_REST_TASK`},
		"registry_metadata":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceTaskRegistryMetadataRepresentation},
		"workspace_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"api_call_mode":            acctest.Representation{RepType: acctest.Optional, Create: `ASYNC_GENERIC`, Update: `ASYNC_GENERIC`},
		"auth_config":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskAuthConfigRepresentation},
		"cancel_rest_call_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskCancelRestCallConfigRepresentation},
		"config_provider_delegate": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskConfigProviderDelegateRepresentation},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"execute_rest_call_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRepresentation},
		"op_config_values":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskOpConfigValuesRepresentation},
		"parameters":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskParametersRepresentation},
		"poll_rest_call_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigRepresentation},
		"typed_expressions":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskTypedExpressionsRepresentation},
	}
	DataintegrationWorkspaceTaskRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_project.test_workspace_project.key}`, Update: `${oci_dataintegration_workspace_project.test_workspace_project.key}`},
		"is_favorite":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":         acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels`}},
	}
	DataintegrationWorkspaceTaskAuthConfigRepresentation = map[string]interface{}{
		"model_type":                acctest.Representation{RepType: acctest.Optional, Create: `OCI_RESOURCE_AUTH_CONFIG`},
		"resource_principal_source": acctest.Representation{RepType: acctest.Optional, Create: `WORKSPACE`, Update: `WORKSPACE`},
	}
	DataintegrationWorkspaceTaskCancelRestCallConfigRepresentation = map[string]interface{}{
		"model_type":      acctest.Representation{RepType: acctest.Optional, Create: `CANCEL_REST_CALL_CONFIG`},
		"config_values":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskCancelRestCallConfigConfigValuesRepresentation},
		"method_type":     acctest.Representation{RepType: acctest.Optional, Create: `DELETE`, Update: `DELETE`},
		"request_headers": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Content-Type": "application/json"}, Update: map[string]string{"Content-Type": "application/json"}},
	}
	DataintegrationWorkspaceTaskConfigProviderDelegateRepresentation = map[string]interface{}{
		"bindings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskConfigProviderDelegateBindingsRepresentation},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigRepresentation = map[string]interface{}{
		"model_type":      acctest.Representation{RepType: acctest.Optional, Create: `REST_CALL_CONFIG`},
		"config_values":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigConfigValuesRepresentation},
		"method_type":     acctest.Representation{RepType: acctest.Optional, Create: `POST`, Update: `POST`},
		"request_headers": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Content-Type": "application/json"}},
	}
	DataintegrationWorkspaceTaskOpConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskOpConfigValuesConfigParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskParametersRepresentation = map[string]interface{}{
		"model_type":    acctest.Representation{RepType: acctest.Required, Create: `PARAMETER`},
		"config_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskParametersConfigValuesRepresentation},
		"default_value": acctest.Representation{RepType: acctest.Optional, Create: `1234`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_input":      acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_output":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `WORKSPACE_ID`},
		"type":          acctest.Representation{RepType: acctest.Optional, Create: `Seeded:/typeSystems/PLATFORM/dataTypes/STRING`},
		"type_name":     acctest.Representation{RepType: acctest.Optional, Create: `STRING`},
	}

	DataintegrationWorkspaceTaskPollRestCallConfigRepresentation = map[string]interface{}{
		"model_type":      acctest.Representation{RepType: acctest.Optional, Create: `POLL_REST_CALL_CONFIG`},
		"config_values":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesRepresentation},
		"method_type":     acctest.Representation{RepType: acctest.Optional, Create: `GET`},
		"request_headers": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Content-Type": "application/json"}, Update: map[string]string{"Content-Type": "application/json"}},
	}
	DataintegrationWorkspaceTaskTypedExpressionsRepresentation = map[string]interface{}{
		"config_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskTypedExpressionsConfigValuesRepresentation},
		"expression":    acctest.Representation{RepType: acctest.Optional, Create: `CAST(json_path(SYS.RESPONSE_PAYLOAD, 'key') AS String)`, Update: `CAST(json_path(SYS.RESPONSE_PAYLOAD, 'key') AS String)`},
		"model_type":    acctest.Representation{RepType: acctest.Optional, Create: `TYPED_EXPRESSION`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `PROJECT_KEY`},
		"type":          acctest.Representation{RepType: acctest.Optional, Create: `Seeded:/typeSystems/PLATFORM/dataTypes/STRING`},
	}
	DataintegrationWorkspaceTaskCancelRestCallConfigConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskCancelRestCallConfigConfigValuesConfigParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskConfigProviderDelegateBindingsRepresentation = map[string]interface{}{
		"key":              acctest.Representation{RepType: acctest.Optional, Create: `PARAMETER_20230920_094229`},
		"parameter_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskConfigProviderDelegateBindingRepresentation},
	}
	DataintegrationWorkspaceTaskConfigProviderDelegateBindingRepresentation = map[string]interface{}{
		"simple_value": acctest.Representation{RepType: acctest.Optional, Create: `12`},
	}

	DataintegrationWorkspaceTaskExecuteRestCallConfigConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigConfigValuesConfigParamValuesRepresentation},
	}

	DataintegrationWorkspaceTaskParametersConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskParametersConfigValuesConfigParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadConfigParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesConfigParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskTypedExpressionsConfigValuesRepresentation = map[string]interface{}{
		"config_param_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskTypedExpressionsConfigValuesConfigParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskCancelRestCallConfigConfigValuesConfigParamValuesRepresentation = map[string]interface{}{
		"request_url":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskCancelRestCallConfigRequestUrlValuesRepresentation},
		"request_payload": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskCancelRestCallConfigRequestPayloadValuesRepresentation},
	}
	DataintegrationWorkspaceTaskCancelRestCallConfigRequestUrlValuesRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `http://den03cyq.us.oracle.com:8086/20200430/workspaces/`},
	}
	DataintegrationWorkspaceTaskCancelRestCallConfigRequestPayloadValuesRepresentation = map[string]interface{}{
		"ref_value": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskCancelRestCallConfigRequestPayloadRefValuesRepresentation},
	}
	DataintegrationWorkspaceTaskCancelRestCallConfigRequestPayloadRefValuesRepresentation = map[string]interface{}{
		"model_type":    acctest.Representation{RepType: acctest.Required, Create: `JSON_TEXT`},
		"config_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadConfigValuesRepresentation},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigConfigValuesConfigParamValuesRepresentation = map[string]interface{}{
		"request_url":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestUrlValuesRepresentation},
		"request_payload": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadValuesRepresentation},
	}

	DataintegrationWorkspaceTaskExecuteRestCallConfigRequestUrlValuesRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `http://den03cyq.us.oracle.com:8086/20200430/workspaces/`},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadValuesRepresentation = map[string]interface{}{
		"ref_value": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadRefValuesRepresentation},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadRefValuesRepresentation = map[string]interface{}{
		"model_type":    acctest.Representation{RepType: acctest.Required, Create: `JSON_TEXT`},
		"config_values": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadConfigValuesRepresentation},
	}
	DataintegrationWorkspaceTaskParametersConfigValuesConfigParamValuesRepresentation = map[string]interface{}{
		"key":                acctest.Representation{RepType: acctest.Optional, Create: `length`, Update: `length`},
		"config_param_value": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskParameterConfigValuesConfigParamValueRepresentation},
	}

	DataintegrationWorkspaceTaskTypedExpressionsConfigValuesConfigParamValuesRepresentation = map[string]interface{}{
		"length": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskTypedExpressionsConfigValuesConfigParamValuesLengthRepresentation},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadConfigParamValuesRepresentation = map[string]interface{}{
		"data_param": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadDataParamValuesRepresentation},
	}
	DataintegrationWorkspaceTaskTypedExpressionsConfigValuesConfigParamValuesLengthRepresentation = map[string]interface{}{
		"int_value": acctest.Representation{RepType: acctest.Optional, Create: `2000`},
	}
	DataintegrationWorkspaceTaskExecuteRestCallConfigRequestPayloadDataParamValuesRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `{\n    \"modelType\": \"USER_PROJECT\",\n    \"name\":\"PROJECT_NAME\",\n    \"identifier\":\"PROJECT_NAME\",\n    \"description\":\"Project created using REST task.\"\n}`},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesConfigParamValuesRepresentation = map[string]interface{}{
		"poll_max_duration":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollMaxDurationRepresentation},
		"poll_max_duration_unit": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollMaxDurationUnitRepresentation},
		"poll_interval":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollIntervalRepresentation},
		"poll_interval_unit":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollIntervalUnitRepresentation},
		"request_url":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesRequestURLRepresentation},
		"poll_condition":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollConditionRepresentation},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollMaxDurationRepresentation = map[string]interface{}{
		"object_value": acctest.Representation{RepType: acctest.Optional, Create: `140`, Update: `140`},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollMaxDurationUnitRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `MINUTES`, Update: `MINUTES`},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollIntervalRepresentation = map[string]interface{}{
		"object_value": acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `2`},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollIntervalUnitRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `MINUTES`, Update: `MINUTES`},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesRequestURLRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `http://den03cyq.us.oracle.com:8086/20200430/workspaces/WORKSPACE_ID/projects/PROJECT_KEY`, Update: `http://den03cyq.us.oracle.com:8086/20200430/workspaces/WORKSPACE_ID/projects/PROJECT_KEY`},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollConditionRepresentation = map[string]interface{}{
		"ref_value": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollConditionRefValueRepresentation},
	}
	DataintegrationWorkspaceTaskPollRestCallConfigConfigValuesPollConditionRefValueRepresentation = map[string]interface{}{
		"model_type":  acctest.Representation{RepType: acctest.Optional, Create: `EXPRESSION`},
		"expr_string": acctest.Representation{RepType: acctest.Optional, Create: `CAST(json_path(SYS.RESPONSE_PAYLOAD, 'name') AS String) != 'PROJECT_TEST'`},
	}
	DataintegrationWorkspaceTaskOpConfigValuesConfigParamValuesRepresentation = map[string]interface{}{
		"key":                acctest.Representation{RepType: acctest.Optional, Create: `successCondition`, Update: `successCondition`},
		"config_param_value": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceTaskOpConfigValuesConfigParamValueRepresentation},
	}

	DataintegrationWorkspaceTaskParameterConfigValuesConfigParamValueRepresentation = map[string]interface{}{
		"int_value": acctest.Representation{RepType: acctest.Optional, Create: `100`, Update: `100`},
	}

	DataintegrationWorkspaceTaskOpConfigValuesConfigParamValueRepresentation = map[string]interface{}{
		"string_value": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DataintegrationWorkspaceTaskRefValueRepresentation = map[string]interface{}{
		"model_type": acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_ADWC_DATA_ASSET`},
		"name":       acctest.Representation{RepType: acctest.Optional, Create: `AA_BICC`},
	}

	DataintegrationWorkspaceTaskResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, DataintegrationWorkspaceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Required, acctest.Create, DataintegrationWorkspaceProjectRepresentation)
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceTaskResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dataintegration_workspace_task.test_workspace_task"
	datasourceName := "data.oci_dataintegration_workspace_tasks.test_workspace_tasks"
	singularDatasourceName := "data.oci_dataintegration_workspace_task.test_workspace_task"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataintegrationWorkspaceTaskResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Optional, acctest.Create, DataintegrationWorkspaceTaskRepresentation), "dataintegration", "workspaceTask", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceTaskDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Required, acctest.Create, DataintegrationWorkspaceTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "name", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceTaskResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Optional, acctest.Create, DataintegrationWorkspaceTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "api_call_mode", "ASYNC_GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auth_config.0.model_type", "OCI_RESOURCE_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "auth_config.0.resource_principal_source", "WORKSPACE"),
				resource.TestCheckResourceAttr(resourceName, "config_provider_delegate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_provider_delegate.0.bindings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.method_type", "POST"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "name", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "op_config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "op_config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.0.parent_ref.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.default_value", "1234"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_input", "true"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_output", "false"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.model_type", "PARAMETER"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "WORKSPACE_ID"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.parent_ref.#", "1"),
				//TODO resource.TestCheckResourceAttr(resourceName, "parameters.0.root_object_default_value", "{\"dummyKey\": \"dummyValue\"}"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.type", "Seeded:/typeSystems/PLATFORM/dataTypes/STRING"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.type_name", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Optional, acctest.Update, DataintegrationWorkspaceTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "api_call_mode", "ASYNC_GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auth_config.0.model_type", "OCI_RESOURCE_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "auth_config.0.resource_principal_source", "WORKSPACE"),
				resource.TestCheckResourceAttr(resourceName, "cancel_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cancel_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cancel_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cancel_rest_call_config.0.method_type", "DELETE"),
				resource.TestCheckResourceAttr(resourceName, "cancel_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_provider_delegate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.method_type", "POST"),
				resource.TestCheckResourceAttr(resourceName, "execute_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "name", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttr(resourceName, "op_config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "op_config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.default_value", "1234"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_input", "false"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_output", "true"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.model_type", "PARAMETER"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "WORKSPACE_ID"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.type_name", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.method_type", "GET"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.expression", "CAST(json_path(SYS.RESPONSE_PAYLOAD, 'key') AS String)"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.name", "PROJECT_KEY"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.type", "Seeded:/typeSystems/PLATFORM/dataTypes/STRING"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_tasks", "test_workspace_tasks", acctest.Optional, acctest.Update, DataintegrationWorkspaceTaskDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Optional, acctest.Update, DataintegrationWorkspaceTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "fields.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "folder_id"),
				resource.TestCheckResourceAttr(datasourceName, "identifier.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "task_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "task_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_task", "test_workspace_task", acctest.Required, acctest.Create, DataintegrationWorkspaceTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "expand_references"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "api_call_mode", "ASYNC_GENERIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auth_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auth_config.0.model_type", "OCI_RESOURCE_AUTH_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auth_config.0.resource_principal_source", "WORKSPACE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cancel_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cancel_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cancel_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cancel_rest_call_config.0.method_type", "DELETE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cancel_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_provider_delegate.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_rest_call_config.0.config_values.0.parent_ref.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_rest_call_config.0.method_type", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "REST_TASK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TERSI_TEST_REST_TASK"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_version"),
				resource.TestCheckResourceAttr(resourceName, "op_config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "op_config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.default_value", "1234"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_input", "false"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_output", "true"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.model_type", "PARAMETER"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "WORKSPACE_ID"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.type_name", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.method_type", "GET"),
				resource.TestCheckResourceAttr(resourceName, "poll_rest_call_config.0.request_headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.config_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.config_values.0.config_param_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.expression", "CAST(json_path(SYS.RESPONSE_PAYLOAD, 'key') AS String)"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.name", "PROJECT_KEY"),
				resource.TestCheckResourceAttr(resourceName, "typed_expressions.0.type", "Seeded:/typeSystems/PLATFORM/dataTypes/STRING"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataintegrationWorkspaceTaskRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceTaskDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_task" {
			noResourceFound = false
			request := oci_dataintegration.GetTaskRequest{}

			if value, ok := rs.Primary.Attributes["expand_references"]; ok {
				request.ExpandReferences = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.TaskKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetTask(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceTask") {
		resource.AddTestSweepers("DataintegrationWorkspaceTask", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceTask",
			Dependencies: acctest.DependencyGraph["workspaceTask"],
			F:            sweepDataintegrationWorkspaceTaskResource,
		})
	}
}

func sweepDataintegrationWorkspaceTaskResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceTaskIds, err := getDataintegrationWorkspaceTaskIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceTaskId := range workspaceTaskIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceTaskId]; !ok {
			deleteTaskRequest := oci_dataintegration.DeleteTaskRequest{}

			deleteTaskRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteTask(context.Background(), deleteTaskRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceTask %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceTaskId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceTaskIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceTaskId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listTasksRequest := oci_dataintegration.ListTasksRequest{}

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceTask resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listTasksRequest.WorkspaceId = &workspaceId

		listTasksResponse, err := dataIntegrationClient.ListTasks(context.Background(), listTasksRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceTask list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceTask := range listTasksResponse.Items {
			id := *workspaceTask.GetKey()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceTaskId", id)
		}

	}
	return resourceIds, nil
}
