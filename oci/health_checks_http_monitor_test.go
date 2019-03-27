// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	HttpMonitorRequiredOnlyResource = HttpMonitorResourceDependencies +
		generateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Required, Create, httpMonitorRepresentation)

	HttpMonitorResourceConfig = HttpMonitorResourceDependencies +
		generateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Optional, Update, httpMonitorRepresentation)

	httpMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"monitor_id": Representation{repType: Required, create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
	}

	httpMonitorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":         RepresentationGroup{Required, httpMonitorDataSourceFilterRepresentation}}
	httpMonitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_health_checks_http_monitor.test_http_monitor.id}`}},
	}

	httpMonitorRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"interval_in_seconds": Representation{repType: Required, create: `10`, update: `30`},
		"protocol":            Representation{repType: Required, create: `HTTP`, update: `HTTPS`},
		"targets":             Representation{repType: Required, create: []string{`www.oracle.com`}, update: []string{`www.google.com`}},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"headers":             Representation{repType: Optional, create: map[string]string{"headers": "headers"}, update: map[string]string{"headers2": "headers2"}},
		"is_enabled":          Representation{repType: Optional, create: `false`, update: `true`},
		"method":              Representation{repType: Optional, create: `GET`},
		"path":                Representation{repType: Optional, create: `/`},
		"port":                Representation{repType: Optional, create: `80`, update: `443`},
		"timeout_in_seconds":  Representation{repType: Optional, create: `10`, update: `30`},
		"vantage_point_names": Representation{repType: Optional, create: []string{`goo-chs`}},
	}

	HttpMonitorResourceDependencies = DefinedTagsDependencies
)

func TestHealthChecksHttpMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksHttpMonitorResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_health_checks_http_monitor.test_http_monitor"
	datasourceName := "data.oci_health_checks_http_monitors.test_http_monitors"
	singularDatasourceName := "data.oci_health_checks_http_monitor.test_http_monitor"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckHealthChecksHttpMonitorDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Required, Create, httpMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "interval_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "targets.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Optional, Create, httpMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + HttpMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Optional, Update, httpMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_health_checks_http_monitors", "test_http_monitors", Optional, Update, httpMonitorDataSourceRepresentation) +
					compartmentIdVariableStr + HttpMonitorResourceDependencies +
					generateResourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Optional, Update, httpMonitorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "http_monitors.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.interval_in_seconds", "30"),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "http_monitors.0.protocol", "HTTPS"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_health_checks_http_monitor", "test_http_monitor", Required, Create, httpMonitorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + HttpMonitorResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "monitor_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "headers.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "interval_in_seconds", "30"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "method", "GET"),
					resource.TestCheckResourceAttr(singularDatasourceName, "path", "/"),
					resource.TestCheckResourceAttr(singularDatasourceName, "port", "443"),
					resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "HTTPS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "targets.#", "1"),
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
		},
	})
}

func testAccCheckHealthChecksHttpMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).healthChecksClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_health_checks_http_monitor" {
			noResourceFound = false
			request := oci_health_checks.GetHttpMonitorRequest{}

			tmp := rs.Primary.ID
			request.MonitorId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "health_checks")

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
