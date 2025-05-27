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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	currentRegion           = utils.GetEnvSettingWithBlankDefault("region")
	tenancyId               = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	userId                  = utils.GetEnvSettingWithBlankDefault("user_id")
	testActiveFleet         = utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	testCatalogPackage      = utils.GetEnvSettingWithBlankDefault("test_catalog_package")
	testCatalogConfig       = utils.GetEnvSettingWithBlankDefault("test_catalog_config")
	testCatalogSimpleConfig = utils.GetEnvSettingWithBlankDefault("test_catalog_simple_config")
	tfCompartmentId         = utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)

	FleetAppsManagementProvisionRequiredOnlyResource = FleetAppsManagementProvisionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Required, acctest.Create, FleetAppsManagementProvisionRepresentation)

	FleetAppsManagementProvisionResourceConfig = FleetAppsManagementProvisionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Optional, acctest.Update, FleetAppsManagementProvisionRepresentation)

	FleetAppsManagementProvisionSingularDataSourceRepresentation = map[string]interface{}{
		"provision_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_provision.test_provision.id}`},
	}

	FleetAppsManagementProvisionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fleet_id":       acctest.Representation{RepType: acctest.Optional, Create: testActiveFleet},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementProvisionDataSourceFilterRepresentation}}
	FleetAppsManagementProvisionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_provision.test_provision.id}`}},
	}

	FleetAppsManagementProvisionRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_catalog_item_id":      acctest.Representation{RepType: acctest.Required, Create: testCatalogConfig},
		"fleet_id":                    acctest.Representation{RepType: acctest.Required, Create: testActiveFleet},
		"package_catalog_item_id":     acctest.Representation{RepType: acctest.Required, Create: testCatalogPackage},
		"tf_variable_region_id":       acctest.Representation{RepType: acctest.Required, Create: currentRegion},
		"tf_variable_tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: tenancyId},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"provision_description":       acctest.Representation{RepType: acctest.Optional, Create: `provisionDescription`, Update: `provisionDescription2`},
		"tf_variable_compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: tfCompartmentId},
		"tf_variable_current_user_id": acctest.Representation{RepType: acctest.Optional, Create: userId},
	}

	FleetAppsManagementProvisionResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementProvisionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementProvisionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_fleet_apps_management_provision.test_provision"
	datasourceName := "data.oci_fleet_apps_management_provisions.test_provisions"
	singularDatasourceName := "data.oci_fleet_apps_management_provision.test_provision"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementProvisionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Optional, acctest.Create, FleetAppsManagementProvisionRepresentation), "fleetappsmanagement", "provision", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementProvisionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementProvisionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetAppsManagementProvisionRepresentation, map[string]interface{}{
						"config_catalog_item_id": acctest.Representation{RepType: acctest.Required, Create: testCatalogSimpleConfig},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_region_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_tenancy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementProvisionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementProvisionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Optional, acctest.Create, FleetAppsManagementProvisionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(resourceName, "provision_description", "provisionDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_current_user_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_region_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_tenancy_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FleetAppsManagementProvisionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetAppsManagementProvisionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(resourceName, "provision_description", "provisionDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_current_user_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_region_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_tenancy_id"),
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
			Config: config + compartmentIdVariableStr + FleetAppsManagementProvisionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Optional, acctest.Update, FleetAppsManagementProvisionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "config_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(resourceName, "provision_description", "provisionDescription2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_current_user_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_region_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tf_variable_tenancy_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_provisions", "test_provisions", acctest.Optional, acctest.Update, FleetAppsManagementProvisionDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementProvisionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Optional, acctest.Update, FleetAppsManagementProvisionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "provision_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "provision_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_provision", "test_provision", acctest.Required, acctest.Create, FleetAppsManagementProvisionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementProvisionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provision_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployed_resources.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_catalog_item_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_catalog_item_listing_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_catalog_item_listing_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provision_description", "provisionDescription2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rms_apply_job_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tf_outputs.#", "4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementProvisionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementProvisionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementProvisionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_provision" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetProvisionRequest{}

			tmp := rs.Primary.ID
			request.ProvisionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetProvision(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.ProvisionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementProvision") {
		resource.AddTestSweepers("FleetAppsManagementProvision", &resource.Sweeper{
			Name:         "FleetAppsManagementProvision",
			Dependencies: acctest.DependencyGraph["provision"],
			F:            sweepFleetAppsManagementProvisionResource,
		})
	}
}

func sweepFleetAppsManagementProvisionResource(compartment string) error {
	fleetAppsManagementProvisionClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementProvisionClient()
	provisionIds, err := getFleetAppsManagementProvisionIds(compartment)
	if err != nil {
		return err
	}
	for _, provisionId := range provisionIds {
		if ok := acctest.SweeperDefaultResourceId[provisionId]; !ok {
			deleteProvisionRequest := oci_fleet_apps_management.DeleteProvisionRequest{}

			deleteProvisionRequest.ProvisionId = &provisionId

			deleteProvisionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementProvisionClient.DeleteProvision(context.Background(), deleteProvisionRequest)
			if error != nil {
				fmt.Printf("Error deleting Provision %s %s, It is possible that the resource is already deleted. Please verify manually \n", provisionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &provisionId, FleetAppsManagementProvisionSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementProvisionSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementProvisionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProvisionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementProvisionClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementProvisionClient()

	listProvisionsRequest := oci_fleet_apps_management.ListProvisionsRequest{}
	listProvisionsRequest.CompartmentId = &compartmentId
	listProvisionsRequest.LifecycleState = oci_fleet_apps_management.ProvisionLifecycleStateActive
	listProvisionsResponse, err := fleetAppsManagementProvisionClient.ListProvisions(context.Background(), listProvisionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Provision list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, provision := range listProvisionsResponse.Items {
		id := *provision.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProvisionId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementProvisionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if provisionResponse, ok := response.Response.(oci_fleet_apps_management.GetProvisionResponse); ok {
		return provisionResponse.LifecycleState != oci_fleet_apps_management.ProvisionLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementProvisionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementProvisionClient().GetProvision(context.Background(), oci_fleet_apps_management.GetProvisionRequest{
		ProvisionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
