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
	emWarehouseEtlRunSingularDataSourceRepresentation = map[string]interface{}{
		"em_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_em_warehouse_em_warehouse.test_em_warehouse.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	emWarehouseEtlRunDataSourceRepresentation = map[string]interface{}{
		"em_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_em_warehouse_em_warehouse.test_em_warehouse.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	EmWarehouseEtlRunResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Required, acctest.Create, emWarehouseRepresentation)
)

// issue-routing-tag: em_warehouse/default
func TestEmWarehouseEmWarehouseEtlRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmWarehouseEmWarehouseEtlRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	operationsInsightWarehouseId := utils.GetEnvSettingWithBlankDefault("operations_insights_warehouse_id")
	operationsInsightWarehouseIdVariableStr := fmt.Sprintf("variable \"operations_insights_warehouse_id\" { default = \"%s\" }\n", operationsInsightWarehouseId)

	emBridgeId := utils.GetEnvSettingWithBlankDefault("em_bridge_id")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"em_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	datasourceName := "data.oci_em_warehouse_em_warehouse_etl_runs.test_em_warehouse_etl_runs"
	singularDatasourceName := "data.oci_em_warehouse_em_warehouse_etl_run.test_em_warehouse_etl_run"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_em_warehouse_em_warehouse_etl_runs", "test_em_warehouse_etl_runs", acctest.Required, acctest.Create, emWarehouseEtlRunDataSourceRepresentation) +
				compartmentIdVariableStr + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr + EmWarehouseEtlRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),*/
				resource.TestCheckResourceAttrSet(datasourceName, "em_warehouse_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "etl_run_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "etl_run_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_em_warehouse_em_warehouse_etl_run", "test_em_warehouse_etl_run", acctest.Required, acctest.Create, emWarehouseEtlRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr + EmWarehouseEtlRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),*/
				resource.TestCheckResourceAttrSet(singularDatasourceName, "em_warehouse_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
