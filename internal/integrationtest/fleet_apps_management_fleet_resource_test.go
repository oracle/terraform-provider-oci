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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementFleetResourceRequiredOnlyResource = FleetAppsManagementFleetResourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Required, acctest.Create, FleetAppsManagementFleetResourceRepresentation)

	FleetAppsManagementFleetResourceResourceConfig = FleetAppsManagementFleetResourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Optional, acctest.Update, FleetAppsManagementFleetResourceRepresentation)

	FleetAppsManagementFleetResourceSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"fleet_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet_resource.test_fleet_resource.id}`},
	}

	FleetAppsManagementFleetResourceDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"fleet_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `Instance`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"tenancy_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetResourceDataSourceFilterRepresentation}}
	FleetAppsManagementFleetResourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `resource_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_fleet_resource.test_fleet_resource.id}`}},
	}

	FleetAppsManagementFleetResourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fleet_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"resource_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"resource_region": acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
		"resource_type":   acctest.Representation{RepType: acctest.Required, Create: `Instance`},
	}

	FleetAppsManagementFleetResourceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	region := utils.GetEnvSettingWithBlankDefault("region")
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_fleet_resource.test_fleet_resource"
	datasourceName := "data.oci_fleet_apps_management_fleet_resources.test_fleet_resources"
	singularDatasourceName := "data.oci_fleet_apps_management_fleet_resource.test_fleet_resource"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementFleetResourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Optional, acctest.Create, FleetAppsManagementFleetResourceRepresentation), "fleetappsmanagement", "fleetResource", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementFleetResourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Required, acctest.Create, FleetAppsManagementFleetResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Optional, acctest.Create, FleetAppsManagementFleetResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_region", region),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "Instance"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Optional, acctest.Update, FleetAppsManagementFleetResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_region", region),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "Instance"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_resources", "test_fleet_resources", acctest.Optional, acctest.Update, FleetAppsManagementFleetResourceDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetResourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Optional, acctest.Update, FleetAppsManagementFleetResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_resource_type", "Instance"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(datasourceName, "fleet_resource_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_resource_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Required, acctest.Create, FleetAppsManagementFleetResourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_resource_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compliance_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_region", region),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "Instance"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func testAccCheckFleetAppsManagementFleetResourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_fleet_resource" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetFleetResourceRequest{}

			if value, ok := rs.Primary.Attributes["fleet_id"]; ok {
				request.FleetId = &value
			}

			tmp := rs.Primary.ID
			request.FleetResourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetFleetResource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.FleetResourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementFleetResource") {
		resource.AddTestSweepers("FleetAppsManagementFleetResource", &resource.Sweeper{
			Name:         "FleetAppsManagementFleetResource",
			Dependencies: acctest.DependencyGraph["fleetResource"],
			F:            sweepFleetAppsManagementFleetResourceResource,
		})
	}
}

func sweepFleetAppsManagementFleetResourceResource(compartment string) error {
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()
	fleetResourceIds, err := getFleetAppsManagementFleetResourceIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetResourceId := range fleetResourceIds {
		if ok := acctest.SweeperDefaultResourceId[fleetResourceId]; !ok {
			deleteFleetResourceRequest := oci_fleet_apps_management.DeleteFleetResourceRequest{}

			deleteFleetResourceRequest.FleetResourceId = &fleetResourceId

			deleteFleetResourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementClient.DeleteFleetResource(context.Background(), deleteFleetResourceRequest)
			if error != nil {
				fmt.Printf("Error deleting FleetResource %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetResourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fleetResourceId, FleetAppsManagementFleetResourceSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementFleetResourceSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementFleetResourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FleetResourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()

	listFleetResourcesRequest := oci_fleet_apps_management.ListFleetResourcesRequest{}
	//listFleetResourcesRequest.CompartmentId = &compartmentId

	fleetIds, error := getFleetAppsManagementFleetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting fleetId required for FleetResource resource requests \n")
	}
	for _, fleetId := range fleetIds {
		listFleetResourcesRequest.FleetId = &fleetId

		listFleetResourcesRequest.LifecycleState = oci_fleet_apps_management.FleetResourceLifecycleStateActive
		listFleetResourcesResponse, err := fleetAppsManagementClient.ListFleetResources(context.Background(), listFleetResourcesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FleetResource list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fleetResource := range listFleetResourcesResponse.Items {
			id := *fleetResource.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FleetResourceId", id)
		}

	}
	return resourceIds, nil
}

func FleetAppsManagementFleetResourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetResourceResponse, ok := response.Response.(oci_fleet_apps_management.GetFleetResourceResponse); ok {
		return fleetResourceResponse.LifecycleState != oci_fleet_apps_management.FleetResourceLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementFleetResourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementClient().GetFleetResource(context.Background(), oci_fleet_apps_management.GetFleetResourceRequest{
		FleetResourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
