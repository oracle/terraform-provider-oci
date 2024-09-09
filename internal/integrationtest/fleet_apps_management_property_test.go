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
	FleetAppsManagementPropertyRequiredOnlyResource = FleetAppsManagementPropertyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Required, acctest.Create, FleetAppsManagementPropertyRepresentation)

	FleetAppsManagementPropertyResourceConfig = FleetAppsManagementPropertyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Optional, acctest.Update, FleetAppsManagementPropertyRepresentation)

	FleetAppsManagementPropertySingularDataSourceRepresentation = map[string]interface{}{
		"property_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_property.test_property.id}`},
	}

	FleetAppsManagementPropertyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"scope":          acctest.Representation{RepType: acctest.Optional, Create: `TAXONOMY`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPropertyDataSourceFilterRepresentation}}
	FleetAppsManagementPropertyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_property.test_property.id}`}},
	}

	FleetAppsManagementPropertyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"selection":      acctest.Representation{RepType: acctest.Required, Create: `SINGLE_CHOICE`, Update: `MULTI_CHOICE`},
		"value_type":     acctest.Representation{RepType: acctest.Required, Create: `STRING`, Update: `NUMERIC`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"values":         acctest.Representation{RepType: acctest.Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}

	FleetAppsManagementPropertyResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementPropertyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementPropertyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_property.test_property"
	datasourceName := "data.oci_fleet_apps_management_properties.test_properties"
	singularDatasourceName := "data.oci_fleet_apps_management_property.test_property"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementPropertyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Optional, acctest.Create, FleetAppsManagementPropertyRepresentation), "fleetappsmanagement", "property", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementPropertyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Required, acctest.Create, FleetAppsManagementPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "selection", "SINGLE_CHOICE"),
				resource.TestCheckResourceAttr(resourceName, "value_type", "STRING"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPropertyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Optional, acctest.Create, FleetAppsManagementPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_region"),
				resource.TestCheckResourceAttr(resourceName, "selection", "SINGLE_CHOICE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "value_type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "values.#", "1"),

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
			Config: config + compartmentIdVariableStr + FleetAppsManagementPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Optional, acctest.Update, FleetAppsManagementPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_region"),
				resource.TestCheckResourceAttr(resourceName, "selection", "MULTI_CHOICE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "value_type", "NUMERIC"),
				resource.TestCheckResourceAttr(resourceName, "values.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_properties", "test_properties", acctest.Optional, acctest.Update, FleetAppsManagementPropertyDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementPropertyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Optional, acctest.Update, FleetAppsManagementPropertyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "scope", "TAXONOMY"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "property_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "property_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_property", "test_property", acctest.Required, acctest.Create, FleetAppsManagementPropertySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementPropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "property_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scope"),
				resource.TestCheckResourceAttr(singularDatasourceName, "selection", "MULTI_CHOICE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "value_type", "NUMERIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "values.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementPropertyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementPropertyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_property" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetPropertyRequest{}

			tmp := rs.Primary.ID
			request.PropertyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetProperty(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.PropertyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementProperty") {
		resource.AddTestSweepers("FleetAppsManagementProperty", &resource.Sweeper{
			Name:         "FleetAppsManagementProperty",
			Dependencies: acctest.DependencyGraph["property"],
			F:            sweepFleetAppsManagementPropertyResource,
		})
	}
}

func sweepFleetAppsManagementPropertyResource(compartment string) error {
	fleetAppsManagementAdminClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementAdminClient()
	propertyIds, err := getFleetAppsManagementPropertyIds(compartment)
	if err != nil {
		return err
	}
	for _, propertyId := range propertyIds {
		if ok := acctest.SweeperDefaultResourceId[propertyId]; !ok {
			deletePropertyRequest := oci_fleet_apps_management.DeletePropertyRequest{}

			deletePropertyRequest.PropertyId = &propertyId

			deletePropertyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementAdminClient.DeleteProperty(context.Background(), deletePropertyRequest)
			if error != nil {
				fmt.Printf("Error deleting Property %s %s, It is possible that the resource is already deleted. Please verify manually \n", propertyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &propertyId, FleetAppsManagementPropertySweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementPropertySweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementPropertyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PropertyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementAdminClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementAdminClient()

	listPropertiesRequest := oci_fleet_apps_management.ListPropertiesRequest{}
	listPropertiesRequest.CompartmentId = &compartmentId
	listPropertiesRequest.LifecycleState = oci_fleet_apps_management.PropertyLifecycleStateActive
	listPropertiesResponse, err := fleetAppsManagementAdminClient.ListProperties(context.Background(), listPropertiesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Property list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, property := range listPropertiesResponse.Items {
		id := *property.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PropertyId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementPropertySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if propertyResponse, ok := response.Response.(oci_fleet_apps_management.GetPropertyResponse); ok {
		return propertyResponse.LifecycleState != oci_fleet_apps_management.PropertyLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementPropertySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementAdminClient().GetProperty(context.Background(), oci_fleet_apps_management.GetPropertyRequest{
		PropertyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
