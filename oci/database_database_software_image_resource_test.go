// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_database "github.com/oracle/oci-go-sdk/v40/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatabaseSoftwareImageResourceConfigForExaccShape = DatabaseSoftwareImageResourceDependenciesForExaccShape +
		generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExaccShape)

	databaseSoftwareImageSingularDataSourceRepresentationForExaccShape = map[string]interface{}{
		"database_software_image_id": Representation{repType: Required, create: `${oci_database_database_software_image.test_database_software_image.id}`},
	}

	databaseSoftwareImageDataSourceRepresentationForExaccShape = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":       Representation{repType: Optional, create: `DB_Image_V19`},
		"image_shape_family": Representation{repType: Optional, create: `EXACC_SHAPE`},
		"image_type":         Representation{repType: Optional, create: `DATABASE_IMAGE`},
		"state":              Representation{repType: Optional, create: `AVAILABLE`},
		"filter":             RepresentationGroup{Required, databaseSoftwareImageDataSourceFilterRepresentationForExaccShape}}
	databaseSoftwareImageDataSourceFilterRepresentationForExaccShape = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_database_software_image.test_database_software_image.id}`}},
	}

	databaseSoftwareImageRepresentationForExaccShape = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"database_version":   Representation{repType: Required, create: `19.0.0.0`},
		"display_name":       Representation{repType: Required, create: `DB_Image_V19`, update: `DB_Image_V19_U1`},
		"patch_set":          Representation{repType: Required, create: `19.7.0.0`},
		"image_shape_family": Representation{repType: Optional, create: `EXACC_SHAPE`},
		"database_software_image_one_off_patches": Representation{repType: Optional, create: []string{"29910218", "31113249"}},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"Department": "Exacc_Finance"}, update: map[string]string{"Department": "Exacc_Accounting"}},

		"image_type": Representation{repType: Optional, create: `DATABASE_IMAGE`},
	}

	DatabaseSoftwareImageResourceDependenciesForExaccShape = DefinedTagsDependencies
)

func TestDatabaseDatabaseSoftwareImageResourceForExaccShape(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResourceForExaccShape")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResourceForExaccShape") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResourceForExaccShape")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_database_database_software_image.test_database_software_image"
	datasourceName := "data.oci_database_database_software_images.test_database_software_images"
	singularDatasourceName := "data.oci_database_database_software_image.test_database_software_image"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDatabaseSoftwareImageDestroyForExaccShape,
		Steps: []resource.TestStep{

			// verify creation of Database Software Image with EXACC_SHAPE
			{
				Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaccShape +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageRepresentationForExaccShape),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
					resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
					resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				),
			},
			// verify Update of Database Software Image created for shape EXACC_SHAPE
			{
				Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaccShape +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExaccShape),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19_U1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
					resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
					resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", Optional, Update, databaseSoftwareImageDataSourceRepresentationForExaccShape) +
					compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaccShape +
					generateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Update, databaseSoftwareImageRepresentationForExaccShape),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "DB_Image_V19"),
					resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "EXACC_SHAPE"),
					resource.TestCheckResourceAttr(datasourceName, "image_type", "DATABASE_IMAGE"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_included_patches.#", "2"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_software_image_one_off_patches.#", "2"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.database_version", "19.0.0.0"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.display_name", "DB_Image_V19_U1"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_shape_family", "EXACC_SHAPE"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_type", "DATABASE_IMAGE"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.is_upgrade_supported"),
					resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.patch_set", "19.7.0.0"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_software_images.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", Optional, Create, databaseSoftwareImageSingularDataSourceRepresentationForExaccShape) +
					compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExaccShape,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_software_image_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_included_patches.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_one_off_patches.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "19.0.0.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "DB_Image_V19_U1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "EXACC_SHAPE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "image_type", "DATABASE_IMAGE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "patch_set", "19.7.0.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroyForExaccShape(s *terraform.State) error {
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseDatabaseSoftwareImageForExaccShape") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImageForExaccShape", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImageForExaccShape",
			Dependencies: DependencyGraph["databaseSoftwareImageForExaccShape"],
			F:            sweepDatabaseDatabaseSoftwareImageResourceForExaccShape,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResourceForExaccShape(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIdsForExaccShape(compartment)
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
			waitTillCondition(testAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitConditionForExacc, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationForExacc, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIdsForExaccShape(compartment string) ([]string, error) {
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
	listDatabaseSoftwareImagesRequest.ImageShapeFamily = oci_database.DatabaseSoftwareImageSummaryImageShapeFamilyExaccShape
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

func databaseSoftwareImageSweepWaitConditionForExacc(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted) && (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateTerminated)
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperationForExacc(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
