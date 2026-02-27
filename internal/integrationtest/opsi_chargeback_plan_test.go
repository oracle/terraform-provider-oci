// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiChargebackPlanRequiredOnlyResource = OpsiChargebackPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Required, acctest.Create, OpsiChargebackPlanRepresentation)

	OpsiChargebackPlanResourceConfig = OpsiChargebackPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Optional, acctest.Update, OpsiChargebackPlanRepresentation)

	OpsiChargebackPlanSingularDataSourceRepresentation = map[string]interface{}{
		"chargebackplan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_chargeback_plan.test_chargeback_plan.id}`},
	}

	OpsiChargebackPlanDataSourceRepresentation = map[string]interface{}{
		"chargebackplan_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_chargeback_plan.test_chargeback_plan.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiChargebackPlanDataSourceFilterRepresentation}}
	OpsiChargebackPlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_chargeback_plan.test_chargeback_plan.id}`}},
	}

	OpsiChargebackPlanRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_source":     acctest.Representation{RepType: acctest.Required, Create: `CHARGEBACK_EXADATA`},
		"plan_name":         acctest.Representation{RepType: acctest.Required, Create: `planNameTF1`, Update: `planNameTF1u`},
		"plan_type":         acctest.Representation{RepType: acctest.Required, Create: `WEIGHTED_ALLOCATION`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"plan_custom_items": OpsiChargebackPlanPlanCustomItemsRepresentation,
		"plan_description":  acctest.Representation{RepType: acctest.Required, Create: `planDescription`, Update: `planDescription2`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesRepresentation},
	}

	ignoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OpsiChargebackPlanPlanCustomItemsRepresentation = []acctest.RepresentationGroup{
		{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"is_customizable": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
				"name":            acctest.Representation{RepType: acctest.Required, Create: `statistic`, Update: `statistic`},
				"value":           acctest.Representation{RepType: acctest.Required, Create: `AVG`, Update: `MAX`},
			},
		},
		{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"name":  acctest.Representation{RepType: acctest.Required, Create: `metricCostSplit`, Update: `metricCostSplit`},
				"value": acctest.Representation{RepType: acctest.Required, Create: `CPU:100`, Update: `CPU:100`},
			},
		},
	}

	OpsiChargebackPlanResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiChargebackPlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiChargebackPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_chargeback_plan.test_chargeback_plan"
	datasourceName := "data.oci_opsi_chargeback_plans.test_chargeback_plans"
	singularDatasourceName := "data.oci_opsi_chargeback_plan.test_chargeback_plan"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpsiChargebackPlanResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Optional, acctest.Create, OpsiChargebackPlanRepresentation), "operationsinsights", "chargebackPlan", t)

	acctest.ResourceTest(t, testAccCheckOpsiChargebackPlanDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OpsiChargebackPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Required, acctest.Create, OpsiChargebackPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "CHARGEBACK_EXADATA"),
				resource.TestCheckResourceAttr(resourceName, "plan_name", "planNameTF1"),
				resource.TestCheckResourceAttr(resourceName, "plan_type", "WEIGHTED_ALLOCATION"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiChargebackPlanResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OpsiChargebackPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Optional, acctest.Create, OpsiChargebackPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "CHARGEBACK_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.is_customizable", "false"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.name", "statistic"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.value", "AVG"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.1.name", "metricCostSplit"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.1.value", "CPU:100"),
				resource.TestCheckResourceAttr(resourceName, "plan_description", "planDescription"),
				resource.TestCheckResourceAttr(resourceName, "plan_name", "planNameTF1"),
				resource.TestCheckResourceAttr(resourceName, "plan_type", "WEIGHTED_ALLOCATION"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OpsiChargebackPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiChargebackPlanRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "CHARGEBACK_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.is_customizable", "false"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.name", "statistic"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.value", "AVG"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.1.name", "metricCostSplit"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.1.value", "CPU:100"),
				resource.TestCheckResourceAttr(resourceName, "plan_description", "planDescription"),
				resource.TestCheckResourceAttr(resourceName, "plan_name", "planNameTF1"),
				resource.TestCheckResourceAttr(resourceName, "plan_type", "WEIGHTED_ALLOCATION"),
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
			Config: config + compartmentIdVariableStr + OpsiChargebackPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Optional, acctest.Update, OpsiChargebackPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "CHARGEBACK_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.is_customizable", "false"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.name", "statistic"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.0.value", "MAX"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.1.name", "metricCostSplit"),
				resource.TestCheckResourceAttr(resourceName, "plan_custom_items.1.value", "CPU:100"),
				resource.TestCheckResourceAttr(resourceName, "plan_description", "planDescription2"),
				resource.TestCheckResourceAttr(resourceName, "plan_name", "planNameTF1u"),
				resource.TestCheckResourceAttr(resourceName, "plan_type", "WEIGHTED_ALLOCATION"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_chargeback_plans", "test_chargeback_plans", acctest.Optional, acctest.Update, OpsiChargebackPlanDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiChargebackPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Optional, acctest.Update, OpsiChargebackPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "chargebackplan_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),

				resource.TestCheckResourceAttr(datasourceName, "chargeback_plan_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "chargeback_plan_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_chargeback_plan", "test_chargeback_plan", acctest.Required, acctest.Create, OpsiChargebackPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiChargebackPlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "chargebackplan_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "CHARGEBACK_EXADATA"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_customizable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_category"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_custom_items.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_custom_items.0.is_customizable", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_custom_items.0.name", "statistic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_custom_items.0.value", "MAX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_custom_items.1.name", "metricCostSplit"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_custom_items.1.value", "CPU:100"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_description", "planDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_name", "planNameTF1u"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plan_type", "WEIGHTED_ALLOCATION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OpsiChargebackPlanRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiChargebackPlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_chargeback_plan" {
			noResourceFound = false
			request := oci_opsi.GetChargebackPlanRequest{}

			if value, ok := rs.Primary.Attributes["chargebackplan_id"]; ok && value != "" {
				request.ChargebackplanId = &value
			} else if id := strings.TrimSpace(rs.Primary.ID); id != "" {
				request.ChargebackplanId = &id
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetChargebackPlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiChargebackPlan") {
		resource.AddTestSweepers("OpsiChargebackPlan", &resource.Sweeper{
			Name:         "OpsiChargebackPlan",
			Dependencies: acctest.DependencyGraph["chargebackPlan"],
			F:            sweepOpsiChargebackPlanResource,
		})
	}
}

func sweepOpsiChargebackPlanResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	chargebackPlanIds, err := getOpsiChargebackPlanIds(compartment)
	if err != nil {
		return err
	}
	for _, chargebackPlanId := range chargebackPlanIds {
		if ok := acctest.SweeperDefaultResourceId[chargebackPlanId]; !ok {
			deleteChargebackPlanRequest := oci_opsi.DeleteChargebackPlanRequest{}
			deleteChargebackPlanRequest.ChargebackplanId = &chargebackPlanId

			deleteChargebackPlanRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteChargebackPlan(context.Background(), deleteChargebackPlanRequest)
			if error != nil {
				fmt.Printf("Error deleting ChargebackPlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", chargebackPlanId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &chargebackPlanId, OpsiChargebackPlanSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiChargebackPlanSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiChargebackPlanIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ChargebackPlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listChargebackPlansRequest := oci_opsi.ListChargebackPlansRequest{}
	listChargebackPlansRequest.CompartmentId = &compartmentId
	listChargebackPlansResponse, err := operationsInsightsClient.ListChargebackPlans(context.Background(), listChargebackPlansRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ChargebackPlan list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, chargebackPlan := range listChargebackPlansResponse.Items {
		id := *chargebackPlan.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ChargebackPlanId", id)
	}
	return resourceIds, nil
}

func OpsiChargebackPlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if chargebackPlanResponse, ok := response.Response.(oci_opsi.GetChargebackPlanResponse); ok {
		return chargebackPlanResponse.LifecycleState != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func OpsiChargebackPlanSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetChargebackPlan(context.Background(), oci_opsi.GetChargebackPlanRequest{
		ChargebackplanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
