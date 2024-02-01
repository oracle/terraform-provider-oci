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
	MediaServicesMediaAssetRequiredOnlyResource = MediaServicesMediaAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Required, acctest.Create, MediaServicesMediaAssetRepresentation)

	MediaServicesMediaAssetResourceConfig = MediaServicesMediaAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Optional, acctest.Update, MediaServicesMediaAssetRepresentation)

	MediaServicesMediaServicesMediaAssetSingularDataSourceRepresentation = map[string]interface{}{
		"media_asset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_asset.test_media_asset.id}`},
	}

	MediaServicesMediaServicesMediaAssetDataSourceRepresentation = map[string]interface{}{
		"bucket":                        acctest.Representation{RepType: acctest.Optional, Create: `bucket`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"master_media_asset_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_asset.test_media_asset.id}`},
		"media_workflow_job_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow_job.test_media_workflow_job.id}`},
		"object":                        acctest.Representation{RepType: acctest.Optional, Create: `object`},
		"parent_media_asset_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_asset.test_media_asset.id}`},
		"source_media_workflow_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow.test_media_workflow.id}`},
		"source_media_workflow_version": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"state":                         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                          acctest.Representation{RepType: acctest.Optional, Create: `AUDIO`, Update: `VIDEO`},
		"filter":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: MediaServicesMediaAssetDataSourceFilterRepresentation}}

	MediaServicesMediaAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_media_services_media_asset.test_media_asset.id}`}},
	}

	MediaServicesMediaAssetRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `AUDIO`, Update: `VIDEO`},
		"is_lock_override":              acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"bucket":                        acctest.Representation{RepType: acctest.Optional, Create: `bucket`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesMediaAssetLocksRepresentation},
		"media_asset_tags":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesMediaAssetMediaAssetTagsRepresentation},
		"media_workflow_job_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow_job.test_media_workflow_job.id}`},
		"metadata":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: MediaServicesMediaAssetMetadataRepresentation},
		"namespace":                     acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
		"object":                        acctest.Representation{RepType: acctest.Optional, Create: `object`},
		"object_etag":                   acctest.Representation{RepType: acctest.Optional, Create: `objectEtag`},
		"segment_range_end_index":       acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"segment_range_start_index":     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"source_media_workflow_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_media_services_media_workflow.test_media_workflow.id}`},
		"source_media_workflow_version": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsAndLocks},
	}

	MediaServicesMediaAssetLocksRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message":        acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	MediaServicesMediaAssetMediaAssetTagsRepresentation = map[string]interface{}{
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
		"type":  acctest.Representation{RepType: acctest.Optional, Create: `USER`, Update: `SYSTEM`},
	}

	MediaServicesMediaAssetMetadataRepresentation = map[string]interface{}{
		"metadata": acctest.Representation{RepType: acctest.Required, Create: `{\"some\":\"json\"}`, Update: `{\"some\":\"json2\"}`},
	}

	MediaServicesMediaAssetResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow_job", "test_media_workflow_job", acctest.Required, acctest.Create, MediaServicesMediaWorkflowJobRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_workflow", "test_media_workflow", acctest.Required, acctest.Create, MediaServicesMediaWorkflowRepresentation)
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesMediaAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_media_services_media_asset.test_media_asset"
	datasourceName := "data.oci_media_services_media_assets.test_media_assets"
	singularDatasourceName := "data.oci_media_services_media_asset.test_media_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MediaServicesMediaAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Optional, acctest.Create, MediaServicesMediaAssetRepresentation), "mediaservices", "mediaAsset", t)

	acctest.ResourceTest(t, testAccCheckMediaServicesMediaAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Required, acctest.Create, MediaServicesMediaAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "AUDIO"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaAssetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MediaServicesMediaAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Optional, acctest.Create, MediaServicesMediaAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.0.type", "USER"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_job_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.metadata", `{"some":"json"}`),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),
				resource.TestCheckResourceAttr(resourceName, "object_etag", "objectEtag"),
				resource.TestCheckResourceAttr(resourceName, "segment_range_end_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "segment_range_start_index", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_media_workflow_id"),
				resource.TestCheckResourceAttr(resourceName, "source_media_workflow_version", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "AUDIO"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MediaServicesMediaAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MediaServicesMediaAssetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.0.type", "USER"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_job_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.metadata", `{"some":"json"}`),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),
				resource.TestCheckResourceAttr(resourceName, "object_etag", "objectEtag"),
				resource.TestCheckResourceAttr(resourceName, "segment_range_end_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "segment_range_start_index", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_media_workflow_id"),
				resource.TestCheckResourceAttr(resourceName, "source_media_workflow_version", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "AUDIO"),

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
			Config: config + compartmentIdVariableStr + MediaServicesMediaAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Optional, acctest.Update, MediaServicesMediaAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.0.type", "SYSTEM"),
				resource.TestCheckResourceAttr(resourceName, "media_asset_tags.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "media_workflow_job_id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.0.metadata", `{"some":"json2"}`),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "object"),
				resource.TestCheckResourceAttr(resourceName, "object_etag", "objectEtag"),
				resource.TestCheckResourceAttr(resourceName, "segment_range_end_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "segment_range_start_index", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "source_media_workflow_id"),
				resource.TestCheckResourceAttr(resourceName, "source_media_workflow_version", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "VIDEO"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_assets", "test_media_assets", acctest.Optional, acctest.Update, MediaServicesMediaServicesMediaAssetDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Optional, acctest.Update, MediaServicesMediaAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", "bucket"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "media_workflow_job_id"),
				resource.TestCheckResourceAttr(datasourceName, "object", "object"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_media_workflow_id"),
				resource.TestCheckResourceAttr(datasourceName, "source_media_workflow_version", "10"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", "VIDEO"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaAssetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "media_asset_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", "bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "media_asset_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "media_asset_tags.0.type", "SYSTEM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "media_asset_tags.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.0.metadata", `{"some":"json2"}`),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "object"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_etag", "objectEtag"),
				resource.TestCheckResourceAttr(singularDatasourceName, "segment_range_end_index", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "segment_range_start_index", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_media_workflow_version", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "VIDEO"),
			),
		},
		// verify resource import
		{
			Config:            config + MediaServicesMediaAssetRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_lock_override",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMediaServicesMediaAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MediaServicesClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_media_services_media_asset" {
			noResourceFound = false
			request := oci_media_services.GetMediaAssetRequest{}

			tmp := rs.Primary.ID
			request.MediaAssetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")

			response, err := client.GetMediaAsset(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_media_services.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MediaServicesMediaAsset") {
		resource.AddTestSweepers("MediaServicesMediaAsset", &resource.Sweeper{
			Name:         "MediaServicesMediaAsset",
			Dependencies: acctest.DependencyGraph["mediaAsset"],
			F:            sweepMediaServicesMediaAssetResource,
		})
	}
}

func sweepMediaServicesMediaAssetResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	mediaAssetIds, err := getMediaServicesMediaAssetIds(compartment)
	if err != nil {
		return err
	}
	for _, mediaAssetId := range mediaAssetIds {
		if ok := acctest.SweeperDefaultResourceId[mediaAssetId]; !ok {
			deleteMediaAssetRequest := oci_media_services.DeleteMediaAssetRequest{}

			deleteMediaAssetRequest.MediaAssetId = &mediaAssetId

			deleteMediaAssetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteMediaAsset(context.Background(), deleteMediaAssetRequest)
			if error != nil {
				fmt.Printf("Error deleting MediaAsset %s %s, It is possible that the resource is already deleted. Please verify manually \n", mediaAssetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mediaAssetId, MediaServicesMediaAssetSweepWaitCondition, time.Duration(3*time.Minute),
				MediaServicesMediaAssetSweepResponseFetchOperation, "media_services", true)
		}
	}
	return nil
}

func getMediaServicesMediaAssetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MediaAssetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listMediaAssetsRequest := oci_media_services.ListMediaAssetsRequest{}
	listMediaAssetsRequest.CompartmentId = &compartmentId
	listMediaAssetsRequest.LifecycleState = oci_media_services.ListMediaAssetsLifecycleStateActive
	listMediaAssetsResponse, err := mediaServicesClient.ListMediaAssets(context.Background(), listMediaAssetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MediaAsset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mediaAsset := range listMediaAssetsResponse.Items {
		id := *mediaAsset.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MediaAssetId", id)
	}
	return resourceIds, nil
}

func MediaServicesMediaAssetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mediaAssetResponse, ok := response.Response.(oci_media_services.GetMediaAssetResponse); ok {
		return mediaAssetResponse.LifecycleState != oci_media_services.LifecycleStateDeleted
	}
	return false
}

func MediaServicesMediaAssetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MediaServicesClient().GetMediaAsset(context.Background(), oci_media_services.GetMediaAssetRequest{
		MediaAssetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
