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
	FleetAppsManagementFleetPropertyResourceConfig = FleetAppsManagementFleetPropertyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_property", "test_fleet_property", acctest.Optional, acctest.Update, FleetAppsManagementFleetPropertyRepresentation)

	FleetAppsManagementFleetPropertySingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"fleet_property_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet_property.test_fleet_property.id}`},
	}

	FleetAppsManagementFleetPropertyDataSourceRepresentation = map[string]interface{}{
		"fleet_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetPropertyDataSourceFilterRepresentation}}
	FleetAppsManagementFleetPropertyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_fleet_property.test_fleet_property.id}`}},
	}

	FleetAppsManagementFleetPropertyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"fleet_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"property_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_property.test_property.id}`},
		"value":          acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	FleetAppsManagementFleetPropertyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Required, acctest.Create, FleetAppsManagementPropertyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetPropertyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetPropertyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_fleet_property.test_fleet_property"
	datasourceName := "data.oci_fleet_apps_management_fleet_properties.test_fleet_properties"
	singularDatasourceName := "data.oci_fleet_apps_management_fleet_property.test_fleet_property"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementFleetPropertyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_property", "test_fleet_property", acctest.Required, acctest.Create, FleetAppsManagementFleetPropertyRepresentation), "fleetappsmanagement", "fleetProperty", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementFleetPropertyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_property", "test_fleet_property", acctest.Required, acctest.Create, FleetAppsManagementFleetPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "property_id"),
				resource.TestCheckResourceAttr(resourceName, "value", "value"),

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
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_property", "test_fleet_property", acctest.Optional, acctest.Update, FleetAppsManagementFleetPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "property_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_properties", "test_fleet_properties", acctest.Optional, acctest.Update, FleetAppsManagementFleetPropertyDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_property", "test_fleet_property", acctest.Optional, acctest.Update, FleetAppsManagementFleetPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "fleet_property_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_property_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_property", "test_fleet_property", acctest.Required, acctest.Create, FleetAppsManagementFleetPropertySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetPropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_property_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "value", "value2"),
			),
		},
	})
}

func testAccCheckFleetAppsManagementFleetPropertyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_fleet_property" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetFleetPropertyRequest{}

			if value, ok := rs.Primary.Attributes["fleet_id"]; ok {
				request.FleetId = &value
			}

			tmp := rs.Primary.ID
			request.FleetPropertyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetFleetProperty(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.FleetPropertyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementFleetProperty") {
		resource.AddTestSweepers("FleetAppsManagementFleetProperty", &resource.Sweeper{
			Name:         "FleetAppsManagementFleetProperty",
			Dependencies: acctest.DependencyGraph["fleetProperty"],
			F:            sweepFleetAppsManagementFleetPropertyResource,
		})
	}
}

func sweepFleetAppsManagementFleetPropertyResource(compartment string) error {
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()
	fleetPropertyIds, err := getFleetAppsManagementFleetPropertyIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetPropertyId := range fleetPropertyIds {
		if ok := acctest.SweeperDefaultResourceId[fleetPropertyId]; !ok {
			deleteFleetPropertyRequest := oci_fleet_apps_management.DeleteFleetPropertyRequest{}

			deleteFleetPropertyRequest.FleetPropertyId = &fleetPropertyId

			deleteFleetPropertyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementClient.DeleteFleetProperty(context.Background(), deleteFleetPropertyRequest)
			if error != nil {
				fmt.Printf("Error deleting FleetProperty %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetPropertyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fleetPropertyId, FleetAppsManagementFleetPropertySweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementFleetPropertySweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementFleetPropertyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FleetPropertyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()

	listFleetPropertiesRequest := oci_fleet_apps_management.ListFleetPropertiesRequest{}
	listFleetPropertiesRequest.CompartmentId = &compartmentId

	fleetIds, error := getFleetAppsManagementFleetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting fleetId required for FleetProperty resource requests \n")
	}
	for _, fleetId := range fleetIds {
		listFleetPropertiesRequest.FleetId = &fleetId

		listFleetPropertiesRequest.LifecycleState = oci_fleet_apps_management.FleetPropertyLifecycleStateActive
		listFleetPropertiesResponse, err := fleetAppsManagementClient.ListFleetProperties(context.Background(), listFleetPropertiesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FleetProperty list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fleetProperty := range listFleetPropertiesResponse.Items {
			id := *fleetProperty.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FleetPropertyId", id)
		}

	}
	return resourceIds, nil
}

func FleetAppsManagementFleetPropertySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetPropertyResponse, ok := response.Response.(oci_fleet_apps_management.GetFleetPropertyResponse); ok {
		return fleetPropertyResponse.LifecycleState != oci_fleet_apps_management.FleetPropertyLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementFleetPropertySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementClient().GetFleetProperty(context.Background(), oci_fleet_apps_management.GetFleetPropertyRequest{
		FleetPropertyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
