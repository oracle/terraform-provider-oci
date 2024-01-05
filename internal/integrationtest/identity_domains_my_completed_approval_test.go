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
	IdentityDomainsMyCompletedApprovalSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_completed_approval_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_my_completed_approvals.test_my_completed_approvals.my_completed_approvals.0.id}`},
	}

	IdentityDomainsMyCompletedApprovalDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_completed_approval_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_completed_approval_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":                  acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyCompletedApprovalResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyCompletedApprovalResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyCompletedApprovalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_completed_approvals.test_my_completed_approvals"
	singularDatasourceName := "data.oci_identity_domains_my_completed_approval.test_my_completed_approval"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_completed_approvals", "test_my_completed_approvals", acctest.Required, acctest.Create, IdentityDomainsMyCompletedApprovalDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyCompletedApprovalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "my_completed_approval_count", "10"),
				//resource.TestCheckResourceAttr(datasourceName, "my_completed_approval_filter", "myCompletedApprovalFilter"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_completed_approvals.#"),
				//resource.TestCheckResourceAttr(datasourceName, "my_completed_approvals.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_completed_approvals", "test_my_completed_approvals", acctest.Required, acctest.Create, IdentityDomainsMyCompletedApprovalDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_completed_approval", "test_my_completed_approval", acctest.Required, acctest.Create, IdentityDomainsMyCompletedApprovalSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyCompletedApprovalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_completed_approval_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "justification"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_created_time"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "request_details"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "request_ocid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "response_time"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "tags.#", "1"),
			),
		},
	})
}
