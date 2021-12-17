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
	ipSecConnectionTunnelErrorRequiredOnlyResource = IpsecConnectionTunnelErrorResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, ipSecConnectionRepresentationCopy)

	ipsecConnectionTunnelErrorSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	//ipSecConnectionRepresentationCopy = GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags"}, ipSecConnectionRepresentation)
	drgRepresentationCopy = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags", "freeform_tags"}, drgRepresentation)
	cpeRepresentationCopy = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags", "freeform_tags"}, cpeRepresentation)

	ipSecConnectionRepresentationCopy = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpe_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"static_routes":  acctest.Representation{RepType: acctest.Required, Create: []string{`10.0.0.0/16`}, Update: []string{`10.1.0.0/16`}},
	}

	IpSecConnectionOptionalResourceCopy = acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Optional, acctest.Create, ipSecConnectionRepresentationCopy) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Required, acctest.Create, cpeRepresentationCopy) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentationCopy)

	IpsecConnectionTunnelErrorResourceConfig = IpSecConnectionOptionalResourceCopy
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
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + IpsecConnectionTunnelErrorResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, ipSecConnectionTunnelDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_error", "test_ipsec_connection_tunnel_error", acctest.Required, acctest.Create, ipsecConnectionTunnelErrorSingularDataSourceRepresentation),
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
