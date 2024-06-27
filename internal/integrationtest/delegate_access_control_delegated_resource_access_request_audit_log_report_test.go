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
	DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportSingularDataSourceRepresentation = map[string]interface{}{
		"delegated_resource_access_request_id": acctest.Representation{RepType: acctest.Required, Create: `${var.accReqId}`},
		"is_process_tree_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_requests", "test_delegated_resource_access_requests", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation)
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegatedResourceAccessRequestAuditLogReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegatedResourceAccessRequestAuditLogReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_delegate_access_control_delegated_resource_access_request_audit_log_report.test_delegated_resource_access_request_audit_log_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_request_audit_log_report", "test_delegated_resource_access_request_audit_log_report", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delegated_resource_access_request_id"),
				// Prakash not in output
				//resource.TestCheckResourceAttr(singularDatasourceName, "is_process_tree_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_report_status"),
				// Prakash for future
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "process_tree"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "report"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_report_generated"),
			),
		},
	})
}
