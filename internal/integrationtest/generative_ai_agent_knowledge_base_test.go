// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentKnowledgeBaseRequiredOnlyResource = GenerativeAiAgentKnowledgeBaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, GenerativeAiAgentBYOKnowledgeBaseRepresentation)

	GenerativeAiAgentKnowledgeBaseResourceConfig = GenerativeAiAgentKnowledgeBaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Update, GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation)

	GenerativeAiAgentKnowledgeBaseSingularDataSourceRepresentation = map[string]interface{}{
		"knowledge_base_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_knowledge_base.test_knowledge_base.id}`},
	}

	GenerativeAiAgentKnowledgeBaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentKnowledgeBaseDataSourceFilterRepresentation}}
	GenerativeAiAgentKnowledgeBaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_knowledge_base.test_knowledge_base.id}`}},
	}

	GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"index_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentKnowledgeBaseDefaultIndexConfigRepresentation},
		// "defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GenerativeAiAgentBYOKnowledgeBaseRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"index_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentKnowledgeBaseBYOIndexConfigRepresentation},
		// "defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GenerativeAiAgentKnowledgeBaseDefaultIndexConfigRepresentation = map[string]interface{}{
		"index_config_type":           acctest.Representation{RepType: acctest.Required, Create: `DEFAULT_INDEX_CONFIG`},
		"should_enable_hybrid_search": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	GenerativeAiAgentKnowledgeBaseBYOIndexConfigRepresentation = map[string]interface{}{
		"index_config_type": acctest.Representation{RepType: acctest.Required, Create: `OCI_OPEN_SEARCH_INDEX_CONFIG`},
		"cluster_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.os_cluster_id}`},
		"indexes":           acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentKnowledgeBaseIndexConfigIndexesRepresentation},
		"secret_detail":     acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentKnowledgeBaseIndexConfigSecretDetailRepresentation},
	}
	GenerativeAiAgentKnowledgeBaseIndexConfigIndexesRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `iaas`},
		"schema": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentKnowledgeBaseIndexConfigIndexesSchemaRepresentation},
	}
	GenerativeAiAgentKnowledgeBaseIndexConfigSecretDetailRepresentation = map[string]interface{}{
		"type":            acctest.Representation{RepType: acctest.Required, Create: `BASIC_AUTH_SECRET`},
		"vault_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	GenerativeAiAgentKnowledgeBaseIndexConfigIndexesSchemaRepresentation = map[string]interface{}{
		"body_key":           acctest.Representation{RepType: acctest.Required, Create: `body`},
		"embedding_body_key": acctest.Representation{RepType: acctest.Optional, Create: `embeddingBodyKey`},
		"title_key":          acctest.Representation{RepType: acctest.Optional, Create: `title`},
		"url_key":            acctest.Representation{RepType: acctest.Optional, Create: `url`},
	}

	GenerativeAiAgentKnowledgeBaseResourceDependencies = ``
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentKnowledgeBaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentKnowledgeBaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	clusterId := utils.GetEnvSettingWithDefault("os_cluster_id_for_byo_kb", compartmentId)
	clusterIdVariableStr := fmt.Sprintf("variable \"os_cluster_id\" { default = \"%s\" }\n", clusterId)

	secretIdU := utils.GetEnvSettingWithDefault("vault_secret_id_for_os_cluster", compartmentId)
	secretIdVariableStr := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretIdU)

	resourceName := "oci_generative_ai_agent_knowledge_base.test_knowledge_base"
	datasourceName := "data.oci_generative_ai_agent_knowledge_bases.test_knowledge_bases"
	singularDatasourceName := "data.oci_generative_ai_agent_knowledge_base.test_knowledge_base"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Create, GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation), "generativeaiagent", "knowledgeBase", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentKnowledgeBaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + clusterIdVariableStr + secretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, GenerativeAiAgentBYOKnowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "index_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "index_config.0.cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.index_config_type", "OCI_OPEN_SEARCH_INDEX_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.indexes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.indexes.0.schema.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.secret_detail.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "index_config.0.secret_detail.0.vault_secret_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + clusterIdVariableStr + secretIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Create, GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "index_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.index_config_type", "DEFAULT_INDEX_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.should_enable_hybrid_search", "false"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "index_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.index_config_type", "DEFAULT_INDEX_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.should_enable_hybrid_search", "false"),
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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Update, GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "index_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.index_config_type", "DEFAULT_INDEX_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "index_config.0.should_enable_hybrid_search", "false"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_knowledge_bases", "test_knowledge_bases", acctest.Optional, acctest.Update, GenerativeAiAgentKnowledgeBaseDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Optional, acctest.Update, GenerativeAiAgentServiceManagedKnowledgeBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "knowledge_base_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "knowledge_base_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, GenerativeAiAgentKnowledgeBaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiAgentKnowledgeBaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "knowledge_base_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "index_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "index_config.0.index_config_type", "DEFAULT_INDEX_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "index_config.0.should_enable_hybrid_search", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiAgentKnowledgeBaseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentKnowledgeBaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_knowledge_base" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetKnowledgeBaseRequest{}

			tmp := rs.Primary.ID
			request.KnowledgeBaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetKnowledgeBase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.KnowledgeBaseLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentKnowledgeBase") {
		resource.AddTestSweepers("GenerativeAiAgentKnowledgeBase", &resource.Sweeper{
			Name:         "GenerativeAiAgentKnowledgeBase",
			Dependencies: acctest.DependencyGraph["knowledgeBase"],
			F:            sweepGenerativeAiAgentKnowledgeBaseResource,
		})
	}
}

func sweepGenerativeAiAgentKnowledgeBaseResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	knowledgeBaseIds, err := getGenerativeAiAgentKnowledgeBaseIds(compartment)
	if err != nil {
		return err
	}
	for _, knowledgeBaseId := range knowledgeBaseIds {
		if ok := acctest.SweeperDefaultResourceId[knowledgeBaseId]; !ok {
			deleteKnowledgeBaseRequest := oci_generative_ai_agent.DeleteKnowledgeBaseRequest{}

			deleteKnowledgeBaseRequest.KnowledgeBaseId = &knowledgeBaseId

			deleteKnowledgeBaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteKnowledgeBase(context.Background(), deleteKnowledgeBaseRequest)
			if error != nil {
				fmt.Printf("Error deleting KnowledgeBase %s %s, It is possible that the resource is already deleted. Please verify manually \n", knowledgeBaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &knowledgeBaseId, GenerativeAiAgentKnowledgeBaseSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentKnowledgeBaseSweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentKnowledgeBaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "KnowledgeBaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listKnowledgeBasesRequest := oci_generative_ai_agent.ListKnowledgeBasesRequest{}
	listKnowledgeBasesRequest.CompartmentId = &compartmentId
	listKnowledgeBasesRequest.LifecycleState = oci_generative_ai_agent.KnowledgeBaseLifecycleStateActive
	listKnowledgeBasesResponse, err := generativeAiAgentClient.ListKnowledgeBases(context.Background(), listKnowledgeBasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting KnowledgeBase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, knowledgeBase := range listKnowledgeBasesResponse.Items {
		id := *knowledgeBase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "KnowledgeBaseId", id)
	}
	return resourceIds, nil
}

func GenerativeAiAgentKnowledgeBaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if knowledgeBaseResponse, ok := response.Response.(oci_generative_ai_agent.GetKnowledgeBaseResponse); ok {
		return knowledgeBaseResponse.LifecycleState != oci_generative_ai_agent.KnowledgeBaseLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentKnowledgeBaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetKnowledgeBase(context.Background(), oci_generative_ai_agent.GetKnowledgeBaseRequest{
		KnowledgeBaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
