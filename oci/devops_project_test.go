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
	oci_devops "github.com/oracle/oci-go-sdk/v48/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DevopsProjectRequiredOnlyResource = DevopsProjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation)

	DevopsProjectResourceConfig = DevopsProjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Optional, Update, devopsProjectRepresentation)

	devopsProjectSingularDataSourceRepresentation = map[string]interface{}{
		"project_id": Representation{repType: Required, create: `${oci_devops_project.test_project.id}`},
	}

	devopsProjectName = randomString(10, charsetWithoutDigits)

	devopsProjectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_devops_project.test_project.id}`},
		"name":           Representation{repType: Optional, create: devopsProjectName},
		"state":          Representation{repType: Optional, create: `Active`},
		"filter":         RepresentationGroup{Required, devopsProjectDataSourceFilterRepresentation}}
	devopsProjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_devops_project.test_project.id}`}},
	}

	devopsProjectRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"name":                Representation{repType: Required, create: devopsProjectName},
		"notification_config": RepresentationGroup{Required, projectNotificationConfigRepresentation},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}
	projectLoggingConfigRepresentation = map[string]interface{}{
		"log_group_id":             Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`},
		"retention_period_in_days": Representation{repType: Required, create: `30`, update: `60`},
		"display_name_prefix":      Representation{repType: Optional, create: `displayNamePrefix`, update: `displayNamePrefix2`},
		"is_archiving_enabled":     Representation{repType: Optional, create: `false`, update: `true`},
	}
	projectNotificationConfigRepresentation = map[string]interface{}{
		"topic_id": Representation{repType: Required, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}

	DevopsProjectResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_devops_project.test_project"
	datasourceName := "data.oci_devops_projects.test_projects"
	singularDatasourceName := "data.oci_devops_project.test_project"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DevopsProjectResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Optional, Create, devopsProjectRepresentation), "devops", "project", t)

	ResourceTest(t, testAccCheckDevopsProjectDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_project", "test_project", Optional, Create, devopsProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DevopsProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_project", "test_project", Optional, Create,
					representationCopyWithNewProperties(devopsProjectRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

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
			Config: config + compartmentIdVariableStr + DevopsProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_project", "test_project", Optional, Update, devopsProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_config.0.topic_id"),

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
				generateDataSourceFromRepresentationMap("oci_devops_projects", "test_projects", Optional, Update, devopsProjectDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_project", "test_project", Optional, Update, devopsProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(datasourceName, "state", "Active"),

				resource.TestCheckResourceAttr(datasourceName, "project_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "project_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsProjectResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", devopsProjectName),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DevopsProjectResourceConfig,
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

func testAccCheckDevopsProjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_project" {
			noResourceFound = false
			request := oci_devops.GetProjectRequest{}

			tmp := rs.Primary.ID
			request.ProjectId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "devops")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DevopsProject") {
		resource.AddTestSweepers("DevopsProject", &resource.Sweeper{
			Name:         "DevopsProject",
			Dependencies: DependencyGraph["project"],
			F:            sweepDevopsProjectResource,
		})
	}
}

func sweepDevopsProjectResource(compartment string) error {
	projectClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	projectIds, err := getDevopsProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, projectId := range projectIds {
		if ok := SweeperDefaultResourceId[projectId]; !ok {
			deleteProjectRequest := oci_devops.DeleteProjectRequest{}

			deleteProjectRequest.ProjectId = &projectId

			deleteProjectRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "devops")
			_, error := projectClient.DeleteProject(context.Background(), deleteProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Project %s %s, It is possible that the resource is already deleted. Please verify manually \n", projectId, error)
				continue
			}
			waitTillCondition(testAccProvider, &projectId, devopsProjectSweepWaitCondition, time.Duration(3*time.Minute),
				devopsProjectSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDevopsProjectIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	projectClient := GetTestClients(&schema.ResourceData{}).devopsClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ProjectId", id)
	}
	return resourceIds, nil
}

func devopsProjectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if projectResponse, ok := response.Response.(oci_devops.GetProjectResponse); ok {
		return projectResponse.LifecycleState != oci_devops.ProjectLifecycleStateDeleted
	}
	return false
}

func devopsProjectSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.devopsClient().GetProject(context.Background(), oci_devops.GetProjectRequest{
		ProjectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
