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
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseSchedulingPlanRequiredOnlyResource = DatabaseSchedulingPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Required, acctest.Create, DatabaseSchedulingPlanRepresentation)

	DatabaseSchedulingPlanResourceConfig = DatabaseSchedulingPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Optional, acctest.Update, DatabaseSchedulingPlanRepresentation)

	DatabaseSchedulingPlanSingularDataSourceRepresentation = map[string]interface{}{
		"scheduling_plan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_plan.test_scheduling_plan.id}`},
	}

	DatabaseSchedulingPlanDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_scheduling_plan.test_scheduling_plan.id}`},
		"resource_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scheduling_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseSchedulingPlanDataSourceFilterRepresentation}}
	DatabaseSchedulingPlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_scheduling_plan.test_scheduling_plan.id}`}},
	}

	DatabaseSchedulingPlanRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"resource_id":                            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scheduling_policy_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
		"service_type":                           acctest.Representation{RepType: acctest.Required, Create: `EXACC`},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_using_recommended_scheduled_actions": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DatabaseSchedulingPlanResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Required, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation) + DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseSchedulingPlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseSchedulingPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_scheduling_plan.test_scheduling_plan"
	datasourceName := "data.oci_database_scheduling_plans.test_scheduling_plans"
	singularDatasourceName := "data.oci_database_scheduling_plan.test_scheduling_plan"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseSchedulingPlanResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Optional, acctest.Create, DatabaseSchedulingPlanRepresentation), "database", "schedulingPlan", t)

	acctest.ResourceTest(t, testAccCheckDatabaseSchedulingPlanDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Required, acctest.Create, DatabaseSchedulingPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACC"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPlanResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Optional, acctest.Create, DatabaseSchedulingPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_using_recommended_scheduled_actions", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACC"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseSchedulingPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseSchedulingPlanRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_using_recommended_scheduled_actions", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACC"),
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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_plans", "test_scheduling_plans", acctest.Optional, acctest.Update, DatabaseSchedulingPlanDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Optional, acctest.Update, DatabaseSchedulingPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Required, acctest.Create, DatabaseSchedulingPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduling_plan_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_time_in_mins"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_using_recommended_scheduled_actions", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_intent"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "EXACC"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseSchedulingPlanRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseSchedulingPlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_scheduling_plan" {
			noResourceFound = false
			request := oci_database.GetSchedulingPlanRequest{}

			tmp := rs.Primary.ID
			request.SchedulingPlanId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetSchedulingPlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.SchedulingPlanLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseSchedulingPlan") {
		resource.AddTestSweepers("DatabaseSchedulingPlan", &resource.Sweeper{
			Name:         "DatabaseSchedulingPlan",
			Dependencies: acctest.DependencyGraph["schedulingPlan"],
			F:            sweepDatabaseSchedulingPlanResource,
		})
	}
}

func sweepDatabaseSchedulingPlanResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	schedulingPlanIds, err := getDatabaseSchedulingPlanIds(compartment)
	if err != nil {
		return err
	}
	for _, schedulingPlanId := range schedulingPlanIds {
		if ok := acctest.SweeperDefaultResourceId[schedulingPlanId]; !ok {
			deleteSchedulingPlanRequest := oci_database.DeleteSchedulingPlanRequest{}

			deleteSchedulingPlanRequest.SchedulingPlanId = &schedulingPlanId

			deleteSchedulingPlanRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteSchedulingPlan(context.Background(), deleteSchedulingPlanRequest)
			if error != nil {
				fmt.Printf("Error deleting SchedulingPlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", schedulingPlanId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &schedulingPlanId, DatabaseSchedulingPlanSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseSchedulingPlanSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseSchedulingPlanIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SchedulingPlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listSchedulingPlansRequest := oci_database.ListSchedulingPlansRequest{}
	listSchedulingPlansRequest.CompartmentId = &compartmentId
	listSchedulingPlansRequest.LifecycleState = oci_database.SchedulingPlanSummaryLifecycleStateAvailable
	listSchedulingPlansResponse, err := databaseClient.ListSchedulingPlans(context.Background(), listSchedulingPlansRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SchedulingPlan list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedulingPlan := range listSchedulingPlansResponse.Items {
		id := *schedulingPlan.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SchedulingPlanId", id)
	}
	return resourceIds, nil
}

func DatabaseSchedulingPlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if schedulingPlanResponse, ok := response.Response.(oci_database.GetSchedulingPlanResponse); ok {
		return schedulingPlanResponse.LifecycleState != oci_database.SchedulingPlanLifecycleStateDeleted
	}
	return false
}

func DatabaseSchedulingPlanSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetSchedulingPlan(context.Background(), oci_database.GetSchedulingPlanRequest{
		SchedulingPlanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
