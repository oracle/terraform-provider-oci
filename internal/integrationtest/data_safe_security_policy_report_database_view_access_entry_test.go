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
	DataSafeSecurityPolicyReportDatabaseViewAccessEntrySingularDataSourceRepresentation = map[string]interface{}{
		"database_view_access_entry_key": acctest.Representation{RepType: acctest.Required, Create: `${var.database_view_access_entry_key}`},
		"security_policy_report_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_report_id}`},
	}

	DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceRepresentation = map[string]interface{}{
		"security_policy_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_report_id}`},
		"scim_query":                acctest.Representation{RepType: acctest.Optional, Create: `scimQuery`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	DataSafeSecurityPolicyReportDatabaseViewAccessEntryResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyReportDatabaseViewAccessEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityPolicyReportDatabaseViewAccessEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyReportId := utils.GetEnvSettingWithBlankDefault("security_policy_report_ocid")
	securityPolicyReportIdVariableStr := fmt.Sprintf("variable \"security_policy_report_id\" { default = \"%s\" }\n", securityPolicyReportId)

	databaseViewAccessEntryKey := utils.GetEnvSettingWithBlankDefault("database_view_access_entrykey")
	databaseViewAccessEntryKeyVariableStr := fmt.Sprintf("variable \"database_view_access_entry_key\" { default = \"%s\" }\n", databaseViewAccessEntryKey)

	datasourceName := "data.oci_data_safe_security_policy_report_database_view_access_entries.test_security_policy_report_database_view_access_entries"
	singularDatasourceName := "data.oci_data_safe_security_policy_report_database_view_access_entry.test_security_policy_report_database_view_access_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_report_database_view_access_entries", "test_security_policy_report_database_view_access_entries", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyReportIdVariableStr + DataSafeSecurityPolicyReportDatabaseViewAccessEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(datasourceName, "scim_query", "scimQuery"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_report_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_view_access_entry_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_report_database_view_access_entry", "test_security_policy_report_database_view_access_entry", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportDatabaseViewAccessEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyReportIdVariableStr + databaseViewAccessEntryKeyVariableStr + DataSafeSecurityPolicyReportDatabaseViewAccessEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_view_access_entry_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_report_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "column_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grant_from_role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grantee"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grantor"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_database_vault"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_real_application_security"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_redaction"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_sql_firewall"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_virtual_private_database"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privilege"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privilege_grantable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privilege_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_schema"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "view_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "view_schema"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "view_text"),
			),
		},
	})
}
