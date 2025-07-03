// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CapacityManagementOccmDemandSignalRequiredOnlyResource = CapacityManagementOccmDemandSignalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalRepresentation)

	CapacityManagementOccmDemandSignalResourceConfig = CapacityManagementOccmDemandSignalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalRepresentation)

	CapacityManagementOccmDemandSignalSingularDataSourceRepresentation = map[string]interface{}{
		"occm_demand_signal_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id}`},
	}

	CapacityManagementOccmDemandSignalDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CapacityManagementOccmDemandSignalDataSourceFilterRepresentation}}
	CapacityManagementOccmDemandSignalDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id}`}},
	}

	CapacityManagementOccmDemandSignalRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"lifecycle_details": acctest.Representation{RepType: acctest.Optional, Create: "CREATED", Update: "SUBMITTED"},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		//"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CapacityManagementOccmDemandSignalResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccmDemandSignalResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccmDemandSignalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("sp_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_capacity_management_occm_demand_signal.test_occm_demand_signal"
	datasourceName := "data.oci_capacity_management_occm_demand_signals.test_occm_demand_signals"
	singularDatasourceName := "data.oci_capacity_management_occm_demand_signal.test_occm_demand_signal"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CapacityManagementOccmDemandSignalResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Optional, acctest.Create, CapacityManagementOccmDemandSignalRepresentation), "capacitymanagement", "occmDemandSignal", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementOccmDemandSignalDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccmDemandSignalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccmDemandSignalResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccmDemandSignalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Optional, acctest.Create, CapacityManagementOccmDemandSignalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CapacityManagementOccmDemandSignalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_details", "SUBMITTED"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
			),
		},
		// verify datasource
		{

			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occm_demand_signals", "test_occm_demand_signals", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalDataSourceRepresentation) +
				compartmentIdVariableStr + CapacityManagementOccmDemandSignalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "occm_demand_signal_collection.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CapacityManagementOccmDemandSignalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occm_demand_signal_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_details", "SUBMITTED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CapacityManagementOccmDemandSignalRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCapacityManagementOccmDemandSignalDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DemandSignalClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_capacity_management_occm_demand_signal" {
			noResourceFound = false
			request := oci_capacity_management.GetOccmDemandSignalRequest{}

			tmp := rs.Primary.ID
			request.OccmDemandSignalId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")

			response, err := client.GetOccmDemandSignal(context.Background(), request)
			fmt.Printf("DEBUG: Checking destroy for resource %s, err: %v, response: %+v\n", tmp, err, response)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_capacity_management.OccmDemandSignalLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
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
	if !acctest.InSweeperExcludeList("CapacityManagementOccmDemandSignal") {
		resource.AddTestSweepers("CapacityManagementOccmDemandSignal", &resource.Sweeper{
			Name:         "CapacityManagementOccmDemandSignal",
			Dependencies: acctest.DependencyGraph["occmDemandSignal"],
			F:            sweepCapacityManagementOccmDemandSignalResource,
		})
	}
}

func sweepCapacityManagementOccmDemandSignalResource(compartment string) error {
	demandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).DemandSignalClient()
	occmDemandSignalIds, err := getCapacityManagementOccmDemandSignalIds(compartment)
	if err != nil {
		return err
	}
	for _, occmDemandSignalId := range occmDemandSignalIds {
		if ok := acctest.SweeperDefaultResourceId[occmDemandSignalId]; !ok {
			deleteOccmDemandSignalRequest := oci_capacity_management.DeleteOccmDemandSignalRequest{}

			deleteOccmDemandSignalRequest.OccmDemandSignalId = &occmDemandSignalId

			deleteOccmDemandSignalRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")
			_, error := demandSignalClient.DeleteOccmDemandSignal(context.Background(), deleteOccmDemandSignalRequest)
			if error != nil {
				fmt.Printf("Error deleting OccmDemandSignal %s %s, It is possible that the resource is already deleted. Please verify manually \n", occmDemandSignalId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occmDemandSignalId, CapacityManagementOccmDemandSignalSweepWaitCondition, time.Duration(3*time.Minute),
				CapacityManagementOccmDemandSignalSweepResponseFetchOperation, "capacity_management", true)
		}
	}
	return nil
}

func getCapacityManagementOccmDemandSignalIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccmDemandSignalId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	demandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).DemandSignalClient()

	listOccmDemandSignalsRequest := oci_capacity_management.ListOccmDemandSignalsRequest{}
	listOccmDemandSignalsRequest.CompartmentId = &compartmentId
	listOccmDemandSignalsResponse, err := demandSignalClient.ListOccmDemandSignals(context.Background(), listOccmDemandSignalsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccmDemandSignal list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occmDemandSignal := range listOccmDemandSignalsResponse.Items {
		id := *occmDemandSignal.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccmDemandSignalId", id)
	}
	return resourceIds, nil
}

func CapacityManagementOccmDemandSignalSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occmDemandSignalResponse, ok := response.Response.(oci_capacity_management.GetOccmDemandSignalResponse); ok {
		return occmDemandSignalResponse.LifecycleState != oci_capacity_management.OccmDemandSignalLifecycleStateDeleted
	}
	return false
}

func CapacityManagementOccmDemandSignalSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DemandSignalClient().GetOccmDemandSignal(context.Background(), oci_capacity_management.GetOccmDemandSignalRequest{
		OccmDemandSignalId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
