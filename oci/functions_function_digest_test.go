// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_functions "github.com/oracle/oci-go-sdk/v45/functions"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

func TestFunctionsFunctionResource_digest(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsFunctionResource_digest")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	imageA1 := getEnvSettingWithBlankDefault("image")
	imageA1Digest := getEnvSettingWithBlankDefault("image_digest")

	imageB := getEnvSettingWithBlankDefault("image_for_update")
	imageBDigest := getEnvSettingWithBlankDefault("image_digest_for_update")

	imageA2 := getEnvSettingWithBlankDefault("image_same_repo")
	imageA2Digest := getEnvSettingWithBlankDefault("image_digest_same_repo")

	resourceName := "oci_functions_function.test_function"

	var resId, resId2 string

	// The following tests all operate in the same way:
	// - reset the function definition to use the image A1 at the digest a1
	// - update to a new set of coordinates
	// - confirm that the update has produced the result intended in the control plane
	// - delete the resource ready for the next test
	// We use three image/digest pairs.
	// The image A1@a1 is used to reset state for each test.
	// The image A2 uses the same base repository name as A1 (eg, iad.ocir.io/foo/bar:1 and iad.ocir.io/foo/bar:2);
	// in this case, the digests of both images are visible under the same namespace. Consequently, it's possible
	// for a user to specify an image and digest that don't appear to correspond. (This is common practice when
	// using the `:latest` tag.)
	// The image B uses a different base respository (eg, iad.ocir.io/foo/baz:1).
	type testCase struct {
		newImage       string
		newDigest      *string
		expectedImage  string
		expectedDigest string
		nonzeroPlan    bool // if the resulting plan after an apply is nonzero
	}
	var steps []resource.TestStep
	nullStr := ""
	for _, tc := range []testCase{
		// Test "no-op" operations - updating to the same image and digest
		{imageA1, &imageA1Digest, imageA1, imageA1Digest, false},
		{imageA1, &nullStr, imageA1, imageA1Digest, true}, // request an explicit lookup from the CP
		{imageA1, nil, imageA1, imageA1Digest, false},     // leave the digest unspecified

		// Test updates to an image under the same repository - with the expected digest
		{imageA2, &imageA2Digest, imageA2, imageA2Digest, false},
		{imageA2, &nullStr, imageA2, imageA2Digest, true}, // request an explicit lookup from the CP
		{imageA2, nil, imageA2, imageA2Digest, false},     // leave the digest unspecified

		// Test udpates to an image in a different repository
		{imageB, &imageBDigest, imageB, imageBDigest, false},
		{imageB, &nullStr, imageB, imageBDigest, true}, // request an explicit lookup from the CP
		{imageB, nil, imageB, imageBDigest, false},     // leave the digest unspecified

		// Finally: test unusual intra-repository updates
		// this is equivalent to updating `:latest` with an explicit digest:
		{imageA1, &imageA2Digest, imageA1, imageA2Digest, false},
		// This scenario triggers a CP lookup. Unlike the test above, there's not really a good reason for a user to do this -
		// we cover the situation with a test but don't expect to see it in practice:
		{imageA2, &imageA1Digest, imageA2, imageA2Digest, true},
	} {
		// Shadow the loop variable so that closures behave correctly
		tc2 := tc

		steps = append(steps, resource.TestStep{
			// Reset the function to A1@a1
			Config: config + compartmentIdVariableStr + FunctionResourceDependencies +
				generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Create, functionBaseRepresentation(imageA1, &imageA1Digest)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "UpdatedImageFunction"),
				resource.TestCheckResourceAttr(resourceName, "image", imageA1),
				resource.TestCheckResourceAttr(resourceName, "image_digest", imageA1Digest),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
				func(s *terraform.State) error {
					fn, err := retrieveFunctionResourceFromControlPlane(resId)
					if err != nil {
						return err
					}
					if *fn.Image != imageA1 {
						return fmt.Errorf("Resource did not have the expected image: %s != %s", *fn.Image, imageA1)
					}
					if *fn.ImageDigest != imageA1Digest {
						return fmt.Errorf("Resource did not have the expected digest: %s != %s", *fn.ImageDigest, imageA1Digest)
					}
					return nil
				},
			),
		})
		steps = append(steps, resource.TestStep{
			// Update the function with the new image coordinates
			Config: config + compartmentIdVariableStr + FunctionResourceDependencies +
				generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Update, functionBaseRepresentation(tc2.newImage, tc2.newDigest)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "UpdatedImageFunction"),
				resource.TestCheckResourceAttr(resourceName, "image", tc2.expectedImage),
				resource.TestCheckResourceAttr(resourceName, "image_digest", tc2.expectedDigest),
				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
				func(s *terraform.State) error {
					fn, err := retrieveFunctionResourceFromControlPlane(resId)
					if err != nil {
						return err
					}
					if *fn.Image != tc2.expectedImage {
						return fmt.Errorf("Resource did not update to the expected image: %s != %s", *fn.Image, tc2.expectedImage)
					}
					if *fn.ImageDigest != tc2.expectedDigest {
						return fmt.Errorf("Resource did not update to the expected digest: %s != %s", *fn.ImageDigest, tc2.expectedDigest)
					}
					return nil
				},
			),
			// If the user asks to force the CP recalculation of the image_digest, the plan will never go to zero.
			ExpectNonEmptyPlan: tc2.nonzeroPlan,
		})
		steps = append(steps, resource.TestStep{
			// delete before next step to reset all state
			Config: config + compartmentIdVariableStr + FunctionResourceDependencies,
		})
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFunctionsFunctionNoneRemaining,
		Steps:        steps,
	})
}

// functionBaseRepresentation returns the bare bones of a function resource definition - with the
// image set to the provided value, and the image_digest optionally set also.
func functionBaseRepresentation(image string, digest *string) map[string]interface{} {
	// Rather than indirect through variables, this is called to inject docker URIs and digests directly
	// Leave digest = nil to omit the field.
	m := map[string]interface{}{
		"application_id": Representation{repType: Required, create: `${oci_functions_application.test_application.id}`},
		"display_name":   Representation{repType: Required, create: `UpdatedImageFunction`},
		"image":          Representation{repType: Required, create: image, update: image},
		"memory_in_mbs":  Representation{repType: Required, create: `128`, update: `128`},
	}
	if digest != nil {
		m["image_digest"] = Representation{repType: Optional, create: *digest, update: *digest}
	}
	return m
}

// We lean on the init() in functions_function_test.go

// testAccCheckFunctionsFunctionNoneRemaining should not find any resources - the test case cleans them up
func testAccCheckFunctionsFunctionNoneRemaining(s *terraform.State) error {
	client := testAccProvider.Meta().(*OracleClients).functionsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_functions_function" {
			request := oci_functions.GetFunctionRequest{}

			tmp := rs.Primary.ID
			request.FunctionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "functions")

			response, err := client.GetFunction(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_functions.FunctionLifecycleStateDeleted): true,
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

	return nil
}

func retrieveFunctionResourceFromControlPlane(id string) (oci_functions.GetFunctionResponse, error) {
	client := testAccProvider.Meta().(*OracleClients).functionsManagementClient()
	request := oci_functions.GetFunctionRequest{}
	request.FunctionId = &id
	request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "functions")
	return client.GetFunction(context.Background(), request)
}

func TestFunctionsFunctionResource_digest_create(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsFunctionResource_digest_create")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	imageA1 := getEnvSettingWithBlankDefault("image")
	imageA1Digest := getEnvSettingWithBlankDefault("image_digest")

	resourceName := "oci_functions_function.test_function"

	var resId string

	// The following tests check the three approaches to creating a function:
	// - specifying the pair (image, image_digest) for a fully-resolved image;
	// - omitting the image_digest value and having the controlplane loko it up;
	// - using the sentinel empty string to force a CP-side lookup.
	// After each create the result should be the same.
	// This just uses the main image. The first two of these tests are morally equivalent to
	// some of the tests performed by TestFunctionsFunctionResource_basic
	type testCase struct {
		newDigest   *string
		nonzeroPlan bool // if the resulting plan after an apply is nonzero
	}
	var steps []resource.TestStep
	nullStr := ""
	for _, tc := range []testCase{
		// Test "no-op" operations - updating to the same image and digest
		{&imageA1Digest, false},
		{&nullStr, true}, // request an explicit lookup from the CP - this won't plan-to-zero
		{nil, false},     // leave the digest unspecified
	} {
		steps = append(steps, resource.TestStep{
			// Create a function at A1@a1, through one means or another
			Config: config + compartmentIdVariableStr + FunctionResourceDependencies +
				generateResourceFromRepresentationMap("oci_functions_function", "test_function", Optional, Create, functionBaseRepresentation(imageA1, tc.newDigest)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "UpdatedImageFunction"),
				resource.TestCheckResourceAttr(resourceName, "image", imageA1),
				resource.TestCheckResourceAttr(resourceName, "image_digest", imageA1Digest),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
				func(s *terraform.State) error {
					fn, err := retrieveFunctionResourceFromControlPlane(resId)
					if err != nil {
						return err
					}
					if *fn.Image != imageA1 {
						return fmt.Errorf("Resource did not have the expected image: %s != %s", *fn.Image, imageA1)
					}
					if *fn.ImageDigest != imageA1Digest {
						return fmt.Errorf("Resource did not have the expected digest: %s != %s", *fn.ImageDigest, imageA1Digest)
					}
					return nil
				},
			),
			ExpectNonEmptyPlan: tc.nonzeroPlan,
		})
		steps = append(steps, resource.TestStep{
			// delete before next step to reset all state
			Config: config + compartmentIdVariableStr + FunctionResourceDependencies,
		})
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFunctionsFunctionNoneRemaining,
		Steps:        steps,
	})
}
