// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ipSecConnectionTunnelSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": Representation{repType: Required, create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	ipSecConnectionTunnelDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	IpSecConnectionTunnelResourceConfig = IpSecConnectionOptionalResource
)

func TestCoreIpSecConnectionTunnelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionTunnelResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels"
	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel.test_ip_sec_connection_tunnel"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
					compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.cpe_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.id"),
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
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel", "test_ip_sec_connection_tunnel", Required, Create, ipSecConnectionTunnelSingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
					compartmentIdVariableStr + IpSecConnectionTunnelResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnel_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_ip"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "routing"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_status_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vpn_ip"),
				),
			},
		},
	})
}
