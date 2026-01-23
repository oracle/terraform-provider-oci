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
	MysqlDbSystemMaintenanceEventDataSourceRepresentation = map[string]interface{}{
		"db_system_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"maintenance_action":               acctest.Representation{RepType: acctest.Optional, Create: `DATABASE`},
		"maintenance_status":               acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"maintenance_type":                 acctest.Representation{RepType: acctest.Optional, Create: `AUTOMATIC`},
		"mysql_version_after_maintenance":  acctest.Representation{RepType: acctest.Optional, Create: `mysqlVersionAfterMaintenance`},
		"mysql_version_before_maintenance": acctest.Representation{RepType: acctest.Optional, Create: `mysqlVersionBeforeMaintenance`},
	}

	MysqlDbSystemMaintenanceEventResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, MysqlMysqlDbSystemRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: mysql/default
func TestMysqlDbSystemMaintenanceEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlDbSystemMaintenanceEventResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_mysql_db_system_maintenance_events.test_db_system_maintenance_events"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_db_system_maintenance_events", "test_db_system_maintenance_events", acctest.Required, acctest.Create, MysqlDbSystemMaintenanceEventDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlDbSystemMaintenanceEventResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_events.#"),
			),
		},
	})
}
