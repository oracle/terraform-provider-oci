// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v56/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AiAnomalyDetectionProjectRequiredOnlyResource = AiAnomalyDetectionProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Required, acctest.Create, aiAnomalyDetectionProjectRepresentation)

	AiAnomalyDetectionProjectResourceConfig = AiAnomalyDetectionProjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Optional, acctest.Update, aiAnomalyDetectionProjectRepresentation)

	aiAnomalyDetectionProjectSingularDataSourceRepresentation = map[string]interface{}{
		"project_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_project.test_project.id}`},
	}

	aiAnomalyDetectionProjectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: aiAnomalyDetectionProjectDataSourceFilterRepresentation}}
	aiAnomalyDetectionProjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_anomaly_detection_project.test_project.id}`}},
	}

	aiAnomalyDetectionProjectRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	AiAnomalyDetectionProjectResourceDependencies = DefinedTagsDependencies
)

func TestAiAnomalyDetectionProjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionProjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_project.test_project"
	datasourceName := "data.oci_ai_anomaly_detection_projects.test_projects"
	singularDatasourceName := "data.oci_ai_anomaly_detection_project.test_project"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiAnomalyDetectionProjectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Optional, acctest.Create, aiAnomalyDetectionProjectRepresentation), "aianomalydetection", "project", t)

	acctest.ResourceTest(t, testAccCheckAiAnomalyDetectionProjectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Required, acctest.Create, aiAnomalyDetectionProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Optional, acctest.Create, aiAnomalyDetectionProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(aiAnomalyDetectionProjectRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Optional, acctest.Update, aiAnomalyDetectionProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_projects", "test_projects", acctest.Optional, acctest.Update, aiAnomalyDetectionProjectDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionProjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Optional, acctest.Update, aiAnomalyDetectionProjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Required, acctest.Create, aiAnomalyDetectionProjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionProjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_project" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetProjectRequest{}

			tmp := rs.Primary.ID
			request.ProjectId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("AiAnomalyDetectionProject") {
		resource.AddTestSweepers("AiAnomalyDetectionProject", &resource.Sweeper{
			Name:         "AiAnomalyDetectionProject",
			Dependencies: acctest.DependencyGraph["project"],
			F:            sweepAiAnomalyDetectionProjectResource,
		})
	}
}

func sweepAiAnomalyDetectionProjectResource(compartment string) error {
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()
	projectIds, err := aiAnomalyDetectionGetProjectIds(compartment)
	if err != nil {
		return err
	}
	for _, projectId := range projectIds {
		if ok := acctest.SweeperDefaultResourceId[projectId]; !ok {
			deleteProjectRequest := oci_ai_anomaly_detection.DeleteProjectRequest{}

			deleteProjectRequest.ProjectId = &projectId

			deleteProjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteProject(context.Background(), deleteProjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Project %s %s, It is possible that the resource is already deleted. Please verify manually \n", projectId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &projectId, aiAnomalyDetectionProjectSweepWaitCondition, time.Duration(3*time.Minute),
				aiAnomalyDetectionProjectSweepResponseFetchOperation, "ai_anomaly_detection", true)
		}
	}
	return nil
}

func aiAnomalyDetectionGetProjectIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProjectId", id)
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

func aiAnomalyDetectionProjectSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnomalyDetectionClient().GetProject(context.Background(), oci_ai_anomaly_detection.GetProjectRequest{
		ProjectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
