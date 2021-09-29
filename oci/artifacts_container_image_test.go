// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_artifacts "github.com/oracle/oci-go-sdk/v48/artifacts"
	"github.com/oracle/oci-go-sdk/v48/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (

	// We currently don't support OCI style container image creation, so we can't do TF resource for image
	// Ticket to track adding the creation endpoint https://jira.oci.oraclecorp.com/browse/OCIR-2136.
	// Therefore, we need to set the env var of the pre-canned container image for testing, i.e. TF_VAR_container_image_ocid

	imageId       = getEnvSettingWithBlankDefault("container_image_ocid")
	compartmentId = getEnvSettingWithBlankDefault("tenancy_ocid")

	containerImageSingularDataSourceRepresentation = map[string]interface{}{
		"image_id": Representation{repType: Required, create: imageId},
	}

	containerImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: compartmentId},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `false`},
		"image_id":                  Representation{repType: Optional, create: imageId},
		"is_versioned":              Representation{repType: Optional, create: `true`},
		"state":                     Representation{repType: Optional, create: `AVAILABLE`},
	}

	ContainerImageResourceConfig = ""
)

// issue-routing-tag: artifacts/default
func TestArtifactsContainerImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestArtifactsContainerImageResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	datasourceName := "data.oci_artifacts_container_images.test_container_images"
	singularDatasourceName := "data.oci_artifacts_container_image.test_container_image"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_artifacts_container_images", "test_container_images", Optional, Create, containerImageDataSourceRepresentation) +
				ContainerImageResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_versioned", "true"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "container_image_collection.0.remaining_items_count", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_artifacts_container_image", "test_container_image", Required, Create, containerImageSingularDataSourceRepresentation) +
				ContainerImageResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			),
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ArtifactsContainerImage") {
		resource.AddTestSweepers("ArtifactsContainerImage", &resource.Sweeper{
			Name:         "ArtifactsContainerImage",
			Dependencies: DependencyGraph["containerImage"],
			F:            sweepArtifactsContainerImageResource,
		})
	}
}

func sweepArtifactsContainerImageResource(compartment string) error {
	artifactsClient := GetTestClients(&schema.ResourceData{}).artifactsClient()
	containerImageIds, err := getContainerImageIds(compartment)
	if err != nil {
		return err
	}
	for _, containerImageId := range containerImageIds {
		if ok := SweeperDefaultResourceId[containerImageId]; !ok {
			deleteContainerImageRequest := oci_artifacts.DeleteContainerImageRequest{}

			deleteContainerImageRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "artifacts")
			_, error := artifactsClient.DeleteContainerImage(context.Background(), deleteContainerImageRequest)
			if error != nil {
				fmt.Printf("Error deleting ContainerImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", containerImageId, error)
				continue
			}
			waitTillCondition(testAccProvider, &containerImageId, containerImageSweepWaitCondition, time.Duration(3*time.Minute),
				containerImageSweepResponseFetchOperation, "artifacts", true)
		}
	}
	return nil
}

func getContainerImageIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ContainerImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	artifactsClient := GetTestClients(&schema.ResourceData{}).artifactsClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ContainerImageId", id)
	}
	return resourceIds, nil
}

func containerImageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if containerImageResponse, ok := response.Response.(oci_artifacts.GetContainerImageResponse); ok {
		return containerImageResponse.LifecycleState != oci_artifacts.ContainerImageLifecycleStateDeleted
	}
	return false
}

func containerImageSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.artifactsClient().GetContainerImage(context.Background(), oci_artifacts.GetContainerImageRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
