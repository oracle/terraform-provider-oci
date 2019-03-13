// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	metricDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `false`},
		"dimension_filters":         Representation{repType: Optional, create: map[string]string{"resourceId": "${oci_load_balancer_load_balancer.test_load_balancer.id}"}},
		"name":                      Representation{repType: Optional, create: `AcceptedConnections`},
		"namespace":                 Representation{repType: Optional, create: `oci_lbaas`},
	}

	MetricResourceConfig = LoadBalancerResourceConfig
)

func TestMonitoringMetricResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_metrics.test_metrics"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_monitoring_metrics", "test_metrics", Optional, Create, metricDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_monitoring_metrics", "test_metrics_with_group_by", Required, Create, representationCopyWithNewProperties(metricDataSourceRepresentation, map[string]interface{}{
						"group_by": Representation{repType: Required, create: []string{`namespace`}},
					})) +
					compartmentIdVariableStr + MetricResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
					resource.TestCheckResourceAttr(datasourceName, "dimension_filters.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "name", "AcceptedConnections"),
					resource.TestCheckResourceAttr(datasourceName, "namespace", "oci_lbaas"),

					resource.TestCheckResourceAttrSet(datasourceName, "metrics.#"),

					resource.TestCheckResourceAttr("data.oci_monitoring_metrics.test_metrics_with_group_by", "compartment_id", compartmentId),
					resource.TestCheckResourceAttr("data.oci_monitoring_metrics.test_metrics_with_group_by", "group_by.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_monitoring_metrics.test_metrics_with_group_by", "metrics.#"),
				),
			},
		},
	})
}
