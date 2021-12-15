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
	operationsInsightsWarehouseResourceUsageSummarySingularDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id": Representation{RepType: Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
	}

	OperationsInsightsWarehouseResourceUsageSummaryResourceConfig = GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseResourceUsageSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseResourceUsageSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_opsi_operations_insights_warehouse_resource_usage_summary.test_operations_insights_warehouse_resource_usage_summary"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_resource_usage_summary", "test_operations_insights_warehouse_resource_usage_summary", Required, Create, operationsInsightsWarehouseResourceUsageSummarySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperationsInsightsWarehouseResourceUsageSummaryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_warehouse_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_used"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_used_in_gbs"),
			),
		},
	})
}
