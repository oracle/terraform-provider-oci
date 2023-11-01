// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

/*
*
	Dependency variables:
	    resource_id_for_baselineable_metric_evaluate = var.resource_id_for_baselineable_metric_evaluate
*/

var (
	StackMonitoringBaselineableMetricsEvaluateResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric_evaluate_baselineable_metric", acctest.Required, acctest.Create, StackMonitoringBaselineableMetricRepresentation)

	StackMonitoringBaselineableMetricsEvaluateSingularDataSourceRepresentation = map[string]interface{}{
		"baselineable_metric_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_baselineable_metric.test_baselineable_metric_evaluate_baselineable_metric.id}`},
		"items":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsSingularDataSourceRepresentation},
		"resource_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.resource_id}`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsSingularDataSourceRepresentation = map[string]interface{}{
		"evaluation_data_points": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsEvaluationDataPointsSingularDataSourceRepresentation},
		"training_data_points": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation0},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation1},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation2},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation3},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation4},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation5},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation6},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation7},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation8},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation9},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation10},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation11},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation12},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation13},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation14},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation15},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation16},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation17},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation18},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation19},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation20},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation21},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation22},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation23},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation24},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation25},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation26},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation27},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation28},
			{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation29},
		},
		"dimensions": acctest.Representation{RepType: acctest.Required, Create: map[string]string{"dimensions": "dimensions"}},
	}
	StackMonitoringBaselineableMetricsEvaluateItemsEvaluationDataPointsSingularDataSourceRepresentation = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-15T05:00:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation0 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:00:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation1 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:01:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation2 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:02:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation3 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:03:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation4 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:04:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation5 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:05:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation6 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:06:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation7 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:07:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation8 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:08:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation9 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:09:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation10 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:10:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation11 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:11:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation12 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:12:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation13 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:13:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation14 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:14:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation15 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:15:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation16 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:16:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation17 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:17:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation18 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:18:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation19 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:19:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation20 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:20:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation21 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:21:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation22 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:22:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation23 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:23:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation24 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:24:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation25 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:25:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation26 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:26:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation27 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:27:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation28 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:28:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}

	StackMonitoringBaselineableMetricsEvaluateItemsTrainingDataPointsSingularDataSourceRepresentation29 = map[string]interface{}{
		"timestamp": acctest.Representation{RepType: acctest.Required, Create: `2023-05-14T05:29:00.001Z`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `1.1`},
	}
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringBaselineableMetricsEvaluateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringBaselineableMetricsEvaluateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceId := utils.GetEnvSettingWithBlankDefault("stack_mon_baselineable_metric_evaluate_resource_id")
	resourceIdVariableStr := fmt.Sprintf("variable \"resource_id\" { default = \"%s\" }\n", resourceId)
	if resourceIdVariableStr == "" {
		t.Skip("Setting environmental variable resource_id_for_baselineable_metric_evaluate that represents ocid of the resource with enterprise edition is pre-requisite for this test")
	}

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_stack_monitoring_baselineable_metrics_evaluate.test_baselineable_metrics_evaluate"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_baselineable_metrics_evaluate", "test_baselineable_metrics_evaluate", acctest.Required, acctest.Create, StackMonitoringBaselineableMetricsEvaluateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + resourceIdVariableStr + StackMonitoringBaselineableMetricsEvaluateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "baselineable_metric_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.evaluation_data_points.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.training_data_points.#", "30"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_id"),
			),
		},
	})
}
