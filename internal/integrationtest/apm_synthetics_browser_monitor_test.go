// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
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
	BrowserMonitorResourceConfig = MonitorResourceDependencies1 +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, browserMonitorRepresentation)

	browserMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"monitor_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_monitor.test_monitor.id}`},
	}

	browserMonitorDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"monitor_type":  acctest.Representation{RepType: acctest.Optional, Create: `BROWSER`},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: browserMonitorDataSourceFilterRepresentationn}}
	browserMonitorDataSourceFilterRepresentationn = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `display_name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_monitor.test_monitor.display_name}`}},
	}

	browserMonitorRepresentation = map[string]interface{}{
		"apm_domain_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"monitor_type":               acctest.Representation{RepType: acctest.Required, Create: `BROWSER`},
		"repeat_interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `600`, Update: `1200`},
		"vantage_points":             acctest.Representation{RepType: acctest.Required, Create: []string{`OraclePublic-us-ashburn-1`}},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                     acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"target":                     acctest.Representation{RepType: acctest.Optional, Create: `https://console.us-ashburn-1.oraclecloud.com`, Update: `https://console.us-phoenix-1.oraclecloud.com`},
		"timeout_in_seconds":         acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `120`},
	}
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsBrowserMonitorResource(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsMonitorResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_monitor.test_monitor"
	datasourceName := "data.oci_apm_synthetics_monitors.test_monitors"
	singularDatasourceName := "data.oci_apm_synthetics_monitor.test_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitorResourceDependencies1+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, browserMonitorRepresentation), "apmsynthetics", "monitor", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmSyntheticsMonitorDestroy,
		Steps: []resource.TestStep{

			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + MonitorResourceDependencies1 +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Create, browserMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "600"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-ashburn-1.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "60"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
					resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0"),

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
				Config: config + compartmentIdVariableStr + MonitorResourceDependencies1 +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, browserMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(resourceName, "repeat_interval_in_seconds", "1200"),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "target", "https://console.us-phoenix-1.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "120"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_point_count"),
					resource.TestCheckResourceAttr(resourceName, "vantage_points.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vantage_points.0"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitors", "test_monitors", acctest.Optional, acctest.Update, browserMonitorDataSourceRepresentation) +
					compartmentIdVariableStr + MonitorResourceDependencies1 +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Optional, acctest.Update, browserMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(datasourceName, "status", "DISABLED"),

					resource.TestCheckResourceAttr(datasourceName, "monitor_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "monitor_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_monitor", "test_monitor", acctest.Required, acctest.Create, browserMonitorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BrowserMonitorResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.0.config_type", "BROWSER_CONFIG"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "monitor_type", "BROWSER"),
					resource.TestCheckResourceAttr(singularDatasourceName, "repeat_interval_in_seconds", "1200"),
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
				Config: config + compartmentIdVariableStr + BrowserMonitorResourceConfig,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
}
