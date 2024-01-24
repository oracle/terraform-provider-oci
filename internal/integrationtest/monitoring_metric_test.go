// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	MonitoringMonitoringMetricDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"dimension_filters":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"resourceId": "${oci_load_balancer_load_balancer.test_load_balancer.id}"}},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `AcceptedConnections`},
		"namespace":                 acctest.Representation{RepType: acctest.Optional, Create: `oci_lbaas`},
		"resource_group":            acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`},
	}

	MonitoringMetricResourceConfig = LoadBalancerResourceConfig
)

// issue-routing-tag: monitoring/default
func TestMonitoringMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_metrics.test_metrics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_metrics", "test_metrics", acctest.Optional, acctest.Create, MonitoringMonitoringMetricDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_metrics", "test_metrics_with_group_by", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(MonitoringMonitoringMetricDataSourceRepresentation, map[string]interface{}{
					"group_by": acctest.Representation{RepType: acctest.Required, Create: []string{`namespace`}},
				})) +
				compartmentIdVariableStr + MonitoringMetricResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "dimension_filters.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "AcceptedConnections"),
				resource.TestCheckResourceAttr(datasourceName, "namespace", "oci_lbaas"),
				resource.TestCheckResourceAttr(datasourceName, "resource_group", "resourceGroup"),

				resource.TestCheckResourceAttrSet(datasourceName, "metrics.#"),

				resource.TestCheckResourceAttr("data.oci_monitoring_metrics.test_metrics_with_group_by", "compartment_id", compartmentId),
				resource.TestCheckResourceAttr("data.oci_monitoring_metrics.test_metrics_with_group_by", "group_by.#", "1"),
				resource.TestCheckResourceAttrSet("data.oci_monitoring_metrics.test_metrics_with_group_by", "metrics.#"),
			),
		},
	})
}
