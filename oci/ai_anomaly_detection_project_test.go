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
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v48/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v48/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AiAnomalyDetectionProjectRequiredOnlyResource = AiAnomalyDetectionProjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Required, Create, aiAnomalyDetectionProjectRepresentation)

	AiAnomalyDetectionProjectResourceConfig = AiAnomalyDetectionProjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Optional, Update, aiAnomalyDetectionProjectRepresentation)

	aiAnomalyDetectionProjectSingularDataSourceRepresentation = map[string]interface{}{
		"project_id": Representation{repType: Required, create: `${oci_ai_anomaly_detection_project.test_project.id}`},
	}

	aiAnomalyDetectionProjectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, aiAnomalyDetectionProjectDataSourceFilterRepresentation}}
	aiAnomalyDetectionProjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_ai_anomaly_detection_project.test_project.id}`}},
	}

	aiAnomalyDetectionProjectRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      RepresentationGroup{Required, ignoreDefinedTagsChangesRep},
	}

	ignoreDefinedTagsChangesRep = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	AiAnomalyDetectionProjectResourceDependencies = DefinedTagsDependencies
)

func TestAiAnomalyDetectionProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_project.test_project"
	datasourceName := "data.oci_ai_anomaly_detection_projects.test_projects"
	singularDatasourceName := "data.oci_ai_anomaly_detection_project.test_project"

	var resId, resId2 string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+AiAnomalyDetectionProjectResourceDependencies+
		generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Optional, Create, aiAnomalyDetectionProjectRepresentation), "aianomalydetection", "project", t)

	ResourceTest(t, testAccCheckAiAnomalyDetectionProjectDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Required, Create, aiAnomalyDetectionProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Optional, Create, aiAnomalyDetectionProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// 					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Optional, Create,
					representationCopyWithNewProperties(aiAnomalyDetectionProjectRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				// 					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Optional, Update, aiAnomalyDetectionProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// 					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				generateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_projects", "test_projects", Optional, Update, aiAnomalyDetectionProjectDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				generateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Optional, Update, aiAnomalyDetectionProjectRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "project_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "project_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", Required, Create, aiAnomalyDetectionProjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionProjectResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				// 					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceConfig,
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

func testAccCheckAiAnomalyDetectionProjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).anomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_project" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetProjectRequest{}

			tmp := rs.Primary.ID
			request.ProjectId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ai_anomaly_detection")

			response, err := client.GetProject(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_anomaly_detection.ProjectLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("AiAnomalyDetectionProject") {
		resource.AddTestSweepers("AiAnomalyDetectionProject", &resource.Sweeper{
			Name:         "AiAnomalyDetectionProject",
			Dependencies: DependencyGraph["project"],
			F:            sweepAiAnomalyDetectionProjectResource,
		})
	}
}

func sweepAiAnomalyDetectionProjectResource(compartment string) error {
	anomalyDetectionClient := GetTestClients(&schema.ResourceData{}).anomalyDetectionClient()
	projectIds, err := aiAnomalyDetectionGetProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, projectId := range projectIds {
		if ok := SweeperDefaultResourceId[projectId]; !ok {
			deleteProjectRequest := oci_ai_anomaly_detection.DeleteProjectRequest{}

			deleteProjectRequest.ProjectId = &projectId

			deleteProjectRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteProject(context.Background(), deleteProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Project %s %s, It is possible that the resource is already deleted. Please verify manually \n", projectId, error)
				continue
			}
			waitTillCondition(testAccProvider, &projectId, aiAnomalyDetectionProjectSweepWaitCondition, time.Duration(3*time.Minute),
				aiAnomalyDetectionProjectSweepResponseFetchOperation, "ai_anomaly_detection", true)
		}
	}
	return nil
}

func aiAnomalyDetectionGetProjectIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := GetTestClients(&schema.ResourceData{}).anomalyDetectionClient()

	listProjectsRequest := oci_ai_anomaly_detection.ListProjectsRequest{}
	listProjectsRequest.CompartmentId = &compartmentId
	listProjectsRequest.LifecycleState = oci_ai_anomaly_detection.ProjectLifecycleStateActive
	listProjectsResponse, err := anomalyDetectionClient.ListProjects(context.Background(), listProjectsRequest)

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

func aiAnomalyDetectionProjectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if projectResponse, ok := response.Response.(oci_ai_anomaly_detection.GetProjectResponse); ok {
		return projectResponse.LifecycleState != oci_ai_anomaly_detection.ProjectLifecycleStateDeleted
	}
	return false
}

func aiAnomalyDetectionProjectSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.anomalyDetectionClient().GetProject(context.Background(), oci_ai_anomaly_detection.GetProjectRequest{
		ProjectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
