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
	FleetAppsManagementPatchRequiredOnlyResource = FleetAppsManagementPatchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Required, acctest.Create, FleetAppsManagementPatchRepresentation)

	FleetAppsManagementPatchResourceConfig = FleetAppsManagementPatchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Optional, acctest.Update, FleetAppsManagementPatchRepresentation)

	FleetAppsManagementPatchSingularDataSourceRepresentation = map[string]interface{}{
		"patch_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_patch.test_patch.id}`},
	}

	FleetAppsManagementPatchDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_patch.test_patch.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `TerraformProviderTestPatch10`},
		"patch_type_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.patch_type_platform_configuration_id}`},
		"product_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.product_platform_configuration_id}`},
		"should_compliance_policy_rules_be_applied": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_released_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2023-01-01T00:00:00.111Z`},
		"time_released_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2024-09-09T00:00:00.111Z`},
		"type":                                   acctest.Representation{RepType: acctest.Optional, Create: `USER_DEFINED`},
		"version":                                acctest.Representation{RepType: acctest.Optional, Create: `1.1`},
		"filter":                                 acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPatchDataSourceFilterRepresentation}}
	FleetAppsManagementPatchDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_patch.test_patch.id}`}},
	}

	FleetAppsManagementPatchRepresentation = map[string]interface{}{
		"artifact_details":  acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPatchArtifactDetailsRepresentation},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `TerraformProviderTestPatch10`},
		"patch_type":        acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPatchPatchTypeRepresentation},
		"product":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPatchProductRepresentation},
		"severity":          acctest.Representation{RepType: acctest.Required, Create: `HIGH`, Update: `CRITICAL`},
		"time_released":     acctest.Representation{RepType: acctest.Required, Create: `2024-01-01T00:00:00.111Z`, Update: `2024-02-02T00:00:00.111Z`},
		"dependent_patches": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPatchDependentPatchesRepresentation},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
	}
	FleetAppsManagementPatchArtifactDetailsRepresentation = map[string]interface{}{
		"category": acctest.Representation{RepType: acctest.Required, Create: `GENERIC`},
		"artifact": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPatchArtifactDetailsArtifactRepresentation},
	}
	FleetAppsManagementPatchPatchTypeRepresentation = map[string]interface{}{
		"platform_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${var.patch_type_platform_configuration_id}`},
	}
	FleetAppsManagementPatchProductRepresentation = map[string]interface{}{
		"platform_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${var.product_platform_configuration_id}`},
		"version":                   acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `1.1`},
	}
	FleetAppsManagementPatchDependentPatchesRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `id`, Update: `id2`},
	}
	FleetAppsManagementPatchArtifactDetailsArtifactRepresentation = map[string]interface{}{
		"content": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementPatchArtifactDetailsArtifactContentRepresentation},
	}
	FleetAppsManagementPatchArtifactDetailsArtifactsRepresentation = map[string]interface{}{
		"architecture": acctest.Representation{RepType: acctest.Optional, Create: `ARM_64`, Update: `X64`},
		"content":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementPatchArtifactDetailsArtifactsContentRepresentation},
		"os_type":      acctest.Representation{RepType: acctest.Optional, Create: `WINDOWS`, Update: `LINUX`},
	}
	FleetAppsManagementPatchArtifactDetailsArtifactContentRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"checksum":    acctest.Representation{RepType: acctest.Required, Create: `checksum`, Update: `checksum2`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `object`, Update: `object2`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`},
	}
	FleetAppsManagementPatchArtifactDetailsArtifactsContentRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"checksum":    acctest.Representation{RepType: acctest.Required, Create: `checksum`, Update: `checksum2`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `object`, Update: `object2`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`},
	}

	FleetAppsManagementPatchResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	patchTypePlatformConfigurationId := utils.GetEnvSettingWithBlankDefault("test_patch_type_platform_configuration_id")
	productPlatformConfigurationId := utils.GetEnvSettingWithBlankDefault("test_product_platform_configuration_id")

	platformConfigurationsStr := fmt.Sprintf(
		"variable \"patch_type_platform_configuration_id\" { default = \"%s\" }\n"+
			"variable \"product_platform_configuration_id\" { default = \"%s\" }\n",
		patchTypePlatformConfigurationId, productPlatformConfigurationId)

	resourceName := "oci_fleet_apps_management_patch.test_patch"
	datasourceName := "data.oci_fleet_apps_management_patches.test_patches"
	singularDatasourceName := "data.oci_fleet_apps_management_patch.test_patch"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementPatchResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Optional, acctest.Create, FleetAppsManagementPatchRepresentation), "fleetappsmanagement", "patch", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementPatchDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + platformConfigurationsStr + compartmentIdVariableStr + FleetAppsManagementPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Required, acctest.Create, FleetAppsManagementPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "artifact_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.checksum", "checksum"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.object", "object"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.category", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "TerraformProviderTestPatch10"),
				resource.TestCheckResourceAttr(resourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_type.0.platform_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "product.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "product.0.platform_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "severity", "HIGH"),
				resource.TestCheckResourceAttr(resourceName, "time_released", "2024-01-01T00:00:00.111Z"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + platformConfigurationsStr + compartmentIdVariableStr + FleetAppsManagementPatchResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + platformConfigurationsStr + compartmentIdVariableStr + FleetAppsManagementPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Optional, acctest.Create, FleetAppsManagementPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "artifact_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.checksum", "checksum"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.object", "object"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.category", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dependent_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dependent_patches.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TerraformProviderTestPatch10"),
				resource.TestCheckResourceAttr(resourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_type.0.platform_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "product.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "product.0.platform_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "product.0.version", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "severity", "HIGH"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_released", "2024-01-01T00:00:00.111Z"),
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
			Config: config + platformConfigurationsStr + compartmentIdVariableStr + FleetAppsManagementPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Optional, acctest.Update, FleetAppsManagementPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "artifact_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.checksum", "checksum2"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.object", "object2"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.artifact.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(resourceName, "artifact_details.0.category", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dependent_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dependent_patches.0.id", "id2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TerraformProviderTestPatch10"),
				resource.TestCheckResourceAttr(resourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "patch_type.0.platform_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "product.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "product.0.platform_configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "product.0.version", "1.1"),
				resource.TestCheckResourceAttr(resourceName, "severity", "CRITICAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_released", "2024-02-02T00:00:00.111Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_patches", "test_patches", acctest.Optional, acctest.Update, FleetAppsManagementPatchDataSourceRepresentation) +
				platformConfigurationsStr + compartmentIdVariableStr + FleetAppsManagementPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Optional, acctest.Update, FleetAppsManagementPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "TerraformProviderTestPatch10"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_type_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "product_id"),
				resource.TestCheckResourceAttr(datasourceName, "should_compliance_policy_rules_be_applied", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_released_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_released_less_than"),
				resource.TestCheckResourceAttr(datasourceName, "type", "USER_DEFINED"),
				resource.TestCheckResourceAttr(datasourceName, "version", "1.1"),

				resource.TestCheckResourceAttr(datasourceName, "patch_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "patch_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_patch", "test_patch", acctest.Required, acctest.Create, FleetAppsManagementPatchSingularDataSourceRepresentation) +
				platformConfigurationsStr + compartmentIdVariableStr + FleetAppsManagementPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.0.content.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.0.content.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.0.content.0.checksum", "checksum2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.0.content.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.0.content.0.object", "object2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.artifact.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_details.0.category", "GENERIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dependent_patches.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dependent_patches.0.id", "id2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TerraformProviderTestPatch10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_type.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product.0.version", "1.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "severity", "CRITICAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementPatchRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementPatchDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementOperationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_patch" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetPatchRequest{}

			tmp := rs.Primary.ID
			request.PatchId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetPatch(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.PatchLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementPatch") {
		resource.AddTestSweepers("FleetAppsManagementPatch", &resource.Sweeper{
			Name:         "FleetAppsManagementPatch",
			Dependencies: acctest.DependencyGraph["patch"],
			F:            sweepFleetAppsManagementPatchResource,
		})
	}
}

func sweepFleetAppsManagementPatchResource(compartment string) error {
	fleetAppsManagementOperationsClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementOperationsClient()
	patchIds, err := getFleetAppsManagementPatchIds(compartment)
	if err != nil {
		return err
	}
	for _, patchId := range patchIds {
		if ok := acctest.SweeperDefaultResourceId[patchId]; !ok {
			deletePatchRequest := oci_fleet_apps_management.DeletePatchRequest{}

			deletePatchRequest.PatchId = &patchId

			deletePatchRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementOperationsClient.DeletePatch(context.Background(), deletePatchRequest)
			if error != nil {
				fmt.Printf("Error deleting Patch %s %s, It is possible that the resource is already deleted. Please verify manually \n", patchId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &patchId, FleetAppsManagementPatchSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementPatchSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementPatchIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PatchId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementOperationsClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementOperationsClient()

	listPatchesRequest := oci_fleet_apps_management.ListPatchesRequest{}
	listPatchesRequest.CompartmentId = &compartmentId
	listPatchesRequest.LifecycleState = oci_fleet_apps_management.PatchLifecycleStateActive
	listPatchesResponse, err := fleetAppsManagementOperationsClient.ListPatches(context.Background(), listPatchesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Patch list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, patch := range listPatchesResponse.Items {
		id := *patch.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PatchId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementPatchSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if patchResponse, ok := response.Response.(oci_fleet_apps_management.GetPatchResponse); ok {
		return patchResponse.LifecycleState != oci_fleet_apps_management.PatchLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementPatchSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementOperationsClient().GetPatch(context.Background(), oci_fleet_apps_management.GetPatchRequest{
		PatchId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
