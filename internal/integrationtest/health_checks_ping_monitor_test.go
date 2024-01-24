// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_health_checks "github.com/oracle/oci-go-sdk/v65/healthchecks"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	HealthChecksPingMonitorRequiredOnlyResource = HealthChecksPingMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Required, acctest.Create, HealthChecksPingMonitorRepresentation)

	HealthChecksPingMonitorResourceConfig = HealthChecksPingMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Optional, acctest.Update, HealthChecksPingMonitorRepresentation)

	HealthChecksHealthChecksPingMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"monitor_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_health_checks_ping_monitor.test_ping_monitor.id}`},
	}

	HealthChecksHealthChecksPingMonitorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"home_region":    acctest.Representation{RepType: acctest.Optional, Create: `${var.region}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: HealthChecksPingMonitorDataSourceFilterRepresentation}}
	HealthChecksPingMonitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_health_checks_ping_monitor.test_ping_monitor.id}`}},
	}

	HealthChecksPingMonitorRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `30`},
		"protocol":            acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"targets":             acctest.Representation{RepType: acctest.Required, Create: []string{`www.oracle.com`}, Update: []string{`www.google.com`}},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"port":                acctest.Representation{RepType: acctest.Optional, Create: `80`},
		"timeout_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `30`},
		"vantage_point_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`goo-chs`}},
	}

	HealthChecksPingMonitorResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: health_checks/default
func TestHealthChecksPingMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksPingMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_health_checks_ping_monitor.test_ping_monitor"
	datasourceName := "data.oci_health_checks_ping_monitors.test_ping_monitors"
	singularDatasourceName := "data.oci_health_checks_ping_monitor.test_ping_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+HealthChecksPingMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Optional, acctest.Create, HealthChecksPingMonitorRepresentation), "healthchecks", "pingMonitor", t)

	acctest.ResourceTest(t, testAccCheckHealthChecksPingMonitorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HealthChecksPingMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Required, acctest.Create, HealthChecksPingMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + HealthChecksPingMonitorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + HealthChecksPingMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Optional, acctest.Create, HealthChecksPingMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + HealthChecksPingMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(HealthChecksPingMonitorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
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
			Config: config + compartmentIdVariableStr + HealthChecksPingMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Optional, acctest.Update, HealthChecksPingMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "port", "80"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_health_checks_ping_monitors", "test_ping_monitors", acctest.Optional, acctest.Update, HealthChecksHealthChecksPingMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + HealthChecksPingMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Optional, acctest.Update, HealthChecksPingMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "home_region"),

				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_monitors.0.home_region"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_monitors.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.0.interval_in_seconds", "30"),
				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "ping_monitors.0.protocol", "TCP"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_monitors.0.results_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_monitors.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Required, acctest.Create, HealthChecksHealthChecksPingMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + HealthChecksPingMonitorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "interval_in_seconds", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "80"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "results_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "targets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vantage_point_names.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + HealthChecksPingMonitorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckHealthChecksPingMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).HealthChecksClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_health_checks_ping_monitor" {
			noResourceFound = false
			request := oci_health_checks.GetPingMonitorRequest{}

			tmp := rs.Primary.ID
			request.MonitorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "health_checks")

			_, err := client.GetPingMonitor(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("HealthChecksPingMonitor") {
		resource.AddTestSweepers("HealthChecksPingMonitor", &resource.Sweeper{
			Name:         "HealthChecksPingMonitor",
			Dependencies: acctest.DependencyGraph["pingMonitor"],
			F:            sweepHealthChecksPingMonitorResource,
		})
	}
}

func sweepHealthChecksPingMonitorResource(compartment string) error {
	healthChecksClient := acctest.GetTestClients(&schema.ResourceData{}).HealthChecksClient()
	pingMonitorIds, err := getHealthChecksPingMonitorIds(compartment)
	if err != nil {
		return err
	}
	for _, pingMonitorId := range pingMonitorIds {
		if ok := acctest.SweeperDefaultResourceId[pingMonitorId]; !ok {
			deletePingMonitorRequest := oci_health_checks.DeletePingMonitorRequest{}

			deletePingMonitorRequest.MonitorId = &pingMonitorId

			deletePingMonitorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "health_checks")
			_, error := healthChecksClient.DeletePingMonitor(context.Background(), deletePingMonitorRequest)
			if error != nil {
				fmt.Printf("Error deleting PingMonitor %s %s, It is possible that the resource is already deleted. Please verify manually \n", pingMonitorId, error)
				continue
			}
		}
	}
	return nil
}

func getHealthChecksPingMonitorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PingMonitorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	healthChecksClient := acctest.GetTestClients(&schema.ResourceData{}).HealthChecksClient()

	listPingMonitorsRequest := oci_health_checks.ListPingMonitorsRequest{}
	listPingMonitorsRequest.CompartmentId = &compartmentId
	listPingMonitorsResponse, err := healthChecksClient.ListPingMonitors(context.Background(), listPingMonitorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PingMonitor list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pingMonitor := range listPingMonitorsResponse.Items {
		id := *pingMonitor.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PingMonitorId", id)
	}
	return resourceIds, nil
}
