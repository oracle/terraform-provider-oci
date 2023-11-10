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

var (
	CoreIpSecConnectionTunnelRouteRequiredOnlyResource = CoreIpsecConnectionTunnelRouteResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, CoreIpSecConnectionRepresentation)

	CoreIpSecConnectionTunnelRepresentationCopy = map[string]interface{}{
		"ipsec_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
		"routing":          acctest.Representation{RepType: acctest.Required, Create: `BGP`},
		"ike_version":      acctest.Representation{RepType: acctest.Optional, Create: `V1`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `MyIPSecConnectionTunnel`},
		"bgp_session_info": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreIpSecConnectionTunnelConfigurationBgpSessionInfoRepresentationCopy},
		"shared_secret":    acctest.Representation{RepType: acctest.Optional, Create: `sharedsecret1`, Update: `sharedsecret2`},
	}

	CoreIpSecConnectionTunnelConfigurationBgpSessionInfoRepresentationCopy = map[string]interface{}{
		"customer_bgp_asn":      acctest.Representation{RepType: acctest.Optional, Create: `1587232876`},
		"customer_interface_ip": acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.16/31`},
		"oracle_interface_ip":   acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.17/31`},
	}

	CoreIpsecConnectionTunnelRouteResourceConfig = CoreIpSecConnectionOptionalResourceCopy
)

// issue-routing-tag: core/default
func TestCoreIpsecConnectionTunnelRouteResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpsecConnectionTunnelRouteResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_ipsec_connection_tunnel_routes.test_ipsec_connection_tunnel_routes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + CoreIpsecConnectionTunnelRouteResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", acctest.Optional, acctest.Create, CoreIpSecConnectionTunnelRepresentationCopy) +
				`data "oci_core_ipsec_connection_tunnel_routes" "test_ipsec_connection_tunnel_routes" {
					advertiser = "ORACLE"
					ipsec_id = "${oci_core_ipsec.test_ip_sec_connection.id}"
					tunnel_id = "${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}"
				}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "advertiser", "ORACLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_id"),

				// expected to be empty
				resource.TestCheckResourceAttr(datasourceName, "tunnel_routes.#", "0"),
			),
		},
	})
}
