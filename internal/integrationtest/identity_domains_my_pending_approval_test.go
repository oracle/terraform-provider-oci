// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsMyPendingApprovalSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_pending_approval_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_my_pending_approvals.test_my_pending_approvals.my_pending_approvals.0.id}`},
	}

	IdentityDomainsMyPendingApprovalDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_pending_approval_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_pending_approval_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":                acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	// Note: steps to create myPendingApprovals (can be done in Postman):
	// 1. create an ApprovalWorkflowStep with my user id in approvers
	// 2. create an ApprovalWorkflow with the created ApprovalWorkflowStep
	// 3. create a requestable group
	// 4. create an ApprovalWorkflowAssignment with the created ApprovalWorkflow and the group
	// 5. create an instance of MyRequest with the **applicant's auth token** to join the group
	IdentityDomainsMyPendingApprovalResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyPendingApprovalResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyPendingApprovalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_pending_approvals.test_my_pending_approvals"
	singularDatasourceName := "data.oci_identity_domains_my_pending_approval.test_my_pending_approval"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_pending_approvals", "test_my_pending_approvals", acctest.Required, acctest.Create, IdentityDomainsMyPendingApprovalDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyPendingApprovalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.expires"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.request_created_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.request_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.resource_display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.resource_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_pending_approvals.0.status"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_pending_approvals", "test_my_pending_approvals", acctest.Required, acctest.Create, IdentityDomainsMyPendingApprovalDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_pending_approval", "test_my_pending_approval", acctest.Required, acctest.Create, IdentityDomainsMyPendingApprovalSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyPendingApprovalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_pending_approval_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_created_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
			),
		},
	})
}
