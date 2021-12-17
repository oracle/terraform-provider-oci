// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ipSecConnectionTunnelRequiredOnlyResource = IpSecConnectionTunnelResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, ipSecConnectionRepresentation)

	ipSecConnectionTunnelSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	ipSecConnectionTunnelDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	ipSecConnectionTunnelRepresentation = map[string]interface{}{
		"ipsec_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
		"routing":          acctest.Representation{RepType: acctest.Required, Create: `STATIC`, Update: `BGP`},
		"ike_version":      acctest.Representation{RepType: acctest.Optional, Create: `V1`, Update: `V2`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `MyIPSecConnectionTunnel`, Update: `displayName2`},
		"bgp_session_info": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationBgpSessionInfoRepresentation},
		"shared_secret":    acctest.Representation{RepType: acctest.Optional, Create: `sharedsecret1`, Update: `sharedsecret2`},
	}

	ipSecConnectionTunnelConfigurationBgpSessionInfoRepresentation = map[string]interface{}{
		"customer_bgp_asn":      acctest.Representation{RepType: acctest.Optional, Update: `1587232876`},
		"customer_interface_ip": acctest.Representation{RepType: acctest.Optional, Update: `10.0.0.16/31`},
		"oracle_interface_ip":   acctest.Representation{RepType: acctest.Optional, Update: `10.0.0.17/31`},
	}

	IpSecConnectionTunnelResourceConfig = IpSecConnectionOptionalResource
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionTunnelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionTunnelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipsec_connection_tunnel_management.test_ip_sec_connection_tunnel_management"
	datasourceName := "data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels"
	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel.test_ip_sec_connection_tunnel"

	var resId, resId2 string

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", acctest.Required, acctest.Create, ipSecConnectionTunnelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "shared_secret"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", acctest.Optional, acctest.Create, ipSecConnectionTunnelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
				resource.TestCheckResourceAttr(resourceName, "ike_version", "V1"),
				resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnectionTunnel"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", acctest.Optional, acctest.Update, ipSecConnectionTunnelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel", "test_ip_sec_connection_tunnel", acctest.Required, acctest.Create, ipSecConnectionTunnelSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
