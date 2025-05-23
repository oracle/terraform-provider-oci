// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExadataInfrastructureFleetMetricSingularDataSourceRepresentation = map[string]interface{}{
		"compare_baseline_time": acctest.Representation{RepType: acctest.Required, Create: `${var.compare_baseline_time}`},
		"compare_target_time":   acctest.Representation{RepType: acctest.Required, Create: `${var.compare_target_time}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compare_type":          acctest.Representation{RepType: acctest.Optional, Create: `HOUR`},
		"filter_by_exadata_infrastructure_deployment_type": acctest.Representation{RepType: acctest.Optional, Create: `ONPREMISE`},
		"filter_by_exadata_infrastructure_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `CREATING`},
	}

	DatabaseManagementExadataInfrastructureFleetMetricResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExadataInfrastructureFleetMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExadataInfrastructureFleetMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compareBaselineTime := utils.GetEnvSettingWithBlankDefault("compare_baseline_time")
	compareBaselineTimeStr := fmt.Sprintf("variable \"compare_baseline_time\" { default = \"%s\" }\n", compareBaselineTime)

	compareTargetTime := utils.GetEnvSettingWithBlankDefault("compare_target_time")
	compareTargetTimeStr := fmt.Sprintf("variable \"compare_target_time\" { default = \"%s\" }\n", compareTargetTime)

	//get from tf vars for Comprebaseline and CompareTargetTime and pass it down

	singularDatasourceName := "data.oci_database_management_exadata_infrastructure_fleet_metric.test_exadata_infrastructure_fleet_metric"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_exadata_infrastructure_fleet_metric", "test_exadata_infrastructure_fleet_metric", acctest.Required, acctest.Create, DatabaseManagementExadataInfrastructureFleetMetricSingularDataSourceRepresentation) +
				compartmentIdVariableStr + compareBaselineTimeStr + compareTargetTimeStr + DatabaseManagementExadataInfrastructureFleetMetricResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compare_baseline_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compare_target_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compare_baseline_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "exadata_infrastructure_fleet_summary.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleet_exadata_infrastructures.#", "6"),
			),
		},
	})
}
