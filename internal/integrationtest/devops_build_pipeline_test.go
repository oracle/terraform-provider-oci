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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BuildPipelineRequiredOnlyResource = BuildPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, buildPipelineRepresentation)

	BuildPipelineResourceConfig = BuildPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Optional, acctest.Update, buildPipelineRepresentation)

	buildPipelineSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}

	buildPipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineDataSourceFilterRepresentation}}
	buildPipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_build_pipeline.test_build_pipeline.id}`}},
	}

	buildPipelineRepresentation = map[string]interface{}{
		"project_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"build_pipeline_parameters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: buildPipelineBuildPipelineParametersRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	buildPipelineBuildPipelineParametersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: buildPipelineBuildPipelineParametersItemsRepresentation},
	}
	buildPipelineBuildPipelineParametersItemsRepresentation = map[string]interface{}{
		"name":          acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"default_value": acctest.Representation{RepType: acctest.Optional, Create: `defaultValue`, Update: `defaultValue2`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}

	BuildPipelineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline.test_build_pipeline"
	datasourceName := "data.oci_devops_build_pipelines.test_build_pipelines"
	singularDatasourceName := "data.oci_devops_build_pipeline.test_build_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Optional, acctest.Create, buildPipelineRepresentation), "devops", "buildPipeline", t)

	acctest.ResourceTest(t, testAccCheckDevopsBuildPipelineDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, buildPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Optional, acctest.Create, buildPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + BuildPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Optional, acctest.Update, buildPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_pipelines", "test_build_pipelines", acctest.Optional, acctest.Update, buildPipelineDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Optional, acctest.Update, buildPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", acctest.Required, acctest.Create, buildPipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_build_pipeline" {
			noResourceFound = false
			request := oci_devops.GetBuildPipelineRequest{}

			tmp := rs.Primary.ID
			request.BuildPipelineId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DevopsBuildPipeline") {
		resource.AddTestSweepers("DevopsBuildPipeline", &resource.Sweeper{
			Name:         "DevopsBuildPipeline",
			Dependencies: acctest.DependencyGraph["buildPipeline"],
			F:            sweepDevopsBuildPipelineResource,
		})
	}
}

func sweepDevopsBuildPipelineResource(compartment string) error {
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	buildPipelineIds, err := getBuildPipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, buildPipelineId := range buildPipelineIds {
		if ok := acctest.SweeperDefaultResourceId[buildPipelineId]; !ok {
			deleteBuildPipelineRequest := oci_devops.DeleteBuildPipelineRequest{}

			deleteBuildPipelineRequest.BuildPipelineId = &buildPipelineId

			deleteBuildPipelineRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteBuildPipeline(context.Background(), deleteBuildPipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting BuildPipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", buildPipelineId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &buildPipelineId, buildPipelineSweepWaitCondition, time.Duration(3*time.Minute),
				buildPipelineSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getBuildPipelineIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BuildPipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BuildPipelineId", id)
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

func buildPipelineSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DevopsClient().GetBuildPipeline(context.Background(), oci_devops.GetBuildPipelineRequest{
		BuildPipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
