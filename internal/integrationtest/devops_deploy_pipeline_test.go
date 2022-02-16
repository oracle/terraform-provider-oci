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
	DeployPipelineRequiredOnlyResource = DeployPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineRepresentation)

	DeployPipelineResourceConfig = DeployPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Optional, acctest.Update, deployPipelineRepresentation)

	deployPipelineSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
	}

	deployPipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deployPipelineDataSourceFilterRepresentation}}
	deployPipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`}},
	}

	deployPipelineRepresentation = map[string]interface{}{
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deploy_pipeline_parameters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployPipelineDeployPipelineParametersRepresentation},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	deployPipelineDeployPipelineParametersRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deployPipelineDeployPipelineParametersItemsRepresentation},
	}
	deployPipelineDeployPipelineParametersItemsRepresentation = map[string]interface{}{
		"name":          acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"default_value": acctest.Representation{RepType: acctest.Optional, Create: `defaultValue`, Update: `defaultValue2`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}

	DeployPipelineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployPipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployPipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_pipeline.test_deploy_pipeline"
	datasourceName := "data.oci_devops_deploy_pipelines.test_deploy_pipelines"
	singularDatasourceName := "data.oci_devops_deploy_pipeline.test_deploy_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployPipelineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Optional, acctest.Create, deployPipelineRepresentation), "devops", "deployPipeline", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployPipelineDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineRepresentation),
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
			Config: config + compartmentIdVariableStr + DeployPipelineResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Optional, acctest.Create, deployPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.0.default_value", "defaultValue"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
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
			Config: config + compartmentIdVariableStr + DeployPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Optional, acctest.Update, deployPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.0.default_value", "defaultValue2"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "deploy_pipeline_parameters.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_pipelines", "test_deploy_pipelines", acctest.Optional, acctest.Update, deployPipelineDataSourceRepresentation) +
				compartmentIdVariableStr + DeployPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Optional, acctest.Update, deployPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_pipeline_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deploy_pipeline_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployPipelineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_pipeline_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_pipeline_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_pipeline_parameters.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_pipeline_parameters.0.items.0.default_value", "defaultValue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_pipeline_parameters.0.items.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_pipeline_parameters.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployPipelineResourceConfig,
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

func testAccCheckDevopsDeployPipelineDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_deploy_pipeline" {
			noResourceFound = false
			request := oci_devops.GetDeployPipelineRequest{}

			tmp := rs.Primary.ID
			request.DeployPipelineId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			response, err := client.GetDeployPipeline(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.DeployPipelineLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DevopsDeployPipeline") {
		resource.AddTestSweepers("DevopsDeployPipeline", &resource.Sweeper{
			Name:         "DevopsDeployPipeline",
			Dependencies: acctest.DependencyGraph["deployPipeline"],
			F:            sweepDevopsDeployPipelineResource,
		})
	}
}

func sweepDevopsDeployPipelineResource(compartment string) error {
	deployPipelineClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	deployPipelineIds, err := getDeployPipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, deployPipelineId := range deployPipelineIds {
		if ok := acctest.SweeperDefaultResourceId[deployPipelineId]; !ok {
			deleteDeployPipelineRequest := oci_devops.DeleteDeployPipelineRequest{}

			deleteDeployPipelineRequest.DeployPipelineId = &deployPipelineId

			deleteDeployPipelineRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := deployPipelineClient.DeleteDeployPipeline(context.Background(), deleteDeployPipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting DeployPipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", deployPipelineId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deployPipelineId, deployPipelineSweepWaitCondition, time.Duration(3*time.Minute),
				deployPipelineSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDeployPipelineIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeployPipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deployPipelineClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listDeployPipelinesRequest := oci_devops.ListDeployPipelinesRequest{}
	listDeployPipelinesRequest.CompartmentId = &compartmentId
	listDeployPipelinesRequest.LifecycleState = oci_devops.DeployPipelineLifecycleStateActive
	listDeployPipelinesResponse, err := deployPipelineClient.ListDeployPipelines(context.Background(), listDeployPipelinesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DeployPipeline list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployPipeline := range listDeployPipelinesResponse.Items {
		id := *deployPipeline.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeployPipelineId", id)
	}
	return resourceIds, nil
}

func deployPipelineSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deployPipelineResponse, ok := response.Response.(oci_devops.GetDeployPipelineResponse); ok {
		return deployPipelineResponse.LifecycleState != oci_devops.DeployPipelineLifecycleStateDeleted
	}
	return false
}

func deployPipelineSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DevopsClient().GetDeployPipeline(context.Background(), oci_devops.GetDeployPipelineRequest{
		DeployPipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
