// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ipSecConnectionTunnelSecurityAssociationRequiredOnlyResource = TunnelSecurityAssociationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, ipSecConnectionRepresentationCopy)

	tunnelSecurityAssociationDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	TunnelSecurityAssociationResourceConfig = IpSecConnectionOptionalResourceCopy
)

// issue-routing-tag: core/default
func TestCoreTunnelSecurityAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreTunnelSecurityAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_tunnel_security_associations.test_tunnel_security_associations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + TunnelSecurityAssociationResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_tunnel_security_associations", "test_tunnel_security_associations", acctest.Required, acctest.Create, tunnelSecurityAssociationDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
