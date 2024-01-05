// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataintegrationWorkspaceFolderRequiredOnlyResource = DataintegrationWorkspaceFolderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Required, acctest.Create, DataintegrationWorkspaceFolderRepresentation)

	DataintegrationWorkspaceFolderResourceConfig = DataintegrationWorkspaceFolderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Optional, acctest.Update, DataintegrationWorkspaceFolderRepresentation)

	DataintegrationWorkspaceFolderSingularDataSourceRepresentation = map[string]interface{}{
		"folder_key":   acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_folder.test_workspace_folder.key}`},
		"workspace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
	}

	DataintegrationWorkspaceFolderDataSourceRepresentation = map[string]interface{}{
		"workspace_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"identifier":    acctest.Representation{RepType: acctest.Optional, Create: []string{`IDENTIFIER`}, Update: []string{`IDENTIFIER2`}},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"name_contains": acctest.Representation{RepType: acctest.Optional, Create: `name2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceFolderDataSourceFilterRepresentation}}
	DataintegrationWorkspaceFolderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_folder.test_workspace_folder.name}`}},
	}

	DataintegrationWorkspaceFolderRepresentation = map[string]interface{}{
		"identifier":        acctest.Representation{RepType: acctest.Required, Create: `IDENTIFIER`, Update: `IDENTIFIER2`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"registry_metadata": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceFolderRegistryMetadataRepresentation},
		"workspace_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"model_version":     acctest.Representation{RepType: acctest.Optional, Create: `20220713`, Update: `20220715`},
		"category_name":     acctest.Representation{RepType: acctest.Optional, Create: `CATEGORY`, Update: `CATEGORY1`},
	}
	DataintegrationWorkspaceFolderRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_project.test_workspace_project.key}`},
		"is_favorite":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":         acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
	}

	DataintegrationWorkspaceFolderResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("is_private_network_enabled", acctest.Representation{RepType: acctest.Required, Create: `false`}, DataintegrationWorkspaceRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Required, acctest.Create, DataintegrationWorkspaceProjectRepresentation)
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceFolderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceFolderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dataintegration_workspace_folder.test_workspace_folder"
	datasourceName := "data.oci_dataintegration_workspace_folders.test_workspace_folders"
	singularDatasourceName := "data.oci_dataintegration_workspace_folder.test_workspace_folder"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataintegrationWorkspaceFolderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Optional, acctest.Create, DataintegrationWorkspaceFolderRepresentation), "dataintegration", "workspaceFolder", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceFolderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Required, acctest.Create, DataintegrationWorkspaceFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceFolderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Optional, acctest.Create, DataintegrationWorkspaceFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20220713"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_status"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.0", "labels"),
				resource.TestCheckResourceAttr(resourceName, "category_name", "CATEGORY"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "FOLDER"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrSet(resourceName, "object_version"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.created_by"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.created_by_name"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.updated_by"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.updated_by_name"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "metadata.0.registry_version"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.labels.#", "1"),
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
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Optional, acctest.Update, DataintegrationWorkspaceFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER2"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20220715"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "object_status"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.aggregator_key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "true"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
				resource.TestCheckResourceAttr(resourceName, "category_name", "CATEGORY1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_folders", "test_workspace_folders", acctest.Optional, acctest.Update, DataintegrationWorkspaceFolderDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceFolderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Optional, acctest.Update, DataintegrationWorkspaceFolderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "identifier.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "name2"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "folder_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "folder_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_folder", "test_workspace_folder", acctest.Required, acctest.Create, DataintegrationWorkspaceFolderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceFolderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "folder_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "IDENTIFIER2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "20220715"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_version"),
			),
		},
		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceFolderRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"registry_metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceFolderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_folder" {
			noResourceFound = false
			request := oci_dataintegration.GetFolderRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.FolderKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetFolder(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceFolder") {
		resource.AddTestSweepers("DataintegrationWorkspaceFolder", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceFolder",
			Dependencies: acctest.DependencyGraph["workspaceFolder"],
			F:            sweepDataintegrationWorkspaceFolderResource,
		})
	}
}

func sweepDataintegrationWorkspaceFolderResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceFolderIds, err := getDataintegrationWorkspaceFolderIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceFolderId := range workspaceFolderIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceFolderId]; !ok {
			deleteFolderRequest := oci_dataintegration.DeleteFolderRequest{}

			deleteFolderRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteFolder(context.Background(), deleteFolderRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceFolder %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceFolderId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceFolderIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceFolderId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listFoldersRequest := oci_dataintegration.ListFoldersRequest{}
	// listFoldersRequest.CompartmentId = &compartmentId

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceFolder resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listFoldersRequest.WorkspaceId = &workspaceId

		listFoldersResponse, err := dataIntegrationClient.ListFolders(context.Background(), listFoldersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceFolder list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceFolder := range listFoldersResponse.Items {
			id := *workspaceFolder.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceFolderId", id)
		}

	}
	return resourceIds, nil
}
