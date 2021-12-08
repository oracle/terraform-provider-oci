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
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_devops "github.com/oracle/oci-go-sdk/v53/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BuildPipelineStageRequiredOnlyResource = BuildPipelineStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineStageRepresentation)

	BuildPipelineStageResourceConfig = BuildPipelineStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineStageRepresentation)

	buildPipelineStageSingularDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_stage_id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
	}

	buildPipelineStageDataSourceRepresentation = map[string]interface{}{
		"build_pipeline_id": Representation{RepType: Optional, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"compartment_id":    Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":      Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":                Representation{RepType: Optional, Create: `${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`},
		"state":             Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":            RepresentationGroup{Required, buildPipelineStageDataSourceFilterRepresentation}}
	buildPipelineStageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_build_pipeline_stage.test_build_pipeline_stage.id}`}},
	}

	buildPipelineStageRepresentation = map[string]interface{}{
		"build_pipeline_id":                           Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
		"build_pipeline_stage_predecessor_collection": RepresentationGroup{Required, buildPipelineStageBuildPipelineStagePredecessorCollectionRepresentation},
		"build_pipeline_stage_type":                   Representation{RepType: Required, Create: `TRIGGER_DEPLOYMENT_PIPELINE`},
		"defined_tags":                                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deploy_pipeline_id":                          Representation{RepType: Required, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"description":                                 Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":                                Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                               Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_pass_all_parameters_enabled":              Representation{RepType: Required, Create: `false`, Update: `true`},
	}
	buildPipelineStageBuildPipelineStagePredecessorCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, buildPipelineStageBuildPipelineStagePredecessorCollectionItemsRepresentation},
	}

	buildPipelineStageBuildPipelineStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": Representation{RepType: Required, Create: `${oci_devops_build_pipeline.test_build_pipeline.id}`},
	}

	BuildPipelineStageResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_build_pipeline", "test_build_pipeline", Required, Create, buildPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsBuildPipelineStageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsBuildPipelineStageResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_build_pipeline_stage.test_build_pipeline_stage"
	datasourceName := "data.oci_devops_build_pipeline_stages.test_build_pipeline_stages"
	singularDatasourceName := "data.oci_devops_build_pipeline_stage.test_build_pipeline_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+BuildPipelineStageResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineStageRepresentation), "devops", "buildPipelineStage", t)

	ResourceTest(t, testAccCheckDevopsBuildPipelineStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "is_pass_all_parameters_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Create, buildPipelineStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_pass_all_parameters_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + BuildPipelineStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "build_pipeline_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_pass_all_parameters_enabled", "true"),
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
				GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stages", "test_build_pipeline_stages", Optional, Update, buildPipelineStageDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Optional, Update, buildPipelineStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "build_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "build_pipeline_stage_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "build_pipeline_stage_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_build_pipeline_stage", "test_build_pipeline_stage", Required, Create, buildPipelineStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BuildPipelineStageResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "build_pipeline_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_pass_all_parameters_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BuildPipelineStageResourceConfig,
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

func testAccCheckDevopsBuildPipelineStageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_build_pipeline_stage" {
			noResourceFound = false
			request := oci_devops.GetBuildPipelineStageRequest{}

			tmp := rs.Primary.ID
			request.BuildPipelineStageId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

			response, err := client.GetBuildPipelineStage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.BuildPipelineStageLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !InSweeperExcludeList("DevopsBuildPipelineStage") {
		resource.AddTestSweepers("DevopsBuildPipelineStage", &resource.Sweeper{
			Name:         "DevopsBuildPipelineStage",
			Dependencies: DependencyGraph["buildPipelineStage"],
			F:            sweepDevopsBuildPipelineStageResource,
		})
	}
}

func sweepDevopsBuildPipelineStageResource(compartment string) error {
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	buildPipelineStageIds, err := getBuildPipelineStageIds(compartment)
	if err != nil {
		return err
	}
	for _, buildPipelineStageId := range buildPipelineStageIds {
		if ok := SweeperDefaultResourceId[buildPipelineStageId]; !ok {
			deleteBuildPipelineStageRequest := oci_devops.DeleteBuildPipelineStageRequest{}

			deleteBuildPipelineStageRequest.BuildPipelineStageId = &buildPipelineStageId

			deleteBuildPipelineStageRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := devopsClient.DeleteBuildPipelineStage(context.Background(), deleteBuildPipelineStageRequest)
			if error != nil {
				fmt.Printf("Error deleting BuildPipelineStage %s %s, It is possible that the resource is already deleted. Please verify manually \n", buildPipelineStageId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &buildPipelineStageId, buildPipelineStageSweepWaitCondition, time.Duration(3*time.Minute),
				buildPipelineStageSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getBuildPipelineStageIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "BuildPipelineStageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	devopsClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listBuildPipelineStagesRequest := oci_devops.ListBuildPipelineStagesRequest{}
	listBuildPipelineStagesRequest.CompartmentId = &compartmentId
	listBuildPipelineStagesRequest.LifecycleState = oci_devops.BuildPipelineStageLifecycleStateActive
	listBuildPipelineStagesResponse, err := devopsClient.ListBuildPipelineStages(context.Background(), listBuildPipelineStagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BuildPipelineStage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, buildPipelineStage := range listBuildPipelineStagesResponse.Items {
		id := *buildPipelineStage.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "BuildPipelineStageId", id)
	}
	return resourceIds, nil
}

func buildPipelineStageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if buildPipelineStageResponse, ok := response.Response.(oci_devops.GetBuildPipelineStageResponse); ok {
		return buildPipelineStageResponse.GetLifecycleState() != oci_devops.BuildPipelineStageLifecycleStateDeleted
	}
	return false
}

func buildPipelineStageSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.devopsClient().GetBuildPipelineStage(context.Background(), oci_devops.GetBuildPipelineStageRequest{
		BuildPipelineStageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
