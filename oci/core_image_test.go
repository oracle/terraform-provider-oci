// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	ImageRequiredOnlyResource = ImageResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_image", "test_image", Required, Create, imageRepresentation)

	imageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":             Representation{repType: Optional, create: `MyCustomImage`, update: `displayName2`},
		"operating_system":         Representation{repType: Optional, create: `operatingSystem`},
		"operating_system_version": Representation{repType: Optional, create: `operatingSystemVersion`},
		"shape":                    Representation{repType: Optional, create: `shape`},
		"state":                    Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                   RepresentationGroup{Required, imageDataSourceFilterRepresentation}}
	imageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_image.test_image.id}`}},
	}

	imageRepresentation = map[string]interface{}{
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         Representation{repType: Optional, create: `MyCustomImage`, update: `displayName2`},
		"freeform_tags":        Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"image_source_details": RepresentationGroup{Optional, imageImageSourceDetailsRepresentation},
		"instance_id":          Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"launch_mode":          Representation{repType: Optional, create: `NATIVE`},
		"timeouts":             RepresentationGroup{Required, timeoutsRepresentation},
	}
	imageImageSourceDetailsRepresentation = map[string]interface{}{
		"source_type":       Representation{repType: Required, create: `objectStorageTuple`},
		"bucket_name":       Representation{repType: Optional, create: `MyBucket`},
		"namespace_name":    Representation{repType: Optional, create: `MyNamespace`},
		"object_name":       Representation{repType: Optional, create: `image-to-import.qcow2`},
		"source_image_type": Representation{repType: Optional, create: `QCOW2`},
		"source_uri":        Representation{repType: Optional, create: `https://objectstorage.us-phoenix-1.oraclecloud.com/n/MyNamespace/b/MyBucket/o/image-to-import.qcow2`},
	}

	timeoutsRepresentation = map[string]interface{}{
		"create": Representation{repType: Required, create: `30m`},
	}

	ImageResourceDependencies = InstanceRequiredOnlyResource
)

func TestCoreImageResource_basic(t *testing.T) {
	t.Skip("Long running test")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_image.test_image"
	datasourceName := "data.oci_core_images.test_images"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreImageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_image", "test_image", Required, Create, imageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ImageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_image", "test_image", Optional, Create, imageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "QCOW2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_image", "test_image", Optional, Update, imageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "QCOW2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_images", "test_images", Optional, Update, imageDataSourceRepresentation) +
					compartmentIdVariableStr + ImageResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_image", "test_image", Optional, Update, imageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "images.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.create_image_allowed"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.launch_mode", "NATIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.time_created"),
				),
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
		},
	})
}

func testAccCheckCoreImageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_image" {
			noResourceFound = false
			request := oci_core.GetImageRequest{}

			tmp := rs.Primary.ID
			request.ImageId = &tmp

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
