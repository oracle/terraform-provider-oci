// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiProjectRequiredOnlyResource = GenerativeAiProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Required, acctest.Create, GenerativeAiProjectRepresentation)

	GenerativeAiProjectResourceConfig = GenerativeAiProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Optional, acctest.Update, GenerativeAiProjectRepresentation)

	GenerativeAiProjectSingularDataSourceRepresentation = map[string]interface{}{
		"project_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_project.test_project.id}`},
	}

	GenerativeAiProjectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiProjectDataSourceFilterRepresentation}}
	GenerativeAiProjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_project.test_project.id}`}},
	}

	GenerativeAiProjectRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"conversation_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectConversationConfigRepresentation},
		//"defined_tags":                          acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description":                           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"long_term_memory_config":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectLongTermMemoryConfigRepresentation},
		"short_term_memory_optimization_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectShortTermMemoryOptimizationConfigRepresentation},
	}
	GenerativeAiProjectConversationConfigRepresentation = map[string]interface{}{
		"conversations_retention_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"responses_retention_in_hours":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	GenerativeAiProjectLongTermMemoryConfigRepresentation = map[string]interface{}{
		"standard_long_term_memory_strategy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyRepresentation},
	}
	GenerativeAiProjectShortTermMemoryOptimizationConfigRepresentation = map[string]interface{}{
		"is_enabled":       acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"condenser_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectShortTermMemoryOptimizationConfigCondenserConfigRepresentation},
	}
	GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyRepresentation = map[string]interface{}{
		"is_enabled":        acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"embedding_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyEmbeddingConfigRepresentation},
		"extraction_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyExtractionConfigRepresentation},
	}
	GenerativeAiProjectShortTermMemoryOptimizationConfigCondenserConfigRepresentation = map[string]interface{}{
		"llm_selection": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiProjectShortTermMemoryOptimizationConfigCondenserConfigLlmSelectionRepresentation},
	}
	GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyEmbeddingConfigRepresentation = map[string]interface{}{
		"llm_selection": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyEmbeddingConfigLlmSelectionRepresentation},
	}
	GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyExtractionConfigRepresentation = map[string]interface{}{
		"llm_selection": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyExtractionConfigLlmSelectionRepresentation},
	}
	GenerativeAiProjectShortTermMemoryOptimizationConfigCondenserConfigLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `GEN_AI_MODEL`},
		"model_id":           acctest.Representation{RepType: acctest.Required, Create: `${local.project_model_id}`},
	}
	GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyEmbeddingConfigLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `GEN_AI_MODEL`},
		"model_id":           acctest.Representation{RepType: acctest.Required, Create: `${local.project_model_id}`},
	}
	GenerativeAiProjectLongTermMemoryConfigStandardLongTermMemoryStrategyExtractionConfigLlmSelectionRepresentation = map[string]interface{}{
		"llm_selection_type": acctest.Representation{RepType: acctest.Required, Create: `GEN_AI_MODEL`},
		"model_id":           acctest.Representation{RepType: acctest.Required, Create: `${local.project_model_id}`},
	}

	GenerativeAiProjectResourceDependencies = projectModelDependencies

	projectModelDependencies = `
	locals {

	  filtered_project_models = [
		for item in data.oci_generative_ai_models.project_models.model_collection[0].items : item
		  if (
			(item.version == "1.0")
			&& length(item.capabilities) == 1
			&& (item.display_name == "cohere.command-a-03-2025")
		  )
		]

	  project_model_id = local.filtered_project_models[0].id
	}

	data "oci_generative_ai_models" "project_models" {
	  compartment_id = var.compartment_id
	  display_name = "cohere.command-a-03-2025"
	}
	`
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_project.test_project"
	datasourceName := "data.oci_generative_ai_projects.test_projects"
	singularDatasourceName := "data.oci_generative_ai_project.test_project"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiProjectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Optional, acctest.Create, GenerativeAiProjectRepresentation), "generativeai", "project", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiProjectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Required, acctest.Create, GenerativeAiProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiProjectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Optional, acctest.Create, GenerativeAiProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.0.conversations_retention_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.0.responses_retention_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiProjectRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.0.conversations_retention_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.0.responses_retention_in_hours", "10"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + GenerativeAiProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Optional, acctest.Update, GenerativeAiProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.0.conversations_retention_in_hours", "11"),
				resource.TestCheckResourceAttr(resourceName, "conversation_config.0.responses_retention_in_hours", "11"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttrSet(resourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "short_term_memory_optimization_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_projects", "test_projects", acctest.Optional, acctest.Update, GenerativeAiProjectDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Optional, acctest.Update, GenerativeAiProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "generative_ai_project_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "generative_ai_project_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_project", "test_project", acctest.Required, acctest.Create, GenerativeAiProjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiProjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "conversation_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conversation_config.0.conversations_retention_in_hours", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conversation_config.0.responses_retention_in_hours", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.embedding_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.extraction_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_term_memory_config.0.standard_long_term_memory_strategy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_term_memory_optimization_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_term_memory_optimization_config.0.condenser_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_term_memory_optimization_config.0.condenser_config.0.llm_selection.0.llm_selection_type", "GEN_AI_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_term_memory_optimization_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + GenerativeAiProjectRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiProjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_project" {
			noResourceFound = false
			request := oci_generative_ai.GetGenerativeAiProjectRequest{}

			tmp := rs.Primary.ID
			request.GenerativeAiProjectId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetGenerativeAiProject(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.GenerativeAiProjectLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiProject") {
		resource.AddTestSweepers("GenerativeAiProject", &resource.Sweeper{
			Name:         "GenerativeAiProject",
			Dependencies: acctest.DependencyGraph["project"],
			F:            sweepGenerativeAiProjectResource,
		})
	}
}

func sweepGenerativeAiProjectResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	projectIds, err := getGenerativeAiProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, projectId := range projectIds {
		if ok := acctest.SweeperDefaultResourceId[projectId]; !ok {
			deleteGenerativeAiProjectRequest := oci_generative_ai.DeleteGenerativeAiProjectRequest{}

			deleteGenerativeAiProjectRequest.GenerativeAiProjectId = &projectId

			deleteGenerativeAiProjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteGenerativeAiProject(context.Background(), deleteGenerativeAiProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Project %s %s, It is possible that the resource is already deleted. Please verify manually \n", projectId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &projectId, GenerativeAiProjectSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiProjectSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiProjectIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listGenerativeAiProjectsRequest := oci_generative_ai.ListGenerativeAiProjectsRequest{}
	listGenerativeAiProjectsRequest.CompartmentId = &compartmentId
	listGenerativeAiProjectsRequest.LifecycleState = oci_generative_ai.GenerativeAiProjectLifecycleStateActive
	listGenerativeAiProjectsResponse, err := generativeAiClient.ListGenerativeAiProjects(context.Background(), listGenerativeAiProjectsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Project list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, project := range listGenerativeAiProjectsResponse.Items {
		id := *project.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProjectId", id)
	}
	return resourceIds, nil
}

func GenerativeAiProjectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if projectResponse, ok := response.Response.(oci_generative_ai.GetGenerativeAiProjectResponse); ok {
		return projectResponse.LifecycleState != oci_generative_ai.GenerativeAiProjectLifecycleStateDeleted
	}
	return false
}

func GenerativeAiProjectSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetGenerativeAiProject(context.Background(), oci_generative_ai.GetGenerativeAiProjectRequest{
		GenerativeAiProjectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
