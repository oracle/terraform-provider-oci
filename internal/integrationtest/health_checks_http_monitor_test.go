// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_health_checks "github.com/oracle/oci-go-sdk/v56/healthchecks"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	HttpMonitorRequiredOnlyResource = HttpMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Required, acctest.Create, httpMonitorRepresentation)

	HttpMonitorResourceConfig = HttpMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Optional, acctest.Update, httpMonitorRepresentation)

	httpMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"monitor_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
	}

	httpMonitorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"home_region":    acctest.Representation{RepType: acctest.Optional, Create: `${var.region}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: httpMonitorDataSourceFilterRepresentation}}
	httpMonitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_health_checks_http_monitor.test_http_monitor.id}`}},
	}

	httpMonitorRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `30`},
		"protocol":            acctest.Representation{RepType: acctest.Required, Create: `HTTP`, Update: `HTTPS`},
		"targets":             acctest.Representation{RepType: acctest.Required, Create: []string{`www.oracle.com`}, Update: []string{`www.google.com`}},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"headers":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"headers": "headers"}, Update: map[string]string{"headers2": "headers2"}},
		"is_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"method":              acctest.Representation{RepType: acctest.Optional, Create: `GET`},
		"path":                acctest.Representation{RepType: acctest.Optional, Create: `/`},
		"port":                acctest.Representation{RepType: acctest.Optional, Create: `80`, Update: `443`},
		"timeout_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `30`},
		"vantage_point_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`goo-chs`}},
	}

	HttpMonitorResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: health_checks/default
func TestHealthChecksHttpMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksHttpMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_health_checks_http_monitor.test_http_monitor"
	datasourceName := "data.oci_health_checks_http_monitors.test_http_monitors"
	singularDatasourceName := "data.oci_health_checks_http_monitor.test_http_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+HttpMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Optional, acctest.Create, httpMonitorRepresentation), "healthchecks", "httpMonitor", t)

	acctest.ResourceTest(t, testAccCheckHealthChecksHttpMonitorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Required, acctest.Create, httpMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Optional, acctest.Create, httpMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "path", "/"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + HttpMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(httpMonitorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "path", "/"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Optional, acctest.Update, httpMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "headers.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "path", "/"),
				resource.TestCheckResourceAttr(resourceName, "port", "443"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "HTTPS"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "vantage_point_names.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_health_checks_http_monitors", "test_http_monitors", acctest.Optional, acctest.Update, httpMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + HttpMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Optional, acctest.Update, httpMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "home_region"),

				resource.TestCheckResourceAttr(datasourceName, "http_monitors.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_monitors.0.home_region"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_monitors.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.interval_in_seconds", "30"),
				resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_monitors.0.results_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_monitors.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", acctest.Required, acctest.Create, httpMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + HttpMonitorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "headers.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "interval_in_seconds", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "method", "GET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "path", "/"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "443"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "HTTPS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "results_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_point_names.#", "1"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + HttpMonitorResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckHealthChecksHttpMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).HealthChecksClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_health_checks_http_monitor" {
			noResourceFound = false
			request := oci_health_checks.GetHttpMonitorRequest{}

			tmp := rs.Primary.ID
			request.MonitorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "health_checks")

			_, err := client.GetHttpMonitor(context.Background(), request)

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("HealthChecksHttpMonitor") {
		resource.AddTestSweepers("HealthChecksHttpMonitor", &resource.Sweeper{
			Name:         "HealthChecksHttpMonitor",
			Dependencies: acctest.DependencyGraph["httpMonitor"],
			F:            sweepHealthChecksHttpMonitorResource,
		})
	}
}

func sweepHealthChecksHttpMonitorResource(compartment string) error {
	healthChecksClient := acctest.GetTestClients(&schema.ResourceData{}).HealthChecksClient()
	httpMonitorIds, err := getHttpMonitorIds(compartment)
	if err != nil {
		return err
	}
	for _, httpMonitorId := range httpMonitorIds {
		if ok := acctest.SweeperDefaultResourceId[httpMonitorId]; !ok {
			deleteHttpMonitorRequest := oci_health_checks.DeleteHttpMonitorRequest{}

			deleteHttpMonitorRequest.MonitorId = &httpMonitorId

			deleteHttpMonitorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "health_checks")
			_, error := healthChecksClient.DeleteHttpMonitor(context.Background(), deleteHttpMonitorRequest)
			if error != nil {
				fmt.Printf("Error deleting HttpMonitor %s %s, It is possible that the resource is already deleted. Please verify manually \n", httpMonitorId, error)
				continue
			}
		}
	}
	return nil
}

func getHttpMonitorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HttpMonitorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	healthChecksClient := acctest.GetTestClients(&schema.ResourceData{}).HealthChecksClient()

	listHttpMonitorsRequest := oci_health_checks.ListHttpMonitorsRequest{}
	listHttpMonitorsRequest.CompartmentId = &compartmentId
	listHttpMonitorsResponse, err := healthChecksClient.ListHttpMonitors(context.Background(), listHttpMonitorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HttpMonitor list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, httpMonitor := range listHttpMonitorsResponse.Items {
		id := *httpMonitor.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HttpMonitorId", id)
	}
	return resourceIds, nil
}
