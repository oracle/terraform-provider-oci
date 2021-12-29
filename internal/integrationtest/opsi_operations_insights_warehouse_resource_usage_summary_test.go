// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	operationsInsightsWarehouseResourceUsageSummarySingularDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
	}

	OperationsInsightsWarehouseResourceUsageSummaryResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, operationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseResourceUsageSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseResourceUsageSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_opsi_operations_insights_warehouse_resource_usage_summary.test_operations_insights_warehouse_resource_usage_summary"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_resource_usage_summary", "test_operations_insights_warehouse_resource_usage_summary", acctest.Required, acctest.Create, operationsInsightsWarehouseResourceUsageSummarySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperationsInsightsWarehouseResourceUsageSummaryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_warehouse_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_used"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_used_in_gbs"),
			),
		},
	})
}
