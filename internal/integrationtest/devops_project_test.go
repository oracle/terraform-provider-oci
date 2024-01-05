// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsProjectRequiredOnlyResource = DevopsProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation)

	DevopsProjectResourceConfig = DevopsProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Optional, acctest.Update, DevopsProjectRepresentation)

	DevopsDevopsProjectSingularDataSourceRepresentation = map[string]interface{}{
		"project_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
	}

	devopsProjectName  = utils.RandomString(10, utils.CharsetWithoutDigits)
	devopsLogGroupName = utils.RandomString(10, utils.CharsetWithoutDigits)

	defaultRegionForOns    = utils.GetEnvSettingWithDefault("default_region_for_project_ons", "iad")
	defaultTopicForProject = "ocid1." + "onstopic.oc1." + defaultRegionForOns + ".aaaaaaaadlp34edn4zeo3qaa2hdsxf6qd43itl5ph72mvpk6n44pvcylrk2q"
	onsTopicForProject     = utils.GetEnvSettingWithDefault("oci_ons_notification_topic_for_project", defaultTopicForProject)

	DevopsDevopsProjectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: devopsProjectName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: oci_devops.ProjectLifecycleStateActive},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsProjectDataSourceFilterRepresentation}}
	DevopsProjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_project.test_project.id}`}},
	}

	DevopsProjectRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: devopsProjectName},
		"notification_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsProjectNotificationConfigRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	DevopsProjectLoggingConfigRepresentation = map[string]interface{}{
		"log_group_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_devops_log_group.id}`},
		"retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `30`, Update: `60`},
		"display_name_prefix":      acctest.Representation{RepType: acctest.Optional, Create: `displayNamePrefix`, Update: `displayNamePrefix2`},
		"is_archiving_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	DevopsProjectNotificationConfigRepresentation = map[string]interface{}{
		"topic_id": acctest.Representation{RepType: acctest.Required, Create: onsTopicForProject},
	}

	DevopsLogGroupRepresentation = acctest.RepresentationCopyWithNewProperties(LoggingLogGroupRepresentation, map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: devopsLogGroupName},
	})

	DevopsProjectResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, DevopsLogGroupRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_devops_project.test_project"
	datasourceName := "data.oci_devops_projects.test_projects"
	singularDatasourceName := "data.oci_devops_project.test_project"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsProjectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Optional, acctest.Create, DevopsProjectRepresentation), "devops", "project", t)

	acctest.ResourceTest(t, testAccCheckDevopsProjectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Optional, acctest.Create, DevopsProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DevopsProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DevopsProjectRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

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
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Optional, acctest.Update, DevopsProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_projects", "test_projects", acctest.Optional, acctest.Update, DevopsDevopsProjectDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Optional, acctest.Update, DevopsProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", devopsProjectName),

				resource.TestCheckResourceAttr(datasourceName, "project_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "project_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsDevopsProjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsProjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsProjectRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDevopsProjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_project" {
			noResourceFound = false
			request := oci_devops.GetProjectRequest{}

			tmp := rs.Primary.ID
			request.ProjectId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			response, err := client.GetProject(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.ProjectLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DevopsProject") {
		resource.AddTestSweepers("DevopsProject", &resource.Sweeper{
			Name:         "DevopsProject",
			Dependencies: acctest.DependencyGraph["project"],
			F:            sweepDevopsProjectResource,
		})
	}
}

func sweepDevopsProjectResource(compartment string) error {
	projectClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	projectIds, err := getDevopsProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, projectId := range projectIds {
		if ok := acctest.SweeperDefaultResourceId[projectId]; !ok {
			deleteProjectRequest := oci_devops.DeleteProjectRequest{}

			deleteProjectRequest.ProjectId = &projectId

			deleteProjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := projectClient.DeleteProject(context.Background(), deleteProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Project %s %s, It is possible that the resource is already deleted. Please verify manually \n", projectId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &projectId, DevopsProjectSweepWaitCondition, time.Duration(3*time.Minute),
				DevopsProjectSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDevopsProjectIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	projectClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listProjectsRequest := oci_devops.ListProjectsRequest{}
	listProjectsRequest.CompartmentId = &compartmentId
	listProjectsRequest.LifecycleState = oci_devops.ProjectLifecycleStateActive
	listProjectsResponse, err := projectClient.ListProjects(context.Background(), listProjectsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Project list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, project := range listProjectsResponse.Items {
		id := *project.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProjectId", id)
	}
	return resourceIds, nil
}

func DevopsProjectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if projectResponse, ok := response.Response.(oci_devops.GetProjectResponse); ok {
		return projectResponse.LifecycleState != oci_devops.ProjectLifecycleStateDeleted
	}
	return false
}

func DevopsProjectSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DevopsClient().GetProject(context.Background(), oci_devops.GetProjectRequest{
		ProjectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
