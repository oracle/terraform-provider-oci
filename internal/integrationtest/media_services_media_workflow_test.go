// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MediaServicesMediaWorkflowRequiredOnlyResource = MediaServicesMediaWorkflowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Required, acctest.Create, MediaServicesMediaWorkflowRepresentation)

	MediaServicesMediaWorkflowResourceConfig = MediaServicesMediaWorkflowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowRepresentation)

	MediaServicesMediaServicesMediaWorkflowSingularDataSourceRepresentation = map[string]interface{}{
		"media_workflow_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_workflow.test_media_workflow.id}`},
	}

	MediaServicesMediaServicesMediaWorkflowDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow.test_media_workflow.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesMediaWorkflowDataSourceFilterRepresentation}}
	MediaServicesMediaWorkflowDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_media_workflow.test_media_workflow.id}`}},
	}

	MediaServicesMediaWorkflowRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_lock_override":                 acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                            acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesMediaWorkflowLocksRepresentation},
		"media_workflow_configuration_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id}`}, Update: []string{`${oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id}`}},
		"parameters":                       acctest.Representation{RepType: acctest.Optional, Create: `{\"inputs\":{\"namespace\":\"axjagzvlc4vi\"}}`, Update: `{\"inputs\":{\"namespace\":\"axjagzvlc4vi\"}}`},
		"tasks":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesMediaWorkflowTasksRepresentation},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsAndLocks},
	}

	MediaServicesMediaWorkflowLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesMediaWorkflowTasksRepresentation = map[string]interface{}{
		"type":          acctest.Representation{RepType: acctest.Required, Create: `getFiles`, Update: `getFiles`},
		"version":       acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"key":           acctest.Representation{RepType: acctest.Required, Create: `move`, Update: `move`},
		"prerequisites": acctest.Representation{RepType: acctest.Optional, Create: []string{}, Update: []string{}},
		"parameters":    acctest.Representation{RepType: acctest.Required, Create: `{\"taskParameters\":[{\"bucketName\":\"tf_testing\",\"namespaceName\":\"axjagzvlc4vi\",\"objectName\":\"$${/videos/inputObject}\",\"storageType\":\"objectStorage\",\"target\":\"video.mp4\"}]}`, Update: `{\"taskParameters\":[{\"bucketName\":\"tf_testing\",\"namespaceName\":\"axjagzvlc4vi\",\"objectName\":\"horseInSuit.mp4\",\"storageType\":\"objectStorage\",\"target\":\"video.mp4\"}]}`},
	}

	MediaServicesMediaWorkflowResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_configuration", "test_media_workflow_configuration", acctest.Required, acctest.Create, MediaServicesMediaWorkflowConfigurationRepresentation)
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaWorkflowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesMediaWorkflowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_media_services_media_workflow.test_media_workflow"
	datasourceName := "data.oci_media_services_media_workflows.test_media_workflows"
	singularDatasourceName := "data.oci_media_services_media_workflow.test_media_workflow"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesMediaWorkflowResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Optional, acctest.Create, MediaServicesMediaWorkflowRepresentation), "mediaservices", "mediaWorkflow", t)

	acctest.ResourceTest(t, testAccCheckMediaServicesMediaWorkflowDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Required, acctest.Create, MediaServicesMediaWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Optional, acctest.Create, MediaServicesMediaWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.key", "move"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.prerequisites.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.type", "getFiles"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.version", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.parameters", "{\"taskParameters\":[{\"bucketName\":\"tf_testing\",\"namespaceName\":\"axjagzvlc4vi\",\"objectName\":\"${/videos/inputObject}\",\"storageType\":\"objectStorage\",\"target\":\"video.mp4\"}]}"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"inputs\":{\"namespace\":\"axjagzvlc4vi\"}}"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MediaServicesMediaWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MediaServicesMediaWorkflowRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.key", "move"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.parameters", "{\"taskParameters\":[{\"bucketName\":\"tf_testing\",\"namespaceName\":\"axjagzvlc4vi\",\"objectName\":\"${/videos/inputObject}\",\"storageType\":\"objectStorage\",\"target\":\"video.mp4\"}]}"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.prerequisites.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.type", "getFiles"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.version", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"inputs\":{\"namespace\":\"axjagzvlc4vi\"}}"),

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
			Config: config + compartmentIdVariableStr + MediaServicesMediaWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters", "{\"inputs\":{\"namespace\":\"axjagzvlc4vi\"}}"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.key", "move"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.parameters", "{\"taskParameters\":[{\"bucketName\":\"tf_testing\",\"namespaceName\":\"axjagzvlc4vi\",\"objectName\":\"horseInSuit.mp4\",\"storageType\":\"objectStorage\",\"target\":\"video.mp4\"}]}"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.prerequisites.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.type", "getFiles"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.version", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflows", "test_media_workflows", acctest.Optional, acctest.Update, MediaServicesMediaServicesMediaWorkflowDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Optional, acctest.Update, MediaServicesMediaWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "media_workflow_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "media_workflow_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaWorkflowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "media_workflow_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "media_workflow_configuration_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.key", "move"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.prerequisites.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.type", "getFiles"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.version", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tasks.0.parameters"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters", "{\"inputs\":{\"namespace\":\"axjagzvlc4vi\"}}"), resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesMediaWorkflowRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMediaServicesMediaWorkflowDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MediaServicesClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_media_services_media_workflow" {
			noResourceFound = false
			request := oci_media_services.GetMediaWorkflowRequest{}

			tmp := rs.Primary.ID
			request.MediaWorkflowId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")

			response, err := client.GetMediaWorkflow(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_media_services.MediaWorkflowLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MediaServicesMediaWorkflow") {
		resource.AddTestSweepers("MediaServicesMediaWorkflow", &resource.Sweeper{
			Name:         "MediaServicesMediaWorkflow",
			Dependencies: acctest.DependencyGraph["mediaWorkflow"],
			F:            sweepMediaServicesMediaWorkflowResource,
		})
	}
}

func sweepMediaServicesMediaWorkflowResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	mediaWorkflowIds, err := getMediaServicesMediaWorkflowIds(compartment)
	if err != nil {
		return err
	}
	for _, mediaWorkflowId := range mediaWorkflowIds {
		if ok := acctest.SweeperDefaultResourceId[mediaWorkflowId]; !ok {
			deleteMediaWorkflowRequest := oci_media_services.DeleteMediaWorkflowRequest{}

			deleteMediaWorkflowRequest.MediaWorkflowId = &mediaWorkflowId

			deleteMediaWorkflowRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteMediaWorkflow(context.Background(), deleteMediaWorkflowRequest)
			if error != nil {
				fmt.Printf("Error deleting MediaWorkflow %s %s, It is possible that the resource is already deleted. Please verify manually \n", mediaWorkflowId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mediaWorkflowId, MediaServicesMediaWorkflowSweepWaitCondition, time.Duration(3*time.Minute),
				MediaServicesMediaWorkflowSweepResponseFetchOperation, "media_services", true)
		}
	}
	return nil
}

func getMediaServicesMediaWorkflowIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MediaWorkflowId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listMediaWorkflowsRequest := oci_media_services.ListMediaWorkflowsRequest{}
	listMediaWorkflowsRequest.CompartmentId = &compartmentId
	listMediaWorkflowsRequest.LifecycleState = oci_media_services.MediaWorkflowLifecycleStateActive
	listMediaWorkflowsResponse, err := mediaServicesClient.ListMediaWorkflows(context.Background(), listMediaWorkflowsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MediaWorkflow list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mediaWorkflow := range listMediaWorkflowsResponse.Items {
		id := *mediaWorkflow.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MediaWorkflowId", id)
	}
	return resourceIds, nil
}

func MediaServicesMediaWorkflowSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mediaWorkflowResponse, ok := response.Response.(oci_media_services.GetMediaWorkflowResponse); ok {
		return mediaWorkflowResponse.LifecycleState != oci_media_services.MediaWorkflowLifecycleStateDeleted
	}
	return false
}

func MediaServicesMediaWorkflowSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MediaServicesClient().GetMediaWorkflow(context.Background(), oci_media_services.GetMediaWorkflowRequest{
		MediaWorkflowId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
