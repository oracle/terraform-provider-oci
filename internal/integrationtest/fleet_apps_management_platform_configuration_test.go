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
	compatibleProduct        = utils.GetEnvSettingWithBlankDefault("compatible_product")
	compatibleProductUpdated = utils.GetEnvSettingWithBlankDefault("compatible_product_updated")
	credential               = utils.GetEnvSettingWithBlankDefault("credential")
	credentialUpdated        = utils.GetEnvSettingWithBlankDefault("credential_updated")
	patchType                = utils.GetEnvSettingWithBlankDefault("patch_type")
	patchTypeUpdated         = utils.GetEnvSettingWithBlankDefault("patch_type_updated")

	FleetAppsManagementPlatformConfigurationRequiredOnlyResource = FleetAppsManagementPlatformConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Required, acctest.Create, FleetAppsManagementPlatformConfigurationRepresentation)

	FleetAppsManagementPlatformConfigurationResourceConfig = FleetAppsManagementPlatformConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Optional, acctest.Update, FleetAppsManagementPlatformConfigurationRepresentation)

	FleetAppsManagementPlatformConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"platform_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_platform_configuration.test_platform_configuration.id}`},
	}

	FleetAppsManagementPlatformConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"config_category": acctest.Representation{RepType: acctest.Optional, Create: `PRODUCT`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_platform_configuration.test_platform_configuration.id}`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPlatformConfigurationDataSourceFilterRepresentation}}
	FleetAppsManagementPlatformConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_platform_configuration.test_platform_configuration.id}`}},
	}

	FleetAppsManagementPlatformConfigurationRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_category_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPlatformConfigurationConfigCategoryDetailsRepresentation},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsRepresentation = map[string]interface{}{
		"config_category":     acctest.Representation{RepType: acctest.Required, Create: `PRODUCT`},
		"compatible_products": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPlatformConfigurationConfigCategoryDetailsCompatibleProductsRepresentation},
		"components":          acctest.Representation{RepType: acctest.Optional, Create: []string{`components`}, Update: []string{`components2`}},
		"credentials":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPlatformConfigurationConfigCategoryDetailsCredentialsRepresentation},
		"patch_types":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPlatformConfigurationConfigCategoryDetailsPatchTypesRepresentation},
		"versions":            acctest.Representation{RepType: acctest.Required, Create: []string{`versions`}, Update: []string{`versions2`}},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsCompatibleProductsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: compatibleProduct, Update: compatibleProductUpdated},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsCredentialsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: credential, Update: credentialUpdated},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsPatchTypesRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: patchType, Update: patchTypeUpdated},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsProductsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: compatibleProduct, Update: compatibleProductUpdated},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsSubCategoryDetailsRepresentation = map[string]interface{}{
		"sub_category": acctest.Representation{RepType: acctest.Required, Create: `PRODUCT_STACK_GENERIC`, Update: `PRODUCT_STACK_AS_PRODUCT`},
		"components":   acctest.Representation{RepType: acctest.Optional, Create: []string{`components`}, Update: []string{`components2`}},
		"credentials":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPlatformConfigurationConfigCategoryDetailsSubCategoryDetailsCredentialsRepresentation},
		"patch_types":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPlatformConfigurationConfigCategoryDetailsSubCategoryDetailsPatchTypesRepresentation},
		"versions":     acctest.Representation{RepType: acctest.Optional, Create: []string{`versions`}, Update: []string{`versions2`}},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsSubCategoryDetailsCredentialsRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `id`, Update: `id2`},
	}
	FleetAppsManagementPlatformConfigurationConfigCategoryDetailsSubCategoryDetailsPatchTypesRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `id`, Update: `id2`},
	}

	FleetAppsManagementPlatformConfigurationResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementPlatformConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementPlatformConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_platform_configuration.test_platform_configuration"
	datasourceName := "data.oci_fleet_apps_management_platform_configurations.test_platform_configurations"
	singularDatasourceName := "data.oci_fleet_apps_management_platform_configuration.test_platform_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementPlatformConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Optional, acctest.Create, FleetAppsManagementPlatformConfigurationRepresentation), "fleetappsmanagement", "platformConfiguration", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementPlatformConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPlatformConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Required, acctest.Create, FleetAppsManagementPlatformConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.config_category", "PRODUCT"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.products.#"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.versions.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPlatformConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPlatformConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Optional, acctest.Create, FleetAppsManagementPlatformConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.compatible_products.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.compatible_products.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.compatible_products.0.id"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.components.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.config_category", "PRODUCT"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.credentials.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.credentials.0.id"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.patch_types.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.patch_types.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.patch_types.0.id"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.products.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.sub_category_details.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.versions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_region"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementPlatformConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Optional, acctest.Update, FleetAppsManagementPlatformConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.compatible_products.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.compatible_products.0.display_name"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.compatible_products.0.id", compatibleProductUpdated),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.components.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.config_category", "PRODUCT"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.credentials.0.display_name"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.credentials.0.id", credentialUpdated),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.patch_types.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "config_category_details.0.patch_types.0.display_name"),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.patch_types.0.id", patchTypeUpdated),
				resource.TestCheckResourceAttr(resourceName, "config_category_details.0.versions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_region"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_platform_configurations", "test_platform_configurations", acctest.Optional, acctest.Update, FleetAppsManagementPlatformConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementPlatformConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Optional, acctest.Update, FleetAppsManagementPlatformConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "config_category", "PRODUCT"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "platform_configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "platform_configuration_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_platform_configuration", "test_platform_configuration", acctest.Required, acctest.Create, FleetAppsManagementPlatformConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementPlatformConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_configuration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.compatible_products.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_category_details.0.compatible_products.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.compatible_products.0.id", compatibleProductUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.components.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.config_category", "PRODUCT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.credentials.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_category_details.0.credentials.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.credentials.0.id", credentialUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.patch_types.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_category_details.0.patch_types.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.patch_types.0.id", patchTypeUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_category_details.0.versions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementPlatformConfigurationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementPlatformConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_platform_configuration" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetPlatformConfigurationRequest{}

			tmp := rs.Primary.ID
			request.PlatformConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetPlatformConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.PlatformConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementPlatformConfiguration") {
		resource.AddTestSweepers("FleetAppsManagementPlatformConfiguration", &resource.Sweeper{
			Name:         "FleetAppsManagementPlatformConfiguration",
			Dependencies: acctest.DependencyGraph["platformConfiguration"],
			F:            sweepFleetAppsManagementPlatformConfigurationResource,
		})
	}
}

func sweepFleetAppsManagementPlatformConfigurationResource(compartment string) error {
	fleetAppsManagementAdminClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementAdminClient()
	platformConfigurationIds, err := getFleetAppsManagementPlatformConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, platformConfigurationId := range platformConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[platformConfigurationId]; !ok {
			deletePlatformConfigurationRequest := oci_fleet_apps_management.DeletePlatformConfigurationRequest{}

			deletePlatformConfigurationRequest.PlatformConfigurationId = &platformConfigurationId

			deletePlatformConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementAdminClient.DeletePlatformConfiguration(context.Background(), deletePlatformConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting PlatformConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", platformConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &platformConfigurationId, FleetAppsManagementPlatformConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementPlatformConfigurationSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementPlatformConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PlatformConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementAdminClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementAdminClient()

	listPlatformConfigurationsRequest := oci_fleet_apps_management.ListPlatformConfigurationsRequest{}
	listPlatformConfigurationsRequest.CompartmentId = &compartmentId
	listPlatformConfigurationsRequest.LifecycleState = oci_fleet_apps_management.PlatformConfigurationLifecycleStateActive
	listPlatformConfigurationsResponse, err := fleetAppsManagementAdminClient.ListPlatformConfigurations(context.Background(), listPlatformConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PlatformConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, platformConfiguration := range listPlatformConfigurationsResponse.Items {
		id := *platformConfiguration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PlatformConfigurationId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementPlatformConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if platformConfigurationResponse, ok := response.Response.(oci_fleet_apps_management.GetPlatformConfigurationResponse); ok {
		return platformConfigurationResponse.LifecycleState != oci_fleet_apps_management.PlatformConfigurationLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementPlatformConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementAdminClient().GetPlatformConfiguration(context.Background(), oci_fleet_apps_management.GetPlatformConfigurationRequest{
		PlatformConfigurationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
