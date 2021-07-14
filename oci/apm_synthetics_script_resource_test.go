// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	JsScriptResourceConfig = ScriptResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Update, jsScriptRepresentation)

	jsScriptSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"script_id":     Representation{repType: Required, create: `${oci_apm_synthetics_script.test_script.id}`},
	}

	jsScriptDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"content_type":  Representation{repType: Optional, create: `JS`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":        RepresentationGroup{Required, jsScriptDataSourceFilterRepresentation}}
	jsScriptDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `display_name`},
		"values": Representation{repType: Required, create: []string{`${oci_apm_synthetics_script.test_script.display_name}`}},
	}

	jsScriptContent        = "var request = require('postman-request'); var options = { 'method': 'GET', 'url': '<ORAP><ON>URL</ON><OV>https://console.us-phoenix-1.oraclecloud.com</OV></ORAP>', 'headers': { } }; request(options, function (error, response) { if (error) throw new Error(error); console.log(response.body); });"
	jsScriptRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"content":       Representation{repType: Required, create: jsScriptContent},
		"content_type":  Representation{repType: Required, create: `JS`},
		"display_name":  Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}
)

func TestApmSyntheticsScriptResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsScriptResource_js")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_script.test_script"
	datasourceName := "data.oci_apm_synthetics_scripts.test_scripts"
	singularDatasourceName := "data.oci_apm_synthetics_script.test_script"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ScriptResourceDependencies+
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Create, jsScriptRepresentation), "apmsynthetics", "script", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmSyntheticsScriptDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ScriptResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(resourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ScriptResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ScriptResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Create, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(resourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ScriptResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Update, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(resourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),

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
					generateDataSourceFromRepresentationMap("oci_apm_synthetics_scripts", "test_scripts", Optional, Update, jsScriptDataSourceRepresentation) +
					compartmentIdVariableStr + ScriptResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Update, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "script_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "script_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, jsScriptSingularDataSourceRepresentation) +
					compartmentIdVariableStr + JsScriptResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "script_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(singularDatasourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_uploaded"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + JsScriptResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apm_domain_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}

}
