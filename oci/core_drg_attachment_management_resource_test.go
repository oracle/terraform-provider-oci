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
	DrgAttachmentManagementRPCRequiredOnlyResource string = DrgAttachmentManagementResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Required, Update, drgAttachmentManagementRepresentationRPC)

	DrgAttachmentManagementIpsecRequiredOnlyResource string = DrgAttachmentManagementResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Required, Update, drgAttachmentManagementRepresentationIpsec)

	ipsecConnectionDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	drgAttachmentManagementRepresentationIpsec = map[string]interface{}{
		"compartment_id":     Representation{RepType: Required, Update: `${var.compartment_id}`},
		"attachment_type":    Representation{RepType: Required, Update: `IPSEC_TUNNEL`},
		"drg_id":             Representation{RepType: Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":       Representation{RepType: Optional, Update: `MyTestDrgAttachmentForTunnel1`},
		"network_id":         Representation{RepType: Optional, Update: `${data.oci_core_ipsec_connection_tunnels.test_ipsec_connection_tunnels.ip_sec_connection_tunnels[0].id}`},
		"drg_route_table_id": Representation{RepType: Optional, Update: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	drgAttachmentManagementRepresentationRPC = map[string]interface{}{
		"compartment_id":     Representation{RepType: Required, Update: `${var.compartment_id}`},
		"attachment_type":    Representation{RepType: Required, Update: `REMOTE_PEERING_CONNECTION`},
		"drg_id":             Representation{RepType: Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":       Representation{RepType: Optional, Update: `MyTestDrgAttachmentForRpc`},
		"network_id":         Representation{RepType: Optional, Update: `${oci_core_remote_peering_connection.test_remote_peering_connection.id}`},
		"drg_route_table_id": Representation{RepType: Optional, Update: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	DrgAttachmentManagementResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Required, Create, cpeRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_core_drg_attachments", "test_drg_attachments", Required, Create, drgAttachmentDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Required, Create, remotePeeringConnectionRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreDrgAttachmentManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment_management.test_drg_attachment_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DrgAttachmentManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Required, Update, drgAttachmentManagementRepresentationRPC), "core", "drgAttachmentManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		//verify, updates to RPC management resource
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Optional, Update, drgAttachmentManagementRepresentationRPC),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyTestDrgAttachmentForRpc"),
				resource.TestCheckResourceAttrSet(resourceName, "network_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies,
		},
		//verify, updates to IPSec management resource
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ipsec_connection_tunnels", Required, Create, ipsecConnectionDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Optional, Update, drgAttachmentManagementRepresentationIpsec),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyTestDrgAttachmentForTunnel1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies,
		},
	})
}
