// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	OperatorAccessControlAccessRequestAuditLogReportSingularDataSourceRepresentation = map[string]interface{}{
		//"access_request_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_operator_access_control_access_request.test_access_request.id}`},
		"access_request_id":   acctest.Representation{RepType: acctest.Required, Create: accessReqId},
		"enable_process_tree": acctest.Representation{RepType: acctest.Required, Create: `10`},
	}

	OperatorAccessControlAccessRequestAuditLogReportResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", acctest.Required, acctest.Create,
		OperatorAccessControlOperatorAccessControlAccessRequestDataSourceRepresentation)
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlAccessRequestAuditLogReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOperatorAccessControlAccessRequestAuditLogReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_operator_access_control_access_request_audit_log_report.test_access_request_audit_log_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request_audit_log_report", "test_access_request_audit_log_report", acctest.Required, acctest.Create, OperatorAccessControlAccessRequestAuditLogReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlAccessRequestAuditLogReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_request_id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "enable_process_tree", "10"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_report_status"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "process_tree"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "report"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_report_generation"),
			),
		},
	})
}
