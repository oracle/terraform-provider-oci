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
	managedDatabaseAttentionLogCountSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"group_by":                      acctest.Representation{RepType: acctest.Optional, Create: `URGENCY`},
		"is_regular_expression":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"log_search_text":               acctest.Representation{RepType: acctest.Optional, Create: `logSearchText`},
		"time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
		"type_filter":                   acctest.Representation{RepType: acctest.Optional, Create: `UNKNOWN`},
		"urgency_filter":                acctest.Representation{RepType: acctest.Optional, Create: `IMMEDIATE`},
	}

	managedDatabaseAttentionLogCountDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"group_by":                      acctest.Representation{RepType: acctest.Optional, Create: `URGENCY`},
		"is_regular_expression":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"log_search_text":               acctest.Representation{RepType: acctest.Optional, Create: `logSearchText`},
		"time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
		"type_filter":                   acctest.Representation{RepType: acctest.Optional, Create: `UNKNOWN`},
		"urgency_filter":                acctest.Representation{RepType: acctest.Optional, Create: `IMMEDIATE`},
	}

	ManagedDatabaseAttentionLogCountResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseAttentionLogCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseAttentionLogCountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testManagedDatabaseId := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", testManagedDatabaseId)

	datasourceName := "data.oci_database_management_managed_database_attention_log_counts.test_managed_database_attention_log_counts"
	singularDatasourceName := "data.oci_database_management_managed_database_attention_log_count.test_managed_database_attention_log_count"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_attention_log_counts", "test_managed_database_attention_log_counts", acctest.Required, acctest.Create, managedDatabaseAttentionLogCountDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + ManagedDatabaseAttentionLogCountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "attention_log_counts_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "attention_log_counts_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_attention_log_count", "test_managed_database_attention_log_count", acctest.Required, acctest.Create, managedDatabaseAttentionLogCountSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + ManagedDatabaseAttentionLogCountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
