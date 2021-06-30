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
	ipSecConnectionTunnelRoutingPolicyRequiredOnlyResource = IpSecConnectionTunnelRoutingPolicyResourceConfig +
		generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation)

	ipSecConnectionTunnelRoutingPolicySingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": Representation{repType: Required, create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	ipSecConnectionTunnelRoutingPolicyDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	ipSecConnectionTunnelRoutingPolicyRepresentation = map[string]interface{}{
		"ipsec_id":                 Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":                Representation{repType: Required, create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
		"routing":                  Representation{repType: Required, create: `STATIC`, update: `POLICY`},
		"ike_version":              Representation{repType: Optional, create: `V1`, update: `V2`},
		"display_name":             Representation{repType: Optional, create: `MyIPSecConnectionTunnel`, update: `displayName2`},
		"encryption_domain_config": RepresentationGroup{Optional, ipSecConnectionTunnelEncryptionDomainConfigRepresentation},
		"shared_secret":            Representation{repType: Optional, create: `sharedsecret1`, update: `sharedsecret2`},
	}

	ipSecConnectionTunnelEncryptionDomainConfigRepresentation = map[string]interface{}{
		"cpe_traffic_selector":    Representation{repType: Optional, create: []string{`192.168.1.0/24`}, update: []string{`10.0.0.0/24`}},
		"oracle_traffic_selector": Representation{repType: Optional, create: []string{`10.0.0.0/24`}, update: []string{`192.168.1.0/24`}},
	}

	IpSecConnectionTunnelRoutingPolicyResourceConfig = IpSecConnectionOptionalResource
)

func TestDataSourceCoreIpSecConnectionTunnelResourceRoutingPolicy_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionTunnelResourceRoutingPolicy_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipsec_connection_tunnel_management.test_ip_sec_connection_tunnel_management"
	datasourceName := "data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels"
	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel.test_ip_sec_connection_tunnel"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionTunnelRoutingPolicyResourceConfig +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelRoutingPolicyDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", Optional, Create, ipSecConnectionTunnelRoutingPolicyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
					resource.TestCheckResourceAttr(resourceName, "ike_version", "V1"),
					resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnectionTunnel"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.cpe_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.oracle_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.cpe_traffic_selector.0", "192.168.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.oracle_traffic_selector.0", "10.0.0.0/24"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionTunnelRoutingPolicyResourceConfig +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelRoutingPolicyDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management", Optional, Update, ipSecConnectionTunnelRoutingPolicyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "shared_secret"),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
					resource.TestCheckResourceAttr(resourceName, "ike_version", "V2"),
					resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.cpe_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.oracle_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.cpe_traffic_selector.0", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "encryption_domain_config.0.oracle_traffic_selector.0", "192.168.1.0/24"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelRoutingPolicyDataSourceRepresentation) +
					compartmentIdVariableStr + IpSecConnectionTunnelRoutingPolicyResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.cpe_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.ike_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.routing"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.status"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.time_status_updated"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.vpn_ip"),
					resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config.0.cpe_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config.0.oracle_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config.0.cpe_traffic_selector.0", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config.0.oracle_traffic_selector.0", "192.168.1.0/24"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel", "test_ip_sec_connection_tunnel", Required, Create, ipSecConnectionTunnelRoutingPolicySingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelRoutingPolicyDataSourceRepresentation) +
					compartmentIdVariableStr + IpSecConnectionTunnelRoutingPolicyResourceConfig,
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
					resource.TestCheckResourceAttr(singularDatasourceName, "encryption_domain_config.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "encryption_domain_config.0.cpe_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "encryption_domain_config.0.oracle_traffic_selector.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "encryption_domain_config.0.cpe_traffic_selector.0", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(singularDatasourceName, "encryption_domain_config.0.oracle_traffic_selector.0", "192.168.1.0/24"),
				),
			},
		},
	})
}
