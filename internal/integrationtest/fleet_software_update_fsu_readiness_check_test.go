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
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetSoftwareUpdateFsuReadinessCheckRequiredOnlyResource = FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuReadinessCheckRepresentation)

	FleetSoftwareUpdateFsuReadinessCheckResourceConfig = FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuReadinessCheckRepresentation)

	FleetSoftwareUpdateFsuReadinessCheckSingularDataSourceRepresentation = map[string]interface{}{
		"fsu_readiness_check_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_software_update_fsu_readiness_check.test_fsu_readiness_check.id}`},
	}

	FleetSoftwareUpdateFsuReadinessCheckDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `TARGET`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuReadinessCheckDataSourceFilterRepresentation}}
	FleetSoftwareUpdateFsuReadinessCheckDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_software_update_fsu_readiness_check.test_fsu_readiness_check.id}`}},
	}

	FleetSoftwareUpdateFsuReadinessCheckRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `TARGET`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"targets":        acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuReadinessCheckTargetsRepresentation},
	}
	FleetSoftwareUpdateFsuReadinessCheckTargetsRepresentation = map[string]interface{}{
		"entity_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_target_1}`},
		"entity_type": acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
	}

	FleetSoftwareUpdateFsuReadinessCheckResourceDependencies = ""
)

// issue-routing-tag: fleet_software_update/default
func TestFleetSoftwareUpdateFsuReadinessCheckResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetSoftwareUpdateFsuReadinessCheckResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	dbTargetId1 := utils.GetEnvSettingWithBlankDefault("fsu_db_23_target_1")
	dbTargetId1VariableStr := fmt.Sprintf("variable \"db_target_1\" { default = \"%s\" }\n", dbTargetId1)

	var variablesStr = compartmentIdVariableStr + dbTargetId1VariableStr

	resourceName := "oci_fleet_software_update_fsu_readiness_check.test_fsu_readiness_check"
	datasourceName := "data.oci_fleet_software_update_fsu_readiness_checks.test_fsu_readiness_checks"
	singularDatasourceName := "data.oci_fleet_software_update_fsu_readiness_check.test_fsu_readiness_check"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	var testConfig = config + variablesStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Optional,
			acctest.Create, FleetSoftwareUpdateFsuReadinessCheckRepresentation)
	acctest.SaveConfigContent(testConfig, "fleetsoftwareupdate", "fsuReadinessCheck", t)
	fmt.Printf("FSU_TEST_LOG CONF:\n%s\n", testConfig)

	acctest.ResourceTest(t, testAccCheckFleetSoftwareUpdateFsuReadinessCheckDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuReadinessCheckRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "TARGET"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbTargetId1VariableStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dbTargetId1VariableStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuReadinessCheckRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "issue_count"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "targets.0.entity_id"),
				resource.TestCheckResourceAttr(resourceName, "targets.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "TARGET"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + dbTargetId1VariableStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetSoftwareUpdateFsuReadinessCheckRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "issue_count"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "targets.0.entity_id"),
				resource.TestCheckResourceAttr(resourceName, "targets.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "TARGET"),

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
			Config: config + compartmentIdVariableStr + dbTargetId1VariableStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuReadinessCheckRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "issue_count"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "targets.0.entity_id"),
				resource.TestCheckResourceAttr(resourceName, "targets.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "TARGET"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_checks", "test_fsu_readiness_checks", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuReadinessCheckDataSourceRepresentation) +
				compartmentIdVariableStr + dbTargetId1VariableStr + FleetSoftwareUpdateFsuReadinessCheckResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuReadinessCheckRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),
				resource.TestCheckResourceAttr(datasourceName, "type", "TARGET"),

				resource.TestCheckResourceAttr(datasourceName, "fsu_readiness_check_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_readiness_check_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_readiness_check", "test_fsu_readiness_check", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuReadinessCheckSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbTargetId1VariableStr + FleetSoftwareUpdateFsuReadinessCheckResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fsu_readiness_check_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "issue_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "issues.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "targets.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "TARGET"),
			),
		},
		// verify resource import
		{
			Config:                  config + variablesStr + FleetSoftwareUpdateFsuReadinessCheckRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetSoftwareUpdateFsuReadinessCheckDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetSoftwareUpdateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_software_update_fsu_readiness_check" {
			noResourceFound = false
			request := oci_fleet_software_update.GetFsuReadinessCheckRequest{}

			tmp := rs.Primary.ID
			request.FsuReadinessCheckId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_software_update")

			response, err := client.GetFsuReadinessCheck(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_software_update.FsuReadinessCheckLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetSoftwareUpdateFsuReadinessCheck") {
		resource.AddTestSweepers("FleetSoftwareUpdateFsuReadinessCheck", &resource.Sweeper{
			Name:         "FleetSoftwareUpdateFsuReadinessCheck",
			Dependencies: acctest.DependencyGraph["fsuReadinessCheck"],
			F:            sweepFleetSoftwareUpdateFsuReadinessCheckResource,
		})
	}
}

func sweepFleetSoftwareUpdateFsuReadinessCheckResource(compartment string) error {
	fleetSoftwareUpdateClient := acctest.GetTestClients(&schema.ResourceData{}).FleetSoftwareUpdateClient()
	fsuReadinessCheckIds, err := getFleetSoftwareUpdateFsuReadinessCheckIds(compartment)
	if err != nil {
		return err
	}
	for _, fsuReadinessCheckId := range fsuReadinessCheckIds {
		if ok := acctest.SweeperDefaultResourceId[fsuReadinessCheckId]; !ok {
			deleteFsuReadinessCheckRequest := oci_fleet_software_update.DeleteFsuReadinessCheckRequest{}

			deleteFsuReadinessCheckRequest.FsuReadinessCheckId = &fsuReadinessCheckId

			deleteFsuReadinessCheckRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_software_update")
			_, err := fleetSoftwareUpdateClient.DeleteFsuReadinessCheck(context.Background(), deleteFsuReadinessCheckRequest)
			if err != nil {
				fmt.Printf("Error deleting FsuReadinessCheck %s %s, It is possible that the resource is already deleted. Please verify manually \n", fsuReadinessCheckId, err)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fsuReadinessCheckId, FleetSoftwareUpdateFsuReadinessCheckSweepWaitCondition, time.Duration(3*time.Minute),
				FleetSoftwareUpdateFsuReadinessCheckSweepResponseFetchOperation, "fleet_software_update", true)
		}
	}
	return nil
}

func getFleetSoftwareUpdateFsuReadinessCheckIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FsuReadinessCheckId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetSoftwareUpdateClient := acctest.GetTestClients(&schema.ResourceData{}).FleetSoftwareUpdateClient()

	// Get NeedsAttention readiness checks
	listFsuReadinessChecksRequest1 := oci_fleet_software_update.ListFsuReadinessChecksRequest{}
	listFsuReadinessChecksRequest1.CompartmentId = &compartmentId
	listFsuReadinessChecksRequest1.LifecycleState = oci_fleet_software_update.FsuReadinessCheckLifecycleStateNeedsAttention
	listFsuReadinessChecksResponse1, err := fleetSoftwareUpdateClient.ListFsuReadinessChecks(context.Background(), listFsuReadinessChecksRequest1)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting FsuReadinessCheck list for compartment id : %s , %s \n", compartmentId, err)
	}

	// Get Succeeded readiness checks
	listFsuReadinessChecksRequest2 := oci_fleet_software_update.ListFsuReadinessChecksRequest{}
	listFsuReadinessChecksRequest2.CompartmentId = &compartmentId
	listFsuReadinessChecksRequest2.LifecycleState = oci_fleet_software_update.FsuReadinessCheckLifecycleStateSucceeded
	listFsuReadinessChecksResponse2, err := fleetSoftwareUpdateClient.ListFsuReadinessChecks(context.Background(), listFsuReadinessChecksRequest2)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting FsuReadinessCheck list for compartment id : %s , %s \n", compartmentId, err)
	}

	// Combine results
	allFsuReadinessChecks := append(listFsuReadinessChecksResponse1.Items, listFsuReadinessChecksResponse2.Items...)

	for _, fsuReadinessCheck := range allFsuReadinessChecks {
		id := *fsuReadinessCheck.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FsuReadinessCheckId", id)
	}
	return resourceIds, nil
}

func FleetSoftwareUpdateFsuReadinessCheckSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fsuReadinessCheckResponse, ok := response.Response.(oci_fleet_software_update.GetFsuReadinessCheckResponse); ok {
		return fsuReadinessCheckResponse.GetLifecycleState() != oci_fleet_software_update.FsuReadinessCheckLifecycleStateDeleted
	}
	return false
}

func FleetSoftwareUpdateFsuReadinessCheckSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetSoftwareUpdateClient().GetFsuReadinessCheck(context.Background(), oci_fleet_software_update.GetFsuReadinessCheckRequest{
		FsuReadinessCheckId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
