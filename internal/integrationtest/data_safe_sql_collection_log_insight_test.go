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
	DataSafeSqlCollectionLogInsightDataSourceRepresentation = map[string]interface{}{
		"sql_collection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sqlCollectionId}`},
		"time_ended":        acctest.Representation{RepType: acctest.Required, Create: `2023-09-14T15:04:05Z`},
		"time_started":      acctest.Representation{RepType: acctest.Required, Create: `2023-08-14T15:04:05Z`},
		"group_by":          acctest.Representation{RepType: acctest.Optional, Create: `clientIp`},
	}

	DataSafeSqlCollectionLogInsightResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlCollectionLogInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSqlCollectionLogInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	sqlCollectionId := utils.GetEnvSettingWithBlankDefault("data_safe_sql_collection_generate_fp_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"sqlCollectionId\" { default = \"%s\" }\n", sqlCollectionId)

	datasourceName := "data.oci_data_safe_sql_collection_log_insights.test_sql_collection_log_insights"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_collection_log_insights", "test_sql_collection_log_insights", acctest.Required, acctest.Create, DataSafeSqlCollectionLogInsightDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlCollectionLogInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(datasourceName, "time_ended"),
				//resource.TestCheckResourceAttrSet(datasourceName, "time_started"),

				resource.TestCheckResourceAttrSet(datasourceName, "sql_collection_log_insights_collection.#"),
				//resource.TestCheckResourceAttr(datasourceName, "sql_collection_log_insights_collection.0.items.#", "1"),
			),
		},
	})
}
