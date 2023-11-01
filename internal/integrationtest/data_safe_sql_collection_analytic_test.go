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
	DataSafeSqlCollectionAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`targetId`}},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `COMPLETED`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"time_ended":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"time_started":              acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
	}

	DataSafeSqlCollectionAnalyticResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlCollectionAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSqlCollectionAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_sql_collection_analytics.test_sql_collection_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_collection_analytics", "test_sql_collection_analytics", acctest.Required, acctest.Create, DataSafeSqlCollectionAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlCollectionAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sql_collection_analytics_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "sql_collection_analytics_collection.0.items.#", "1"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_collection_analytics", "test_sql_collection_analytics", acctest.Optional, acctest.Create, DataSafeSqlCollectionAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSqlCollectionAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sql_collection_analytics_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "sql_collection_analytics_collection.0.items.#", "1"),
			),
		},
	})
}
