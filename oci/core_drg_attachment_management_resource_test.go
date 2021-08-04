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
	DrgAttachmentManagementRPCRequiredOnlyResource string = DrgAttachmentManagementResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Required, Update, drgAttachmentManagementRepresentationRPC)

	DrgAttachmentManagementIpsecRequiredOnlyResource string = DrgAttachmentManagementResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Required, Update, drgAttachmentManagementRepresentationIpsec)

	ipsecConnectionDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	drgAttachmentManagementRepresentationIpsec = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, update: `${var.compartment_id}`},
		"attachment_type":    Representation{repType: Required, update: `IPSEC_TUNNEL`},
		"drg_id":             Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"display_name":       Representation{repType: Optional, update: `MyTestDrgAttachmentForTunnel1`},
		"network_id":         Representation{repType: Optional, update: `${data.oci_core_ipsec_connection_tunnels.test_ipsec_connection_tunnels.ip_sec_connection_tunnels[0].id}`},
		"drg_route_table_id": Representation{repType: Optional, update: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	drgAttachmentManagementRepresentationRPC = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, update: `${var.compartment_id}`},
		"attachment_type":    Representation{repType: Required, update: `REMOTE_PEERING_CONNECTION`},
		"drg_id":             Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"display_name":       Representation{repType: Optional, update: `MyTestDrgAttachmentForRpc`},
		"network_id":         Representation{repType: Optional, update: `${oci_core_remote_peering_connection.test_remote_peering_connection.id}`},
		"drg_route_table_id": Representation{repType: Optional, update: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	DrgAttachmentManagementResourceDependencies = generateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", Required, Create, drgRouteTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", Required, Create, drgAttachmentRepresentation) +
		generateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", Required, Create, ipSecConnectionRepresentation) +
		generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		generateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Required, Create, cpeRepresentation) +
		generateDataSourceFromRepresentationMap("oci_core_drg_attachments", "test_drg_attachments", Required, Create, drgAttachmentDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Required, Create, remotePeeringConnectionRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreDrgAttachmentManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment_management.test_drg_attachment_management"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DrgAttachmentManagementResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Required, Update, drgAttachmentManagementRepresentationRPC), "core", "drgAttachmentManagement", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//verify, updates to RPC management resource
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Optional, Update, drgAttachmentManagementRepresentationRPC),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyTestDrgAttachmentForRpc"),
					resource.TestCheckResourceAttrSet(resourceName, "network_id"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies,
			},
			//verify, updates to IPSec management resource
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ipsec_connection_tunnels", Required, Create, ipsecConnectionDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", Optional, Update, drgAttachmentManagementRepresentationIpsec),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyTestDrgAttachmentForTunnel1"),
					resource.TestCheckResourceAttrSet(resourceName, "network_id"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies,
			},
		},
	})
}
