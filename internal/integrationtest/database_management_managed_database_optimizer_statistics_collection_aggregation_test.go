// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	managedDatabaseOptimizerStatisticsCollectionAggregationDataSourceRepresentation = map[string]interface{}{
		"group_type":                          acctest.Representation{RepType: acctest.Required, Create: `TASK_STATUS`},
		"managed_database_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
		"start_time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Required, Create: time.Now().UTC().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")},
		"end_time_less_than_or_equal_to":      acctest.Representation{RepType: acctest.Required, Create: time.Now().UTC().Format("2006-01-02T15:04:05.000Z")},
		"limit":                               acctest.Representation{RepType: acctest.Required, Create: `1000`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_optimizer_statistics_collection_aggregations.test_managed_database_optimizer_statistics_collection_aggregations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_optimizer_statistics_collection_aggregations", "test_managed_database_optimizer_statistics_collection_aggregations", acctest.Required, acctest.Create, managedDatabaseOptimizerStatisticsCollectionAggregationDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "group_type", "TASK_STATUS"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "optimizer_statistics_collection_aggregations_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "optimizer_statistics_collection_aggregations_collection.0.items.#"),
			),
		},
	})
}
