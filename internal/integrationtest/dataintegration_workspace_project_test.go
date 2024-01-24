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
	DataintegrationWorkspaceProjectRequiredOnlyResource = DataintegrationWorkspaceProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Required, acctest.Create, DataintegrationWorkspaceProjectRepresentation)

	DataintegrationWorkspaceProjectResourceConfig = DataintegrationWorkspaceProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Optional, acctest.Update, DataintegrationWorkspaceProjectRepresentation)

	DataintegrationWorkspaceProjectSingularDataSourceRepresentation = map[string]interface{}{
		"project_key":  acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_project.test_workspace_project.key}`},
		"workspace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
	}

	DataintegrationWorkspaceProjectDataSourceRepresentation = map[string]interface{}{
		"workspace_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"identifier":    acctest.Representation{RepType: acctest.Optional, Create: []string{`IDENTIFIER`}, Update: []string{`IDENTIFIER2`}},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"name_contains": acctest.Representation{RepType: acctest.Optional, Create: `nameContains`},
		"fields":        acctest.Representation{RepType: acctest.Optional, Update: []string{`metadata`}},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceProjectDataSourceFilterRepresentation}}
	DataintegrationWorkspaceProjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_project.test_workspace_project.name}`}},
	}

	DataintegrationWorkspaceProjectRepresentation = map[string]interface{}{
		"identifier":        acctest.Representation{RepType: acctest.Required, Create: `IDENTIFIER`, Update: `IDENTIFIER2`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"workspace_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"model_version":     acctest.Representation{RepType: acctest.Optional, Create: `20200901`, Update: `20200902`},
		"registry_metadata": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataintegrationWorkspaceProjectRegistryMetadataRepresentation},
	}
	DataintegrationWorkspaceProjectRegistryMetadataRepresentation = map[string]interface{}{
		"is_favorite": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":      acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
	}

	DataintegrationWorkspaceProjectResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("is_private_network_enabled", acctest.Representation{RepType: acctest.Required, Create: `false`}, DataintegrationWorkspaceRepresentation))
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dataintegration_workspace_project.test_workspace_project"
	datasourceName := "data.oci_dataintegration_workspace_projects.test_workspace_projects"
	singularDatasourceName := "data.oci_dataintegration_workspace_project.test_workspace_project"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataintegrationWorkspaceProjectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Optional, acctest.Create, DataintegrationWorkspaceProjectRepresentation), "dataintegration", "workspaceProject", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceProjectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Required, acctest.Create, DataintegrationWorkspaceProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceProjectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Optional, acctest.Create, DataintegrationWorkspaceProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20200901"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_status"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "key_map"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "USER_PROJECT"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
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
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Optional, acctest.Update, DataintegrationWorkspaceProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER2"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20200902"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "object_status"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "true"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.labels.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_projects", "test_workspace_projects", acctest.Optional, acctest.Update, DataintegrationWorkspaceProjectDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Optional, acctest.Update, DataintegrationWorkspaceProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "identifier.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "nameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),
				resource.TestCheckResourceAttr(datasourceName, "project_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "project_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_project", "test_workspace_project", acctest.Required, acctest.Create, DataintegrationWorkspaceProjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceProjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "IDENTIFIER2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "20200902"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_version"),
			),
		},

		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceProjectRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"registry_metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceProjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_project" {
			noResourceFound = false
			request := oci_dataintegration.GetProjectRequest{}

			if value, ok := rs.Primary.Attributes["project_key"]; ok {
				request.ProjectKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetProject(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceProject") {
		resource.AddTestSweepers("DataintegrationWorkspaceProject", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceProject",
			Dependencies: acctest.DependencyGraph["workspaceProject"],
			F:            sweepDataintegrationWorkspaceProjectResource,
		})
	}
}

func sweepDataintegrationWorkspaceProjectResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceProjectIds, err := getDataintegrationWorkspaceProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceProjectId := range workspaceProjectIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceProjectId]; !ok {
			deleteProjectRequest := oci_dataintegration.DeleteProjectRequest{}

			deleteProjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteProject(context.Background(), deleteProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceProject %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceProjectId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceProjectIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listProjectsRequest := oci_dataintegration.ListProjectsRequest{}
	//listProjectsRequest.CompartmentId = &compartmentId

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceProject resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listProjectsRequest.WorkspaceId = &workspaceId

		listProjectsResponse, err := dataIntegrationClient.ListProjects(context.Background(), listProjectsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceProject list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceProject := range listProjectsResponse.Items {
			id := *workspaceProject.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceProjectId", id)
		}

	}
	return resourceIds, nil
}
