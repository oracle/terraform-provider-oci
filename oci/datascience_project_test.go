// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v48/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ProjectRequiredOnlyResource = ProjectResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation)

	ProjectResourceConfig = ProjectResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Optional, Update, projectRepresentation)

	projectSingularDataSourceRepresentation = map[string]interface{}{
		"project_id": Representation{RepType: Required, Create: `${oci_datascience_project.test_project.id}`},
	}

	projectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, projectDataSourceFilterRepresentation}}
	projectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_datascience_project.test_project.id}`}},
	}

	projectRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ProjectResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_project.test_project"
	datasourceName := "data.oci_datascience_projects.test_projects"
	singularDatasourceName := "data.oci_datascience_project.test_project"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ProjectResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Optional, Create, projectRepresentation), "datascience", "project", t)

	ResourceTest(t, testAccCheckDatascienceProjectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ProjectResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ProjectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ProjectResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Optional, Create, projectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ProjectResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Optional, Create,
					RepresentationCopyWithNewProperties(projectRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ProjectResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Optional, Update, projectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_datascience_projects", "test_projects", Optional, Update, projectDataSourceRepresentation) +
				compartmentIdVariableStr + ProjectResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Optional, Update, projectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "projects.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "projects.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "projects.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "projects.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "projects.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "projects.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "projects.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "projects.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "projects.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "projects.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ProjectResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ProjectResourceConfig,
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

func testAccCheckDatascienceProjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_project" {
			noResourceFound = false
			request := oci_datascience.GetProjectRequest{}

			tmp := rs.Primary.ID
			request.ProjectId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datascience")

			response, err := client.GetProject(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ProjectLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("DatascienceProject") {
		resource.AddTestSweepers("DatascienceProject", &resource.Sweeper{
			Name:         "DatascienceProject",
			Dependencies: DependencyGraph["project"],
			F:            sweepDatascienceProjectResource,
		})
	}
}

func sweepDatascienceProjectResource(compartment string) error {
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()
	projectIds, err := getProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, projectId := range projectIds {
		if ok := SweeperDefaultResourceId[projectId]; !ok {
			deleteProjectRequest := oci_datascience.DeleteProjectRequest{}

			deleteProjectRequest.ProjectId = &projectId

			deleteProjectRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteProject(context.Background(), deleteProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Project %s %s, It is possible that the resource is already deleted. Please verify manually \n", projectId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &projectId, projectSweepWaitCondition, time.Duration(3*time.Minute),
				projectSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getProjectIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()

	listProjectsRequest := oci_datascience.ListProjectsRequest{}
	listProjectsRequest.CompartmentId = &compartmentId
	listProjectsRequest.LifecycleState = oci_datascience.ListProjectsLifecycleStateActive
	listProjectsResponse, err := dataScienceClient.ListProjects(context.Background(), listProjectsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Project list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, project := range listProjectsResponse.Items {
		id := *project.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ProjectId", id)
	}
	return resourceIds, nil
}

func projectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if projectResponse, ok := response.Response.(oci_datascience.GetProjectResponse); ok {
		return projectResponse.LifecycleState != oci_datascience.ProjectLifecycleStateDeleted
	}
	return false
}

func projectSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataScienceClient().GetProject(context.Background(), oci_datascience.GetProjectRequest{
		ProjectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
