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
	ipSecConnectionTunnelErrorRequiredOnlyResource = IpsecConnectionTunnelErrorResourceConfig +
		GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentationCopy)

	ipsecConnectionTunnelErrorSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": Representation{RepType: Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	//ipSecConnectionRepresentationCopy = GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags"}, ipSecConnectionRepresentation)
	drgRepresentationCopy = GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags", "freeform_tags"}, drgRepresentation)
	cpeRepresentationCopy = GetRepresentationCopyWithMultipleRemovedProperties([]string{"defined_tags", "freeform_tags"}, cpeRepresentation)

	ipSecConnectionRepresentationCopy = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpe_id":         Representation{RepType: Required, Create: `${oci_core_cpe.test_cpe.id}`},
		"drg_id":         Representation{RepType: Required, Create: `${oci_core_drg.test_drg.id}`},
		"static_routes":  Representation{RepType: Required, Create: []string{`10.0.0.0/16`}, Update: []string{`10.1.0.0/16`}},
	}

	IpSecConnectionOptionalResourceCopy = GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Optional, Create, ipSecConnectionRepresentationCopy) +
		GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Required, Create, cpeRepresentationCopy) +
		GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentationCopy)

	IpsecConnectionTunnelErrorResourceConfig = IpSecConnectionOptionalResourceCopy
)

// issue-routing-tag: core/default
func TestCoreIpsecConnectionTunnelErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpsecConnectionTunnelErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel_error.test_ipsec_connection_tunnel_error"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + IpsecConnectionTunnelErrorResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", Required, Create, ipSecConnectionTunnelDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_error", "test_ipsec_connection_tunnel_error", Required, Create, ipsecConnectionTunnelErrorSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
