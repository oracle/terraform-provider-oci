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
	ipSecConnectionTunnelSecurityAssociationRequiredOnlyResource = TunnelSecurityAssociationResourceConfig +
		GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentationCopy)

	tunnelSecurityAssociationDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": Representation{RepType: Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	TunnelSecurityAssociationResourceConfig = IpSecConnectionOptionalResourceCopy
)

// issue-routing-tag: core/default
func TestCoreTunnelSecurityAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreTunnelSecurityAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_tunnel_security_associations.test_tunnel_security_associations"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + TunnelSecurityAssociationResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_tunnel_security_associations", "test_tunnel_security_associations", Required, Create, tunnelSecurityAssociationDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_security_associations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_security_associations.0.time"),
				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_security_associations.0.tunnel_sa_error_info"),
				resource.TestCheckResourceAttrSet(datasourceName, "tunnel_security_associations.0.tunnel_sa_status"),
			),
		},
	})
}
