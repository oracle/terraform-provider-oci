// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (

	// We currently don't support OCI style container image creation, so we can't do TF resource for image
	// Ticket to track adding the creation endpoint https://jira.oci.oraclecorp.com/browse/OCIR-2136.
	// Therefore, we need to set the env var of the pre-canned container image for testing, i.e. TF_VAR_container_image_ocid

	imageId       = utils.GetEnvSettingWithBlankDefault("container_image_ocid")
	compartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	ArtifactsArtifactscontainerImageSingularDataSourceRepresentation = map[string]interface{}{
		"image_id": acctest.Representation{RepType: acctest.Required, Create: imageId},
	}

	ArtifactsArtifactscontainerImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: compartmentId},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"image_id":                  acctest.Representation{RepType: acctest.Optional, Create: imageId},
		"is_versioned":              acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	ArtifactsContainerImageResourceConfig = ""
)

// issue-routing-tag: artifacts/default
func TestArtifactsContainerImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsContainerImageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_artifacts_container_images.test_container_images"
	singularDatasourceName := "data.oci_artifacts_container_image.test_container_image"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_images", "test_container_images", acctest.Optional, acctest.Create, ArtifactsArtifactscontainerImageDataSourceRepresentation) +
				ArtifactsContainerImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_versioned", "true"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.0.items.0.defined_tags.%", "2"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.0.items.0.freeform_tags.%", "3"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.0.remaining_items_count", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_artifacts_container_image", "test_container_image", acctest.Required, acctest.Create, ArtifactsArtifactscontainerImageSingularDataSourceRepresentation) +
				ArtifactsContainerImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digest"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "layers.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "layers_size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "manifest_size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pull_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "versions.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "2"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ArtifactsContainerImage") {
		resource.AddTestSweepers("ArtifactsContainerImage", &resource.Sweeper{
			Name:         "ArtifactsContainerImage",
			Dependencies: acctest.DependencyGraph["containerImage"],
			F:            sweepArtifactsContainerImageResource,
		})
	}
}

func sweepArtifactsContainerImageResource(compartment string) error {
	artifactsClient := acctest.GetTestClients(&schema.ResourceData{}).ArtifactsClient()
	containerImageIds, err := getArtifactsContainerImageIds(compartment)
	if err != nil {
		return err
	}
	for _, containerImageId := range containerImageIds {
		if ok := acctest.SweeperDefaultResourceId[containerImageId]; !ok {
			deleteContainerImageRequest := oci_artifacts.DeleteContainerImageRequest{}

			deleteContainerImageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "artifacts")
			_, error := artifactsClient.DeleteContainerImage(context.Background(), deleteContainerImageRequest)
			if error != nil {
				fmt.Printf("Error deleting ContainerImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", containerImageId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &containerImageId, ArtifactscontainerImagesSweepWaitCondition, time.Duration(3*time.Minute),
				ArtifactscontainerImagesSweepResponseFetchOperation, "artifacts", true)
		}
	}
	return nil
}

func getArtifactsContainerImageIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ContainerImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	artifactsClient := acctest.GetTestClients(&schema.ResourceData{}).ArtifactsClient()

	listContainerImagesRequest := oci_artifacts.ListContainerImagesRequest{}
	listContainerImagesRequest.CompartmentId = &compartmentId
	var containerImageLifecycleStateAvailable = string(oci_artifacts.ContainerImageLifecycleStateAvailable)
	listContainerImagesRequest.LifecycleState = &containerImageLifecycleStateAvailable
	listContainerImagesResponse, err := artifactsClient.ListContainerImages(context.Background(), listContainerImagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ContainerImage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, containerImage := range listContainerImagesResponse.Items {
		id := *containerImage.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ContainerImageId", id)
	}
	return resourceIds, nil
}

func ArtifactscontainerImagesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if containerImageResponse, ok := response.Response.(oci_artifacts.GetContainerImageResponse); ok {
		return containerImageResponse.LifecycleState != oci_artifacts.ContainerImageLifecycleStateDeleted
	}
	return false
}

func ArtifactscontainerImagesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ArtifactsClient().GetContainerImage(context.Background(), oci_artifacts.GetContainerImageRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
