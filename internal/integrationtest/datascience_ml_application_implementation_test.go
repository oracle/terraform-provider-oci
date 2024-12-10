// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"archive/zip"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceMlApplicationImplementationRequiredOnlyResource = DatascienceMlApplicationImplementationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Required, acctest.Create, DatascienceMlApplicationImplementationRepresentation)

	DatascienceMlApplicationImplementationResourceConfig = DatascienceMlApplicationImplementationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Update, DatascienceMlApplicationImplementationRepresentation)

	DatascienceMlApplicationImplementationSingularDataSourceRepresentation = map[string]interface{}{
		"ml_application_implementation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application_implementation.test_ml_application_implementation.id}`},
	}

	DatascienceMlApplicationImplementationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ml_application_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"ml_application_implementation_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_ml_application_implementation.test_ml_application_implementation.id}`},
		"name":                             acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMlApplicationImplementationDataSourceFilterRepresentation}}
	DatascienceMlApplicationImplementationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_ml_application_implementation.test_ml_application_implementation.id}`}},
	}

	DatascienceMlApplicationImplementationRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ml_application_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"ml_application_package":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"source_type": "local", "path": "file://ml-app-package-v1.zip"}, Update: map[string]string{"source_type": "local", "path": "file://ml-app-package-v1.zip"}},
		"opc_ml_app_package_args":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bucket_namespace": "idtlxnfdweil"}},
		"allowed_migration_destinations": acctest.Representation{RepType: acctest.Optional, Update: []string{`allowedMigrationDestinations2`}},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatascienceMlApplicationImplementationDefinedTagsChangesRepresentation},
	}

	DatascienceMlApplicationImplementationWithoutPackagePathRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ml_application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `name`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatascienceMlApplicationImplementationDefinedTagsChangesRepresentation},
	}

	DatascienceMlApplicationImplementationRemoteURLRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ml_application_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"ml_application_package":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"source_type": "object_storage_download", "uri": "https://objectstorage.us-ashburn-1.oraclecloud.com/n/ociodscdev/b/Artifact/o/windows.zip"}, Update: map[string]string{"source_type": "object_storage_download", "path": "https://objectstorage.us-ashburn-1.oraclecloud.com/n/ociodscdev/b/Artifact/o/ml-app-package-1.8.zip"}},
		"opc_ml_app_package_args":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bucket_namespace": "idtlxnfdweil"}},
		"allowed_migration_destinations": acctest.Representation{RepType: acctest.Optional, Update: []string{`allowedMigrationDestinations2`}},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatascienceMlApplicationImplementationDefinedTagsChangesRepresentation},
	}

	ignoreDatascienceMlApplicationImplementationDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	DatascienceMlApplicationImplementationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Required, acctest.Create, DatascienceMlApplicationRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceMlApplicationImplementationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceMlApplicationImplementationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_ml_application_implementation.test_ml_application_implementation"
	datasourceName := "data.oci_datascience_ml_application_implementations.test_ml_application_implementations"
	singularDatasourceName := "data.oci_datascience_ml_application_implementation.test_ml_application_implementation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceMlApplicationImplementationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Create, DatascienceMlApplicationImplementationRepresentation), "datascience", "mlApplicationImplementation", t)

	acctest.ResourceTest(t, testAccCheckDatascienceMlApplicationImplementationDestroy, []resource.TestStep{
		// verify Create
		{
			PreConfig: func() {
				generateMLAppPackage("ml-app-package-v1", "1.1")
			},
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Required, acctest.Create, DatascienceMlApplicationImplementationWithoutPackagePathRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies,
		},

		// verify Create with optionals (Package with Remote URL)
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Create, DatascienceMlApplicationImplementationRemoteURLRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "allowed_migration_destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.source_type", "object_storage_download"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Create, DatascienceMlApplicationImplementationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "allowed_migration_destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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

		// verify upload package use case 1 : No change in TF, No change in Cloud
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatascienceMlApplicationImplementationRepresentation, map[string]interface{}{
						"ml_application_package": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"source_type": "local", "path": "file://ml-app-package-v1.zip"}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.source_type", "local"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.path", "file://ml-app-package-v1.zip::1.1"),
			),
		},

		// verify upload package use case 2 : Change in TF only(Change path + Change Version) , No change in Cloud
		{
			PreConfig: func() {
				generateMLAppPackage("ml-app-package-v2", "1.2")
			},
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatascienceMlApplicationImplementationRepresentation, map[string]interface{}{
						"ml_application_package": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"source_type": "local", "path": "file://ml-app-package-v2.zip"}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.source_type", "local"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.path", "file://ml-app-package-v2.zip::1.2"),
			),
		},

		//verify upload package usecase 3 : No change in TF (Version change in descriptor file) , No change in Cloud
		{
			PreConfig: func() {
				generateMLAppPackage("ml-app-package-v2", "1.3")
			},
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatascienceMlApplicationImplementationRepresentation, map[string]interface{}{
						"ml_application_package": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"source_type": "local", "path": "file://ml-app-package-v2.zip"}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.source_type", "local"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.path", "file://ml-app-package-v2.zip::1.3"),
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceMlApplicationImplementationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Update, DatascienceMlApplicationImplementationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
				resource.TestCheckResourceAttr(resourceName, "ml_application_package.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_implementations", "test_ml_application_implementations", acctest.Optional, acctest.Update, DatascienceMlApplicationImplementationDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Update, DatascienceMlApplicationImplementationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ml_application_implementation_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ml_application_implementation_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ml_application_implementation_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Required, acctest.Create, DatascienceMlApplicationImplementationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlApplicationImplementationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_implementation_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "application_components.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_schema.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceMlApplicationImplementationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"ml_application_package", "opc_ml_app_package_args"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceMlApplicationImplementationDestroy(s *terraform.State) error {
	deleteMLPackage("ml-app-package-v1")
	deleteMLPackage("ml-app-package-v2")
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_ml_application_implementation" {
			noResourceFound = false
			request := oci_datascience.GetMlApplicationImplementationRequest{}

			tmp := rs.Primary.ID
			request.MlApplicationImplementationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			_, err := client.GetMlApplicationImplementation(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DatascienceMlApplicationImplementation") {
		resource.AddTestSweepers("DatascienceMlApplicationImplementation", &resource.Sweeper{
			Name:         "DatascienceMlApplicationImplementation",
			Dependencies: acctest.DependencyGraph["mlApplicationImplementation"],
			F:            sweepDatascienceMlApplicationImplementationResource,
		})
	}
}

func sweepDatascienceMlApplicationImplementationResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	mlApplicationImplementationIds, err := getDatascienceMlApplicationImplementationIds(compartment)
	if err != nil {
		return err
	}
	for _, mlApplicationImplementationId := range mlApplicationImplementationIds {
		if ok := acctest.SweeperDefaultResourceId[mlApplicationImplementationId]; !ok {
			deleteMlApplicationImplementationRequest := oci_datascience.DeleteMlApplicationImplementationRequest{}

			deleteMlApplicationImplementationRequest.MlApplicationImplementationId = &mlApplicationImplementationId

			deleteMlApplicationImplementationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteMlApplicationImplementation(context.Background(), deleteMlApplicationImplementationRequest)
			if error != nil {
				fmt.Printf("Error deleting MlApplicationImplementation %s %s, It is possible that the resource is already deleted. Please verify manually \n", mlApplicationImplementationId, error)
				continue
			}
		}
	}
	return nil
}

func getDatascienceMlApplicationImplementationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MlApplicationImplementationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listMlApplicationImplementationsRequest := oci_datascience.ListMlApplicationImplementationsRequest{}
	listMlApplicationImplementationsRequest.CompartmentId = &compartmentId
	listMlApplicationImplementationsResponse, err := dataScienceClient.ListMlApplicationImplementations(context.Background(), listMlApplicationImplementationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MlApplicationImplementation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mlApplicationImplementation := range listMlApplicationImplementationsResponse.Items {
		id := *mlApplicationImplementation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MlApplicationImplementationId", id)
	}
	return resourceIds, nil
}

func generateMLAppPackage(zipName string, pVersion string) error {
	// Create a new zip archive
	archive, err := os.Create(zipName + ".zip")
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer archive.Close()

	// Create a new zip writer
	writer := zip.NewWriter(archive)
	defer writer.Close()

	// Load the descriptor file contents
	descriptorBytes, err := ioutil.ReadFile("../../examples/datascience/ml_applications/descriptor.yaml")
	if err != nil {
		return fmt.Errorf("failed to read descriptor file: %w", err)
	}

	// Convert the contents to a string
	descriptorContent := string(descriptorBytes)

	// Find and replace the packageVersion value
	newContent := regexp.MustCompile(`packageVersion: \d+\.\d+`).ReplaceAllString(descriptorContent, fmt.Sprintf("packageVersion: %s", pVersion))
	// Add descriptor file to the zip archive
	f, err := writer.Create("descriptor.yaml")
	if err != nil {
		return fmt.Errorf("failed to create file in zip: %w", err)
	}

	// Write the updated descriptor content to the zip
	_, err = f.Write([]byte(newContent))
	if err != nil {
		return fmt.Errorf("failed to write descriptor to zip: %w", err)
	}

	fmt.Printf("Zip file '%s' created successfully with updated packageVersion.\n", zipName)
	return nil
}

// deleteMLPackage deletes the zip file with the specified name
func deleteMLPackage(name string) error {
	err := os.Remove(fmt.Sprintf("%s.zip", name))
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("ML application package '%s.zip' not found.\n", name)
			return nil
		}
		return fmt.Errorf("failed to delete zip file: %w", err)
	}
	fmt.Printf("ML application package '%s.zip' deleted.\n", name)
	return nil
}
