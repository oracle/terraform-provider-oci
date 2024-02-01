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
	MediaServicesStreamPackagingConfigRequiredOnlyResource = MediaServicesStreamPackagingConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Required, acctest.Create, MediaServicesStreamPackagingConfigRepresentation)

	MediaServicesStreamPackagingConfigResourceConfig = MediaServicesStreamPackagingConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Optional, acctest.Update, MediaServicesStreamPackagingConfigRepresentation)

	MediaServicesMediaServicesStreamPackagingConfigSingularDataSourceRepresentation = map[string]interface{}{
		"stream_packaging_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_packaging_config.test_stream_packaging_config.id}`},
	}

	MediaServicesMediaServicesStreamPackagingConfigDataSourceRepresentation = map[string]interface{}{
		"distribution_channel_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"stream_packaging_config_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_stream_packaging_config.test_stream_packaging_config.id}`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesStreamPackagingConfigDataSourceFilterRepresentation}}
	MediaServicesStreamPackagingConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_stream_packaging_config.test_stream_packaging_config.id}`}},
	}

	MediaServicesStreamPackagingConfigRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"distribution_channel_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
		"segment_time_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"stream_packaging_format": acctest.Representation{RepType: acctest.Required, Create: `HLS`},
		"is_lock_override":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesStreamPackagingConfigLocksRepresentation},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsSystemTagsAndLocks},
	}

	ignoreDefinedTagsSystemTagsAndLocks = map[string]interface{}{"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`, `locks`}}}

	MediaServicesStreamPackagingConfigLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesStreamPackagingConfigResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Required, acctest.Create, MediaServicesStreamDistributionChannelRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: media_services/default
func TestMediaServicesStreamPackagingConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesStreamPackagingConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_media_services_stream_packaging_config.test_stream_packaging_config"
	datasourceName := "data.oci_media_services_stream_packaging_configs.test_stream_packaging_configs"
	singularDatasourceName := "data.oci_media_services_stream_packaging_config.test_stream_packaging_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesStreamPackagingConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Optional, acctest.Create, MediaServicesStreamPackagingConfigRepresentation), "mediaservices", "streamPackagingConfig", t)

	acctest.ResourceTest(t, testAccCheckMediaServicesStreamPackagingConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamPackagingConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Required, acctest.Create, MediaServicesStreamPackagingConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttr(resourceName, "segment_time_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "stream_packaging_format", "HLS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamPackagingConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamPackagingConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Optional, acctest.Create, MediaServicesStreamPackagingConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "segment_time_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "stream_packaging_format", "HLS"),

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
			Config: config + compartmentIdVariableStr + MediaServicesStreamPackagingConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Optional, acctest.Update, MediaServicesStreamPackagingConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "segment_time_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "stream_packaging_format", "HLS"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_stream_packaging_configs", "test_stream_packaging_configs", acctest.Optional, acctest.Update, MediaServicesMediaServicesStreamPackagingConfigDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesStreamPackagingConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Optional, acctest.Update, MediaServicesStreamPackagingConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_packaging_config_id"),

				resource.TestCheckResourceAttr(datasourceName, "stream_packaging_config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_packaging_config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_stream_packaging_config", "test_stream_packaging_config", acctest.Required, acctest.Create, MediaServicesMediaServicesStreamPackagingConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesStreamPackagingConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_packaging_config_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "segment_time_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_packaging_format", "HLS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesStreamPackagingConfigRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMediaServicesStreamPackagingConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MediaServicesClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_media_services_stream_packaging_config" {
			noResourceFound = false
			request := oci_media_services.GetStreamPackagingConfigRequest{}

			tmp := rs.Primary.ID
			request.StreamPackagingConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")

			response, err := client.GetStreamPackagingConfig(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_media_services.StreamPackagingConfigLifecycleStateDeleted): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MediaServicesStreamPackagingConfig") {
		resource.AddTestSweepers("MediaServicesStreamPackagingConfig", &resource.Sweeper{
			Name:         "MediaServicesStreamPackagingConfig",
			Dependencies: acctest.DependencyGraph["streamPackagingConfig"],
			F:            sweepMediaServicesStreamPackagingConfigResource,
		})
	}
}

func sweepMediaServicesStreamPackagingConfigResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	streamPackagingConfigIds, err := getMediaServicesStreamPackagingConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, streamPackagingConfigId := range streamPackagingConfigIds {
		if ok := acctest.SweeperDefaultResourceId[streamPackagingConfigId]; !ok {
			deleteStreamPackagingConfigRequest := oci_media_services.DeleteStreamPackagingConfigRequest{}

			deleteStreamPackagingConfigRequest.StreamPackagingConfigId = &streamPackagingConfigId

			deleteStreamPackagingConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteStreamPackagingConfig(context.Background(), deleteStreamPackagingConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamPackagingConfig %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamPackagingConfigId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamPackagingConfigId, MediaServicesStreamPackagingConfigSweepWaitCondition, time.Duration(3*time.Minute),
				MediaServicesStreamPackagingConfigSweepResponseFetchOperation, "media_services", true)
		}
	}
	return nil
}

func getMediaServicesStreamPackagingConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamPackagingConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listStreamPackagingConfigsRequest := oci_media_services.ListStreamPackagingConfigsRequest{}
	//listStreamPackagingConfigsRequest.CompartmentId = &compartmentId

	distributionChannelIds, error := getMediaServicesStreamDistributionChannelIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting distributionChannelId required for StreamPackagingConfig resource requests \n")
	}
	for _, distributionChannelId := range distributionChannelIds {
		listStreamPackagingConfigsRequest.DistributionChannelId = &distributionChannelId

		listStreamPackagingConfigsRequest.LifecycleState = oci_media_services.StreamPackagingConfigLifecycleStateActive
		listStreamPackagingConfigsResponse, err := mediaServicesClient.ListStreamPackagingConfigs(context.Background(), listStreamPackagingConfigsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting StreamPackagingConfig list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, streamPackagingConfig := range listStreamPackagingConfigsResponse.Items {
			id := *streamPackagingConfig.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamPackagingConfigId", id)
		}

	}
	return resourceIds, nil
}

func MediaServicesStreamPackagingConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamPackagingConfigResponse, ok := response.Response.(oci_media_services.GetStreamPackagingConfigResponse); ok {
		return streamPackagingConfigResponse.GetLifecycleState() != oci_media_services.StreamPackagingConfigLifecycleStateDeleted
	}
	return false
}

func MediaServicesStreamPackagingConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MediaServicesClient().GetStreamPackagingConfig(context.Background(), oci_media_services.GetStreamPackagingConfigRequest{
		StreamPackagingConfigId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
