// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ipSecConnectionTunnelRequiredOnlyResource = IpSecConnectionTunnelResourceConfig +
		GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation)

	ipSecConnectionTunnelSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": Representation{RepType: Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	ipSecConnectionTunnelDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	ipSecConnectionTunnelRepresentation = map[string]interface{}{
		"ipsec_id":         Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":        Representation{RepType: Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
		"routing":          Representation{RepType: Required, Create: `STATIC`, Update: `BGP`},
		"ike_version":      Representation{RepType: Optional, Create: `V1`, Update: `V2`},
		"display_name":     Representation{RepType: Optional, Create: `MyIPSecConnectionTunnel`, Update: `displayName2`},
		"bgp_session_info": RepresentationGroup{Optional, ipSecConnectionTunnelConfigurationBgpSessionInfoRepresentation},
		"shared_secret":    Representation{RepType: Optional, Create: `sharedsecret1`, Update: `sharedsecret2`},
	}

	ipSecConnectionTunnelConfigurationBgpSessionInfoRepresentation = map[string]interface{}{
		"customer_bgp_asn":      Representation{RepType: Optional, Update: `1587232876`},
		"customer_interface_ip": Representation{RepType: Optional, Update: `10.0.0.16/31`},
		"oracle_interface_ip":   Representation{RepType: Optional, Update: `10.0.0.17/31`},
	}

	IpSecConnectionTunnelResourceConfig = IpSecConnectionOptionalResource
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionTunnelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionTunnelResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipsec_connection_tunnel_management.test_ip_sec_connection_tunnel_management"
	datasourceName := "data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels"
	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel.test_ip_sec_connection_tunnel"

	var resId, resId2 string

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", Required, Create, ipSecConnectionTunnelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "shared_secret"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", Optional, Create, ipSecConnectionTunnelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
				resource.TestCheckResourceAttr(resourceName, "ike_version", "V1"),
				resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnectionTunnel"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", Optional, Update, ipSecConnectionTunnelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shared_secret"),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
				resource.TestCheckResourceAttr(resourceName, "ike_version", "V2"),
				resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.customer_bgp_asn", "1587232876"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.cpe_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.dpd_mode"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.dpd_timeout_in_sec"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.ike_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.nat_translation_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.oracle_can_initiate"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.routing"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.time_status_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.vpn_ip"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel", "test_ip_sec_connection_tunnel", Required, Create, ipSecConnectionTunnelSingularDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnel_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ike_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "routing"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_status_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vpn_ip"),
			),
		},
	})
}
