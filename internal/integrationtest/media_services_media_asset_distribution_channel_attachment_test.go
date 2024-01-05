// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MediaServicesMediaServicesMediaAssetDistributionChannelAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"distribution_channel_id": acctest.Representation{RepType: acctest.Required, Create: `//`},
		"media_asset_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_media_services_media_asset.test_media_asset.id}`},
		"version":                 acctest.Representation{RepType: acctest.Optional, Create: `{}`},
	}

	MediaServicesMediaAssetDistributionChannelAttachmentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_media_services_media_asset", "test_media_asset", acctest.Required, acctest.Create, MediaServicesMediaAssetRepresentation)
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaAssetDistributionChannelAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesMediaAssetDistributionChannelAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_media_services_media_asset_distribution_channel_attachment.test_media_asset_distribution_channel_attachment"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_asset_distribution_channel_attachment", "test_media_asset_distribution_channel_attachment", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaAssetDistributionChannelAttachmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaAssetDistributionChannelAttachmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "distribution_channel_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "media_asset_id"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MediaServicesMediaAssetDistributionChannelAttachment") {
		resource.AddTestSweepers("MediaServicesMediaAssetDistributionChannelAttachment", &resource.Sweeper{
			Name:         "MediaServicesMediaAssetDistributionChannelAttachment",
			Dependencies: acctest.DependencyGraph["mediaAssetDistributionChannelAttachment"],
			F:            sweepMediaServicesMediaAssetDistributionChannelAttachmentResource,
		})
	}
}

func sweepMediaServicesMediaAssetDistributionChannelAttachmentResource(compartment string) error {
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()
	mediaAssetDistributionChannelAttachmentIds, err := getMediaServicesMediaAssetDistributionChannelAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, mediaAssetDistributionChannelAttachmentId := range mediaAssetDistributionChannelAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[mediaAssetDistributionChannelAttachmentId]; !ok {
			deleteMediaAssetDistributionChannelAttachmentRequest := oci_media_services.DeleteMediaAssetDistributionChannelAttachmentRequest{}

			deleteMediaAssetDistributionChannelAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "media_services")
			_, error := mediaServicesClient.DeleteMediaAssetDistributionChannelAttachment(context.Background(), deleteMediaAssetDistributionChannelAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting MediaAssetDistributionChannelAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", mediaAssetDistributionChannelAttachmentId, error)
				continue
			}
		}
	}
	return nil
}

func getMediaServicesMediaAssetDistributionChannelAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MediaAssetDistributionChannelAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mediaServicesClient := acctest.GetTestClients(&schema.ResourceData{}).MediaServicesClient()

	listMediaAssetDistributionChannelAttachmentsRequest := oci_media_services.ListMediaAssetDistributionChannelAttachmentsRequest{}
	//listMediaAssetDistributionChannelAttachmentsRequest.CompartmentId = &compartmentId

	mediaAssetIds, error := getMediaServicesMediaAssetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting mediaAssetId required for MediaAssetDistributionChannelAttachment resource requests \n")
	}
	for _, mediaAssetId := range mediaAssetIds {
		listMediaAssetDistributionChannelAttachmentsRequest.MediaAssetId = &mediaAssetId

		listMediaAssetDistributionChannelAttachmentsResponse, err := mediaServicesClient.ListMediaAssetDistributionChannelAttachments(context.Background(), listMediaAssetDistributionChannelAttachmentsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting MediaAssetDistributionChannelAttachment list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, mediaAssetDistributionChannelAttachment := range listMediaAssetDistributionChannelAttachmentsResponse.Items {
			id := *mediaAssetDistributionChannelAttachment.DistributionChannelId
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MediaAssetDistributionChannelAttachmentId", id)
		}

	}
	return resourceIds, nil
}
