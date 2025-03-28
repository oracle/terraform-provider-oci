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
	TenantmanagercontrolplaneRecipientInvitationSingularDataSourceRepresentation = map[string]interface{}{
		"recipient_invitation_id": acctest.Representation{RepType: acctest.Required, Create: `${var.recipient_invitation_id}`},
	}

	TenantmanagercontrolplaneRecipientInvitationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.recipient_invitation_compartment_id}`},
		"sender_tenancy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.sender_tenancy_id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `${var.state}`},
		"status":            acctest.Representation{RepType: acctest.Optional, Create: `${var.status}`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: TenantmanagercontrolplaneRecipientInvitationDataSourceFilterRepresentation}}

	TenantmanagercontrolplaneRecipientInvitationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.recipient_invitation_id}`}},
	}
)

// issue-routing-tag: tenantmanagercontrolplane/default
func TestTenantmanagercontrolplaneRecipientInvitationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestTenantmanagercontrolplaneRecipientInvitationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("recipient_invitation_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"recipient_invitation_compartment_id\" { default = \"%s\" }\n", compartmentId)

	recipientInvitationId := utils.GetEnvSettingWithBlankDefault("recipient_invitation_id")
	recipientInvitationIdVariableStr := fmt.Sprintf("variable \"recipient_invitation_id\" { default = \"%s\" }\n", recipientInvitationId)

	senderTenancyId := utils.GetEnvSettingWithBlankDefault("sender_tenancy_id")
	senderTenancyIdVariableStr := fmt.Sprintf("variable \"sender_tenancy_id\" { default = \"%s\" }\n", senderTenancyId)

	state := utils.GetEnvSettingWithBlankDefault("recipient_invitation_state")
	stateVariableStr := fmt.Sprintf("variable \"state\" { default = \"%s\" }\n", state)

	status := utils.GetEnvSettingWithBlankDefault("recipient_invitation_status")
	statusVariableStr := fmt.Sprintf("variable \"status\" { default = \"%s\" }\n", status)

	datasourceName := "data.oci_tenantmanagercontrolplane_recipient_invitations.test_recipient_invitations"
	singularDatasourceName := "data.oci_tenantmanagercontrolplane_recipient_invitation.test_recipient_invitation"

	acctest.SaveConfigContent("", "", "", t)

	dataSourceConfig := config + compartmentIdVariableStr + recipientInvitationIdVariableStr + senderTenancyIdVariableStr + stateVariableStr + statusVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_recipient_invitations", "test_recipient_invitations", acctest.Optional, acctest.Create, TenantmanagercontrolplaneRecipientInvitationDataSourceRepresentation)
	singularDataSourceConfig := config + recipientInvitationIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_tenantmanagercontrolplane_recipient_invitation", "test_recipient_invitation", acctest.Required, acctest.Create, TenantmanagercontrolplaneRecipientInvitationSingularDataSourceRepresentation)

	fmt.Printf("Data Source Config: %s\n", dataSourceConfig)
	fmt.Printf("Singular Data Source Config: %s\n", singularDataSourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: dataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sender_tenancy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),
				resource.TestCheckResourceAttrSet(datasourceName, "status"),

				resource.TestCheckResourceAttrSet(datasourceName, "recipient_invitation_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recipient_invitation_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recipient_email_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sender_invitation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
