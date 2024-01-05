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
	managedDatabaseOptimizerStatisticsCollectionOperationSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
		"optimizer_statistics_collection_operation_id": acctest.Representation{RepType: acctest.Required, Create: `${element(element(data.oci_database_management_managed_database_optimizer_statistics_collection_operations.test_managed_database_optimizer_statistics_collection_operations.optimizer_statistics_collection_operations_collection, 0).items, 0).id}`},
	}
	managedDatabaseOptimizerStatisticsCollectionOperationDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
		"start_time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Required, Create: time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02T15:04:05.000Z")},
		"end_time_less_than_or_equal_to":      acctest.Representation{RepType: acctest.Required, Create: time.Now().UTC().Format("2006-01-02T15:04:05.000Z")},
		"filter_by":                           acctest.Representation{RepType: acctest.Required, Create: `completedCount >= 5`},
		"limit":                               acctest.Representation{RepType: acctest.Required, Create: `1000`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_optimizer_statistics_collection_operations.test_managed_database_optimizer_statistics_collection_operations"
	singularDatasourceName := "data.oci_database_management_managed_database_optimizer_statistics_collection_operation.test_managed_database_optimizer_statistics_collection_operation"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_optimizer_statistics_collection_operations", "test_managed_database_optimizer_statistics_collection_operations", acctest.Required, acctest.Create, managedDatabaseOptimizerStatisticsCollectionOperationDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "end_time_less_than_or_equal_to"),
				resource.TestCheckResourceAttr(datasourceName, "filter_by", `completedCount >= 5`),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "start_time_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "optimizer_statistics_collection_operations_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_optimizer_statistics_collection_operations", "test_managed_database_optimizer_statistics_collection_operations", acctest.Required, acctest.Create, managedDatabaseOptimizerStatisticsCollectionOperationDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_optimizer_statistics_collection_operation", "test_managed_database_optimizer_statistics_collection_operation", acctest.Required, acctest.Create, managedDatabaseOptimizerStatisticsCollectionOperationSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tasks.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "optimizer_statistics_collection_operation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "completed_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "duration_in_seconds"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "end_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "failed_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "in_progress_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "start_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "timed_out_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_objects_count"),
			),
		},
	})
}
