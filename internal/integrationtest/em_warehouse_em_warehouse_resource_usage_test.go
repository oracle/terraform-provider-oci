// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

var (
	emWarehouseResourceUsageSingularDataSourceRepresentation = map[string]interface{}{
		"em_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_em_warehouse_em_warehouse.test_em_warehouse.id}`},
	}

	EmWarehouseResourceUsageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Required, acctest.Create, emWarehouseRepresentation)
)

// issue-routing-tag: em_warehouse/default
func TestEmWarehouseEmWarehouseResourceUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmWarehouseEmWarehouseResourceUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	operationsInsightWarehouseId := utils.GetEnvSettingWithBlankDefault("operations_insights_warehouse_id")
	operationsInsightWarehouseIdVariableStr := fmt.Sprintf("variable \"operations_insights_warehouse_id\" { default = \"%s\" }\n", operationsInsightWarehouseId)

	emBridgeId := utils.GetEnvSettingWithBlankDefault("em_bridge_id")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"em_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	singularDatasourceName := "data.oci_em_warehouse_em_warehouse_resource_usage.test_em_warehouse_resource_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_em_warehouse_em_warehouse_resource_usage", "test_em_warehouse_resource_usage", acctest.Required, acctest.Create, emWarehouseResourceUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr + EmWarehouseResourceUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "em_warehouse_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "em_instance_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "em_instances.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "targets_count"),
			),
		},
	})
}
