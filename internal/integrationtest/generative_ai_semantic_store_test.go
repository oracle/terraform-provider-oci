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
	GenerativeAiSemanticStoreRequiredOnlyResource = GenerativeAiSemanticStoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Required, acctest.Create, GenerativeAiSemanticStoreRepresentation)

	GenerativeAiSemanticStoreResourceConfig = GenerativeAiSemanticStoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Optional, acctest.Update, GenerativeAiSemanticStoreRepresentation)

	GenerativeAiSemanticStoreSingularDataSourceRepresentation = map[string]interface{}{
		"semantic_store_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_semantic_store.test_semantic_store.id}`},
	}

	GenerativeAiSemanticStoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"data_source_querying_connection_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection2.id}`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_semantic_store.test_semantic_store.id}`},
		"state":                              acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":                             acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiSemanticStoreDataSourceFilterRepresentation}}
	GenerativeAiSemanticStoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_semantic_store.test_semantic_store.id}`}},
	}
	IgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	GenerativeAiSemanticStoreRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiSemanticStoreDatabaseConnectionRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"schemas":        acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiSemanticStoreSchemasRepresentation},
		//"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"refresh_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: GenerativeAiSemanticStoreRefreshScheduleRepresentation},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreChangesRepresentation},
	}
	GenerativeAiSemanticStoreDatabaseConnectionRepresentation = map[string]interface{}{
		"connection_type":          acctest.Representation{RepType: acctest.Required, Create: `DATABASE_TOOLS_CONNECTION`},
		"enrichment_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
		"querying_connection_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection2.id}`},
	}
	GenerativeAiSemanticStoreSchemasRepresentation = map[string]interface{}{
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `DATABASE_TOOLS_CONNECTION`},
		"schemas":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiSemanticStoreSchemasSchemasRepresentation},
	}
	GenerativeAiSemanticStoreRefreshScheduleRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `ON_CREATE`, Update: `NONE`},
		//"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	GenerativeAiSemanticStoreSchemasSchemasRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}

	GenerativeAiSemanticStoreResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection2", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionRepresentation)
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiSemanticStoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiSemanticStoreResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	secretId := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("secret_id"))

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_semantic_store.test_semantic_store"
	datasourceName := "data.oci_generative_ai_semantic_stores.test_semantic_stores"
	singularDatasourceName := "data.oci_generative_ai_semantic_store.test_semantic_store"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+secretId+GenerativeAiSemanticStoreResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Optional, acctest.Create, GenerativeAiSemanticStoreRepresentation), "generativeai", "semanticStore", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiSemanticStoreDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + secretId + GenerativeAiSemanticStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Required, acctest.Create, GenerativeAiSemanticStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.enrichment_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.querying_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.0.name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + secretId + GenerativeAiSemanticStoreResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + secretId + GenerativeAiSemanticStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Optional, acctest.Create, GenerativeAiSemanticStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.enrichment_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.querying_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "refresh_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "refresh_schedule.0.type", "ON_CREATE"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + secretId + compartmentIdUVariableStr + GenerativeAiSemanticStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiSemanticStoreRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.enrichment_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.querying_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "refresh_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "refresh_schedule.0.type", "ON_CREATE"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + secretId + GenerativeAiSemanticStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Optional, acctest.Update, GenerativeAiSemanticStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.enrichment_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.querying_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "refresh_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "refresh_schedule.0.type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.0.schemas.0.name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_semantic_stores", "test_semantic_stores", acctest.Optional, acctest.Update, GenerativeAiSemanticStoreDataSourceRepresentation) +
				compartmentIdVariableStr + secretId + GenerativeAiSemanticStoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Optional, acctest.Update, GenerativeAiSemanticStoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "data_source_querying_connection_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "semantic_store_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "semantic_store_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_semantic_store", "test_semantic_store", acctest.Required, acctest.Create, GenerativeAiSemanticStoreSingularDataSourceRepresentation) +
				compartmentIdVariableStr + secretId + GenerativeAiSemanticStoreResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "semantic_store_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "refresh_schedule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "refresh_schedule.0.type", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.0.connection_type", "DATABASE_TOOLS_CONNECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.0.schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.0.schemas.0.name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiSemanticStoreRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiSemanticStoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_semantic_store" {
			noResourceFound = false
			request := oci_generative_ai.GetSemanticStoreRequest{}

			tmp := rs.Primary.ID
			request.SemanticStoreId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetSemanticStore(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.SemanticStoreLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiSemanticStore") {
		resource.AddTestSweepers("GenerativeAiSemanticStore", &resource.Sweeper{
			Name:         "GenerativeAiSemanticStore",
			Dependencies: acctest.DependencyGraph["semanticStore"],
			F:            sweepGenerativeAiSemanticStoreResource,
		})
	}
}

func sweepGenerativeAiSemanticStoreResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	semanticStoreIds, err := getGenerativeAiSemanticStoreIds(compartment)
	if err != nil {
		return err
	}
	for _, semanticStoreId := range semanticStoreIds {
		if ok := acctest.SweeperDefaultResourceId[semanticStoreId]; !ok {
			deleteSemanticStoreRequest := oci_generative_ai.DeleteSemanticStoreRequest{}

			deleteSemanticStoreRequest.SemanticStoreId = &semanticStoreId

			deleteSemanticStoreRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteSemanticStore(context.Background(), deleteSemanticStoreRequest)
			if error != nil {
				fmt.Printf("Error deleting SemanticStore %s %s, It is possible that the resource is already deleted. Please verify manually \n", semanticStoreId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &semanticStoreId, GenerativeAiSemanticStoreSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiSemanticStoreSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiSemanticStoreIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SemanticStoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listSemanticStoresRequest := oci_generative_ai.ListSemanticStoresRequest{}
	listSemanticStoresRequest.CompartmentId = &compartmentId
	//listSemanticStoresRequest.LifecycleState = oci_generative_ai.ListSemanticStoresLifecycleStateActive
	listSemanticStoresResponse, err := generativeAiClient.ListSemanticStores(context.Background(), listSemanticStoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SemanticStore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, semanticStore := range listSemanticStoresResponse.Items {
		id := *semanticStore.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SemanticStoreId", id)
	}
	return resourceIds, nil
}

func GenerativeAiSemanticStoreSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if semanticStoreResponse, ok := response.Response.(oci_generative_ai.GetSemanticStoreResponse); ok {
		return semanticStoreResponse.LifecycleState != oci_generative_ai.SemanticStoreLifecycleStateDeleted
	}
	return false
}

func GenerativeAiSemanticStoreSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetSemanticStore(context.Background(), oci_generative_ai.GetSemanticStoreRequest{
		SemanticStoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
