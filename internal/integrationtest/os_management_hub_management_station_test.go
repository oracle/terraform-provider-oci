// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagementStationRequiredOnlyResource = OsManagementHubManagementStationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Required, acctest.Create, OsManagementHubManagementStationRepresentation)

	OsManagementHubManagementStationResourceConfig = OsManagementHubManagementStationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Optional, acctest.Update, OsManagementHubManagementStationRepresentation)

	OsManagementHubManagementStationSingularDataSourceRepresentation = map[string]interface{}{
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
	}

	OsManagementHubManagementStationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagementStationDataSourceFilterRepresentation}}
	OsManagementHubManagementStationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_management_station.test_management_station.id}`}},
	}

	OsManagementHubManagementStationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"hostname":       acctest.Representation{RepType: acctest.Required, Create: `hostname`, Update: `hostname2`},
		"mirror":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagementStationMirrorRepresentation},
		"proxy":          acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagementStationProxyRepresentation},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	OsManagementHubManagementStationMirrorRepresentation = map[string]interface{}{
		"directory": acctest.Representation{RepType: acctest.Required, Create: `/directory`, Update: `/directory2`},
		"port":      acctest.Representation{RepType: acctest.Required, Create: `50001`, Update: `50011`},
		"sslport":   acctest.Representation{RepType: acctest.Required, Create: `50002`, Update: `50012`},
		"sslcert":   acctest.Representation{RepType: acctest.Optional, Create: `/sslcert`, Update: `/sslcert2`},
	}
	OsManagementHubManagementStationProxyRepresentation = map[string]interface{}{
		"hosts":      acctest.Representation{RepType: acctest.Required, Create: []string{`host`}},
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	OsManagementHubManagementStationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagementStationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagementStationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_management_station.test_management_station"
	datasourceName := "data.oci_os_management_hub_management_stations.test_management_stations"
	singularDatasourceName := "data.oci_os_management_hub_management_station.test_management_station"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagementStationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Optional, acctest.Create, OsManagementHubManagementStationRepresentation), "osmanagementhub", "managementStation", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubManagementStationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagementStationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Required, acctest.Create, OsManagementHubManagementStationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "mirror.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.directory", "/directory"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.port", "50001"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.sslport", "50002"),
				resource.TestCheckResourceAttr(resourceName, "proxy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy.0.is_enabled", "true"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagementStationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagementStationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Optional, acctest.Create, OsManagementHubManagementStationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "mirror.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.directory", "/directory"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.port", "50001"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.sslcert", "/sslcert"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.sslport", "50002"),
				resource.TestCheckResourceAttr(resourceName, "proxy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy.0.is_enabled", "true"),

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
			Config: config + compartmentIdVariableStr + OsManagementHubManagementStationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Optional, acctest.Update, OsManagementHubManagementStationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "mirror.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.directory", "/directory2"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.port", "50011"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.sslcert", "/sslcert2"),
				resource.TestCheckResourceAttr(resourceName, "mirror.0.sslport", "50012"),
				resource.TestCheckResourceAttr(resourceName, "proxy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy.0.is_enabled", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_management_stations", "test_management_stations", acctest.Optional, acctest.Update, OsManagementHubManagementStationDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubManagementStationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Optional, acctest.Update, OsManagementHubManagementStationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "management_station_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_station_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Required, acctest.Create, OsManagementHubManagementStationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubManagementStationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_station_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "hostname", "hostname2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror.0.directory", "/directory2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror.0.port", "50011"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror.0.sslcert", "/sslcert2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror.0.sslport", "50012"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mirror_capacity"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror_sync_status.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "overall_percentage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "overall_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_mirrors"),
			),
		},
		// verify resource import
		{
			Config:                  config + OsManagementHubManagementStationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOsManagementHubManagementStationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementStationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_management_station" {
			noResourceFound = false
			request := oci_os_management_hub.GetManagementStationRequest{}

			tmp := rs.Primary.ID
			request.ManagementStationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetManagementStation(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.ManagementStationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsManagementHubManagementStation") {
		resource.AddTestSweepers("OsManagementHubManagementStation", &resource.Sweeper{
			Name:         "OsManagementHubManagementStation",
			Dependencies: acctest.DependencyGraph["managementStation"],
			F:            sweepOsManagementHubManagementStationResource,
		})
	}
}

func sweepOsManagementHubManagementStationResource(compartment string) error {
	managementStationClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementStationClient()
	managementStationIds, err := getOsManagementHubManagementStationIds(compartment)
	if err != nil {
		return err
	}
	for _, managementStationId := range managementStationIds {
		if ok := acctest.SweeperDefaultResourceId[managementStationId]; !ok {
			deleteManagementStationRequest := oci_os_management_hub.DeleteManagementStationRequest{}

			deleteManagementStationRequest.ManagementStationId = &managementStationId

			deleteManagementStationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := managementStationClient.DeleteManagementStation(context.Background(), deleteManagementStationRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementStation %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementStationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managementStationId, OsManagementHubManagementStationSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubManagementStationSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubManagementStationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementStationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementStationClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementStationClient()

	listManagementStationsRequest := oci_os_management_hub.ListManagementStationsRequest{}
	listManagementStationsRequest.CompartmentId = &compartmentId
	listManagementStationsRequest.LifecycleState = oci_os_management_hub.ManagementStationLifecycleStateActive
	listManagementStationsResponse, err := managementStationClient.ListManagementStations(context.Background(), listManagementStationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementStation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementStation := range listManagementStationsResponse.Items {
		id := *managementStation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementStationId", id)
	}
	return resourceIds, nil
}

func OsManagementHubManagementStationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementStationResponse, ok := response.Response.(oci_os_management_hub.GetManagementStationResponse); ok {
		return managementStationResponse.LifecycleState != oci_os_management_hub.ManagementStationLifecycleStateDeleted
	}
	return false
}

func OsManagementHubManagementStationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementStationClient().GetManagementStation(context.Background(), oci_os_management_hub.GetManagementStationRequest{
		ManagementStationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
