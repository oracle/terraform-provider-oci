// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OpsiExadataInsightRequiredOnlyResource = OpsiExadataInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Required, acctest.Create, OpsiExadataInsightRepresentation)

	OpsiExadataInsightResourceConfig = OpsiExadataInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update, OpsiExadataInsightRepresentation)

	OpsiOpsiExadataInsightSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_insight_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_exadata_insight.test_exadata_insight.id}`},
	}

	OpsiOpsiExadataInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"enterprise_manager_bridge_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.enterprise_manager_bridge_id}`},
		"exadata_type":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`DBMACHINE`}},
		"id":                           acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_exadata_insight.test_exadata_insight.id}`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiExadataInsightDataSourceFilterRepresentation}}

	OpsiExadataInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_exadata_insight.test_exadata_insight.id}`}},
	}

	OpsiExadataInsightRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"enterprise_manager_bridge_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.enterprise_manager_bridge_id}`},
		"enterprise_manager_entity_identifier": acctest.Representation{RepType: acctest.Required, Create: `${var.em_exadata_enterprise_manager_entity_id}`},
		"enterprise_manager_identifier":        acctest.Representation{RepType: acctest.Required, Create: `${var.em_exadata_enterprise_manager_id}`},
		"entity_source":                        acctest.Representation{RepType: acctest.Required, Create: `EM_MANAGED_EXTERNAL_EXADATA`, Update: `EM_MANAGED_EXTERNAL_EXADATA`},
		"status":                               acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":                         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesexadataInsightRepresentation},
		"is_auto_sync_enabled":                 acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}

	ignoreChangesexadataInsightRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OpsiExadataInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiExadataInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiExadataInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	emBridgeId := utils.GetEnvSettingWithBlankDefault("enterprise_manager_bridge_ocid")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"enterprise_manager_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	enterpriseManagerId := utils.GetEnvSettingWithBlankDefault("em_exadata_enterprise_manager_id")
	enterpriseManagerIdVariableStr := fmt.Sprintf("variable \"em_exadata_enterprise_manager_id\" { default = \"%s\" }\n", enterpriseManagerId)

	enterpriseManagerEntityId := utils.GetEnvSettingWithBlankDefault("em_exadata_enterprise_manager_entity_id")
	enterpriseManagerEntityIdVariableStr := fmt.Sprintf("variable \"em_exadata_enterprise_manager_entity_id\" { default = \"%s\" }\n", enterpriseManagerEntityId)

	resourceName := "oci_opsi_exadata_insight.test_exadata_insight"
	datasourceName := "data.oci_opsi_exadata_insights.test_exadata_insights"
	singularDatasourceName := "data.oci_opsi_exadata_insight.test_exadata_insight"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+emBridgeIdVariableStr+enterpriseManagerIdVariableStr+enterpriseManagerEntityIdVariableStr+OpsiExadataInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Create, OpsiExadataInsightRepresentation), "opsi", "exadataInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiExadataInsightDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + OpsiExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Create, OpsiExadataInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + compartmentIdUVariableStr + OpsiExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiExadataInsightRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + OpsiExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update, OpsiExadataInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_exadata_insights", "test_exadata_insights", acctest.Optional, acctest.Update, OpsiOpsiExadataInsightDataSourceRepresentation) +
				compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + OpsiExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update, OpsiExadataInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "enterprise_manager_bridge_id"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "exadata_insight_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_insight_summary_collection.0.items.#", "1"),
			),
		},
		//verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Required, acctest.Create, OpsiOpsiExadataInsightSingularDataSourceRepresentation) +
				compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + OpsiExadataInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		//verify enable
		{
			Config: config + compartmentIdVariableStr + emBridgeIdVariableStr + enterpriseManagerIdVariableStr + enterpriseManagerEntityIdVariableStr + OpsiExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(OpsiExadataInsightRepresentation, map[string]interface{}{
						"status": acctest.Representation{RepType: acctest.Required, Update: `ENABLED`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + OpsiExadataInsightRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiExadataInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_exadata_insight" {
			noResourceFound = false
			request := oci_opsi.GetExadataInsightRequest{}

			tmp := rs.Primary.ID
			request.ExadataInsightId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OpsiExadataInsight") {
		resource.AddTestSweepers("OpsiExadataInsight", &resource.Sweeper{
			Name:         "OpsiExadataInsight",
			Dependencies: acctest.DependencyGraph["exadataInsight"],
			F:            sweepOpsiExadataInsightResource,
		})
	}
}

func sweepOpsiExadataInsightResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	exadataInsightIds, err := getOpsiExadataInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, exadataInsightId := range exadataInsightIds {
		if ok := acctest.SweeperDefaultResourceId[exadataInsightId]; !ok {
			deleteExadataInsightRequest := oci_opsi.DeleteExadataInsightRequest{}

			deleteExadataInsightRequest.ExadataInsightId = &exadataInsightId

			deleteExadataInsightRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteExadataInsight(context.Background(), deleteExadataInsightRequest)
			if error != nil {
				fmt.Printf("Error deleting ExadataInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", exadataInsightId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &exadataInsightId, OpsiExadataInsightSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiExadataInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiExadataInsightIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExadataInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExadataInsightId", id)
	}
	return resourceIds, nil
}

func OpsiExadataInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exadataInsightResponse, ok := response.Response.(oci_opsi.GetExadataInsightResponse); ok {
		return exadataInsightResponse.GetLifecycleState() != oci_opsi.ExadataInsightLifecycleStateDeleted
	}
	return false
}

func OpsiExadataInsightSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetExadataInsight(context.Background(), oci_opsi.GetExadataInsightRequest{
		ExadataInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
