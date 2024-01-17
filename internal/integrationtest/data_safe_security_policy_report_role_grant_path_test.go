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
	DataSafeSecurityPolicyReportRoleGrantPathDataSourceRepresentation = map[string]interface{}{
		"granted_role":              acctest.Representation{RepType: acctest.Required, Create: `IMP_FULL_DATABASE`},
		"grantee":                   acctest.Representation{RepType: acctest.Required, Create: `ADMIN`},
		"security_policy_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_report_id}`},
	}

	DataSafeSecurityPolicyReportRoleGrantPathResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_reports", "test_security_policy_reports", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyReportRoleGrantPathResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityPolicyReportRoleGrantPathResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyReportId := utils.GetEnvSettingWithBlankDefault("security_policy_report_ocid")
	securityPolicyReportIdVariableStr := fmt.Sprintf("variable \"security_policy_report_id\" { default = \"%s\" }\n", securityPolicyReportId)

	datasourceName := "data.oci_data_safe_security_policy_report_role_grant_paths.test_security_policy_report_role_grant_paths"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_report_role_grant_paths", "test_security_policy_report_role_grant_paths", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportRoleGrantPathDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyReportIdVariableStr + DataSafeSecurityPolicyReportRoleGrantPathResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "granted_role", "IMP_FULL_DATABASE"),
				resource.TestCheckResourceAttr(datasourceName, "grantee", "ADMIN"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_report_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "role_grant_path_collection.#"),
			),
		},
	})
}
