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
	DrgAttachmentManagementRPCRequiredOnlyResource string = DrgAttachmentManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", acctest.Required, acctest.Update, drgAttachmentManagementRepresentationRPC)

	DrgAttachmentManagementIpsecRequiredOnlyResource string = DrgAttachmentManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", acctest.Required, acctest.Update, drgAttachmentManagementRepresentationIpsec)

	ipsecConnectionDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	drgAttachmentManagementRepresentationIpsec = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Update: `${var.compartment_id}`},
		"attachment_type":    acctest.Representation{RepType: acctest.Required, Update: `IPSEC_TUNNEL`},
		"drg_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Update: `MyTestDrgAttachmentForTunnel1`},
		"network_id":         acctest.Representation{RepType: acctest.Optional, Update: `${data.oci_core_ipsec_connection_tunnels.test_ipsec_connection_tunnels.ip_sec_connection_tunnels[0].id}`},
		"drg_route_table_id": acctest.Representation{RepType: acctest.Optional, Update: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	drgAttachmentManagementRepresentationRPC = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"attachment_type":    acctest.Representation{RepType: acctest.Required, Create: `REMOTE_PEERING_CONNECTION`, Update: `REMOTE_PEERING_CONNECTION`},
		"drg_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `MyTestDrgAttachmentForRpc`, Update: `MyTestDrgAttachmentForRpc`},
		"network_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_remote_peering_connection.test_remote_peering_connection.id}`, Update: `${oci_core_remote_peering_connection.test_remote_peering_connection.id}`},
		"drg_route_table_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg_route_table.test_drg_route_table.id}`, Update: `${oci_core_drg_route_table.test_drg_route_table.id}`},
	}

	DrgAttachmentManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, drgRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, drgAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, ipSecConnectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", acctest.Required, acctest.Create, cpeRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_drg_attachments", "test_drg_attachments", acctest.Required, acctest.Create, drgAttachmentDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Required, acctest.Create, remotePeeringConnectionRepresentation)

	// Dependencies for RPC attachment
	DrgAttachmentManagementResourceDependencies2 = acctest.GenerateResourceFromRepresentationMap("oci_core_drg_route_table", "test_drg_route_table", acctest.Required, acctest.Create, drgRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment", "test_drg_attachment", acctest.Required, acctest.Create, drgAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, drgRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Required, acctest.Create, remotePeeringConnectionRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreDrgAttachmentManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment_management.test_drg_attachment_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DrgAttachmentManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", acctest.Required, acctest.Update, drgAttachmentManagementRepresentationRPC), "core", "drgAttachmentManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify, updates to RPC management resource
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", acctest.Optional, acctest.Update, drgAttachmentManagementRepresentationRPC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ipsec_connection_tunnels", acctest.Required, acctest.Create, ipsecConnectionDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "test_drg_attachment_management", acctest.Optional, acctest.Update, drgAttachmentManagementRepresentationIpsec),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

// Verifying Create and Update Requests for RPC attachment
func TestCoreDrgAttachmentManagementRPC(t *testing.T) {
	httpreplay.SetScenario("TestCoreDrgAttachmentManagementRPC")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg_attachment_management.testDrgAttachmentManagement"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DrgAttachmentManagementResourceDependencies2+
		acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "testDrgAttachmentManagement", acctest.Required, acctest.Update, drgAttachmentManagementRepresentationRPC), "core", "drgAttachmentManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify updates to RPC management resource
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies2 +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "testDrgAttachmentManagement", acctest.Optional, acctest.Update, drgAttachmentManagementRepresentationRPC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyTestDrgAttachmentForRpc"),
				resource.TestCheckResourceAttrSet(resourceName, "network_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies2,
		},

		//verify creating an RPC management resource
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies2 +
				acctest.GenerateResourceFromRepresentationMap("oci_core_drg_attachment_management", "testDrgAttachmentManagement", acctest.Optional, acctest.Create, drgAttachmentManagementRepresentationRPC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyTestDrgAttachmentForRpc"),
				resource.TestCheckResourceAttrSet(resourceName, "network_id"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_route_table_id"),
			),
		},

		// delete finally
		{
			Config: config + compartmentIdVariableStr + DrgAttachmentManagementResourceDependencies2,
		},
	})
}
