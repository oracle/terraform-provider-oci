// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	TenantmanagercontrolplaneSenderInvitationSingularDataSourceRepresentation = map[string]interface{}{
		"sender_invitation_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sender_invitation_id}`},
	}

	TenantmanagercontrolplaneSenderInvitationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.sender_invitation_compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `${var.display_name}`},
		"recipient_tenancy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.recipient_tenancy_id}`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `${var.state}`},
		"status":               acctest.Representation{RepType: acctest.Optional, Create: `${var.status}`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneSenderInvitationDataSourceFilterRepresentation}}
	TenantmanagercontrolplaneSenderInvitationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.sender_invitation_id}`}},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneSenderInvitationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneSenderInvitationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("sender_invitation_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"sender_invitation_compartment_id\" { default = \"%s\" }\n", compartmentId)

	senderInvitationId := utils.GetEnvSettingWithBlankDefault("sender_invitation_id")
	senderInvitationIdVariableStr := fmt.Sprintf("variable \"sender_invitation_id\" { default = \"%s\" }\n", senderInvitationId)

	displayName := utils.GetEnvSettingWithBlankDefault("sender_invitation_display_name")
	displayNameVariableStr := fmt.Sprintf("variable \"display_name\" { default = \"%s\" }\n", displayName)

	recipientTenancyId := utils.GetEnvSettingWithBlankDefault("recipient_tenancy_id")
	recipientTenancyIdVariableStr := fmt.Sprintf("variable \"recipient_tenancy_id\" { default = \"%s\" }\n", recipientTenancyId)

	state := utils.GetEnvSettingWithBlankDefault("sender_invitation_state")
	stateVariableStr := fmt.Sprintf("variable \"state\" { default = \"%s\" }\n", state)

	status := utils.GetEnvSettingWithBlankDefault("sender_invitation_status")
	statusVariableStr := fmt.Sprintf("variable \"status\" { default = \"%s\" }\n", status)

	datasourceName := "data.oci_tenantmanagercontrolplane_sender_invitations.test_sender_invitations"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_sender_invitation.test_sender_invitation"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + senderInvitationIdVariableStr + compartmentIdVariableStr + displayNameVariableStr + recipientTenancyIdVariableStr + stateVariableStr + statusVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_sender_invitations", "test_sender_invitations", acctest.Optional, acctest.Create, TenantmanagercontrolplaneSenderInvitationDataSourceRepresentation)
	singularDataSourceConfig := config + senderInvitationIdVariableStr +
		acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_sender_invitation", "test_sender_invitation", acctest.Required, acctest.Create, TenantmanagercontrolplaneSenderInvitationSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "recipient_tenancy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),
				resource.TestCheckResourceAttrSet(datasourceName, "status"),

				resource.TestCheckResourceAttrSet(datasourceName, "sender_invitation_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sender_invitation_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recipient_invitation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subjects.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
