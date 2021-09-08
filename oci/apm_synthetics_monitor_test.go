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
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v47/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v47/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MonitorRequiredOnlyResource = MonitorResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Required, Create, monitorRepresentation)

	MonitorResourceConfig = MonitorResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Update, monitorRepresentation)

	monitorSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"monitor_id":    Representation{repType: Required, create: `${oci_apm_synthetics_monitor.test_monitor.id}`},
	}

	monitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"monitor_type":  Representation{repType: Optional, create: `SCRIPTED_BROWSER`},
		"script_id":     Representation{repType: Optional, create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":        Representation{repType: Optional, create: `ENABLED`, update: `DISABLED`},
		"filter":        RepresentationGroup{Required, monitorDataSourceFilterRepresentation}}
	monitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `display_name`},
		"values": Representation{repType: Required, create: []string{`${oci_apm_synthetics_monitor.test_monitor.display_name}`}},
	}

	monitorRepresentation = map[string]interface{}{
		"apm_domain_id":              Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"monitor_type":               Representation{repType: Required, create: `SCRIPTED_BROWSER`},
		"repeat_interval_in_seconds": Representation{repType: Required, create: `600`, update: `1200`},
		"vantage_points":             Representation{repType: Required, create: []string{`OraclePublic-us-ashburn-1`}},
		"defined_tags":               Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"script_id":                  Representation{repType: Optional, create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":                     Representation{repType: Optional, create: `ENABLED`, update: `DISABLED`},
		"target":                     Representation{repType: Optional, create: `https://console.us-ashburn-1.oraclecloud.com`, update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":         Representation{repType: Optional, create: `60`, update: `120`},
		"configuration":              RepresentationGroup{Optional, monitorConfigurationRepresentation},
		"script_parameters":          RepresentationGroup{Optional, monitorScriptParametersRepresentation},
	}

	monitorConfigurationRepresentation = map[string]interface{}{
		"config_type":                       Representation{repType: Optional, create: `SCRIPTED_BROWSER_CONFIG`},
		"is_certificate_validation_enabled": Representation{repType: Optional, create: `false`, update: `true`},
		"is_failure_retried":                Representation{repType: Optional, create: `false`, update: `true`},
	}
	monitorScriptParametersRepresentation = map[string]interface{}{
		"param_name":  Representation{repType: Required, create: `testName`, update: `testName`},
		"param_value": Representation{repType: Required, create: `myTest`, update: `myTest1`},
	}

	MonitorResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation) +
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Optional, Create, scriptRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+MonitorResourceDependencies+
		generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Create, monitorRepresentation), "apmsynthetics", "monitor", t)

	ResourceTest(t, testAccCheckApmSyntheticsMonitorDestroy, []resource.TestStep{

		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies +
				generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Create, monitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "false"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
				resource.TestCheckResourceAttrSet(resourceName, "script_id"),
				resource.TestCheckResourceAttrSet(resourceName, "script_name"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_name", "testName"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_value", "myTest"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-ashburn-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0"),

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
			Config: config + compartmentIdVariableStr + MonitorResourceDependencies +
				generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Update, monitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "configuration.0.is_failure_retried", "true"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttrSet(resourceName, "script_id"),
				resource.TestCheckResourceAttrSet(resourceName, "script_name"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(resourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_name", "testName"),
				resource.TestCheckResourceAttr(resourceName, "script_parameters.0.param_value", "myTest1"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0"),

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
				generateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", Optional, Update, monitorDataSourceRepresentation) +
				compartmentIdVariableStr + MonitorResourceDependencies +
				generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Update, monitorRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttrSet(datasourceName, "script_id"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Required, Create, monitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitorResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "SCRIPTED_BROWSER_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.is_certificate_validation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "SCRIPTED_BROWSER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "script_parameters.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_parameters.0.is_overwritten"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script_parameters.0.is_secret"),
				resource.TestCheckResourceAttr(singularDatasourceName, "script_parameters.0.monitor_script_parameter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "120"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_point_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_points.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vantage_points.0"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + MonitorResourceConfig,
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

func testAccCheckApmSyntheticsMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).apmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_monitor" {
			noResourceFound = false
			request := oci_apm_synthetics.GetMonitorRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.MonitorId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apm_synthetics")

			_, err := client.GetMonitor(context.Background(), request)

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
	if !inSweeperExcludeList("ApmSyntheticsMonitor") {
		resource.AddTestSweepers("ApmSyntheticsMonitor", &resource.Sweeper{
			Name:         "ApmSyntheticsMonitor",
			Dependencies: DependencyGraph["monitor"],
			F:            sweepApmSyntheticsMonitorResource,
		})
	}
}

func sweepApmSyntheticsMonitorResource(compartment string) error {
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()
	monitorIds, err := getMonitorIds(compartment)
	if err != nil {
		return err
	}
	for _, monitorId := range monitorIds {
		if ok := SweeperDefaultResourceId[monitorId]; !ok {
			deleteMonitorRequest := oci_apm_synthetics.DeleteMonitorRequest{}

			deleteMonitorRequest.MonitorId = &monitorId

			deleteMonitorRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteMonitor(context.Background(), deleteMonitorRequest)
			if error != nil {
				fmt.Printf("Error deleting Monitor %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitorId, error)
				continue
			}
		}
	}
	return nil
}

func getMonitorIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "MonitorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := GetTestClients(&schema.ResourceData{}).apmSyntheticClient()

	listMonitorsRequest := oci_apm_synthetics.ListMonitorsRequest{}

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for Monitor resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listMonitorsRequest.ApmDomainId = &apmDomainId

		listMonitorsResponse, err := apmSyntheticClient.ListMonitors(context.Background(), listMonitorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Monitor list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, monitor := range listMonitorsResponse.Items {
			id := *monitor.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "MonitorId", id)
		}

	}
	return resourceIds, nil
}
