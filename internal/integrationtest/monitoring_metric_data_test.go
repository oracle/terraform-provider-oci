// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	metricDataEndTimeStr   string
	metricDataStartTimeStr string

	MonitoringMonitoringMetricDataDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":                 acctest.Representation{RepType: acctest.Required, Create: `oci_vcn`},
		"query":                     acctest.Representation{RepType: acctest.Required, Create: `VnicToNetworkPackets[4m].max()`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"end_time":                  acctest.Representation{RepType: acctest.Optional, Create: metricDataEndTimeStr},
		"resolution":                acctest.Representation{RepType: acctest.Optional, Create: `2m`},
		"resource_group":            acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`},
		"start_time":                acctest.Representation{RepType: acctest.Optional, Create: metricDataStartTimeStr},
	}

	MonitoringMetricDataResourceConfig = AvailabilityDomainConfig
)

func generateMetricDataRepresentationWithCurrentTimeInputs() map[string]interface{} {
	endTime := time.Now()
	startTime := endTime.Add(-2 * time.Hour)
	metricDataEndTimeStr = endTime.Format(time.RFC3339Nano)
	metricDataStartTimeStr = startTime.Format(time.RFC3339Nano)
	return acctest.RepresentationCopyWithNewProperties(MonitoringMonitoringMetricDataDataSourceRepresentation, map[string]interface{}{
		"end_time":   acctest.Representation{RepType: acctest.Optional, Create: metricDataEndTimeStr},
		"start_time": acctest.Representation{RepType: acctest.Optional, Create: metricDataStartTimeStr},
	})
}

// issue-routing-tag: monitoring/default
func TestMonitoringMetricDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringMetricDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_metric_data.test_metric_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_metric_data", "test_metric_data", acctest.Optional, acctest.Create, generateMetricDataRepresentationWithCurrentTimeInputs()) +
				compartmentIdVariableStr + MonitoringMetricDataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
