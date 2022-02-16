// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ImageRequiredOnlyResource = ImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Required, acctest.Create, imageRepresentation)

	ImageResourceConfig = ImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Optional, acctest.Update, imageRepresentation)

	imageSingularDataSourceRepresentation = map[string]interface{}{
		"image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_image.test_image.id}`},
	}

	imageDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyCustomImage`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: imageDataSourceFilterRepresentation}}
	imageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_image.test_image.id}`}},
	}

	imageRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyCustomImage`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"instance_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"launch_mode":    acctest.Representation{RepType: acctest.Optional, Create: `NATIVE`},
		"timeouts":       acctest.RepresentationGroup{RepType: acctest.Required, Group: timeoutsRepresentation},
	}

	timeoutsRepresentation = map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `30m`},
	}

	ImageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeImaging
func TestCoreImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreImageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_image.test_image"
	datasourceName := "data.oci_core_images.test_images"
	singularDatasourceName := "data.oci_core_image.test_image"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ImageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Optional, acctest.Create, imageRepresentation), "core", "image", t)

	acctest.ResourceTest(t, testAccCheckCoreImageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Required, acctest.Create, imageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ImageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Optional, acctest.Create, imageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
				resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(imageRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
				resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Optional, acctest.Update, imageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
				resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
				resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_images", "test_images", acctest.Optional, acctest.Update, imageDataSourceRepresentation) +
				compartmentIdVariableStr + ImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Optional, acctest.Update, imageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "images.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "images.0.agent_features.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.base_image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.billable_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "images.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.create_image_allowed"),
				resource.TestCheckResourceAttr(datasourceName, "images.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "images.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "images.0.launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(datasourceName, "images.0.launch_options.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.size_in_mbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "images.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_image", "test_image", acctest.Required, acctest.Create, imageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "agent_features.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "base_image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "billable_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "create_image_allowed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_mode", "NATIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "launch_options.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operating_system"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operating_system_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_mbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ImageResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"image_source_details",
				"instance_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreImageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_image" {
			noResourceFound = false
			request := oci_core.GetImageRequest{}

			tmp := rs.Primary.ID
			request.ImageId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetImage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.ImageLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CoreImage") {
		resource.AddTestSweepers("CoreImage", &resource.Sweeper{
			Name:         "CoreImage",
			Dependencies: acctest.DependencyGraph["image"],
			F:            sweepCoreImageResource,
		})
	}
}

func sweepCoreImageResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	imageIds, err := getImageIds(compartment)
	if err != nil {
		return err
	}
	for _, imageId := range imageIds {
		if ok := acctest.SweeperDefaultResourceId[imageId]; !ok {
			deleteImageRequest := oci_core.DeleteImageRequest{}

			deleteImageRequest.ImageId = &imageId

			deleteImageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteImage(context.Background(), deleteImageRequest)
			if error != nil {
				fmt.Printf("Error deleting Image %s %s, It is possible that the resource is already deleted. Please verify manually \n", imageId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &imageId, imageSweepWaitCondition, time.Duration(3*time.Minute),
				imageSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getImageIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listImagesRequest := oci_core.ListImagesRequest{}
	listImagesRequest.CompartmentId = &compartmentId
	listImagesRequest.LifecycleState = oci_core.ImageLifecycleStateAvailable
	listImagesResponse, err := computeClient.ListImages(context.Background(), listImagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Image list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, image := range listImagesResponse.Items {
		if image.CompartmentId != nil && *image.CompartmentId == compartment && image.BaseImageId != nil {
			id := *image.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ImageId", id)
		}
	}
	return resourceIds, nil
}

func imageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if imageResponse, ok := response.Response.(oci_core.GetImageResponse); ok {
		return imageResponse.LifecycleState != oci_core.ImageLifecycleStateDeleted
	}
	return false
}

func imageSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeClient().GetImage(context.Background(), oci_core.GetImageRequest{
		ImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
