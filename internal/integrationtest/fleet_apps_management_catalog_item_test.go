// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
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
	accessUri       = utils.GetEnvSettingWithBlankDefault("test_access_uri")
	maskedAccessUri = MaskUri(accessUri)

	FleetAppsManagementCatalogItemRequiredOnlyResource = FleetAppsManagementCatalogItemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Required, acctest.Create, FleetAppsManagementCatalogItemRepresentation)

	FleetAppsManagementCatalogItemResourceConfig = FleetAppsManagementCatalogItemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Optional, acctest.Update, FleetAppsManagementCatalogItemRepresentation)

	FleetAppsManagementCatalogItemSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_item_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_catalog_item.test_catalog_item.id}`},
	}

	FleetAppsManagementCatalogItemDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// "catalog_listing_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_listing.test_listing.id}`},
		"catalog_listing_version_criteria": acctest.Representation{RepType: acctest.Optional, Create: `LIST_ALL_VERSIONS`},
		"config_source_type":               acctest.Representation{RepType: acctest.Optional, Create: `PAR_CATALOG_SOURCE`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"package_type":                     acctest.Representation{RepType: acctest.Optional, Create: `TF_PACKAGE`},
		"should_list_public_items":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementCatalogItemDataSourceFilterRepresentation}}
	FleetAppsManagementCatalogItemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_catalog_item.test_catalog_item.id}`}},
	}

	FleetAppsManagementCatalogItemRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_source_type":     acctest.Representation{RepType: acctest.Required, Create: `PAR_CATALOG_SOURCE`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"short_description":      acctest.Representation{RepType: acctest.Required, Create: `shortDescription`, Update: `shortDescription2`},
		"description":            acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"version_description":    acctest.Representation{RepType: acctest.Required, Create: `V1`, Update: `V2`},
		"package_type":           acctest.Representation{RepType: acctest.Required, Create: `TF_PACKAGE`},
		"catalog_source_payload": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementCatalogItemCatalogSourcePayloadRepresentation},
		// "listing_version":        acctest.Representation{RepType: acctest.Required, Create: `1.0.0`},

		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("Oracle-Tags.CreatedBy", "value")}`, Update: `${map("Oracle-Tags.CreatedBy", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		// "is_item_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		// "listing_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_listing.test_listing.id}`},
		"time_released": acctest.Representation{RepType: acctest.Optional, Create: `2025-10-27T00:00:00.000Z`},
		// "clone_catalog_item_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "configure_trigger":          acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	FleetAppsManagementCatalogItemCatalogSourcePayloadRepresentation = map[string]interface{}{
		"config_source_type": acctest.Representation{RepType: acctest.Required, Create: `PAR_CATALOG_SOURCE`},
		"bucket":             acctest.Representation{RepType: acctest.Required, Create: `test-catalog-bucket`},
		"namespace":          acctest.Representation{RepType: acctest.Required, Create: `axfaohqonho7`},
		"object":             acctest.Representation{RepType: acctest.Required, Create: `ObjectStorageCatalog.zip`},

		"access_uri":        acctest.Representation{RepType: acctest.Optional, Create: accessUri},
		"time_expires":      acctest.Representation{RepType: acctest.Optional, Create: `2029-12-31T00:00:00Z`},
		"working_directory": acctest.Representation{RepType: acctest.Optional, Create: `workingDirectory`},
	}

	FleetAppsManagementCatalogItemResourceDependencies = "" // TODO Temp removed: DefinedTagsDependencies

)

func MaskUri(uriString string) string {
	re, err := regexp.Compile("/p/[A-Za-z0-9-]+/n/")
	if err != nil {
		return ""
		// Handle error, the pattern is invalid
	}
	return re.ReplaceAllString(uriString, "/p/***/n/")
}

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementCatalogItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementCatalogItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_fleet_apps_management_catalog_item.test_catalog_item"
	datasourceName := "data.oci_fleet_apps_management_catalog_items.test_catalog_items"
	singularDatasourceName := "data.oci_fleet_apps_management_catalog_item.test_catalog_item"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementCatalogItemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Optional, acctest.Create, FleetAppsManagementCatalogItemRepresentation), "fleetappsmanagement", "catalogItem", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementCatalogItemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementCatalogItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Required, acctest.Create, FleetAppsManagementCatalogItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "package_type", "TF_PACKAGE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementCatalogItemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementCatalogItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Optional, acctest.Create, FleetAppsManagementCatalogItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.access_uri", maskedAccessUri),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.bucket", "test-catalog-bucket"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.namespace", "axfaohqonho7"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.object", "ObjectStorageCatalog.zip"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.time_expires", "2029-12-31T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.working_directory", "workingDirectory"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttr(resourceName, "is_item_locked", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttr(resourceName, "listing_version", "1.0.0"),
				resource.TestCheckResourceAttr(resourceName, "package_type", "TF_PACKAGE"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_released"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "version_description", "V1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FleetAppsManagementCatalogItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetAppsManagementCatalogItemRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.access_uri", maskedAccessUri),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.bucket", "test-catalog-bucket"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.namespace", "axfaohqonho7"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.object", "ObjectStorageCatalog.zip"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.time_expires", "2029-12-31T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.working_directory", "workingDirectory"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttr(resourceName, "is_item_locked", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttr(resourceName, "listing_version", "1.0.0"),
				resource.TestCheckResourceAttr(resourceName, "package_type", "TF_PACKAGE"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_released"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "version_description", "V1"),

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
			Config: config + compartmentIdVariableStr + FleetAppsManagementCatalogItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Optional, acctest.Update, FleetAppsManagementCatalogItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.access_uri", maskedAccessUri),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.bucket", "test-catalog-bucket"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.namespace", "axfaohqonho7"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.object", "ObjectStorageCatalog.zip"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.time_expires", "2029-12-31T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "catalog_source_payload.0.working_directory", "workingDirectory"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttr(resourceName, "is_item_locked", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttr(resourceName, "listing_version", "1.0.0"),
				resource.TestCheckResourceAttr(resourceName, "package_type", "TF_PACKAGE"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_released"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "version_description", "V2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_catalog_items", "test_catalog_items", acctest.Optional, acctest.Update, FleetAppsManagementCatalogItemDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementCatalogItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Optional, acctest.Update, FleetAppsManagementCatalogItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "package_type", "TF_PACKAGE"),
				resource.TestCheckResourceAttr(datasourceName, "should_list_public_items", "false"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_item_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_item_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_catalog_item", "test_catalog_item", acctest.Required, acctest.Create, FleetAppsManagementCatalogItemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementCatalogItemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_item_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_result_payload.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.0.access_uri", maskedAccessUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.0.bucket", "test-catalog-bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.0.config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.0.namespace", "axfaohqonho7"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.0.object", "ObjectStorageCatalog.zip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_source_payload.0.time_expires"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_source_payload.0.working_directory", "workingDirectory"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_source_type", "PAR_CATALOG_SOURCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "is_item_locked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listing_version", "1.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "package_type", "TF_PACKAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_backfill_last_checked"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version_description", "V2"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementCatalogItemRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementCatalogItemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_catalog_item" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetCatalogItemRequest{}

			tmp := rs.Primary.ID
			request.CatalogItemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetCatalogItem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.CatalogItemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementCatalogItem") {
		resource.AddTestSweepers("FleetAppsManagementCatalogItem", &resource.Sweeper{
			Name:         "FleetAppsManagementCatalogItem",
			Dependencies: acctest.DependencyGraph["catalogItem"],
			F:            sweepFleetAppsManagementCatalogItemResource,
		})
	}
}

func sweepFleetAppsManagementCatalogItemResource(compartment string) error {
	fleetAppsManagementCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementCatalogClient()
	catalogItemIds, err := getFleetAppsManagementCatalogItemIds(compartment)
	if err != nil {
		return err
	}
	for _, catalogItemId := range catalogItemIds {
		if ok := acctest.SweeperDefaultResourceId[catalogItemId]; !ok {
			deleteCatalogItemRequest := oci_fleet_apps_management.DeleteCatalogItemRequest{}

			deleteCatalogItemRequest.CatalogItemId = &catalogItemId

			deleteCatalogItemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementCatalogClient.DeleteCatalogItem(context.Background(), deleteCatalogItemRequest)
			if error != nil {
				fmt.Printf("Error deleting CatalogItem %s %s, It is possible that the resource is already deleted. Please verify manually \n", catalogItemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &catalogItemId, FleetAppsManagementCatalogItemSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementCatalogItemSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementCatalogItemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CatalogItemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementCatalogClient()

	listCatalogItemsRequest := oci_fleet_apps_management.ListCatalogItemsRequest{}
	listCatalogItemsRequest.CompartmentId = &compartmentId
	listCatalogItemsRequest.LifecycleState = oci_fleet_apps_management.CatalogItemLifecycleStateActive
	listCatalogItemsResponse, err := fleetAppsManagementCatalogClient.ListCatalogItems(context.Background(), listCatalogItemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CatalogItem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, catalogItem := range listCatalogItemsResponse.Items {
		id := *catalogItem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CatalogItemId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementCatalogItemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if catalogItemResponse, ok := response.Response.(oci_fleet_apps_management.GetCatalogItemResponse); ok {
		return catalogItemResponse.LifecycleState != oci_fleet_apps_management.CatalogItemLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementCatalogItemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementCatalogClient().GetCatalogItem(context.Background(), oci_fleet_apps_management.GetCatalogItemRequest{
		CatalogItemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
