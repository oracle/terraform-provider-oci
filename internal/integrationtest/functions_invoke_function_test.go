// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InvokeFunctionRequiredOnlyResource = InvokeFunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Required, acctest.Create, invokeFunctionRepresentation)

	InvokeFunctionResourceConfig = InvokeFunctionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Update, invokeFunctionRepresentation)

	invokeFunctionSingularDataSourceRepresentation = map[string]interface{}{}

	invokeFunctionRepresentation = map[string]interface{}{
		"function_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
		"invoke_function_body": acctest.Representation{RepType: acctest.Optional, Create: `{\"name\":\"Bob\"}`},
		"fn_intent":            acctest.Representation{RepType: acctest.Optional, Create: `httprequest`},
		"fn_invoke_type":       acctest.Representation{RepType: acctest.Optional, Create: `sync`},
	}

	invokeApplicationDisplayName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	InvokeFunctionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("display_name", acctest.Representation{RepType: acctest.Required, Create: invokeApplicationDisplayName}, applicationRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
		AvailabilityDomainConfig +
		DhcpOptionsRequiredOnlyResource +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig +
		`
	resource "oci_core_security_list" "test_security_list" {
		compartment_id = "${var.compartment_id}"
		egress_security_rules {
    		destination = "0.0.0.0/0"
    		protocol    = "6"
  		}
		ingress_security_rules {
			protocol = "1"
			source = "10.0.1.0/24"
		}
		vcn_id = "${oci_core_vcn.test_vcn.id}"
	}

	resource "oci_core_subnet" "test_subnet" {
		availability_domain = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}"
		dhcp_options_id = "${oci_core_dhcp_options.test_dhcp_options.id}"
		display_name = "tf-subnet"
		dns_label = "dnslabel"
		freeform_tags = {
			"Department" = "Accounting"
		}
		prohibit_public_ip_on_vnic = "false"
		route_table_id = "${oci_core_route_table.test_route_table.id}"
		security_list_ids = ["${oci_core_security_list.test_security_list.id}"]
		vcn_id = "${oci_core_vcn.test_vcn.id}"
	}
	`
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
	t.Skip("Skipping test until functions support async life cycle state transitions.")

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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InvokeFunctionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create, invokeFunctionRepresentation), "functions", "invokeFunction", t)

	sourceFilePath, err := createTmpSourceFile()
	if err != nil {
		t.Fatalf("Unable to Create files for invocation. Error: %q", err)
	}

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Required, acctest.Create, invokeFunctionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 World\"}\n"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create, invokeFunctionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
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
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("fn_intent", acctest.Representation{RepType: acctest.Optional, Create: `cloudevent`}, invokeFunctionRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("fn_invoke_type", acctest.Representation{RepType: acctest.Optional, Create: `detached`}, invokeFunctionRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("fn_intent", acctest.Representation{RepType: acctest.Optional, Create: `cloudevent`}, acctest.GetUpdatedRepresentationCopy("fn_invoke_type", acctest.Representation{RepType: acctest.Optional, Create: `detached`}, invokeFunctionRepresentation))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "function_id"),
			),
		},
		// verify Create with source path
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(invokeFunctionRepresentation, []string{"invoke_function_body"}), map[string]interface{}{
						"input_body_source_path": acctest.Representation{RepType: acctest.Optional, Create: sourceFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
			),
		},
		// verify Create with base64 encoded input
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(invokeFunctionRepresentation, []string{"invoke_function_body"}), map[string]interface{}{
						"invoke_function_body_base64_encoded": acctest.Representation{RepType: acctest.Optional, Create: "eyJuYW1lIjoiQm9iIn0="},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "{\"message\":\"Hello v3 Bob\"}\n"),
			),
		},
		// verify base64 encoded content
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + imageDigestVariableStr + InvokeFunctionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_functions_invoke_function", "test_invoke_function", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(invokeFunctionRepresentation, map[string]interface{}{
						"base64_encode_content": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", "eyJtZXNzYWdlIjoiSGVsbG8gdjMgQm9iIn0K"),
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
