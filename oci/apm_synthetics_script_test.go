// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v46/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v46/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ScriptRequiredOnlyResource = ScriptResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, scriptRepresentation)

	ScriptResourceConfig = ScriptResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Update, scriptRepresentation)

	scriptSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"script_id":     Representation{repType: Required, create: `${oci_apm_synthetics_script.test_script.id}`},
	}

	scriptDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"content_type":  Representation{repType: Optional, create: `SIDE`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":        RepresentationGroup{Required, scriptDataSourceFilterRepresentation}}
	scriptDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `display_name`},
		"values": Representation{repType: Required, create: []string{`${oci_apm_synthetics_script.test_script.display_name}`}},
	}

	scriptRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"content":       Representation{repType: Required, create: `{ \"id\":\"f672ea8c-9508-483e-a123-878920eee73c\", \"version\":\"2.0\", \"name\":\"Sample Project\", \"url\":\"https://console.us-ashburn-1.oraclecloud.com\", \"tests\": [  { \"id\":\"b4522766-e382-40c2-ab01-452cf62e1cec\", \"name\":\"<ORAP><ON>testName</ON><OV>myTest</OV><OS>false</OS></ORAP>\", \"commands\":[ { \"id\":\"d1bc2093-bb61-4919-a554-38ef2653ac02\", \"comment\":\"comment\", \"command\":\"open\", \"target\":\"/\", \"targets\":[[\"css=td.bodytext\",\"css\"]], \"value\":\"xyz\"  } ] } ], \"suites\": [ { \"id\":\"a86b2934-7aa3-4838-b389-93c8aea2af05\",  \"name\":\"Default Suite\",  \"persistSession\":false, \"parallel\":false, \"timeout\":600,  \"tests\":  [  \"b4522766-e382-40c2-ab01-452cf62e1cec\" ] } ], \"urls\": [ \"https://console.us-ashburn-1.oraclecloud.com/\"  ], \"plugins\":[\"xxx\"] }`, update: `{ \"id\":\"f672ea8c-9508-483e-a123-878920eee73c\", \"version\":\"2.0\", \"name\":\"Sample Project 1\", \"url\":\"https://console.us-phoenix-1.oraclecloud.com\", \"tests\": [  { \"id\":\"b4522766-e382-40c2-ab01-452cf62e1cec\", \"name\":\"<ORAP><ON>testName</ON><OV>myTest1</OV><OS>false</OS></ORAP>\", \"commands\":[ { \"id\":\"d1bc2093-bb61-4919-a554-38ef2653ac02\", \"comment\":\"comment\", \"command\":\"open\", \"target\":\"/\", \"targets\":[[\"css=td.bodytext\",\"css\"]], \"value\":\"xyz\"  } ] } ], \"suites\": [ { \"id\":\"a86b2934-7aa3-4838-b389-93c8aea2af05\",  \"name\":\"Default Suite\",  \"persistSession\":false, \"parallel\":false, \"timeout\":600,  \"tests\":  [  \"b4522766-e382-40c2-ab01-452cf62e1cec\" ] } ], \"urls\": [ \"https://console.us-phoenix-1.oraclecloud.com//\"  ], \"plugins\":[\"xxx\"] }`},
		"content_type":  Representation{repType: Required, create: `SIDE`},
		"display_name":  Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"parameters":    RepresentationGroup{Optional, scriptParametersRepresentation},
	}

	scriptParametersRepresentation = map[string]interface{}{
		"param_name": Representation{repType: Required, create: `testName`},
		"is_secret":  Representation{repType: Optional, create: `false`, update: `true`},
	}

	ScriptResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsScriptResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsScriptResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_script.test_script"
	datasourceName := "data.oci_apm_synthetics_scripts.test_scripts"
	singularDatasourceName := "data.oci_apm_synthetics_script.test_script"
	scriptContent := "{ \"id\":\"f672ea8c-9508-483e-a123-878920eee73c\", \"version\":\"2.0\", \"name\":\"Sample Project\", \"url\":\"https://console.us-ashburn-1.oraclecloud.com\", \"tests\": [  { \"id\":\"b4522766-e382-40c2-ab01-452cf62e1cec\", \"name\":\"<ORAP><ON>testName</ON><OV>myTest</OV><OS>false</OS></ORAP>\", \"commands\":[ { \"id\":\"d1bc2093-bb61-4919-a554-38ef2653ac02\", \"comment\":\"comment\", \"command\":\"open\", \"target\":\"/\", \"targets\":[[\"css=td.bodytext\",\"css\"]], \"value\":\"xyz\"  } ] } ], \"suites\": [ { \"id\":\"a86b2934-7aa3-4838-b389-93c8aea2af05\",  \"name\":\"Default Suite\",  \"persistSession\":false, \"parallel\":false, \"timeout\":600,  \"tests\":  [  \"b4522766-e382-40c2-ab01-452cf62e1cec\" ] } ], \"urls\": [ \"https://console.us-ashburn-1.oraclecloud.com/\"  ], \"plugins\":[\"xxx\"] }"
	scriptContentUpdate := "{ \"id\":\"f672ea8c-9508-483e-a123-878920eee73c\", \"version\":\"2.0\", \"name\":\"Sample Project 1\", \"url\":\"https://console.us-phoenix-1.oraclecloud.com\", \"tests\": [  { \"id\":\"b4522766-e382-40c2-ab01-452cf62e1cec\", \"name\":\"<ORAP><ON>testName</ON><OV>myTest1</OV><OS>false</OS></ORAP>\", \"commands\":[ { \"id\":\"d1bc2093-bb61-4919-a554-38ef2653ac02\", \"comment\":\"comment\", \"command\":\"open\", \"target\":\"/\", \"targets\":[[\"css=td.bodytext\",\"css\"]], \"value\":\"xyz\"  } ] } ], \"suites\": [ { \"id\":\"a86b2934-7aa3-4838-b389-93c8aea2af05\",  \"name\":\"Default Suite\",  \"persistSession\":false, \"parallel\":false, \"timeout\":600,  \"tests\":  [  \"b4522766-e382-40c2-ab01-452cf62e1cec\" ] } ], \"urls\": [ \"https://console.us-phoenix-1.oraclecloud.com//\"  ], \"plugins\":[\"xxx\"] }"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ScriptResourceDependencies+
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Create, scriptRepresentation), "apmsynthetics", "script", t)

	ResourceTest(t, testAccCheckApmSyntheticsScriptDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ScriptResourceDependencies +
				generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, scriptRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "content", scriptContent),
				resource.TestCheckResourceAttr(resourceName, "content_type", "SIDE"),
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
				generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Create, scriptRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "content", scriptContent),
				resource.TestCheckResourceAttr(resourceName, "content_type", "SIDE"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "parameters.0.is_overwritten"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_secret", "false"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.param_name", "testName"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.param_value", "myTest"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.script_parameter.#", "1"),

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
				generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Update, scriptRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "content", scriptContentUpdate),
				resource.TestCheckResourceAttr(resourceName, "content_type", "SIDE"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_status_count_map.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "parameters.0.is_overwritten"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.is_secret", "true"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.param_name", "testName"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.script_parameter.#", "1"),

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
				generateDataSourceFromRepresentationMap("oci_apm_synthetics_scripts", "test_scripts", Optional, Update, scriptDataSourceRepresentation) +
				compartmentIdVariableStr + ScriptResourceDependencies +
				generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Update, scriptRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "content_type", "SIDE"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "script_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "script_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, scriptSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ScriptResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "content", scriptContentUpdate),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_type", "SIDE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parameters.0.is_overwritten"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.0.script_parameter.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_uploaded"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ScriptResourceConfig,
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
	})
}

func testAccCheckApmSyntheticsScriptDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).apmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_script" {
			noResourceFound = false
			request := oci_apm_synthetics.GetScriptRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.ScriptId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apm_synthetics")

			_, err := client.GetScript(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ApmSyntheticsScript") {
		resource.AddTestSweepers("ApmSyntheticsScript", &resource.Sweeper{
			Name:         "ApmSyntheticsScript",
			Dependencies: DependencyGraph["script"],
			F:            sweepApmSyntheticsScriptResource,
		})
	}
}

func sweepApmSyntheticsScriptResource(compartment string) error {
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()
	scriptIds, err := getScriptIds(compartment)
	if err != nil {
		return err
	}
	for _, scriptId := range scriptIds {
		if ok := SweeperDefaultResourceId[scriptId]; !ok {
			deleteScriptRequest := oci_apm_synthetics.DeleteScriptRequest{}

			deleteScriptRequest.ScriptId = &scriptId

			deleteScriptRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteScript(context.Background(), deleteScriptRequest)
			if error != nil {
				fmt.Printf("Error deleting Script %s %s, It is possible that the resource is already deleted. Please verify manually \n", scriptId, error)
				continue
			}
		}
	}
	return nil
}

func getScriptIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ScriptId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()

	listScriptsRequest := oci_apm_synthetics.ListScriptsRequest{}

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for Script resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listScriptsRequest.ApmDomainId = &apmDomainId

		listScriptsResponse, err := apmSyntheticClient.ListScripts(context.Background(), listScriptsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Script list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, script := range listScriptsResponse.Items {
			id := *script.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "ScriptId", id)
		}

	}
	return resourceIds, nil
}
