// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v43/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatabaseInsightRequiredOnlyResource = DatabaseInsightResourceDependencies +
		generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Required, Create, databaseInsightRepresentation)

	DatabaseInsightResourceConfig = DatabaseInsightResourceDependencies +
		generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Update, databaseInsightRepresentation)

	databaseInsightSingularDataSourceRepresentation = map[string]interface{}{
		"database_insight_id": Representation{repType: Required, create: `${oci_opsi_database_insight.test_database_insight.id}`},
	}

	databaseInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               Representation{repType: Optional, create: `${var.compartment_id}`},
		"database_type":                Representation{repType: Optional, create: []string{`EXTERNAL-NONCDB`}},
		"enterprise_manager_bridge_id": Representation{repType: Optional, create: `${var.enterprise_manager_bridge_id}`},
		"fields":                       Representation{repType: Optional, create: []string{`databaseName`, `databaseType`, `compartmentId`, `databaseDisplayName`, `freeformTags`, `definedTags`, `systemTags`}},
		"id":                           Representation{repType: Optional, create: `${oci_opsi_database_insight.test_database_insight.id}`},
		"state":                        Representation{repType: Optional, create: []string{`ACTIVE`}},
		"status":                       Representation{repType: Optional, create: []string{`ENABLED`}, update: []string{`DISABLED`}},
		"filter":                       RepresentationGroup{Required, databaseInsightDataSourceFilterRepresentation},
	}

	databaseInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_opsi_database_insight.test_database_insight.id}`}},
	}

	databaseInsightRepresentation = map[string]interface{}{
		"compartment_id":                       Representation{repType: Required, create: `${var.compartment_id}`},
		"enterprise_manager_bridge_id":         Representation{repType: Required, create: `${var.enterprise_manager_bridge_id}`},
		"enterprise_manager_entity_identifier": Representation{repType: Required, create: `${var.enterprise_manager_entity_id}`},
		"enterprise_manager_identifier":        Representation{repType: Required, create: `${var.enterprise_manager_id}`},
		"status":                               Representation{repType: Optional, create: `ENABLED`, update: `DISABLED`},
		"entity_source":                        Representation{repType: Required, create: `EM_MANAGED_EXTERNAL_DATABASE`, update: `EM_MANAGED_EXTERNAL_DATABASE`},
		"defined_tags":                         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                        Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                            RepresentationGroup{Required, ignoreChangesdatabaseInsightRepresentation},
	}

	ignoreChangesdatabaseInsightRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	DatabaseInsightResourceDependencies = DefinedTagsDependencies
)

func TestOpsiDatabaseInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiDatabaseInsightResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	emBridgeId := getEnvSettingWithBlankDefault("enterprise_manager_bridge_ocid")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	enterpriseManagerId := getEnvSettingWithBlankDefault("enterprise_manager_id")
	enterpriseManagerIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_id\" { default = \"%s\" }\n", enterpriseManagerId)

	enterpriseManagerEntityId := getEnvSettingWithBlankDefault("enterprise_manager_entity_id")
	enterpriseManagerEntityIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_entity_id\" { default = \"%s\" }\n", enterpriseManagerEntityId)

	resourceName := "oci_opsi_database_insight.test_database_insight"
	datasourceName := "data.oci_opsi_database_insights.test_database_insights"
	singularDatasourceName := "data.oci_opsi_database_insight.test_database_insight"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+emBridgeIdVariableStr+enterpriseManagerIdVariableStr+enterpriseManagerEntityIdVariableStr+DatabaseInsightResourceDependencies+
		generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Create, databaseInsightRepresentation), "opsi", "databaseInsight", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOpsiDatabaseInsightDestroy,
		Steps: []resource.TestStep{
			// verify create with optional
			{
				Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Create, databaseInsightRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttrSet(resourceName, "database_id"), // Won't be available for EM managed databases
					//resource.TestCheckResourceAttrSet(resourceName, "database_name"),
					//resource.TestCheckResourceAttrSet(resourceName, "database_resource_type"),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_bridge_id"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
					resource.TestCheckResourceAttr(resourceName, "entity_source", "EM_MANAGED_EXTERNAL_DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + compartmentIdUVariableStr + DatabaseInsightResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Create,
						representationCopyWithNewProperties(databaseInsightRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					//resource.TestCheckResourceAttrSet(resourceName, "database_id"), // Won't be available for EM managed databases
					//resource.TestCheckResourceAttrSet(resourceName, "database_name"),
					//resource.TestCheckResourceAttrSet(resourceName, "database_resource_type"),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_bridge_id"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
					resource.TestCheckResourceAttr(resourceName, "entity_source", "EM_MANAGED_EXTERNAL_DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Update, databaseInsightRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttrSet(resourceName, "database_id"), // Won't be available for EM managed databases
					//resource.TestCheckResourceAttrSet(resourceName, "database_name"),
					//resource.TestCheckResourceAttrSet(resourceName, "database_resource_type"),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_bridge_id"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
					resource.TestCheckResourceAttr(resourceName, "entity_source", "EM_MANAGED_EXTERNAL_DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_opsi_database_insights", "test_database_insights", Optional, Update, databaseInsightDataSourceRepresentation) +
					compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Update, databaseInsightRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(datasourceName, "database_id.#", "1"), // Won't be available for EM managed databases
					resource.TestCheckResourceAttr(datasourceName, "database_type.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "enterprise_manager_bridge_id"),
					resource.TestCheckResourceAttr(datasourceName, "fields.#", "7"),
					//resource.TestCheckResourceAttr(datasourceName, "id.#", "1"), // id is no more list. It is a string
					resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

					resource.TestCheckResourceAttr(datasourceName, "database_insights_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_insights_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Required, Create, databaseInsightSingularDataSourceRepresentation) +
					compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(singularDatasourceName, "connection_credential_details.#", "1"), //Won't be available for EM managed databses
					//resource.TestCheckResourceAttr(singularDatasourceName, "connection_details.#", "1"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "connector_id"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "database_display_name"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "database_name"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "database_resource_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "database_version"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "db_additional_details"),
					//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_entity_display_name"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
					resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
					resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
					resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "EM_MANAGED_EXTERNAL_DATABASE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_id"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "processor_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceConfig,
			},
			// verify enable
			{
				Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + DatabaseInsightResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_database_insight", "test_database_insight", Optional, Update,
						representationCopyWithNewProperties(databaseInsightRepresentation, map[string]interface{}{
							"status": Representation{repType: Required, update: `ENABLED`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckOpsiDatabaseInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).operationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_database_insight" {
			noResourceFound = false
			request := oci_opsi.GetDatabaseInsightRequest{}

			tmp := rs.Primary.ID
			request.DatabaseInsightId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "opsi")

			response, err := client.GetDatabaseInsight(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("OpsiDatabaseInsight") {
		resource.AddTestSweepers("OpsiDatabaseInsight", &resource.Sweeper{
			Name:         "OpsiDatabaseInsight",
			Dependencies: DependencyGraph["databaseInsight"],
			F:            sweepOpsiDatabaseInsightResource,
		})
	}
}

func sweepOpsiDatabaseInsightResource(compartment string) error {
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()
	databaseInsightIds, err := getDatabaseInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseInsightId := range databaseInsightIds {
		if ok := SweeperDefaultResourceId[databaseInsightId]; !ok {
			deleteDatabaseInsightRequest := oci_opsi.DeleteDatabaseInsightRequest{}

			deleteDatabaseInsightRequest.DatabaseInsightId = &databaseInsightId

			deleteDatabaseInsightRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteDatabaseInsight(context.Background(), deleteDatabaseInsightRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseInsightId, error)
				continue
			}
			waitTillCondition(testAccProvider, &databaseInsightId, databaseInsightSweepWaitCondition, time.Duration(3*time.Minute),
				databaseInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getDatabaseInsightIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DatabaseInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()

	listDatabaseInsightsRequest := oci_opsi.ListDatabaseInsightsRequest{}
	listDatabaseInsightsRequest.CompartmentId = &compartmentId
	listDatabaseInsightsRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listDatabaseInsightsResponse, err := operationsInsightsClient.ListDatabaseInsights(context.Background(), listDatabaseInsightsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseInsight list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseInsight := range listDatabaseInsightsResponse.Items {
		id := *databaseInsight.GetId()
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseInsightId", id)
	}
	return resourceIds, nil
}

func databaseInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseInsightResponse, ok := response.Response.(oci_opsi.GetDatabaseInsightResponse); ok {
		return databaseInsightResponse.GetLifecycleState() != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func databaseInsightSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.operationsInsightsClient().GetDatabaseInsight(context.Background(), oci_opsi.GetDatabaseInsightRequest{
		DatabaseInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
