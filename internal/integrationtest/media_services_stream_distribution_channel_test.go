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
	MediaServicesStreamDistributionChannelRequiredOnlyResource = MediaServicesStreamDistributionChannelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Required, acctest.Create, MediaServicesStreamDistributionChannelRepresentation)

	MediaServicesStreamDistributionChannelResourceConfig = MediaServicesStreamDistributionChannelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Optional, acctest.Update, MediaServicesStreamDistributionChannelRepresentation)

	MediaServicesMediaServicesStreamDistributionChannelSingularDataSourceRepresentation = map[string]interface{}{
		"stream_distribution_channel_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
	}

	MediaServicesMediaServicesStreamDistributionChannelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesStreamDistributionChannelDataSourceFilterRepresentation}}
	MediaServicesStreamDistributionChannelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`}},
	}

	MediaServicesStreamDistributionChannelRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_lock_override": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesStreamDistributionChannelLocksRepresentation},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsAndLocks},
	}

	MediaServicesStreamDistributionChannelLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesStreamDistributionChannelResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: media_services/default
func TestMediaServicesStreamDistributionChannelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesStreamDistributionChannelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_media_services_stream_distribution_channel.test_stream_distribution_channel"
	datasourceName := "data.oci_media_services_stream_distribution_channels.test_stream_distribution_channels"
	singularDatasourceName := "data.oci_media_services_stream_distribution_channel.test_stream_distribution_channel"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesStreamDistributionChannelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Optional, acctest.Create, MediaServicesStreamDistributionChannelRepresentation), "mediaservices", "streamDistributionChannel", t)

	acctest.ResourceTest(t, testAccCheckMediaServicesStreamDistributionChannelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamDistributionChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Required, acctest.Create, MediaServicesStreamDistributionChannelRepresentation),
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
			Config: config + compartmentIdVariableStr + MediaServicesStreamDistributionChannelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamDistributionChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Optional, acctest.Create, MediaServicesStreamDistributionChannelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MediaServicesStreamDistributionChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MediaServicesStreamDistributionChannelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + MediaServicesStreamDistributionChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Optional, acctest.Update, MediaServicesStreamDistributionChannelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_stream_distribution_channels", "test_stream_distribution_channels", acctest.Optional, acctest.Update, MediaServicesMediaServicesStreamDistributionChannelDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesStreamDistributionChannelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Optional, acctest.Update, MediaServicesStreamDistributionChannelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "stream_distribution_channel_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_distribution_channel_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Required, acctest.Create, MediaServicesMediaServicesStreamDistributionChannelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesStreamDistributionChannelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_distribution_channel_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesStreamDistributionChannelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMediaServicesStreamDistributionChannelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MediaServicesClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_media_services_stream_distribution_channel" {
			noResourceFound = false
			request := oci_media_services.GetStreamDistributionChannelRequest{}

			tmp := rs.Primary.ID
			request.StreamDistributionChannelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")

			response, err := client.GetStreamDistributionChannel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_media_services.StreamDistributionChannelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MediaServicesStreamDistributionChannel") {
		resource.AddTestSweepers("MediaServicesStreamDistributionChannel", &resource.Sweeper{
			Name:         "MediaServicesStreamDistributionChannel",
			Dependencies: acctest.DependencyGraph["streamDistributionChannel"],
			F:            sweepMediaServicesStreamDistributionChannelResource,
		})
	}
}

func sweepMediaServicesStreamDistributionChannelResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	streamDistributionChannelIds, err := getMediaServicesStreamDistributionChannelIds(compartment)
	if err != nil {
		return err
	}
	for _, streamDistributionChannelId := range streamDistributionChannelIds {
		if ok := acctest.SweeperDefaultResourceId[streamDistributionChannelId]; !ok {
			deleteStreamDistributionChannelRequest := oci_media_services.DeleteStreamDistributionChannelRequest{}

			deleteStreamDistributionChannelRequest.StreamDistributionChannelId = &streamDistributionChannelId

			deleteStreamDistributionChannelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteStreamDistributionChannel(context.Background(), deleteStreamDistributionChannelRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamDistributionChannel %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamDistributionChannelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamDistributionChannelId, MediaServicesStreamDistributionChannelSweepWaitCondition, time.Duration(3*time.Minute),
				MediaServicesStreamDistributionChannelSweepResponseFetchOperation, "media_services", true)
		}
	}
	return nil
}

func getMediaServicesStreamDistributionChannelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamDistributionChannelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listStreamDistributionChannelsRequest := oci_media_services.ListStreamDistributionChannelsRequest{}
	listStreamDistributionChannelsRequest.CompartmentId = &compartmentId
	listStreamDistributionChannelsRequest.LifecycleState = oci_media_services.StreamDistributionChannelLifecycleStateActive
	listStreamDistributionChannelsResponse, err := mediaServicesClient.ListStreamDistributionChannels(context.Background(), listStreamDistributionChannelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting StreamDistributionChannel list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, streamDistributionChannel := range listStreamDistributionChannelsResponse.Items {
		id := *streamDistributionChannel.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamDistributionChannelId", id)
	}
	return resourceIds, nil
}

func MediaServicesStreamDistributionChannelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamDistributionChannelResponse, ok := response.Response.(oci_media_services.GetStreamDistributionChannelResponse); ok {
		return streamDistributionChannelResponse.LifecycleState != oci_media_services.StreamDistributionChannelLifecycleStateDeleted
	}
	return false
}

func MediaServicesStreamDistributionChannelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MediaServicesClient().GetStreamDistributionChannel(context.Background(), oci_media_services.GetStreamDistributionChannelRequest{
		StreamDistributionChannelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
