// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	CoreIpsecConnectionTunnelErrorRequiredOnlyResource = CoreIpsecConnectionTunnelErrorResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, CoreIpSecConnectionRepresentationCopy)

	CoreCoreIpsecConnectionTunnelErrorSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	//CoreIpSecConnectionRepresentationCopy = GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags"}, CoreIpSecConnectionRepresentation)
	drgRepresentationCopy = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags", "freeform_tags"}, CoreDrgRepresentation)
	cpeRepresentationCopy = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags", "freeform_tags"}, CoreCpeRepresentation)

	CoreIpSecConnectionRepresentationCopy = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpe_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"static_routes":  acctest.Representation{RepType: acctest.Required, Create: []string{`10.0.0.0/16`}, Update: []string{`10.1.0.0/16`}},
	}

	CoreIpSecConnectionOptionalResourceCopy = acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Create, CoreIpSecConnectionRepresentationCopy) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Required, acctest.Create, cpeRepresentationCopy) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentationCopy)

	CoreIpsecConnectionTunnelErrorResourceConfig = CoreIpSecConnectionOptionalResourceCopy
)

// issue-routing-tag: core/default
func TestCoreIpsecConnectionTunnelErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpsecConnectionTunnelErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel_error.test_ipsec_connection_tunnel_error"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + CoreIpsecConnectionTunnelErrorResourceConfig,
			Check: func(s *terraform.State) (err error) {
				log.Printf("Wait for ipsec tunnel to provision, for status to recognize down, and tunnel error to be set")
				time.Sleep(30 * time.Second)
				return nil
			},
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + CoreIpsecConnectionTunnelErrorResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_error", "test_ipsec_connection_tunnel_error", acctest.Required, acctest.Create, CoreCoreIpsecConnectionTunnelErrorSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnel_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "error_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oci_resources_link"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "solution"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "timestamp"),
			),
		},
	})
}
