// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InvokeFunctionRequiredOnlyResource = InvokeFunctionResourceDependencies +
		generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Required, Create, invokeFunctionRepresentation)

	InvokeFunctionResourceConfig = InvokeFunctionResourceDependencies +
		generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Update, invokeFunctionRepresentation)

	invokeFunctionSingularDataSourceRepresentation = map[string]interface{}{}

	invokeFunctionRepresentation = map[string]interface{}{
		"function_id":          Representation{repType: Required, create: `${oci_functions_function.test_function.id}`},
		"invoke_function_body": Representation{repType: Optional, create: `{\"name\":\"Bob\"}`},
		"fn_intent":            Representation{repType: Optional, create: `httprequest`},
		"fn_invoke_type":       Representation{repType: Optional, create: `sync`},
	}

	InvokeFunctionResourceDependencies = FunctionRequiredOnlyResource
	sourceFile                         *os.File
)

func createTmpSourceFile() (string, error) {
	sourceFile, err := ioutil.TempFile(os.TempDir(), "source-")
	if err != nil {
		return "", err
	}

	text := []byte("{\"name\":\"Bob\"}")
	if _, err = sourceFile.Write(text); err != nil {
		return "", err
	}

	// Close the file
	if err := sourceFile.Close(); err != nil {
		return "", err
	}

	return sourceFile.Name(), nil
}

func TestFunctionsInvokeFunctionResource_basic(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skipping TestFunctionsInvokeFunctionResource_basic in HttpReplay mode till json encoding is fixed.")
	}

	httpreplay.SetScenario("TestFunctionsInvokeFunctionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := getEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	imageDigest := getEnvSettingWithBlankDefault("image_digest")
	imageDigestVariableStr := fmt.Sprintf("variable \"image_digest\" { default = \"%s\" }\n", imageDigest)

	resourceName := "oci_functions_invoke_function.test_invoke_function"
	sourceFilePath, err := createTmpSourceFile()
	if err != nil {
		t.Fatalf("Unable to create files for invocation. Error: %q", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckFunctionsInvokeFunctionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Required, Create, invokeFunctionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "function_id"),
					resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 World\"}\n"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create, invokeFunctionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
				),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create,
						getUpdatedRepresentationCopy("fn_intent", Representation{repType: Optional, create: `cloudevent`}, invokeFunctionRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
				),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create,
						getUpdatedRepresentationCopy("fn_invoke_type", Representation{repType: Optional, create: `detached`}, invokeFunctionRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "function_id"),
				),
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create,
						getUpdatedRepresentationCopy("fn_intent", Representation{repType: Optional, create: `cloudevent`}, getUpdatedRepresentationCopy("fn_invoke_type", Representation{repType: Optional, create: `detached`}, invokeFunctionRepresentation))),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "function_id"),
				),
			},
			// verify create with source path
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(invokeFunctionRepresentation, []string{"invoke_function_body"}), map[string]interface{}{
							"input_body_source_path": Representation{repType: Optional, create: sourceFilePath},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
				),
			},
			// verify create with base64 encoded input
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(invokeFunctionRepresentation, []string{"invoke_function_body"}), map[string]interface{}{
							"invoke_function_body_base64_encoded": Representation{repType: Optional, create: "eyJuYW1lIjoiQm9iIn0="},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
				),
			},
			// verify base64 encoded content
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
					generateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", Optional, Create,
						representationCopyWithNewProperties(invokeFunctionRepresentation, map[string]interface{}{
							"base64_encode_content": Representation{repType: Optional, create: `true`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content", "eyJtZXNzYWdlIjoiSGVsbG8gdjMgQm9iIn0K"),
				),
			},
		},
	})
}

func testAccCheckFunctionsInvokeFunctionDestroy(s *terraform.State) error {
	if sourceFile != nil {
		if _, err := os.Stat(sourceFile.Name()); err == nil {
			os.Remove(sourceFile.Name())
		}
	}
	return nil
}
