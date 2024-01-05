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
	CoreVirtualCircuitAssociatedTunnelDataSourceRepresentation = map[string]interface{}{
		"virtual_circuit_id": acctest.Representation{RepType: acctest.Required, Create: `${var.ipsec_over_fc_vc_id}`},
	}
)

// issue-routing-tag: core/default
func TestCoreVirtualCircuitAssociatedTunnelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitAssociatedTunnelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	virtualCircuitId := utils.GetEnvSettingWithBlankDefault("ipsec_over_fc_vc_id")
	virtualCircuitIdVariableStr := fmt.Sprintf("variable \"ipsec_over_fc_vc_id\" { default = \"%s\" }\n", virtualCircuitId)

	datasourceName := "data.oci_core_virtual_circuit_associated_tunnels.test_virtual_circuit_associated_tunnels"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuit_associated_tunnels", "test_virtual_circuit_associated_tunnels", acctest.Required, acctest.Create, CoreVirtualCircuitAssociatedTunnelDataSourceRepresentation) +
				virtualCircuitIdVariableStr,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_associated_tunnel_details.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_associated_tunnel_details.0.ipsec_connection_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_associated_tunnel_details.0.tunnel_id"),
			),
		},
	})
}
