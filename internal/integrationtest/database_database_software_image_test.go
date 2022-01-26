// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	DatabaseSoftwareImageRequiredOnlyResource = DatabaseSoftwareImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageRepresentation)

	DatabaseSoftwareImageResourceConfig = DatabaseSoftwareImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentation)

	databaseSoftwareImageSingularDataSourceRepresentation = map[string]interface{}{
		"database_software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database_software_image.test_database_software_image.id}`},
	}

	databaseSoftwareImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `image1`, Update: `displayName2`},
		"image_shape_family": acctest.Representation{RepType: acctest.Optional, Create: `VM_BM_SHAPE`},
		"image_type":         acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_IMAGE`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseSoftwareImageDataSourceFilterRepresentation}}
	databaseSoftwareImageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_database_software_image.test_database_software_image.id}`}},
	}

	databaseSoftwareImageRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version": acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `image1`, Update: `displayName2`},
		"patch_set":        acctest.Representation{RepType: acctest.Required, Create: `19.6.0.0`},
		"database_software_image_one_off_patches": acctest.Representation{RepType: acctest.Optional, Create: []string{"31113249", "27929509"}},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_shape_family": acctest.Representation{RepType: acctest.Optional, Create: `VM_BM_SHAPE`},
		"image_type":         acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_IMAGE`},
		"ls_inventory":       acctest.Representation{RepType: acctest.Optional, Create: `lsInventory`},
	}

	DatabaseSoftwareImageResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseSoftwareImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResource_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResource_basic") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResource_basic")
	}

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_database_software_image.test_database_software_image"
	datasourceName := "data.oci_database_database_software_images.test_database_software_images"
	singularDatasourceName := "data.oci_database_database_software_image.test_database_software_image"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseSoftwareImageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create, databaseSoftwareImageRepresentation), "database", "databaseSoftwareImage", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseSoftwareImageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create, databaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "VM_BM_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(databaseSoftwareImageRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "VM_BM_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "VM_BM_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0"),
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
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, databaseSoftwareImageWaitTillAvailableConditionExa, time.Duration(20*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationExa, "database", true),
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", acctest.Optional, acctest.Update, databaseSoftwareImageDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "VM_BM_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "database_software_images.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_shape_family", "VM_BM_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.is_upgrade_supported"),
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.patch_set", "19.6.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_software_image_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "VM_BM_SHAPE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_upgrade_supported"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_set", "19.6.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database_software_image" {
			noResourceFound = false
			request := oci_database.GetDatabaseSoftwareImageRequest{}

			tmp := rs.Primary.ID
			request.DatabaseSoftwareImageId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetDatabaseSoftwareImage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseSoftwareImageLifecycleStateDeleted): true, string(oci_database.DatabaseSoftwareImageLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseDatabaseSoftwareImage") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImage", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImage",
			Dependencies: acctest.DependencyGraph["databaseSoftwareImage"],
			F:            sweepDatabaseDatabaseSoftwareImageResource,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseSoftwareImageId := range databaseSoftwareImageIds {
		if ok := acctest.SweeperDefaultResourceId[databaseSoftwareImageId]; !ok {
			deleteDatabaseSoftwareImageRequest := oci_database.DeleteDatabaseSoftwareImageRequest{}

			deleteDatabaseSoftwareImageRequest.DatabaseSoftwareImageId = &databaseSoftwareImageId

			deleteDatabaseSoftwareImageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabaseSoftwareImage(context.Background(), deleteDatabaseSoftwareImageRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseSoftwareImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseSoftwareImageId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitCondition, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseSoftwareImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listDatabaseSoftwareImagesRequest := oci_database.ListDatabaseSoftwareImagesRequest{}
	listDatabaseSoftwareImagesRequest.CompartmentId = &compartmentId
	listDatabaseSoftwareImagesRequest.LifecycleState = oci_database.DatabaseSoftwareImageSummaryLifecycleStateAvailable
	listDatabaseSoftwareImagesResponse, err := databaseClient.ListDatabaseSoftwareImages(context.Background(), listDatabaseSoftwareImagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseSoftwareImage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseSoftwareImage := range listDatabaseSoftwareImagesResponse.Items {
		id := *databaseSoftwareImage.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseSoftwareImageId", id)
	}
	return resourceIds, nil
}

func databaseSoftwareImageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted) && (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateTerminated)
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
