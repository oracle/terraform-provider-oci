// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	metricDataEndTimeStr   string
	metricDataStartTimeStr string

	metricDataDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"namespace":                 Representation{RepType: Required, Create: `oci_vcn`},
		"query":                     Representation{RepType: Required, Create: `VnicToNetworkPackets[4m].max()`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`},
		"end_time":                  Representation{RepType: Optional, Create: metricDataEndTimeStr},
		"resolution":                Representation{RepType: Optional, Create: `2m`},
		"resource_group":            Representation{RepType: Optional, Create: `resourceGroup`},
		"start_time":                Representation{RepType: Optional, Create: metricDataStartTimeStr},
	}

	MetricDataResourceConfig = AvailabilityDomainConfig
)

func generateMetricDataRepresentationWithCurrentTimeInputs() map[string]interface{} {
	endTime := time.Now()
	startTime := endTime.Add(-2 * time.Hour)
	metricDataEndTimeStr = endTime.Format(time.RFC3339Nano)
	metricDataStartTimeStr = startTime.Format(time.RFC3339Nano)
	return RepresentationCopyWithNewProperties(metricDataDataSourceRepresentation, map[string]interface{}{
		"end_time":   Representation{RepType: Optional, Create: metricDataEndTimeStr},
		"start_time": Representation{RepType: Optional, Create: metricDataStartTimeStr},
	})
}

// issue-routing-tag: monitoring/default
func TestMonitoringMetricDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringMetricDataResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_metric_data.test_metric_data"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_monitoring_metric_data", "test_metric_data", Optional, Create, generateMetricDataRepresentationWithCurrentTimeInputs()) +
				compartmentIdVariableStr + MetricDataResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "end_time", metricDataEndTimeStr),
				resource.TestCheckResourceAttr(datasourceName, "namespace", "oci_vcn"),
				resource.TestCheckResourceAttr(datasourceName, "query", "VnicToNetworkPackets[4m].max()"),
				resource.TestCheckResourceAttr(datasourceName, "resolution", "2m"),
				resource.TestCheckResourceAttr(datasourceName, "resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(datasourceName, "start_time", metricDataStartTimeStr),

				resource.TestCheckResourceAttrSet(datasourceName, "metric_data.#"),
			),
		},
	})
}
