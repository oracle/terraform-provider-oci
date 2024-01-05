// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ApmSyntheticsScriptResourceRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Required, acctest.Create, jsScriptRepresentation)

	JsScriptResourceConfig = ApmSyntheticsScriptResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Optional, acctest.Update, jsScriptRepresentation)

	jsScriptSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"script_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_script.test_script.id}`},
	}

	jsScriptDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"content_type":  acctest.Representation{RepType: acctest.Optional, Create: `JS`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: jsScriptDataSourceFilterRepresentation}}
	jsScriptDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `display_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_script.test_script.display_name}`}},
	}

	jsScriptContent        = "var request = require('postman-request'); var options = { 'method': 'GET', 'url': '<ORAP><ON>URL</ON><OV>https://console.us-phoenix-1.oraclecloud.com</OV></ORAP>', 'headers': { } }; request(options, function (error, response) { if (error) throw new Error(error); console.log(response.body); });"
	jsScriptRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"content":       acctest.Representation{RepType: acctest.Required, Create: jsScriptContent},
		"content_type":  acctest.Representation{RepType: acctest.Required, Create: `JS`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsScriptResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsScriptResource_js")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_script.test_script"
	datasourceName := "data.oci_apm_synthetics_scripts.test_scripts"
	singularDatasourceName := "data.oci_apm_synthetics_script.test_script"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsScriptResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Optional, acctest.Create, jsScriptRepresentation), "apmsynthetics", "script", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmSyntheticsScriptDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + ApmSyntheticsScriptResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Required, acctest.Create, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(resourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + ApmSyntheticsScriptResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + ApmSyntheticsScriptResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Optional, acctest.Create, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(resourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),

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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApmSyntheticsScriptResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Optional, acctest.Update, jsScriptRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(resourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_scripts", "test_scripts", acctest.Optional, acctest.Update, jsScriptDataSourceRepresentation) +
					compartmentIdVariableStr + ApmSyntheticsScriptResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Optional, acctest.Update, jsScriptRepresentation),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", acctest.Required, acctest.Create, jsScriptSingularDataSourceRepresentation) +
					compartmentIdVariableStr + JsScriptResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "script_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "content", jsScriptContent),
					resource.TestCheckResourceAttr(singularDatasourceName, "content_type", "JS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_uploaded"),
				),
			},
			// verify resource import
			{
				Config:            config + ApmSyntheticsScriptResourceRequiredOnlyResource,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}

}
