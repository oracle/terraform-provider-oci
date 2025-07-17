// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiVisionStreamGroupRequiredOnlyResource = AiVisionStreamGroupResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Required, acctest.Create, AiVisionStreamGroupRepresentation)

	AiVisionStreamGroupResourceConfig = AiVisionStreamGroupResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Optional, acctest.Update, AiVisionStreamGroupRepresentation)

	AiVisionStreamGroupSingularDataSourceRepresentation = map[string]interface{}{
		"stream_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_stream_group.test_stream_group.id}`},
	}

	AiVisionStreamGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_vision_stream_group.test_stream_group.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamGroupDataSourceFilterRepresentation}}
	AiVisionStreamGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_vision_stream_group.test_stream_group.id}`}},
	}

	AiVisionStreamGroupRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"is_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"stream_source_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_vision_stream_source.test_stream_source.id}`}, Update: []string{`${oci_ai_vision_stream_source.test_stream_source.id}`}},
	}

	AiVisionStreamGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_vision_private_endpoint", "test_vision_private_endpoint", acctest.Required, acctest.Create, AiVisionVisionPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Required, acctest.Create, AiVisionStreamSourceRepresentation)
)

// issue-routing-tag: ai_vision/default
func TestAiVisionStreamGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiVisionStreamGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	//compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	resourceName := "oci_ai_vision_stream_group.test_stream_group"
	datasourceName := "data.oci_ai_vision_stream_groups.test_stream_groups"
	singularDatasourceName := "data.oci_ai_vision_stream_group.test_stream_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+AiVisionStreamGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Optional, acctest.Create, AiVisionStreamGroupRepresentation), "aivision", "streamGroup", t)

	acctest.ResourceTest(t, testAccCheckAiVisionStreamGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Required, acctest.Create, AiVisionStreamGroupRepresentation),
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Optional, acctest.Create, AiVisionStreamGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_ids.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Optional, acctest.Update, AiVisionStreamGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_ids.#", "1"),

				func(s *terraform.State) (err error) {
					log.Printf("[INFO] Create 2...")
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_stream_groups", "test_stream_groups", acctest.Optional, acctest.Update, AiVisionStreamGroupDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Optional, acctest.Update, AiVisionStreamGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "stream_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_stream_group", "test_stream_group", acctest.Required, acctest.Create, AiVisionStreamGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_source_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiVisionStreamGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiVisionStreamGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceVisionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_vision_stream_group" {
			noResourceFound = false
			request := oci_ai_vision.GetStreamGroupRequest{}

			tmp := rs.Primary.ID
			request.StreamGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")

			response, err := client.GetStreamGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_vision.StreamGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiVisionStreamGroup") {
		resource.AddTestSweepers("AiVisionStreamGroup", &resource.Sweeper{
			Name:         "AiVisionStreamGroup",
			Dependencies: acctest.DependencyGraph["streamGroup"],
			F:            sweepAiVisionStreamGroupResource,
		})
	}
}

func sweepAiVisionStreamGroupResource(compartment string) error {
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()
	streamGroupIds, err := getAiVisionStreamGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, streamGroupId := range streamGroupIds {
		if ok := acctest.SweeperDefaultResourceId[streamGroupId]; !ok {
			deleteStreamGroupRequest := oci_ai_vision.DeleteStreamGroupRequest{}

			deleteStreamGroupRequest.StreamGroupId = &streamGroupId

			deleteStreamGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")
			_, error := aiServiceVisionClient.DeleteStreamGroup(context.Background(), deleteStreamGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamGroupId, AiVisionStreamGroupSweepWaitCondition, time.Duration(3*time.Minute),
				AiVisionStreamGroupSweepResponseFetchOperation, "ai_vision", true)
		}
	}
	return nil
}

func getAiVisionStreamGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()

	listStreamGroupsRequest := oci_ai_vision.ListStreamGroupsRequest{}
	listStreamGroupsRequest.CompartmentId = &compartmentId
	listStreamGroupsResponse, err := aiServiceVisionClient.ListStreamGroups(context.Background(), listStreamGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting StreamGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, streamGroup := range listStreamGroupsResponse.Items {
		id := *streamGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamGroupId", id)
	}
	return resourceIds, nil
}

func AiVisionStreamGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamGroupResponse, ok := response.Response.(oci_ai_vision.GetStreamGroupResponse); ok {
		return streamGroupResponse.LifecycleState != oci_ai_vision.StreamGroupLifecycleStateDeleted
	}
	return false
}

func AiVisionStreamGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceVisionClient().GetStreamGroup(context.Background(), oci_ai_vision.GetStreamGroupRequest{
		StreamGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
