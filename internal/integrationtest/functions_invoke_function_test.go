// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	FunctionsInvokeFunctionRequiredOnlyResource = FunctionsInvokeFunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Required, acctest.Create, FunctionsInvokeFunctionRepresentation)

	InvokeFunctionResourceConfig = FunctionsInvokeFunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Update, FunctionsInvokeFunctionRepresentation)

	FunctionsInvokeFunctionSingularDataSourceRepresentation = map[string]interface{}{}

	FunctionsInvokeFunctionRepresentation = map[string]interface{}{
		"function_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
		"invoke_function_body": acctest.Representation{RepType: acctest.Optional, Create: `{\"name\":\"Bob\"}`},
		"fn_intent":            acctest.Representation{RepType: acctest.Optional, Create: `httprequest`},
		"is_dry_run":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"fn_invoke_type":       acctest.Representation{RepType: acctest.Optional, Create: `sync`},
	}

	invokeApplicationDisplayName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	FunctionsInvokeFunctionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("display_name", acctest.Representation{RepType: acctest.Required, Create: invokeApplicationDisplayName}, FunctionsApplicationRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
		AvailabilityDomainConfig +
		CoreDhcpOptionsRequiredOnlyResource +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig

	sourceFile *os.File
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

// issue-routing-tag: functions/default
func TestFunctionsInvokeFunctionResource_basic(t *testing.T) {

	if httpreplay.ModeRecordReplay() {
		t.Skip("Skipping TestFunctionsInvokeFunctionResource_basic in HttpReplay mode till json encoding is fixed.")
	}

	httpreplay.SetScenario("TestFunctionsInvokeFunctionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	imageDigest := utils.GetEnvSettingWithBlankDefault("image_digest")
	imageDigestVariableStr := fmt.Sprintf("variable \"image_digest\" { default = \"%s\" }\n", imageDigest)

	resourceName := "oci_functions_invoke_function.test_invoke_function"

	var resId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FunctionsInvokeFunctionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create, FunctionsInvokeFunctionRepresentation), "functions", "invokeFunction", t)

	sourceFilePath, err := createTmpSourceFile()
	if err != nil {
		t.Fatalf("Unable to Create files for invocation. Error: %q", err)
	}

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Required, acctest.Create, FunctionsInvokeFunctionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\": \"Hello World\"}"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create, FunctionsInvokeFunctionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "invoke_function_body", "{\"name\":\"Bob\"}"),
				resource.TestCheckResourceAttr(resourceName, "fn_intent", "httprequest"),
				resource.TestCheckResourceAttr(resourceName, "fn_invoke_type", "sync"),
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
				resource.TestCheckResourceAttr(resourceName, "is_dry_run", "false"),

				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\": \"Hello Bob\"}"),
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
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("fn_intent", acctest.Representation{RepType: acctest.Optional, Create: `cloudevent`}, FunctionsInvokeFunctionRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\": \"Hello Bob\"}"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("fn_invoke_type", acctest.Representation{RepType: acctest.Optional, Create: `detached`}, FunctionsInvokeFunctionRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("fn_intent", acctest.Representation{RepType: acctest.Optional, Create: `cloudevent`}, acctest.GetUpdatedRepresentationCopy("fn_invoke_type", acctest.Representation{RepType: acctest.Optional, Create: `detached`}, FunctionsInvokeFunctionRepresentation))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
			),
		},
		// verify Create with source path
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(FunctionsInvokeFunctionRepresentation, []string{"invoke_function_body"}), map[string]interface{}{
						"input_body_source_path": acctest.Representation{RepType: acctest.Optional, Create: sourceFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\": \"Hello Bob\"}"),
			),
		},
		// verify Create with base64 encoded input
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(FunctionsInvokeFunctionRepresentation, []string{"invoke_function_body"}), map[string]interface{}{
						"invoke_function_body_base64_encoded": acctest.Representation{RepType: acctest.Optional, Create: "eyJuYW1lIjoiQm9iIn0="},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\": \"Hello Bob\"}"),
			),
		},
		// verify base64 encoded content
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + FunctionsInvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FunctionsInvokeFunctionRepresentation, map[string]interface{}{
						"base64_encode_content": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "eyJtZXNzYWdlIjogIkhlbGxvIEJvYiJ9"),
			),
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
