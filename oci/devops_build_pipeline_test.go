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
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_devops "github.com/oracle/oci-go-sdk/v55/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BuildPipelineRequiredOnlyResource = BuildPipelineResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineRepresentation)

	BuildPipelineResourceConfig = BuildPipelineResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Optional, Update, buildPipelineRepresentation)

	buildPipelineSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}

	buildPipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"project_id":     Representation{RepType: Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, buildPipelineDataSourceFilterRepresentation}}
	buildPipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_build_pipeline.test_build_pipeline.id}`}},
	}

	buildPipelineRepresentation = map[string]interface{}{
		"project_id":                Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"build_pipeline_parameters": RepresentationGroup{Optional, buildPipelineBuildPipelineParametersRepresentation},
		"defined_tags":              Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":               Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":              Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":             Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	buildPipelineBuildPipelineParametersRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineBuildPipelineParametersItemsRepresentation},
	}
	buildPipelineBuildPipelineParametersItemsRepresentation = map[string]interface{}{
		"name":          Representation{RepType: Required, Create: `name`, Update: `name2`},
		"default_value": Representation{RepType: Optional, Create: `defaultValue`, Update: `defaultValue2`},
		"description":   Representation{RepType: Optional, Create: `description`, Update: `description2`},
	}

	BuildPipelineResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline.test_build_pipeline"
	datasourceName := "data.oci_devops_build_pipelines.test_build_pipelines"
	singularDatasourceName := "data.oci_devops_build_pipeline.test_build_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Optional, Create, buildPipelineRepresentation), "devops", "buildPipeline", t)

	ResourceTest(t, testAccCheckDevopsBuildPipelineDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BuildPipelineResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Optional, Create, buildPipelineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.0.default_value", "defaultValue"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BuildPipelineResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Optional, Update, buildPipelineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.0.default_value", "defaultValue2"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_parameters.0.items.0.name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_devops_build_pipelines", "test_build_pipelines", Optional, Update, buildPipelineDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Optional, Update, buildPipelineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "build_pipeline_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "build_pipeline_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "build_pipeline_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_parameters.0.items.0.default_value", "defaultValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_parameters.0.items.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_parameters.0.items.0.name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
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
			Config: config + compartmentIdVariableStr + BuildPipelineResourceConfig,
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

func testAccCheckDevopsBuildPipelineDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_build_pipeline" {
			noResourceFound = false
			request := oci_devops.GetBuildPipelineRequest{}

			tmp := rs.Primary.ID
			request.BuildPipelineId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

			response, err := client.GetBuildPipeline(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.BuildPipelineLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("DevopsBuildPipeline") {
		resource.AddTestSweepers("DevopsBuildPipeline", &resource.Sweeper{
			Name:         "DevopsBuildPipeline",
			Dependencies: DependencyGraph["buildPipeline"],
			F:            sweepDevopsBuildPipelineResource,
		})
	}
}

func sweepDevopsBuildPipelineResource(compartment string) error {
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	buildPipelineIds, err := getBuildPipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, buildPipelineId := range buildPipelineIds {
		if ok := SweeperDefaultResourceId[buildPipelineId]; !ok {
			deleteBuildPipelineRequest := oci_devops.DeleteBuildPipelineRequest{}

			deleteBuildPipelineRequest.BuildPipelineId = &buildPipelineId

			deleteBuildPipelineRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteBuildPipeline(context.Background(), deleteBuildPipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting BuildPipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", buildPipelineId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &buildPipelineId, buildPipelineSweepWaitCondition, time.Duration(3*time.Minute),
				buildPipelineSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getBuildPipelineIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "BuildPipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listBuildPipelinesRequest := oci_devops.ListBuildPipelinesRequest{}
	listBuildPipelinesRequest.CompartmentId = &compartmentId
	listBuildPipelinesRequest.LifecycleState = oci_devops.BuildPipelineLifecycleStateActive
	listBuildPipelinesResponse, err := devopsClient.ListBuildPipelines(context.Background(), listBuildPipelinesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BuildPipeline list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, buildPipeline := range listBuildPipelinesResponse.Items {
		id := *buildPipeline.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "BuildPipelineId", id)
	}
	return resourceIds, nil
}

func buildPipelineSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if buildPipelineResponse, ok := response.Response.(oci_devops.GetBuildPipelineResponse); ok {
		return buildPipelineResponse.LifecycleState != oci_devops.BuildPipelineLifecycleStateDeleted
	}
	return false
}

func buildPipelineSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.devopsClient().GetBuildPipeline(context.Background(), oci_devops.GetBuildPipelineRequest{
		BuildPipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
