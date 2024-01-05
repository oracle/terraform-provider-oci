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
	DatabasemaintenanceRunHistorySingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_run_history_id": acctest.Representation{RepType: acctest.Required, Create: `${var.mr_history_ocid}`},
	}

	DatabasemaintenanceRunHistoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"maintenance_type":     acctest.Representation{RepType: acctest.Optional, Create: `PLANNED`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_resource_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_target_resource.test_target_resource.id}`},
		"target_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `EXACC_INFRASTRUCTURE`},
	}

	DatabaseMaintenanceRunHistoryResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseMaintenanceRunHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMaintenanceRunHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	mrHistoryId := utils.GetEnvSettingWithBlankDefault("mr_history_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	mrHistoryIdVariableStr := fmt.Sprintf("variable \"mr_history_ocid\" { default = \"%s\" }\n", mrHistoryId)

	datasourceName := "data.oci_database_maintenance_run_histories.test_maintenance_run_histories"
	singularDatasourceName := "data.oci_database_maintenance_run_history.test_maintenance_run_history"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_maintenance_run_histories", "test_maintenance_run_histories", acctest.Required, acctest.Create, DatabasemaintenanceRunHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseMaintenanceRunHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"), //This field is not needed in response for maintenance run history
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_run_histories.0.maintenance_run_details.0.maintenance_type", "PLANNED"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_run_histories.0.maintenance_run_details.0.state", "FAILED"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_run_histories.0.maintenance_run_details.0.target_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_run_histories.0.maintenance_run_details.0.target_resource_type", "EXACC_INFRASTRUCTURE"),

				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_run_histories.#"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_run_histories.0.db_servers_history_details.#", "8"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_run_histories.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_run_histories.0.maintenance_run_details.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_maintenance_run_history", "test_maintenance_run_history", acctest.Required, acctest.Create, DatabasemaintenanceRunHistorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + mrHistoryIdVariableStr + DatabaseMaintenanceRunHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_run_history_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "db_servers_history_details.#", "8"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_run_details.#", "1"),
			),
		},
	})
}
