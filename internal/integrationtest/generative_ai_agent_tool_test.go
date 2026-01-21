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
	GenerativeAiAgentToolToolConfigRepresentation = map[string]interface{}{
		"tool_config_type":              acctest.Representation{RepType: acctest.Required, Create: `SQL_TOOL_CONFIG`, Update: `RAG_TOOL_CONFIG`},
		"agent_endpoint_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_agent_agent_endpoint.test_agent_endpoint.id}`},
		"api_schema":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigApiSchemaRepresentation},
		"dialect":                       acctest.Representation{RepType: acctest.Optional, Create: `SQL_LITE`, Update: `ORACLE_SQL`},
		"embedding_llm_customization":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigEmbeddingLlmCustomizationRepresentation},
		"function":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigFunctionRepresentation},
		"generation_llm_customization":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigGenerationLlmCustomizationRepresentation},
		"knowledge_base_configs":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigKnowledgeBaseConfigsRepresentation},
		"model_size":                    acctest.Representation{RepType: acctest.Optional, Create: `SMALL`, Update: `LARGE`},
		"reasoning_llm_customization":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigReasoningLlmCustomizationRepresentation},
		"reranking_llm_customization":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigRerankingLlmCustomizationRepresentation},
		"runtime_version":               acctest.Representation{RepType: acctest.Optional, Create: `runtimeVersion`, Update: `runtimeVersion2`},
		"should_enable_self_correction": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_enable_sql_execution":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"table_and_column_description":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigTableAndColumnDescriptionRepresentation},
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
	GenerativeAiAgentToolToolConfigEmbeddingLlmCustomizationRepresentation = map[string]interface{}{
		"instruction":          acctest.Representation{RepType: acctest.Optional, Create: `instruction`, Update: `instruction2`},
		"llm_hyper_parameters": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"llmHyperParameters": "llmHyperParameters"}, Update: map[string]string{"llmHyperParameters2": "llmHyperParameters2"}},
		"llm_selection":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigEmbeddingLlmCustomizationLlmSelectionRepresentation},
	}
	GenerativeAiAgentToolToolConfigFunctionRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"parameters":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"parameters": "parameters"}, Update: map[string]string{"parameters2": "parameters2"}},
	}
	GenerativeAiAgentToolToolConfigGenerationLlmCustomizationRepresentation = map[string]interface{}{
		"instruction":          acctest.Representation{RepType: acctest.Optional, Create: `instruction`, Update: `instruction2`},
		"llm_hyper_parameters": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"llmHyperParameters": "llmHyperParameters"}, Update: map[string]string{"llmHyperParameters2": "llmHyperParameters2"}},
		"llm_selection":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigGenerationLlmCustomizationLlmSelectionRepresentation},
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
	GenerativeAiAgentToolToolConfigKnowledgeBaseConfigsRepresentation = map[string]interface{}{
		"knowledge_base_id": acctest.Representation{RepType: acctest.Required, Create: `${var.knowledge_base_id}`},
	}
	GenerativeAiAgentToolToolConfigReasoningLlmCustomizationRepresentation = map[string]interface{}{
		"instruction":          acctest.Representation{RepType: acctest.Optional, Create: `instruction`, Update: `instruction2`},
		"llm_hyper_parameters": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"llmHyperParameters": "llmHyperParameters"}, Update: map[string]string{"llmHyperParameters2": "llmHyperParameters2"}},
		"llm_selection":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigReasoningLlmCustomizationLlmSelectionRepresentation},
	}
	GenerativeAiAgentToolToolConfigRerankingLlmCustomizationRepresentation = map[string]interface{}{
		"instruction":          acctest.Representation{RepType: acctest.Optional, Create: `instruction`, Update: `instruction2`},
		"llm_hyper_parameters": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"llmHyperParameters": "llmHyperParameters"}, Update: map[string]string{"llmHyperParameters2": "llmHyperParameters2"}},
		"llm_selection":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigRerankingLlmCustomizationLlmSelectionRepresentation},
	}
	GenerativeAiAgentToolToolConfigTableAndColumnDescriptionRepresentation = map[string]interface{}{
		"input_location_type": acctest.Representation{RepType: acctest.Required, Create: `INLINE`, Update: `OBJECT_STORAGE_PREFIX`},
		"bucket":              acctest.Representation{RepType: acctest.Optional, Create: `bucket`, Update: `bucket2`},
		"content":             acctest.Representation{RepType: acctest.Optional, Create: `content`, Update: `content2`},
		"namespace":           acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
		"prefix":              acctest.Representation{RepType: acctest.Optional, Create: `prefix`, Update: `prefix2`},
	}
	GenerativeAiAgentToolToolConfigEmbeddingLlmCustomizationLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`, Update: `CUSTOM_GEN_AI_MODEL`},
		"endpoint_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
		"model_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
	}
	GenerativeAiAgentToolToolConfigGenerationLlmCustomizationLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`, Update: `CUSTOM_GEN_AI_MODEL`},
		"endpoint_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
		"model_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
	}
	GenerativeAiAgentToolToolConfigHttpEndpointAuthConfigHttpEndpointAuthSourcesRepresentation = map[string]interface{}{
		"http_endpoint_auth_scope":        acctest.Representation{RepType: acctest.Optional, Create: `AGENT`},
		"http_endpoint_auth_scope_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiAgentToolToolConfigHttpEndpointAuthConfigHttpEndpointAuthSourcesHttpEndpointAuthScopeConfigRepresentation},
	}
	GenerativeAiAgentToolToolConfigReasoningLlmCustomizationLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`, Update: `CUSTOM_GEN_AI_MODEL`},
		"endpoint_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
		"model_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
	}
	GenerativeAiAgentToolToolConfigRerankingLlmCustomizationLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`, Update: `CUSTOM_GEN_AI_MODEL`},
		"endpoint_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
		"model_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
	}
	GenerativeAiAgentToolToolConfigHttpEndpointAuthConfigHttpEndpointAuthSourcesHttpEndpointAuthScopeConfigRepresentation = map[string]interface{}{
		"http_endpoint_auth_scope_config_type": acctest.Representation{RepType: acctest.Required, Create: `HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG`, Update: `HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG`},
		"client_id":                            acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_agent_client.test_client.id}`},
		"idcs_url":                             acctest.Representation{RepType: acctest.Optional, Create: `idcsUrl`, Update: `idcsUrl2`},
		"key_location":                         acctest.Representation{RepType: acctest.Optional, Create: `HEADER`, Update: `QUERY_PARAMETER`},
		"key_name":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key.test_key.name}`},
		"scope_url":                            acctest.Representation{RepType: acctest.Optional, Create: `scopeUrl`, Update: `scopeUrl2`},
		"vault_secret_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
	}
	GenerativeAiAgentToolResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Required, acctest.Create, AiDocumentModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_document_project", "test_project", acctest.Required, acctest.Create, AiDocumentProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Required, acctest.Create, AiLanguageEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_project", "test_project", acctest.Required, acctest.Create, AiLanguageProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, DatabaseDbHomeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, DatabaseMigrationConnectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent_endpoint", "test_agent_endpoint", acctest.Required, acctest.Create, GenerativeAiAgentAgentEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_agent", "test_agent", acctest.Required, acctest.Create, GenerativeAiAgentAgentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_provisioned_capacity", "test_provisioned_capacity", acctest.Required, acctest.Create, GenerativeAiAgentProvisionedCapacityRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
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
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.agent_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.content", "content"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.object", "object"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.database_connection.0.connection_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_connection.0.connection_type", "DATABASE_TOOL_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.content", "content"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.dialect", "SQL_LITE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.instruction", "instruction"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.llm_selection_type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.instruction", "instruction"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.llm_selection_type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope", "AGENT"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.client_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.http_endpoint_auth_scope_config_type", "HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.idcs_url", "idcsUrl"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.key_location", "HEADER"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.key_name"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.scope_url", "scopeUrl"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.vault_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.content", "content"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.knowledge_base_configs.0.knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.model_size", "SMALL"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.instruction", "instruction"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.llm_selection_type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.instruction", "instruction"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.llm_selection_type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.runtime_version", "runtimeVersion"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_self_correction", "false"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_sql_execution", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.content", "content"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.input_location_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.tool_config_type", "SQL_TOOL_CONFIG"),
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
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.agent_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "OBJECT_STORAGE_LOCATION"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.content", "content2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.api_schema.0.object", "object2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.database_connection.0.connection_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_connection.0.connection_type", "DATABASE_TOOL_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.content", "content2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.input_location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.database_schema.0.prefix", "prefix2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.dialect", "ORACLE_SQL"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.function.0.parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope", "AGENT"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.client_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.http_endpoint_auth_scope_config_type", "HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.idcs_url", "idcsUrl2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.key_location", "QUERY_PARAMETER"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.key_name"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.scope_url", "scopeUrl2"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.vault_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.content", "content2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.input_location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.icl_examples.0.prefix", "prefix2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.knowledge_base_configs.0.knowledge_base_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.model_size", "LARGE"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.runtime_version", "runtimeVersion2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_self_correction", "true"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.should_enable_sql_execution", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "tool_config.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.content", "content2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.input_location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "tool_config.0.table_and_column_description.0.prefix", "prefix2"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.api_schema.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.api_schema.0.api_schema_input_location_type", "OBJECT_STORAGE_LOCATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.api_schema.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.api_schema.0.content", "content2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.api_schema.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.api_schema.0.object", "object2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_connection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_connection.0.connection_type", "DATABASE_TOOL_CONNECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_schema.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_schema.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_schema.0.content", "content2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_schema.0.input_location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_schema.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.database_schema.0.prefix", "prefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.dialect", "ORACLE_SQL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.embedding_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.embedding_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.embedding_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.embedding_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.function.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.function.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.function.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.function.0.parameters.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.generation_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope", "AGENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.http_endpoint_auth_scope_config_type", "HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.idcs_url", "idcsUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.key_location", "QUERY_PARAMETER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.http_endpoint_auth_config.0.http_endpoint_auth_sources.0.http_endpoint_auth_scope_config.0.scope_url", "scopeUrl2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.icl_examples.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.icl_examples.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.icl_examples.0.content", "content2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.icl_examples.0.input_location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.icl_examples.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.icl_examples.0.prefix", "prefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.knowledge_base_configs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.model_size", "LARGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reasoning_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reasoning_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reasoning_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reasoning_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reranking_llm_customization.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reranking_llm_customization.0.instruction", "instruction2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reranking_llm_customization.0.llm_hyper_parameters.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.reranking_llm_customization.0.llm_selection.0.llm_selection_type", "CUSTOM_GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.runtime_version", "runtimeVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.should_enable_self_correction", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.should_enable_sql_execution", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.table_and_column_description.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.table_and_column_description.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.table_and_column_description.0.content", "content2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.table_and_column_description.0.input_location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.table_and_column_description.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tool_config.0.table_and_column_description.0.prefix", "prefix2"),
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
