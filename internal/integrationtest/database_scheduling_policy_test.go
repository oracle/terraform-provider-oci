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
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseSchedulingPolicyRequiredOnlyResource = DatabaseSchedulingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRepresentation)

	DatabaseSchedulingPolicyResourceConfig = DatabaseSchedulingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Optional, acctest.Update, DatabaseSchedulingPolicyRepresentation)

	DatabaseSchedulingPolicySingularDataSourceRepresentation = map[string]interface{}{
		"scheduling_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
	}

	DatabaseSchedulingPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `testSchedulingPolicy`, Update: `displayName8`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `NEEDS_ATTENTION`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseSchedulingPolicyDataSourceFilterRepresentation}}
	DatabaseSchedulingPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_scheduling_policy.test_scheduling_policy.id}`}},
	}

	DatabaseSchedulingPolicyRepresentation = map[string]interface{}{
		"cadence":             acctest.Representation{RepType: acctest.Required, Create: `HALFYEARLY`, Update: `QUARTERLY`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `testSchedulingPolicy`, Update: `displayName8`},
		"cadence_start_month": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseSchedulingPolicyCadenceStartMonthRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	DatabaseSchedulingPolicyCadenceStartMonthRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JANUARY`, Update: `FEBRUARY`},
	}

	DatabaseSchedulingPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseSchedulingPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseSchedulingPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_scheduling_policy.test_scheduling_policy"
	datasourceName := "data.oci_database_scheduling_policies.test_scheduling_policies"
	singularDatasourceName := "data.oci_database_scheduling_policy.test_scheduling_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseSchedulingPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Optional, acctest.Create, DatabaseSchedulingPolicyRepresentation), "database", "schedulingPolicy", t)

	acctest.ResourceTest(t, testAccCheckDatabaseSchedulingPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cadence", "HALFYEARLY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testSchedulingPolicy"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Optional, acctest.Create, DatabaseSchedulingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cadence", "HALFYEARLY"),
				resource.TestCheckResourceAttr(resourceName, "cadence_start_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cadence_start_month.0.name", "JANUARY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testSchedulingPolicy"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseSchedulingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseSchedulingPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cadence", "HALFYEARLY"),
				resource.TestCheckResourceAttr(resourceName, "cadence_start_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cadence_start_month.0.name", "JANUARY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testSchedulingPolicy"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Optional, acctest.Update, DatabaseSchedulingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cadence", "QUARTERLY"),
				resource.TestCheckResourceAttr(resourceName, "cadence_start_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cadence_start_month.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName8"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_policies", "test_scheduling_policies", acctest.Optional, acctest.Update, DatabaseSchedulingPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Optional, acctest.Update, DatabaseSchedulingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName8"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NEEDS_ATTENTION"),

				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.0.cadence", "QUARTERLY"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.0.cadence_start_month.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.0.cadence_start_month.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.0.display_name", "displayName8"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policies.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policies.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policies.0.time_created"),
				//resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policies.0.time_next_window_starts"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policies.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduling_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cadence", "QUARTERLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cadence_start_month.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cadence_start_month.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_next_window_starts"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseSchedulingPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseSchedulingPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_scheduling_policy" {
			noResourceFound = false
			request := oci_database.GetSchedulingPolicyRequest{}

			tmp := rs.Primary.ID
			request.SchedulingPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetSchedulingPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.SchedulingPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseSchedulingPolicy") {
		resource.AddTestSweepers("DatabaseSchedulingPolicy", &resource.Sweeper{
			Name:         "DatabaseSchedulingPolicy",
			Dependencies: acctest.DependencyGraph["schedulingPolicy"],
			F:            sweepDatabaseSchedulingPolicyResource,
		})
	}
}

func sweepDatabaseSchedulingPolicyResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	schedulingPolicyIds, err := getDatabaseSchedulingPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, schedulingPolicyId := range schedulingPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[schedulingPolicyId]; !ok {
			deleteSchedulingPolicyRequest := oci_database.DeleteSchedulingPolicyRequest{}

			deleteSchedulingPolicyRequest.SchedulingPolicyId = &schedulingPolicyId

			deleteSchedulingPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteSchedulingPolicy(context.Background(), deleteSchedulingPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting SchedulingPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", schedulingPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &schedulingPolicyId, DatabaseSchedulingPolicySweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseSchedulingPolicySweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseSchedulingPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SchedulingPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listSchedulingPoliciesRequest := oci_database.ListSchedulingPoliciesRequest{}
	listSchedulingPoliciesRequest.CompartmentId = &compartmentId
	listSchedulingPoliciesRequest.LifecycleState = oci_database.SchedulingPolicySummaryLifecycleStateNeedsAttention
	listSchedulingPoliciesResponse, err := databaseClient.ListSchedulingPolicies(context.Background(), listSchedulingPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SchedulingPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedulingPolicy := range listSchedulingPoliciesResponse.Items {
		id := *schedulingPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SchedulingPolicyId", id)
	}
	return resourceIds, nil
}

func DatabaseSchedulingPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if schedulingPolicyResponse, ok := response.Response.(oci_database.GetSchedulingPolicyResponse); ok {
		return schedulingPolicyResponse.LifecycleState != oci_database.SchedulingPolicyLifecycleStateDeleted
	}
	return false
}

func DatabaseSchedulingPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetSchedulingPolicy(context.Background(), oci_database.GetSchedulingPolicyRequest{
		SchedulingPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
