// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentToolRagRequiredOnlyResource             = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentRagToolRepresentation)
	GenerativeAiAgentToolSqlRequiredOnlyResource             = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentSqlToolRepresentation)
	GenerativeAiAgentToolFunctionCallingRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentFunctionCallingToolRepresentation)
	GenerativeAiAgentToolHttpEndpointRequiredOnlyResource    = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentHttpEndpointToolRepresentation)
	GenerativeAiAgentToolAgentRequiredOnlyResource           = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentAgentToolRepresentation)

	GenerativeAiAgentToolRagResourceConfig             = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentRagToolRepresentation)
	GenerativeAiAgentToolSqlResourceConfig             = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentSqlToolRepresentation)
	GenerativeAiAgentToolFunctionCallingResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentFunctionCallingToolRepresentation)
	GenerativeAiAgentToolHttpEndpointResourceConfig    = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentHttpEndpointToolRepresentation)
	GenerativeAiAgentToolAgentResourceConfig           = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentAgentToolRepresentation)

	GenerativeAiAgentToolSingularDataSourceRepresentation = map[string]interface{}{
		"tool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_tool.test_tool.id}`},
	}

	GenerativeAiAgentToolDataSourceRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolDataSourceFilterRepresentation}}
	GenerativeAiAgentToolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_tool.test_tool.id}`}},
	}

	GenerativeAiAgentRagToolRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"tool_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolRagToolConfigRepresentation},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiAgentSqlToolRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"tool_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolSqlToolConfigRepresentation},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiAgentFunctionCallingToolRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"tool_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolFunctionCallingToolConfigRepresentation},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiAgentHttpEndpointToolRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"tool_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentTooHttpEndpointToolConfigRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiAgentAgentToolRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"tool_config":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentAgentToolConfigRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	GenerativeAiAgentToolRagToolConfigRepresentation = map[string]interface{}{
		"tool_config_type":             acctest.Representation{RepType: acctest.Required, Create: `RAG_TOOL_CONFIG`},
		"generation_llm_customization": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigGenerationLlmCustomizationRepresentation},
		"knowledge_base_configs":       acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigKnowledgeBaseConfigsRepresentation},
	}
	GenerativeAiAgentToolFunctionCallingToolConfigRepresentation = map[string]interface{}{
		"tool_config_type": acctest.Representation{RepType: acctest.Required, Create: `FUNCTION_CALLING_TOOL_CONFIG`},
		"function":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigFunctionRepresentation},
	}
	GenerativeAiAgentTooHttpEndpointToolConfigRepresentation = map[string]interface{}{
		"tool_config_type":          acctest.Representation{RepType: acctest.Required, Create: `HTTP_ENDPOINT_TOOL_CONFIG`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"api_schema":                acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigApiSchemaRepresentation},
		"http_endpoint_auth_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigAuthConfigRepresentation},
	}
	GenerativeAiAgentAgentToolConfigRepresentation = map[string]interface{}{
		"tool_config_type":  acctest.Representation{RepType: acctest.Required, Create: `AGENT_TOOL_CONFIG`},
		"agent_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.agent_endpoint_id}`, Update: `${var.agent_endpoint_id_for_update}`},
	}
	GenerativeAiAgentToolToolConfigAuthConfigRepresentation = map[string]interface{}{
		"http_endpoint_auth_sources": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigAuthSourcesRepresentation},
	}
	GenerativeAiAgentToolToolConfigAuthSourcesRepresentation = map[string]interface{}{
		"http_endpoint_auth_scope":        acctest.Representation{RepType: acctest.Required, Create: `AGENT`},
		"http_endpoint_auth_scope_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigAuthScopeConfigRepresentation},
	}
	GenerativeAiAgentToolToolConfigAuthScopeConfigRepresentation = map[string]interface{}{
		"http_endpoint_auth_scope_config_type": acctest.Representation{RepType: acctest.Required, Create: `HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG`},
	}
	GenerativeAiAgentToolSqlToolConfigRepresentation = map[string]interface{}{
		"tool_config_type":              acctest.Representation{RepType: acctest.Required, Create: `SQL_TOOL_CONFIG`},
		"database_schema":               acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentToolToolConfigDatabaseInlineSchemaRepresentation},
		"dialect":                       acctest.Representation{RepType: acctest.Required, Create: `SQL_LITE`, Update: `ORACLE_SQL`},
		"model_size":                    acctest.Representation{RepType: acctest.Optional, Create: `SMALL`},
		"should_enable_self_correction": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"should_enable_sql_execution":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	GenerativeAiAgentToolToolConfigGenerationLlmCustomizationRepresentation = map[string]interface{}{
		"instruction": acctest.Representation{RepType: acctest.Optional, Create: `instruction`, Update: `instruction2`},
	}
	GenerativeAiAgentToolToolConfigKnowledgeBaseConfigsRepresentation = map[string]interface{}{
		"knowledge_base_id": acctest.Representation{RepType: acctest.Required, Create: `${var.knowledge_base_id}`},
	}
	GenerativeAiAgentToolToolConfigDatabaseInlineSchemaRepresentation = map[string]interface{}{
		"input_location_type": acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"content":             acctest.Representation{RepType: acctest.Required, Create: `CREATE TABLE example ();`, Update: `CREATE TABLE example2 ();`},
	}
	//GenerativeAiAgentToolToolConfigDatabaseOssSchemaRepresentation = map[string]interface{}{
	//	"input_location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_PREFIX`},
	//	"bucket":              acctest.Representation{RepType: acctest.Optional, Create: `bucket`, Update: `bucket2`},
	//	"namespace":           acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
	//	"prefix":              acctest.Representation{RepType: acctest.Optional, Create: `prefix`, Update: `prefix2`},
	//}
	GenerativeAiAgentToolToolConfigFunctionRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"parameters":  acctest.Representation{RepType: acctest.Required, Create: map[string]string{"parameters": "parameters"}, Update: map[string]string{"parameters2": "parameters2"}},
	}
	GenerativeAiAgentToolToolConfigApiSchemaRepresentation = map[string]interface{}{
		"api_schema_input_location_type": acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"content":                        acctest.Representation{RepType: acctest.Required, Create: `{\"openapi\": \"3.0.0\",\"info\": {\"title\": \"Minimal API\",\"version\": \"1.0\"},\"servers\": [{\"url\": \"https://example.com/api\"}],\"paths\": {\"/ping\": {\"get\": {\"summary\": \"Ping for health check\",\"responses\": {\"200\": {\"description\": \"OK\"}}}}}}`, Update: `{\"openapi\": \"3.0.0\",\"info\": {\"title\": \"Minimal API\",\"version\": \"1.0\"},\"servers\": [{\"url\": \"https://updated.com/api\"}],\"paths\": {\"/ping\": {\"get\": {\"summary\": \"Ping for health check\",\"responses\": {\"200\": {\"description\": \"OK\"}}}}}}`},
	}
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentToolResource_rag(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentToolResource_rag")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// To set the agent id for creating agent endpoint add TF_VAR env var for agent_id
	agentId := utils.GetEnvSettingWithBlankDefault("agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)
	knowledgeBaseId := utils.GetEnvSettingWithBlankDefault("knowledge_base_id")
	knowledgeBaseIdVariableStr := fmt.Sprintf("variable \"knowledge_base_id\" { default = \"%s\" }\n", knowledgeBaseId)

	resourceName := "oci_generative_ai_agent_tool.test_tool"
	datasourceName := "data.oci_generative_ai_agent_tools.test_tools"
	singularDatasourceName := "data.oci_generative_ai_agent_tool.test_tool"

	// dependencies: agentIdVariableStr

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+knowledgeBaseIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentRagToolRepresentation), "generativeaiagent", "tool", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentToolDestroy, []resource.TestStep{
		// verify RAG tool create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentRagToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.knowledge_base_configs.0.knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "RAG_TOOL_CONFIG"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete tool before next create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr,
		},
		// verify RAG tool create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentRagToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.instruction", "instruction"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.knowledge_base_configs.0.knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "RAG_TOOL_CONFIG"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify RAG tool updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentRagToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.knowledge_base_configs.0.knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "RAG_TOOL_CONFIG"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify RAG tool datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tools", "test_tools", acctest.Optional, acctest.Update, GenerativeAiAgentToolDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentRagToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tool_collection.0.items.#", "1"),
			),
		},
		// verify RAG tool singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentToolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr + GenerativeAiAgentToolRagResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tool_config.0.knowledge_base_configs.0.knowledge_base_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.tool_config_type", "RAG_TOOL_CONFIG"),
			),
		},
		// verify RAG resource import
		{
			Config:                  config + GenerativeAiAgentToolRagRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
func TestGenerativeAiAgentToolResource_sql(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentToolResource_sql")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// To set the agent id for creating agent endpoint add TF_VAR env var for agent_id
	agentId := utils.GetEnvSettingWithBlankDefault("agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)
	knowledgeBaseId := utils.GetEnvSettingWithBlankDefault("knowledge_base_id")
	knowledgeBaseIdVariableStr := fmt.Sprintf("variable \"knowledge_base_id\" { default = \"%s\" }\n", knowledgeBaseId)

	resourceName := "oci_generative_ai_agent_tool.test_tool"
	datasourceName := "data.oci_generative_ai_agent_tools.test_tools"
	singularDatasourceName := "data.oci_generative_ai_agent_tool.test_tool"

	// dependencies: agentIdVariableStr

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.

	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+knowledgeBaseIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentSqlToolRepresentation), "generativeaiagent", "tool", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentToolDestroy, []resource.TestStep{
		// verify SQL tool create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "SQL_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.dialect", "SQL_LITE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete tool before next create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr,
		},
		// verify SQL tool create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "SQL_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.content", "CREATE TABLE example ();"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.model_size", "SMALL"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_self_correction", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_sql_execution", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.dialect", "SQL_LITE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify SQL tool updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "SQL_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.content", "CREATE TABLE example2 ();"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.model_size", "SMALL"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_self_correction", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_sql_execution", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.dialect", "ORACLE_SQL"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify SQL tool datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tools", "test_tools", acctest.Optional, acctest.Update, GenerativeAiAgentToolDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentSqlToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tool_collection.0.items.#", "1"),
			),
		},
		// verify SQL tool singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentToolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr + GenerativeAiAgentToolSqlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tool_id"),

				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "SQL_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.content", "CREATE TABLE example2 ();"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.model_size", "SMALL"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_self_correction", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_sql_execution", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.dialect", "ORACLE_SQL"),
			),
		},
		// verify SQL resource import
		{
			Config:                  config + GenerativeAiAgentToolSqlRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
func TestGenerativeAiAgentToolResource_fc(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentToolResource_fc")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// To set the agent id for creating agent endpoint add TF_VAR env var for agent_id
	agentId := utils.GetEnvSettingWithBlankDefault("agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)
	knowledgeBaseId := utils.GetEnvSettingWithBlankDefault("knowledge_base_id")
	knowledgeBaseIdVariableStr := fmt.Sprintf("variable \"knowledge_base_id\" { default = \"%s\" }\n", knowledgeBaseId)

	resourceName := "oci_generative_ai_agent_tool.test_tool"
	datasourceName := "data.oci_generative_ai_agent_tools.test_tools"
	singularDatasourceName := "data.oci_generative_ai_agent_tool.test_tool"

	// dependencies: agentIdVariableStr

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+knowledgeBaseIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentFunctionCallingToolRepresentation), "generativeaiagent", "tool", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentToolDestroy, []resource.TestStep{
		// verify FC tool create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentFunctionCallingToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "FUNCTION_CALLING_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.parameters.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete tool before next create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr,
		},
		// verify FC tool create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentFunctionCallingToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "FUNCTION_CALLING_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.parameters.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// verify FC tool updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentFunctionCallingToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "FUNCTION_CALLING_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.parameters.%", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify FC tool datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tools", "test_tools", acctest.Optional, acctest.Update, GenerativeAiAgentToolDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentFunctionCallingToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tool_collection.0.items.#", "1"),
			),
		},
		// verify FC tool singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentToolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + knowledgeBaseIdVariableStr + GenerativeAiAgentToolFunctionCallingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tool_id"),

				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "FUNCTION_CALLING_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.parameters.%", "1"),
			),
		},
		// verify FC resource import
		{
			Config:                  config + GenerativeAiAgentToolFunctionCallingRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestGenerativeAiAgentToolResource_http(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentToolResource_http")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// To set the agent id for creating agent endpoint add TF_VAR env var for agent_id
	agentId := utils.GetEnvSettingWithBlankDefault("agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)
	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	resourceName := "oci_generative_ai_agent_tool.test_tool"
	datasourceName := "data.oci_generative_ai_agent_tools.test_tools"
	singularDatasourceName := "data.oci_generative_ai_agent_tool.test_tool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+subnetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentHttpEndpointToolRepresentation), "generativeaiagent", "tool", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentToolDestroy, []resource.TestStep{
		// verify http endpoint tool create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentHttpEndpointToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "HTTP_ENDPOINT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.content", `{"openapi": "3.0.0","info": {"title": "Minimal API","version": "1.0"},"servers": [{"url": "https://example.com/api"}],"paths": {"/ping": {"get": {"summary": "Ping for health check","responses": {"200": {"description": "OK"}}}}}}`),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete tool before next create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + subnetIdVariableStr,
		},
		// verify http endpoint tool create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentHttpEndpointToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "HTTP_ENDPOINT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.content", `{"openapi": "3.0.0","info": {"title": "Minimal API","version": "1.0"},"servers": [{"url": "https://example.com/api"}],"paths": {"/ping": {"get": {"summary": "Ping for health check","responses": {"200": {"description": "OK"}}}}}}`),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify http endpoint tool updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentHttpEndpointToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "HTTP_ENDPOINT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.content", `{"openapi": "3.0.0","info": {"title": "Minimal API","version": "1.0"},"servers": [{"url": "https://updated.com/api"}],"paths": {"/ping": {"get": {"summary": "Ping for health check","responses": {"200": {"description": "OK"}}}}}}`),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify http endpoint tool datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tools", "test_tools", acctest.Optional, acctest.Update, GenerativeAiAgentToolDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentHttpEndpointToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tool_collection.0.items.#", "1"),
			),
		},
		// verify http endpoint tool singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentToolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + subnetIdVariableStr + GenerativeAiAgentToolHttpEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tool_id"),

				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "HTTP_ENDPOINT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.content", `{"openapi": "3.0.0","info": {"title": "Minimal API","version": "1.0"},"servers": [{"url": "https://updated.com/api"}],"paths": {"/ping": {"get": {"summary": "Ping for health check","responses": {"200": {"description": "OK"}}}}}}`),
			),
		},
		// verify http endpoint resource import
		{
			Config:                  config + GenerativeAiAgentToolHttpEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestGenerativeAiAgentToolResource_agent(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentToolResource_agent")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// To set the agent id for creating agent endpoint add TF_VAR env var for agent_id
	agentId := utils.GetEnvSettingWithBlankDefault("agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)
	agentEndpointId := utils.GetEnvSettingWithBlankDefault("agent_endpoint_id")
	agentEndpointIdVariableStr := fmt.Sprintf("variable \"agent_endpoint_id\" { default = \"%s\" }\n", agentEndpointId)
	agentEndpointIdForUpdate := utils.GetEnvSettingWithBlankDefault("agent_endpoint_id_for_update")
	agentEndpointIdForUpdateVariableStr := fmt.Sprintf("variable \"agent_endpoint_id_for_update\" { default = \"%s\" }\n", agentEndpointIdForUpdate)

	resourceName := "oci_generative_ai_agent_tool.test_tool"
	datasourceName := "data.oci_generative_ai_agent_tools.test_tools"
	singularDatasourceName := "data.oci_generative_ai_agent_tool.test_tool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+agentEndpointIdVariableStr+agentEndpointIdForUpdateVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentAgentToolRepresentation), "generativeaiagent", "tool", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentToolDestroy, []resource.TestStep{
		// verify agent tool create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + agentEndpointIdVariableStr + agentEndpointIdForUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentAgentToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "AGENT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.agent_endpoint_id", agentEndpointId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete tool before next create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + agentEndpointIdVariableStr + agentEndpointIdForUpdateVariableStr,
		},
		// verify http endpoint tool create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + agentEndpointIdVariableStr + agentEndpointIdForUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Create, GenerativeAiAgentAgentToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "AGENT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.agent_endpoint_id", agentEndpointId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify http endpoint tool updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + agentEndpointIdVariableStr + agentEndpointIdForUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentAgentToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "AGENT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.agent_endpoint_id", agentEndpointIdForUpdate),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify http endpoint tool datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tools", "test_tools", acctest.Optional, acctest.Update, GenerativeAiAgentToolDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + agentEndpointIdVariableStr + agentEndpointIdForUpdateVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Optional, acctest.Update, GenerativeAiAgentAgentToolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "agent_id", agentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "tool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tool_collection.0.items.#", "1"),
			),
		},
		// verify http endpoint tool singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_tool", "test_tool", acctest.Required, acctest.Create, GenerativeAiAgentToolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + agentEndpointIdVariableStr + agentEndpointIdForUpdateVariableStr + GenerativeAiAgentToolAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tool_id"),

				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "AGENT_TOOL_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.agent_endpoint_id", agentEndpointIdForUpdate),
			),
		},
		// verify http endpoint resource import
		{
			Config:                  config + GenerativeAiAgentToolAgentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentToolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_tool" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetToolRequest{}

			tmp := rs.Primary.ID
			request.ToolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetTool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.ToolLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentTool") {
		resource.AddTestSweepers("GenerativeAiAgentTool", &resource.Sweeper{
			Name:         "GenerativeAiAgentTool",
			Dependencies: acctest.DependencyGraph["tool"],
			F:            sweepGenerativeAiAgentToolResource,
		})
	}
}

func sweepGenerativeAiAgentToolResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	toolIds, err := getGenerativeAiAgentToolIds(compartment)
	if err != nil {
		return err
	}
	for _, toolId := range toolIds {
		if ok := acctest.SweeperDefaultResourceId[toolId]; !ok {
			deleteToolRequest := oci_generative_ai_agent.DeleteToolRequest{}

			deleteToolRequest.ToolId = &toolId

			deleteToolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteTool(context.Background(), deleteToolRequest)
			if error != nil {
				fmt.Printf("Error deleting Tool %s %s, It is possible that the resource is already deleted. Please verify manually \n", toolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &toolId, GenerativeAiAgentToolSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentToolSweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentToolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ToolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listToolsRequest := oci_generative_ai_agent.ListToolsRequest{}
	listToolsRequest.CompartmentId = &compartmentId
	listToolsRequest.LifecycleState = oci_generative_ai_agent.ToolLifecycleStateActive
	listToolsResponse, err := generativeAiAgentClient.ListTools(context.Background(), listToolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Tool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, tool := range listToolsResponse.Items {
		id := *tool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ToolId", id)
	}
	return resourceIds, nil
}

func GenerativeAiAgentToolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if toolResponse, ok := response.Response.(oci_generative_ai_agent.GetToolResponse); ok {
		return toolResponse.LifecycleState != oci_generative_ai_agent.ToolLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentToolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetTool(context.Background(), oci_generative_ai_agent.GetToolRequest{
		ToolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
