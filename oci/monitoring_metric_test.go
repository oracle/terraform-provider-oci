// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	metricDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`},
		"dimension_filters":         Representation{RepType: Optional, Create: map[string]string{"resourceId": "${oci_load_balancer_load_balancer.test_load_balancer.id}"}},
		"name":                      Representation{RepType: Optional, Create: `AcceptedConnections`},
		"namespace":                 Representation{RepType: Optional, Create: `oci_lbaas`},
		"resource_group":            Representation{RepType: Optional, Create: `resourceGroup`},
	}

	MetricResourceConfig = LoadBalancerResourceConfig
)

// issue-routing-tag: monitoring/default
func TestMonitoringMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_metrics.test_metrics"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_monitoring_metrics", "test_metrics", Optional, Create, metricDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_monitoring_metrics", "test_metrics_with_group_by", Required, Create, RepresentationCopyWithNewProperties(metricDataSourceRepresentation, map[string]interface{}{
					"group_by": Representation{RepType: Required, Create: []string{`namespace`}},
				})) +
				compartmentIdVariableStr + MetricResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
