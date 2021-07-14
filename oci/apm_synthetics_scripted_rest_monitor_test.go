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
	ScriptedRestMonitorResourceConfig = ScriptedRestMonitorResourceDependencies +
		generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Update, scriptedRestMonitorRepresentation)

	scriptedRestMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"monitor_id":    Representation{repType: Required, create: `${oci_apm_synthetics_monitor.test_monitor.id}`},
	}

	scriptedRestMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"monitor_type":  Representation{repType: Optional, create: `SCRIPTED_REST`},
		"script_id":     Representation{repType: Optional, create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":        Representation{repType: Optional, create: `ENABLED`, update: `DISABLED`},
		"filter":        RepresentationGroup{Required, scriptedRestMonitorDataSourceFilterRepresentationn}}
	scriptedRestMonitorDataSourceFilterRepresentationn = map[string]interface{}{
		"name":   Representation{repType: Required, create: `display_name`},
		"values": Representation{repType: Required, create: []string{`${oci_apm_synthetics_monitor.test_monitor.display_name}`}},
	}

	scriptedRestMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":              Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"monitor_type":               Representation{repType: Required, create: `SCRIPTED_REST`},
		"repeat_interval_in_seconds": Representation{repType: Required, create: `600`, update: `1200`},
		"vantage_points":             Representation{repType: Required, create: []string{`OraclePublic-us-ashburn-1`}},
		"defined_tags":               Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"script_id":                  Representation{repType: Optional, create: `${oci_apm_synthetics_script.test_script.id}`},
		"status":                     Representation{repType: Optional, create: `ENABLED`, update: `DISABLED`},
		"target":                     Representation{repType: Optional, create: `https://console.us-ashburn-1.oraclecloud.com`, update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":         Representation{repType: Optional, create: `60`, update: `120`},
	}
	ScriptedRestMonitorResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation) +
		generateResourceFromRepresentationMap("oci_apm_synthetics_script", "test_script", Required, Create, jsScriptRepresentation)
)

func TestApmSyntheticsScriptedRestMonitorResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsMonitorResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ScriptedRestMonitorResourceDependencies+
		generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Create, scriptedRestMonitorRepresentation), "apmsynthetics", "monitor", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmSyntheticsMonitorDestroy,
		Steps: []resource.TestStep{

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ScriptedRestMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Create, scriptedRestMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_REST_CONFIG"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_REST"),
					resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
					resource.TestCheckResourceAttrSet(resourceName, "script_id"),
					resource.TestCheckResourceAttrSet(resourceName, "script_name"),
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
				Config: config + compartmentIdVariableStr + ScriptedRestMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Update, scriptedRestMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "SCRIPTED_REST_CONFIG"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_type", "SCRIPTED_REST"),
					resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
					resource.TestCheckResourceAttrSet(resourceName, "script_id"),
					resource.TestCheckResourceAttrSet(resourceName, "script_name"),
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
					generateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", Optional, Update, scriptedRestMonitorDataSourceRepresentation) +
					compartmentIdVariableStr + ScriptedRestMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Optional, Update, scriptedRestMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "monitor_type", "SCRIPTED_REST"),
					resource.TestCheckResourceAttrSet(datasourceName, "script_id"),
					resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

					resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", Required, Create, scriptedRestMonitorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ScriptedRestMonitorResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "SCRIPTED_REST_CONFIG"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "SCRIPTED_REST"),
					resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "script_name"),
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
				Config: config + compartmentIdVariableStr + ScriptedRestMonitorResourceConfig,
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
