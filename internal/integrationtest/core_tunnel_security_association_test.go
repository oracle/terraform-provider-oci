// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ipSecConnectionTunnelSecurityAssociationRequiredOnlyResource = CoreTunnelSecurityAssociationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, CoreIpSecConnectionRepresentationCopy)

	CoreCoreTunnelSecurityAssociationDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	CoreTunnelSecurityAssociationResourceConfig = CoreIpSecConnectionOptionalResourceCopy
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
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + CoreTunnelSecurityAssociationResourceConfig,
			Check: func(s *terraform.State) (err error) {
				log.Printf("Wait for ipsec tunnel to create security associations")
				time.Sleep(1 * time.Minute)
				return nil
			},
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + CoreTunnelSecurityAssociationResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_tunnel_security_associations", "test_tunnel_security_associations", acctest.Required, acctest.Create, CoreCoreTunnelSecurityAssociationDataSourceRepresentation),
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
