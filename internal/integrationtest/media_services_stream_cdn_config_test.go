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
	MediaServicesStreamCdnConfigRequiredOnlyResource = MediaServicesStreamCdnConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Required, acctest.Create, MediaServicesStreamCdnConfigRepresentation)

	MediaServicesStreamCdnConfigResourceConfig = MediaServicesStreamCdnConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Optional, acctest.Update, MediaServicesStreamCdnConfigRepresentation)

	MediaServicesMediaServicesStreamCdnConfigSingularDataSourceRepresentation = map[string]interface{}{
		"stream_cdn_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_cdn_config.test_stream_cdn_config.id}`},
	}

	MediaServicesMediaServicesStreamCdnConfigDataSourceRepresentation = map[string]interface{}{
		"distribution_channel_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                      acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_stream_cdn_config.test_stream_cdn_config.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesStreamCdnConfigDataSourceFilterRepresentation}}
	MediaServicesStreamCdnConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_stream_cdn_config.test_stream_cdn_config.id}`}},
	}

	MediaServicesStreamCdnConfigRepresentation = map[string]interface{}{
		"config":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesStreamCdnConfigConfigRepresentation},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"distribution_channel_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
		"is_lock_override":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"locks":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesStreamCdnConfigLocksRepresentation},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsAndLocks},
	}

	MediaServicesStreamCdnConfigConfigRepresentation = map[string]interface{}{
		"type":                           acctest.Representation{RepType: acctest.Required, Create: `AKAMAI_MANUAL`, Update: `AKAMAI_MANUAL`},
		"edge_hostname":                  acctest.Representation{RepType: acctest.Optional, Create: `edgeHostname`, Update: `edgeHostname2`},
		"edge_path_prefix":               acctest.Representation{RepType: acctest.Optional, Create: `edgePathPrefix`, Update: `edgePathPrefix2`},
		"edge_token_key":                 acctest.Representation{RepType: acctest.Optional, Create: `edgeTokenKey`, Update: `edgeTokenKey2`},
		"edge_token_salt":                acctest.Representation{RepType: acctest.Optional, Create: `edgeTokenSalt`, Update: `edgeTokenSalt2`},
		"is_edge_token_auth":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"origin_auth_secret_key_a":       acctest.Representation{RepType: acctest.Optional, Create: `originAuthSecretKeyA`, Update: `originAuthSecretKeyA2`},
		"origin_auth_secret_key_b":       acctest.Representation{RepType: acctest.Optional, Create: `originAuthSecretKeyB`, Update: `originAuthSecretKeyB2`},
		"origin_auth_secret_key_nonce_a": acctest.Representation{RepType: acctest.Optional, Create: `originAuthSecretKeyNonceA`, Update: `originAuthSecretKeyNonceA2`},
		"origin_auth_secret_key_nonce_b": acctest.Representation{RepType: acctest.Optional, Create: `originAuthSecretKeyNonceB`, Update: `originAuthSecretKeyNonceB2`},
		"origin_auth_sign_encryption":    acctest.Representation{RepType: acctest.Optional, Create: `SHA256-HMAC`},
		"origin_auth_sign_type":          acctest.Representation{RepType: acctest.Optional, Create: `ForwardURL`},
	}

	MediaServicesStreamCdnConfigLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesStreamCdnConfigResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_distribution_channel", "test_stream_distribution_channel", acctest.Required, acctest.Create, MediaServicesStreamDistributionChannelRepresentation)

	MediaServicesStreamCdnConfigRepresentationWithEdge = map[string]interface{}{
		"config":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesStreamCdnConfigConfigRepresentationWithEdge},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"distribution_channel_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id}`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	MediaServicesStreamCdnConfigConfigRepresentationWithEdge = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `EDGE`, Update: `EDGE`},
	}
)

// issue-routing-tag: media_services/default
func TestMediaServicesStreamCdnConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesStreamCdnConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_media_services_stream_cdn_config.test_stream_cdn_config"
	datasourceName := "data.oci_media_services_stream_cdn_configs.test_stream_cdn_configs"
	singularDatasourceName := "data.oci_media_services_stream_cdn_config.test_stream_cdn_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesStreamCdnConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Optional, acctest.Create, MediaServicesStreamCdnConfigRepresentationWithEdge), "mediaservices", "streamCdnConfig", t)

	acctest.ResourceTest(t, testAccCheckMediaServicesStreamCdnConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamCdnConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Required, acctest.Create, MediaServicesStreamCdnConfigRepresentationWithEdge),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config.0.type", "EDGE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "distribution_channel_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamCdnConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesStreamCdnConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Optional, acctest.Create, MediaServicesStreamCdnConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_hostname", "edgeHostname"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_path_prefix", "edgePathPrefix"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_token_key", "edgeTokenKey"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_token_salt", "edgeTokenSalt"),
				resource.TestCheckResourceAttr(resourceName, "config.0.is_edge_token_auth", "false"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_a", "originAuthSecretKeyA"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_b", "originAuthSecretKeyB"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_nonce_a", "originAuthSecretKeyNonceA"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_nonce_b", "originAuthSecretKeyNonceB"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_sign_encryption", "SHA256-HMAC"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_sign_type", "ForwardURL"),
				resource.TestCheckResourceAttr(resourceName, "config.0.type", "AKAMAI_MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),

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
			Config: config + compartmentIdVariableStr + MediaServicesStreamCdnConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Optional, acctest.Update, MediaServicesStreamCdnConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_hostname", "edgeHostname2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_path_prefix", "edgePathPrefix2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_token_key", "edgeTokenKey2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.edge_token_salt", "edgeTokenSalt2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.is_edge_token_auth", "true"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_a", "originAuthSecretKeyA2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_b", "originAuthSecretKeyB2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_nonce_a", "originAuthSecretKeyNonceA2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_secret_key_nonce_b", "originAuthSecretKeyNonceB2"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_sign_encryption", "SHA256-HMAC"),
				resource.TestCheckResourceAttr(resourceName, "config.0.origin_auth_sign_type", "ForwardURL"),
				resource.TestCheckResourceAttr(resourceName, "config.0.type", "AKAMAI_MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_stream_cdn_configs", "test_stream_cdn_configs", acctest.Optional, acctest.Update, MediaServicesMediaServicesStreamCdnConfigDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesStreamCdnConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Optional, acctest.Update, MediaServicesStreamCdnConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "stream_cdn_config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_cdn_config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_stream_cdn_config", "test_stream_cdn_config", acctest.Required, acctest.Create, MediaServicesMediaServicesStreamCdnConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesStreamCdnConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_cdn_config_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.edge_hostname", "edgeHostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.edge_path_prefix", "edgePathPrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.edge_token_key", "edgeTokenKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.edge_token_salt", "edgeTokenSalt2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.is_edge_token_auth", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.origin_auth_secret_key_a", "originAuthSecretKeyA2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.origin_auth_secret_key_b", "originAuthSecretKeyB2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.origin_auth_secret_key_nonce_a", "originAuthSecretKeyNonceA2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.origin_auth_secret_key_nonce_b", "originAuthSecretKeyNonceB2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.origin_auth_sign_encryption", "SHA256-HMAC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.origin_auth_sign_type", "ForwardURL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "config.0.type", "AKAMAI_MANUAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesStreamCdnConfigRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMediaServicesStreamCdnConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MediaServicesClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_media_services_stream_cdn_config" {
			noResourceFound = false
			request := oci_media_services.GetStreamCdnConfigRequest{}

			tmp := rs.Primary.ID
			request.StreamCdnConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")

			response, err := client.GetStreamCdnConfig(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_media_services.StreamCdnConfigLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MediaServicesStreamCdnConfig") {
		resource.AddTestSweepers("MediaServicesStreamCdnConfig", &resource.Sweeper{
			Name:         "MediaServicesStreamCdnConfig",
			Dependencies: acctest.DependencyGraph["streamCdnConfig"],
			F:            sweepMediaServicesStreamCdnConfigResource,
		})
	}
}

func sweepMediaServicesStreamCdnConfigResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	streamCdnConfigIds, err := getMediaServicesStreamCdnConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, streamCdnConfigId := range streamCdnConfigIds {
		if ok := acctest.SweeperDefaultResourceId[streamCdnConfigId]; !ok {
			deleteStreamCdnConfigRequest := oci_media_services.DeleteStreamCdnConfigRequest{}

			deleteStreamCdnConfigRequest.StreamCdnConfigId = &streamCdnConfigId

			deleteStreamCdnConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteStreamCdnConfig(context.Background(), deleteStreamCdnConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamCdnConfig %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamCdnConfigId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamCdnConfigId, MediaServicesStreamCdnConfigSweepWaitCondition, time.Duration(3*time.Minute),
				MediaServicesStreamCdnConfigSweepResponseFetchOperation, "media_services", true)
		}
	}
	return nil
}

func getMediaServicesStreamCdnConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamCdnConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listStreamCdnConfigsRequest := oci_media_services.ListStreamCdnConfigsRequest{}
	//listStreamCdnConfigsRequest.CompartmentId = &compartmentId

	distributionChannelIds, error := getMediaServicesStreamDistributionChannelIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting distributionChannelId required for StreamCdnConfig resource requests \n")
	}
	for _, distributionChannelId := range distributionChannelIds {
		listStreamCdnConfigsRequest.DistributionChannelId = &distributionChannelId

		listStreamCdnConfigsRequest.LifecycleState = oci_media_services.StreamCdnConfigLifecycleStateActive
		listStreamCdnConfigsResponse, err := mediaServicesClient.ListStreamCdnConfigs(context.Background(), listStreamCdnConfigsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting StreamCdnConfig list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, streamCdnConfig := range listStreamCdnConfigsResponse.Items {
			id := *streamCdnConfig.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamCdnConfigId", id)
		}

	}
	return resourceIds, nil
}

func MediaServicesStreamCdnConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamCdnConfigResponse, ok := response.Response.(oci_media_services.GetStreamCdnConfigResponse); ok {
		return streamCdnConfigResponse.LifecycleState != oci_media_services.StreamCdnConfigLifecycleStateDeleted
	}
	return false
}

func MediaServicesStreamCdnConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MediaServicesClient().GetStreamCdnConfig(context.Background(), oci_media_services.GetStreamCdnConfigRequest{
		StreamCdnConfigId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
