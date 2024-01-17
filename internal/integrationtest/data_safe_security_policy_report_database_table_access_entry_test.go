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
	DataSafeSecurityPolicyReportDatabaseTableAccessEntrySingularDataSourceRepresentation = map[string]interface{}{
		"database_table_access_entry_key": acctest.Representation{RepType: acctest.Required, Create: `${var.database_table_access_entry_key}`},
		"security_policy_report_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_report_id}`},
	}

	DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceRepresentation = map[string]interface{}{
		"security_policy_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_report_id}`},
		"scim_query":                acctest.Representation{RepType: acctest.Optional, Create: `scimQuery`},
	}

	DataSafeSecurityPolicyReportDatabaseTableAccessEntryResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyReportDatabaseTableAccessEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityPolicyReportDatabaseTableAccessEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyReportId := utils.GetEnvSettingWithBlankDefault("security_policy_report_ocid")
	securityPolicyReportIdVariableStr := fmt.Sprintf("variable \"security_policy_report_id\" { default = \"%s\" }\n", securityPolicyReportId)

	databaseTableAccessEntryKey := utils.GetEnvSettingWithBlankDefault("database_table_access_entrykey")
	databaseTableAccessEntryKeyVariableStr := fmt.Sprintf("variable \"database_table_access_entry_key\" { default = \"%s\" }\n", databaseTableAccessEntryKey)

	datasourceName := "data.oci_data_safe_security_policy_report_database_table_access_entries.test_security_policy_report_database_table_access_entries"
	singularDatasourceName := "data.oci_data_safe_security_policy_report_database_table_access_entry.test_security_policy_report_database_table_access_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_report_database_table_access_entries", "test_security_policy_report_database_table_access_entries", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyReportIdVariableStr + DataSafeSecurityPolicyReportDatabaseTableAccessEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(datasourceName, "scim_query", "scimQuery"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_report_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_table_access_entry_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_report_database_table_access_entry", "test_security_policy_report_database_table_access_entry", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportDatabaseTableAccessEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyReportIdVariableStr + databaseTableAccessEntryKeyVariableStr + DataSafeSecurityPolicyReportDatabaseTableAccessEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_table_access_entry_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_report_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_through_object"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "are_all_tables_accessible"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grantee"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grantor"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_database_vault"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_label_security"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_real_application_security"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_redaction"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_sql_firewall"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_view"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_access_constrained_by_virtual_private_database"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_sensitive"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privilege"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privilege_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "table_schema"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_id"),
			),
		},
	})
}
