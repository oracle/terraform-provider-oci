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
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_devops "github.com/oracle/oci-go-sdk/v46/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployStageRequiredOnlyResource = DeployStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployStageRepresentation)

	DeployStageResourceConfig = DeployStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployStageRepresentation)

	deployStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": Representation{repType: Required, create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployStageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Optional, create: `${var.compartment_id}`},
		"deploy_pipeline_id": Representation{repType: Optional, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":                 Representation{repType: Optional, create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
		"filter":             RepresentationGroup{Required, deployStageDataSourceFilterRepresentation}}
	deployStageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_devops_deploy_stage.test_deploy_stage.id}`}},
	}

	deployStageRepresentation = map[string]interface{}{
		"deploy_pipeline_id":                  Representation{repType: Required, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deploy_stage_predecessor_collection": RepresentationGroup{Required, deployStageDeployStagePredecessorCollectionRepresentation},
		"deploy_stage_type":                   Representation{repType: Required, create: `WAIT`},
		"defined_tags":                        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                         Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":                        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                       Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"wait_criteria":                       RepresentationGroup{Required, deployStageWaitCriteriaRepresentation},
	}

	deployStageDeployStagePredecessorCollectionRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, deployStageDeployStagePredecessorCollectionItemsRepresentation},
	}
	deployStageWaitCriteriaRepresentation = map[string]interface{}{
		"wait_duration": Representation{repType: Required, create: `PT5S`},
		"wait_type":     Representation{repType: Required, create: `ABSOLUTE_WAIT`},
	}
	deployStageDeployStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
	}

	DeployStageResourceDependencies = generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployArtifactRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DeployStageResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployStageRepresentation), "devops", "deployStage", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDevopsDeployStageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DeployStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "WAIT"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.0.wait_duration", "PT5S"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.0.wait_type", "ABSOLUTE_WAIT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DeployStageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DeployStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "WAIT"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.0.wait_duration", "PT5S"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.0.wait_type", "ABSOLUTE_WAIT"),

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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DeployStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.0.wait_duration", "PT5S"),
					resource.TestCheckResourceAttr(resourceName, "wait_criteria.0.wait_type", "ABSOLUTE_WAIT"),

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
					generateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", Optional, Update, deployStageDataSourceRepresentation) +
					compartmentIdVariableStr + DeployStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),

					resource.TestCheckResourceAttr(datasourceName, "deploy_stage_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployStageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DeployStageResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "WAIT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "wait_criteria.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "wait_criteria.0.wait_duration", "PT5S"),
					resource.TestCheckResourceAttr(singularDatasourceName, "wait_criteria.0.wait_type", "ABSOLUTE_WAIT"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DeployStageResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDevopsDeployStageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_deploy_stage" {
			noResourceFound = false
			request := oci_devops.GetDeployStageRequest{}

			tmp := rs.Primary.ID
			request.DeployStageId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "devops")

			response, err := client.GetDeployStage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.DeployStageLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("DevopsDeployStage") {
		resource.AddTestSweepers("DevopsDeployStage", &resource.Sweeper{
			Name:         "DevopsDeployStage",
			Dependencies: DependencyGraph["deployStage"],
			F:            sweepDevopsDeployStageResource,
		})
	}
}

func sweepDevopsDeployStageResource(compartment string) error {
	deployStageClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	deployStageIds, err := getDeployStageIds(compartment)
	if err != nil {
		return err
	}
	for _, deployStageId := range deployStageIds {
		if ok := SweeperDefaultResourceId[deployStageId]; !ok {
			deleteDeployStageRequest := oci_devops.DeleteDeployStageRequest{}

			deleteDeployStageRequest.DeployStageId = &deployStageId

			deleteDeployStageRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "devops")
			_, error := deployStageClient.DeleteDeployStage(context.Background(), deleteDeployStageRequest)
			if error != nil {
				fmt.Printf("Error deleting DeployStage %s %s, It is possible that the resource is already deleted. Please verify manually \n", deployStageId, error)
				continue
			}
			waitTillCondition(testAccProvider, &deployStageId, deployStageSweepWaitCondition, time.Duration(3*time.Minute),
				deployStageSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDeployStageIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DeployStageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deployStageClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listDeployStagesRequest := oci_devops.ListDeployStagesRequest{}
	listDeployStagesRequest.CompartmentId = &compartmentId
	listDeployStagesRequest.LifecycleState = oci_devops.DeployStageLifecycleStateActive
	listDeployStagesResponse, err := deployStageClient.ListDeployStages(context.Background(), listDeployStagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DeployStage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployStage := range listDeployStagesResponse.Items {
		id := *deployStage.GetId()
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DeployStageId", id)
	}
	return resourceIds, nil
}

func deployStageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deployStageResponse, ok := response.Response.(oci_devops.GetDeployStageResponse); ok {
		return deployStageResponse.GetLifecycleState() != oci_devops.DeployStageLifecycleStateDeleted
	}
	return false
}

func deployStageSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.devopsClient().GetDeployStage(context.Background(), oci_devops.GetDeployStageRequest{
		DeployStageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
