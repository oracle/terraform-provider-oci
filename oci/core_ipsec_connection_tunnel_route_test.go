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
	ipSecConnectionTunnelRouteRequiredOnlyResource = IpsecConnectionTunnelRouteResourceConfig +
		GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation)

	ipSecConnectionTunnelRepresentationCopy = map[string]interface{}{
		"ipsec_id":         Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":        Representation{RepType: Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
		"routing":          Representation{RepType: Required, Create: `BGP`},
		"ike_version":      Representation{RepType: Optional, Create: `V1`},
		"display_name":     Representation{RepType: Optional, Create: `MyIPSecConnectionTunnel`},
		"bgp_session_info": RepresentationGroup{Optional, ipSecConnectionTunnelConfigurationBgpSessionInfoRepresentationCopy},
		"shared_secret":    Representation{RepType: Optional, Create: `sharedsecret1`, Update: `sharedsecret2`},
	}

	ipSecConnectionTunnelConfigurationBgpSessionInfoRepresentationCopy = map[string]interface{}{
		"customer_bgp_asn":      Representation{RepType: Optional, Create: `1587232876`},
		"customer_interface_ip": Representation{RepType: Optional, Create: `10.0.0.16/31`},
		"oracle_interface_ip":   Representation{RepType: Optional, Create: `10.0.0.17/31`},
	}

	IpsecConnectionTunnelRouteResourceConfig = IpSecConnectionOptionalResourceCopy
)

// issue-routing-tag: core/default
func TestCoreIpsecConnectionTunnelRouteResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpsecConnectionTunnelRouteResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_ipsec_connection_tunnel_routes.test_ipsec_connection_tunnel_routes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + IpsecConnectionTunnelRouteResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", Optional, Create, ipSecConnectionTunnelRepresentationCopy) +
				`data "oci_core_ipsec_connection_tunnel_routes" "test_ipsec_connection_tunnel_routes" {
					advertiser = "ORACLE"
					ipsec_id = "${oci_core_ipsec.test_ip_sec_connection.id}"
					tunnel_id = "${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}"
				}`,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "advertiser", "ORACLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_id"),

				// expected to be empty
				resource.TestCheckResourceAttr(datasourceName, "tunnel_routes.#", "0"),
			),
		},
	})
}
