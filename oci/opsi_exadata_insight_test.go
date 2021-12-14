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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v54/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExadataInsightRequiredOnlyResource = ExadataInsightResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Required, Create, exadataInsightRepresentation)

	ExadataInsightResourceConfig = ExadataInsightResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Update, exadataInsightRepresentation)

	exadataInsightSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_insight_id": Representation{RepType: Required, Create: `${oci_opsi_exadata_insight.test_exadata_insight.id}`},
	}

	exadataInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":    Representation{RepType: Optional, Create: `false`},
		"enterprise_manager_bridge_id": Representation{RepType: Optional, Create: `${var.enterprise_manager_bridge_id}`},
		"exadata_type":                 Representation{RepType: Optional, Create: []string{`DBMACHINE`}},
		"id":                           Representation{RepType: Optional, Create: `${oci_opsi_exadata_insight.test_exadata_insight.id}`},
		"state":                        Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"status":                       Representation{RepType: Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                       RepresentationGroup{Required, exadataInsightDataSourceFilterRepresentation}}

	exadataInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_opsi_exadata_insight.test_exadata_insight.id}`}},
	}

	exadataInsightRepresentation = map[string]interface{}{
		"compartment_id":                       Representation{RepType: Required, Create: `${var.compartment_id}`},
		"enterprise_manager_bridge_id":         Representation{RepType: Required, Create: `${var.enterprise_manager_bridge_id}`},
		"enterprise_manager_entity_identifier": Representation{RepType: Required, Create: `${var.enterprise_manager_entity_id}`},
		"enterprise_manager_identifier":        Representation{RepType: Required, Create: `${var.enterprise_manager_id}`},
		"entity_source":                        Representation{RepType: Required, Create: `EM_MANAGED_EXTERNAL_EXADATA`, Update: `EM_MANAGED_EXTERNAL_EXADATA`},
		"status":                               Representation{RepType: Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":                         Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                        Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                            RepresentationGroup{Required, ignoreChangesexadataInsightRepresentation},
		"is_auto_sync_enabled":                 Representation{RepType: Optional, Create: `true`, Update: `false`},
	}

	ignoreChangesexadataInsightRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	ExadataInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiExadataInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiExadataInsightResource_basic")
	defer httpreplay.SaveScenario()

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

	resourceName := "oci_opsi_exadata_insight.test_exadata_insight"
	datasourceName := "data.oci_opsi_exadata_insights.test_exadata_insights"
	singularDatasourceName := "data.oci_opsi_exadata_insight.test_exadata_insight"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+emBridgeIdVariableStr+enterpriseManagerIdVariableStr+enterpriseManagerEntityIdVariableStr+ExadataInsightResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Create, exadataInsightRepresentation), "opsi", "exadataInsight", t)

	ResourceTest(t, testAccCheckOpsiExadataInsightDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Create, exadataInsightRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_bridge_id"),
				resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
				resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "EM_MANAGED_EXTERNAL_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_sync_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + compartmentIdUVariableStr + ExadataInsightResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Create,
					RepresentationCopyWithNewProperties(exadataInsightRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_bridge_id"),
				resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
				resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "EM_MANAGED_EXTERNAL_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_sync_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Update, exadataInsightRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_bridge_id"),
				resource.TestCheckResourceAttr(resourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_name"),
				resource.TestCheckResourceAttrSet(resourceName, "enterprise_manager_entity_type"),
				resource.TestCheckResourceAttr(resourceName, "enterprise_manager_identifier", enterpriseManagerId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "EM_MANAGED_EXTERNAL_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_auto_sync_enabled"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_opsi_exadata_insights", "test_exadata_insights", Optional, Update, exadataInsightDataSourceRepresentation) +
				compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Update, exadataInsightRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "enterprise_manager_bridge_id"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_type.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "id.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "exadata_insight_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_insight_summary_collection.0.items.#", "1"),
			),
		},
		//verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Required, Create, exadataInsightSingularDataSourceRepresentation) +
				compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_insight_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_entity_display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enterprise_manager_entity_identifier", enterpriseManagerEntityId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_entity_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_entity_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enterprise_manager_identifier", enterpriseManagerId),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "EM_MANAGED_EXTERNAL_EXADATA"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_rack_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_sync_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_virtualized_exadata"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		//remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceConfig,
		},
		//verify enable
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + ExadataInsightResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", Optional, Update,
					RepresentationCopyWithNewProperties(exadataInsightRepresentation, map[string]interface{}{
						"status": Representation{RepType: Required, Update: `ENABLED`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
	})
}

func testAccCheckOpsiExadataInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).operationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_exadata_insight" {
			noResourceFound = false
			request := oci_opsi.GetExadataInsightRequest{}

			tmp := rs.Primary.ID
			request.ExadataInsightId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "opsi")

			response, err := client.GetExadataInsight(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.ExadataInsightLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("OpsiExadataInsight") {
		resource.AddTestSweepers("OpsiExadataInsight", &resource.Sweeper{
			Name:         "OpsiExadataInsight",
			Dependencies: DependencyGraph["exadataInsight"],
			F:            sweepOpsiExadataInsightResource,
		})
	}
}

func sweepOpsiExadataInsightResource(compartment string) error {
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()
	exadataInsightIds, err := getExadataInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, exadataInsightId := range exadataInsightIds {
		if ok := SweeperDefaultResourceId[exadataInsightId]; !ok {
			deleteExadataInsightRequest := oci_opsi.DeleteExadataInsightRequest{}

			deleteExadataInsightRequest.ExadataInsightId = &exadataInsightId

			deleteExadataInsightRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteExadataInsight(context.Background(), deleteExadataInsightRequest)
			if error != nil {
				fmt.Printf("Error deleting ExadataInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", exadataInsightId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &exadataInsightId, exadataInsightSweepWaitCondition, time.Duration(3*time.Minute),
				exadataInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getExadataInsightIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ExadataInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()

	listExadataInsightsRequest := oci_opsi.ListExadataInsightsRequest{}
	listExadataInsightsRequest.CompartmentId = &compartmentId
	listExadataInsightsRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listExadataInsightsResponse, err := operationsInsightsClient.ListExadataInsights(context.Background(), listExadataInsightsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExadataInsight list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, exadataInsight := range listExadataInsightsResponse.Items {
		id := *exadataInsight.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ExadataInsightId", id)
	}
	return resourceIds, nil
}

func exadataInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exadataInsightResponse, ok := response.Response.(oci_opsi.GetExadataInsightResponse); ok {
		return exadataInsightResponse.GetLifecycleState() != oci_opsi.ExadataInsightLifecycleStateDeleted
	}
	return false
}

func exadataInsightSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.operationsInsightsClient().GetExadataInsight(context.Background(), oci_opsi.GetExadataInsightRequest{
		ExadataInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
