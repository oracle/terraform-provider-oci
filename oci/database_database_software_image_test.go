// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatabaseSoftwareImageRequiredOnlyResource = DatabaseSoftwareImageResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Required, Create, databaseSoftwareImageRepresentation)

	DatabaseSoftwareImageResourceConfig = DatabaseSoftwareImageResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentation)

	databaseSoftwareImageSingularDataSourceRepresentation = map[string]interface{}{
		"database_software_image_id": Representation{repType: Required, create: `${oci_database_database_software_image.test_database_software_image.id}`},
	}

	databaseSoftwareImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":       Representation{repType: Optional, create: `image1`, update: `displayName2`},
		"image_shape_family": Representation{repType: Optional, create: `imageShapeFamily`},
		"image_type":         Representation{repType: Optional, create: `imageType`},
		"state":              Representation{repType: Optional, create: `AVAILABLE`},
		"filter":             RepresentationGroup{Required, databaseSoftwareImageDataSourceFilterRepresentation}}
	databaseSoftwareImageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_database_software_image.test_database_software_image.id}`}},
	}

	databaseSoftwareImageRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"database_version": Representation{repType: Required, create: `19c`},
		"display_name":     Representation{repType: Required, create: `image1`, update: `displayName2`},
		"patch_set":        Representation{repType: Required, create: `19.6.0.0.0`},
		"database_software_image_one_off_patches": Representation{repType: Optional, create: []string{`databaseSoftwareImageOneOffPatches`}},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"image_shape_family": Representation{repType: Optional, create: `imageShapeFamily`},
		"image_type":         Representation{repType: Optional, create: `imageType`},
		"ls_inventory":       Representation{repType: Optional, create: `lsInventory`},
	}

	DatabaseSoftwareImageResourceDependencies = DefinedTagsDependencies
)

func TestDatabaseDatabaseSoftwareImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResource_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResource_basic") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResource_basic")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_database_software_image.test_database_software_image"
	datasourceName := "data.oci_database_database_software_images.test_database_software_images"
	singularDatasourceName := "data.oci_database_database_software_image.test_database_software_image"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDatabaseSoftwareImageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Required, Create, databaseSoftwareImageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
					resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0.0"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_shape_family", "imageShapeFamily"),
					resource.TestCheckResourceAttr(resourceName, "image_type", "imageType"),
					resource.TestCheckResourceAttr(resourceName, "ls_inventory", "lsInventory"),
					resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0.0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseSoftwareImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create,
						representationCopyWithNewProperties(databaseSoftwareImageRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_shape_family", "imageShapeFamily"),
					resource.TestCheckResourceAttr(resourceName, "image_type", "imageType"),
					resource.TestCheckResourceAttr(resourceName, "ls_inventory", "lsInventory"),
					resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0.0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_shape_family", "imageShapeFamily"),
					resource.TestCheckResourceAttr(resourceName, "image_type", "imageType"),
					resource.TestCheckResourceAttr(resourceName, "ls_inventory", "lsInventory"),
					resource.TestCheckResourceAttr(resourceName, "patch_set", "19.6.0.0.0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", Optional, Update, databaseSoftwareImageDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseSoftwareImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "imageShapeFamily"),
					resource.TestCheckResourceAttr(datasourceName, "image_type", "imageType"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "database_software_images.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_included_patches.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_one_off_patches.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_version", "19c"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_shape_family", "imageShapeFamily"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_type", "imageType"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.included_patches_summary"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.ls_inventory", "lsInventory"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.patch_set", "19.6.0.0.0"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Required, Create, databaseSoftwareImageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseSoftwareImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_software_image_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_included_patches.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_one_off_patches.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "imageShapeFamily"),
					resource.TestCheckResourceAttr(singularDatasourceName, "image_type", "imageType"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "included_patches_summary"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ls_inventory", "lsInventory"),
					resource.TestCheckResourceAttr(singularDatasourceName, "patch_set", "19.6.0.0.0"),
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
		},
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database_software_image" {
			noResourceFound = false
			request := oci_database.GetDatabaseSoftwareImageRequest{}

			tmp := rs.Primary.ID
			request.DatabaseSoftwareImageId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetDatabaseSoftwareImage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseSoftwareImageLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseDatabaseSoftwareImage") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImage", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImage",
			Dependencies: DependencyGraph["databaseSoftwareImage"],
			F:            sweepDatabaseDatabaseSoftwareImageResource,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseSoftwareImageId := range databaseSoftwareImageIds {
		if ok := SweeperDefaultResourceId[databaseSoftwareImageId]; !ok {
			deleteDatabaseSoftwareImageRequest := oci_database.DeleteDatabaseSoftwareImageRequest{}

			deleteDatabaseSoftwareImageRequest.DatabaseSoftwareImageId = &databaseSoftwareImageId

			deleteDatabaseSoftwareImageRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabaseSoftwareImage(context.Background(), deleteDatabaseSoftwareImageRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseSoftwareImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseSoftwareImageId, error)
				continue
			}
			waitTillCondition(testAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitCondition, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DatabaseSoftwareImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseSoftwareImageId", id)
	}
	return resourceIds, nil
}

func databaseSoftwareImageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
