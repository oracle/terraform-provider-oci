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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	DatabaseSoftwareImageResourceConfigForExaccShape = DatabaseSoftwareImageResourceDependenciesForExaShape +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentationForExaccShape)

	DatabaseSoftwareImageResourceConfigForExadataShape = DatabaseSoftwareImageResourceDependenciesForExaShape +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentationForExadataShape)

	databaseSoftwareImageSingularDataSourceRepresentationForExaccShape = map[string]interface{}{
		"database_software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database_software_image.test_database_software_image.id}`},
	}

	databaseSoftwareImageDataSourceRepresentationForExaccShape = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `DB_Image_V19`},
		"image_shape_family": acctest.Representation{RepType: acctest.Optional, Create: `EXACC_SHAPE`},
		"image_type":         acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_IMAGE`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseSoftwareImageDataSourceFilterRepresentationForExaccShape}}
	databaseSoftwareImageDataSourceFilterRepresentationForExaccShape = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_database_software_image.test_database_software_image.id}`}},
	}

	databaseSoftwareImageDataSourceRepresentationForExadataShape = acctest.GetUpdatedRepresentationCopy("image_shape_family",
		acctest.Representation{RepType: acctest.Required, Create: `EXADATA_SHAPE`},
		databaseSoftwareImageDataSourceRepresentationForExaccShape)

	databaseSoftwareImageDataSourceRepresentationForVmBmShape = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `VMBM_DB_Image_V19`},
		"database_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"patch_set":          acctest.Representation{RepType: acctest.Required, Create: `19.7.0.0`},
		"image_shape_family": acctest.Representation{RepType: acctest.Optional, Create: `VMBM_SHAPE`},
		"image_type":         acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_IMAGE`},
	}

	databaseSoftwareImageRepresentationForExaccShape = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `DB_Image_V19`, Update: `DB_Image_V19_U1`},
		"patch_set":          acctest.Representation{RepType: acctest.Required, Create: `19.7.0.0`},
		"image_shape_family": acctest.Representation{RepType: acctest.Optional, Create: `EXACC_SHAPE`},
		"database_software_image_one_off_patches": acctest.Representation{RepType: acctest.Optional, Create: []string{"29910218", "31113249"}},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Exacc_Finance"}, Update: map[string]string{"Department": "Exacc_Accounting"}},

		"image_type": acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_IMAGE`},
	}

	databaseSoftwareImageRepresentationForExadataShape = acctest.GetUpdatedRepresentationCopy("image_shape_family",
		acctest.Representation{RepType: acctest.Required, Create: `EXADATA_SHAPE`},
		databaseSoftwareImageRepresentationForExaccShape)

	DatabaseSoftwareImageResourceDependenciesForExaShape = DefinedTagsDependencies

	DatabaseSoftwareImageResourceDependenciesForSourceDbHome = acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentation) +
		ExaBaseDependencies + AvailabilityDomainConfig +
		DatabaseSoftwareImageResourceDependenciesForExaShape

	databaseSoftwareImageRepresentationForSourceDbHome = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version": acctest.Representation{RepType: acctest.Optional, Create: `12.1.0.2`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `image1`, Update: `displayName2`},
		"patch_set":        acctest.Representation{RepType: acctest.Optional, Create: `12.1.0.2.210119`},
		"database_software_image_one_off_patches": acctest.Representation{RepType: acctest.Optional, Create: []string{"31113249", "27929509"}},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image_shape_family": acctest.Representation{RepType: acctest.Optional, Create: `EXADATA_SHAPE`},
		"image_type":         acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_IMAGE`},
		"ls_inventory":       acctest.Representation{RepType: acctest.Optional, Create: nil},
		"source_db_home_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseSoftwareImageResourceForExaccShape(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResourceForExaccShape")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResourceForExaccShape") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResourceForExaccShape")
	}

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_database_database_software_image.test_database_software_image"
	datasourceName := "data.oci_database_database_software_images.test_database_software_images"
	singularDatasourceName := "data.oci_database_database_software_image.test_database_software_image"

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseSoftwareImageDestroyForExaccShape, []resource.TestStep{

		// verify creation of Database Software Image with EXACC_SHAPE
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create, databaseSoftwareImageRepresentationForExaccShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentationForExaccShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", acctest.Optional, acctest.Update, databaseSoftwareImageDataSourceRepresentationForExaccShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentationForExaccShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create, databaseSoftwareImageSingularDataSourceRepresentationForExaccShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExaccShape,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroyForExaccShape(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("DatabaseDatabaseSoftwareImageForExaccShape") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImageForExaccShape", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImageForExaccShape",
			Dependencies: acctest.DependencyGraph["databaseSoftwareImageForExaccShape"],
			F:            sweepDatabaseDatabaseSoftwareImageResourceForExaccShape,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResourceForExaccShape(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIdsForExaccShape(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitConditionForExacc, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationForExacc, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIdsForExaccShape(compartment string) ([]string, error) {
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
	listDatabaseSoftwareImagesRequest.ImageShapeFamily = oci_database.DatabaseSoftwareImageSummaryImageShapeFamilyExaccShape
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

func databaseSoftwareImageSweepWaitConditionForExacc(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted) && (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateTerminated)
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperationForExacc(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// Exadata test function

// issue-routing-tag: database/default
func TestDatabaseDatabaseSoftwareImageResourceExadata_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSoftwareImageResourceExadata_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "DatabaseSoftwareImageResourceExadata_basic") {
		t.Skip("Skipping suppressed TestDatabaseDatabaseSoftwareImageResourceExadata_basic")
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseSoftwareImageResourceDependenciesForExaShape+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create, databaseSoftwareImageRepresentationForExadataShape), "database", "databaseSoftwareImage", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseSoftwareImageDestroyExa, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageRepresentationForExadataShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape,
		},

		//verify Create with source_db_home
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForSourceDbHome +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageRepresentationForSourceDbHome),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_db_home_id"),
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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create, databaseSoftwareImageRepresentationForExadataShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(databaseSoftwareImageRepresentationForExadataShape, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
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
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentationForExadataShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "patch_set", "19.7.0.0"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_software_images", "test_database_software_images", acctest.Optional, acctest.Update, databaseSoftwareImageDataSourceRepresentationForExadataShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceDependenciesForExaShape +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Optional, acctest.Update, databaseSoftwareImageRepresentationForExadataShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "DB_Image_V19"),
				resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "EXADATA_SHAPE"),
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
				resource.TestCheckResourceAttr(datasourceName, "database_software_images.0.image_shape_family", "EXADATA_SHAPE"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_software_image", "test_database_software_image", acctest.Required, acctest.Create, databaseSoftwareImageSingularDataSourceRepresentationForExaccShape) +
				compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExadataShape,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_included_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_software_image_one_off_patches.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "DB_Image_V19_U1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_type", "DATABASE_IMAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_set", "19.7.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DatabaseSoftwareImageResourceConfigForExadataShape,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{
				//	"source_db_home_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseDatabaseSoftwareImageDestroyExa(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("DatabaseDatabaseSoftwareImageExa") {
		resource.AddTestSweepers("DatabaseDatabaseSoftwareImageExa", &resource.Sweeper{
			Name:         "DatabaseDatabaseSoftwareImageExa",
			Dependencies: acctest.DependencyGraph["databaseSoftwareImageExa"],
			F:            sweepDatabaseDatabaseSoftwareImageResourceExa,
		})
	}
}

func sweepDatabaseDatabaseSoftwareImageResourceExa(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	databaseSoftwareImageIds, err := getDatabaseSoftwareImageIdsExa(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseSoftwareImageId, databaseSoftwareImageSweepWaitConditionExa, time.Duration(3*time.Minute),
				databaseSoftwareImageSweepResponseFetchOperationExa, "database", true)
		}
	}
	return nil
}

func getDatabaseSoftwareImageIdsExa(compartment string) ([]string, error) {
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

func databaseSoftwareImageSweepWaitConditionExa(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateDeleted) && (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateTerminated)
	}
	return false
}

func databaseSoftwareImageWaitTillAvailableConditionExa(response common.OCIOperationResponse) bool {
	if databaseSoftwareImageResponse, ok := response.Response.(oci_database.GetDatabaseSoftwareImageResponse); ok {
		fmt.Print("Checking whether the state of resource is Available: ", databaseSoftwareImageResponse.LifecycleState == oci_database.DatabaseSoftwareImageLifecycleStateAvailable, "\n")
		return (databaseSoftwareImageResponse.LifecycleState != oci_database.DatabaseSoftwareImageLifecycleStateAvailable)
	}
	return false
}

func databaseSoftwareImageSweepResponseFetchOperationExa(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDatabaseSoftwareImage(context.Background(), oci_database.GetDatabaseSoftwareImageRequest{
		DatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
