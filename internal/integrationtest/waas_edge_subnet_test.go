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
	WaasWaasEdgeSubnetDataSourceRepresentation = map[string]interface{}{}

	WaasEdgeSubnetResourceConfig = ""
)

// issue-routing-tag: waas/default
func TestWaasEdgeSubnetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasEdgeSubnetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_waas_edge_subnets.test_edge_subnets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_edge_subnets", "test_edge_subnets", acctest.Required, acctest.Create, WaasWaasEdgeSubnetDataSourceRepresentation) +
				compartmentIdVariableStr + WaasEdgeSubnetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.0.cidr"),
				resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.0.region"),
				resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.0.time_modified"),
			),
		},
	})
}
